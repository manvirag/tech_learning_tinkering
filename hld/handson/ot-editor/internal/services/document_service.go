package services

import (
	"fmt"
	"log"
	"sync"
	"time"

	"ot-editor/internal/models"
	"ot-editor/internal/ot"
	"ot-editor/internal/storage"
)

// DocumentService handles document operations with OT
type DocumentService struct {
	docManager    *models.DocumentManager
	clientManager *ClientManager
	storage       storage.DocumentStorage
	autoSave      *storage.AutoSaveManager
	backup        *storage.BackupManager
	mutex         sync.RWMutex
}

// NewDocumentServiceWithStorage creates a new document service with storage
func NewDocumentServiceWithStorage() *DocumentService {
	// Initialize storage
	fileStorage := storage.NewFileStorage("storage/ot-documents")
	autoSave := storage.NewAutoSaveManager(fileStorage, 30*time.Second) // Auto-save every 30 seconds
	backup := storage.NewBackupManager(fileStorage, "storage/ot-backups")

	service := &DocumentService{
		docManager:    models.NewDocumentManager(),
		clientManager: NewClientManager(),
		storage:       fileStorage,
		autoSave:      autoSave,
		backup:        backup,
	}

	// Load existing documents from storage
	service.loadExistingDocuments()

	// Start auto-save
	autoSave.Start()

	return service
}

// loadExistingDocuments loads all documents from storage on startup
func (ds *DocumentService) loadExistingDocuments() {
	if ds.storage == nil {
		return // No storage configured
	}

	documentIDs, err := ds.storage.ListDocuments()
	if err != nil {
		log.Printf("Error listing OT documents from storage: %v", err)
		return
	}

	for _, id := range documentIDs {
		var doc models.Document
		if err := ds.storage.LoadDocument(id, &doc); err != nil {
			log.Printf("Error loading OT document %s: %v", id, err)
			continue
		}

		// Add to document manager
		ds.docManager.AddDocumentFromStorage(id, &doc)
		log.Printf("Loaded OT document %s from storage (version: %d)", id, doc.Version)
	}

	log.Printf("Loaded %d OT documents from storage", len(documentIDs))
}

// saveDocument saves a document to storage
func (ds *DocumentService) saveDocument(doc *models.Document) {
	if ds.storage == nil {
		return // No storage configured
	}

	if err := ds.storage.SaveDocument(doc.ID, doc); err != nil {
		log.Printf("Error saving OT document %s: %v", doc.ID, err)
	}
}

// JoinDocument adds a client to a document
func (ds *DocumentService) JoinDocument(clientID, documentID, clientName string) (*models.Document, *models.Client, error) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	// Get or create document
	doc := ds.docManager.GetOrCreateDocument(documentID, fmt.Sprintf("Document %s", documentID))

	// Create or get client (without websocket connection - that's handled by websocket handler)
	client := models.NewClient(clientID, clientName, nil)

	// Add client to document
	doc.AddClient(client)

	// Register client in client manager
	ds.clientManager.AddClient(client)

	log.Printf("Client %s (%s) joined OT document %s", clientID, clientName, documentID)

	return doc, client, nil
}

// LeaveDocument removes a client from a document
func (ds *DocumentService) LeaveDocument(clientID, documentID string) error {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	// Get document
	doc, exists := ds.docManager.GetDocument(documentID)
	if !exists {
		return fmt.Errorf("document %s not found", documentID)
	}

	// Remove client from document
	doc.RemoveClient(clientID)

	// Unregister client from client manager
	ds.clientManager.RemoveClient(clientID)

	log.Printf("Client %s left OT document %s", clientID, documentID)

	return nil
}

// ApplyOperation applies an operation to a document using OT
func (ds *DocumentService) ApplyOperation(clientID, documentID string, operation *models.Operation) (*models.Operation, error) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	// Get document
	doc, exists := ds.docManager.GetDocument(documentID)
	if !exists {
		return nil, fmt.Errorf("document %s not found", documentID)
	}

	// Validate client
	_, clientExists := doc.GetClient(clientID)
	if !clientExists {
		return nil, fmt.Errorf("client %s not found in document %s", clientID, documentID)
	}

	// Set operation metadata
	operation.Author = clientID
	operation.Version = doc.Version

	// Transform operation against concurrent operations
	// Get operations that happened after the client's version
	concurrentOps := ds.getConcurrentOperations(doc, operation.Version)

	transformedOp := operation
	var err error

	if len(concurrentOps) > 0 {
		transformedOp, err = ot.TransformAgainstOperationList(operation, concurrentOps)
		if err != nil {
			return nil, fmt.Errorf("failed to transform operation: %w", err)
		}
		log.Printf("Operation transformed: %s -> %s", operation.String(), transformedOp.String())
	}

	// Apply the transformed operation to the document
	err = doc.ApplyOperation(transformedOp)
	if err != nil {
		return nil, fmt.Errorf("failed to apply operation: %w", err)
	}

	// Save to storage after operation
	ds.saveDocument(doc)

	log.Printf("Applied OT operation: %s to document %s (new version: %d)",
		transformedOp.String(), documentID, doc.Version)

	return transformedOp, nil
}

// UpdateCursor updates a client's cursor position
func (ds *DocumentService) UpdateCursor(clientID, documentID string, position int) error {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	// Get document
	doc, exists := ds.docManager.GetDocument(documentID)
	if !exists {
		return fmt.Errorf("document %s not found", documentID)
	}

	// Get client
	client, clientExists := doc.GetClient(clientID)
	if !clientExists {
		return fmt.Errorf("client %s not found in document %s", clientID, documentID)
	}

	// Update cursor position
	client.CursorPos = position
	client.LastSeen = time.Now()

	return nil
}

// GetDocumentState returns the current state of a document
func (ds *DocumentService) GetDocumentState(documentID string) (*models.Document, error) {
	ds.mutex.RLock()
	defer ds.mutex.RUnlock()

	doc, exists := ds.docManager.GetDocument(documentID)
	if !exists {
		return nil, fmt.Errorf("document %s not found", documentID)
	}

	return doc, nil
}

// GetDocumentClients returns all clients in a document
func (ds *DocumentService) GetDocumentClients(documentID string) ([]*models.Client, error) {
	ds.mutex.RLock()
	defer ds.mutex.RUnlock()

	doc, exists := ds.docManager.GetDocument(documentID)
	if !exists {
		return nil, fmt.Errorf("document %s not found", documentID)
	}

	return doc.GetAllClients(), nil
}

// BroadcastToDocument sends a message to all clients in a document except the sender
func (ds *DocumentService) BroadcastToDocument(documentID, senderID string, message *models.WebSocketMessage) error {
	clients, err := ds.GetDocumentClients(documentID)
	if err != nil {
		return err
	}

	for _, client := range clients {
		if client.ID != senderID {
			// Send message to client through their connection
			ds.clientManager.SendToClient(client.ID, message)
		}
	}

	return nil
}

// SendToClient sends a message to a specific client
func (ds *DocumentService) SendToClient(clientID string, message *models.WebSocketMessage) error {
	return ds.clientManager.SendToClient(clientID, message)
}

// CreateDocument creates a new document
func (ds *DocumentService) CreateDocument(documentID, title string) *models.Document {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	doc := ds.docManager.CreateDocument(documentID, title)

	// Save to storage immediately
	ds.saveDocument(doc)

	return doc
}

// ListDocuments returns all documents
func (ds *DocumentService) ListDocuments() []*models.Document {
	ds.mutex.RLock()
	defer ds.mutex.RUnlock()

	return ds.docManager.ListDocuments()
}

// DeleteDocument deletes a document
func (ds *DocumentService) DeleteDocument(documentID string) error {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	// Create backup before deletion
	if doc, exists := ds.docManager.GetDocument(documentID); exists && ds.backup != nil {
		ds.backup.BackupDocument(documentID, doc)
	}

	// Get all clients and disconnect them
	if doc, exists := ds.docManager.GetDocument(documentID); exists {
		clients := doc.GetAllClients()
		for _, client := range clients {
			ds.clientManager.RemoveClient(client.ID)
		}
	}

	// Delete from storage
	if ds.storage != nil {
		if err := ds.storage.DeleteDocument(documentID); err != nil {
			log.Printf("Error deleting OT document %s from storage: %v", documentID, err)
		}
	}

	// Delete the document
	if ds.docManager.DeleteDocument(documentID) {
		log.Printf("Deleted OT document %s", documentID)
		return nil
	}

	return fmt.Errorf("document %s not found", documentID)
}

// GetClientManager returns the client manager
func (ds *DocumentService) GetClientManager() *ClientManager {
	return ds.clientManager
}

// Helper function to get concurrent operations
func (ds *DocumentService) getConcurrentOperations(doc *models.Document, clientVersion int) []models.Operation {
	var concurrentOps []models.Operation

	// Get operations that happened after the client's version
	for _, op := range doc.Operations {
		if op.Version > clientVersion {
			concurrentOps = append(concurrentOps, op)
		}
	}

	return concurrentOps
}

// ComposeOperations optimizes a list of operations
func (ds *DocumentService) ComposeOperations(documentID string) error {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	doc, exists := ds.docManager.GetDocument(documentID)
	if !exists {
		return fmt.Errorf("document %s not found", documentID)
	}

	if len(doc.Operations) < 2 {
		return nil // Nothing to compose
	}

	composedOps, err := ot.ComposeOperations(doc.Operations)
	if err != nil {
		return fmt.Errorf("failed to compose operations: %w", err)
	}

	doc.Operations = composedOps

	// Save to storage after composition
	ds.saveDocument(doc)

	log.Printf("Composed operations for OT document %s: %d -> %d operations",
		documentID, len(doc.Operations), len(composedOps))

	return nil
}

// GetOperationHistory returns the operation history for a document
func (ds *DocumentService) GetOperationHistory(documentID string) ([]models.Operation, error) {
	ds.mutex.RLock()
	defer ds.mutex.RUnlock()

	doc, exists := ds.docManager.GetDocument(documentID)
	if !exists {
		return nil, fmt.Errorf("document %s not found", documentID)
	}

	return doc.Operations, nil
}

// ForceSync forces a document to be saved to storage and creates a backup
func (ds *DocumentService) ForceSync(documentID string) error {
	if ds.storage == nil {
		return fmt.Errorf("no storage configured")
	}

	ds.mutex.RLock()
	doc, exists := ds.docManager.GetDocument(documentID)
	ds.mutex.RUnlock()

	if !exists {
		return fmt.Errorf("document %s not found", documentID)
	}

	// Save to storage
	if err := ds.storage.SaveDocument(documentID, doc); err != nil {
		return fmt.Errorf("failed to save document: %v", err)
	}

	// Create backup
	if ds.backup != nil {
		if err := ds.backup.BackupDocument(documentID, doc); err != nil {
			log.Printf("Warning: Failed to backup OT document %s: %v", documentID, err)
		}
	}

	log.Printf("Force synced OT document %s to storage", documentID)
	return nil
}

// Shutdown gracefully shuts down the service
func (ds *DocumentService) Shutdown() {
	log.Println("Shutting down OT document service...")

	if ds.autoSave != nil {
		// Stop auto-save
		ds.autoSave.Stop()
	}

	// Save all documents
	documents := ds.docManager.ListDocuments()
	for _, doc := range documents {
		ds.saveDocument(doc)
	}

	// Cleanup old backups (older than 7 days)
	if ds.backup != nil {
		if err := ds.backup.CleanupOldBackups(7 * 24 * time.Hour); err != nil {
			log.Printf("Error cleaning up old OT backups: %v", err)
		}
	}

	log.Printf("Saved %d OT documents to storage", len(documents))
}
