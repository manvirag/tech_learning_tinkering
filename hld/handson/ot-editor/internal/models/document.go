package models

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Document represents a collaborative document
type Document struct {
	ID         string             `json:"id"`
	Title      string             `json:"title"`
	Content    string             `json:"content"`
	Version    int                `json:"version"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	Operations []Operation        `json:"operations,omitempty"` // History of operations
	Clients    map[string]*Client `json:"-"`                    // Connected clients
	mutex      sync.RWMutex       `json:"-"`                    // Thread safety
}

// Client represents a connected client/user
type Client struct {
	ID        string          `json:"id"`
	Username  string          `json:"username"`
	Color     string          `json:"color"`
	CursorPos int             `json:"cursor_pos"`
	LastSeen  time.Time       `json:"last_seen"`
	Conn      *websocket.Conn `json:"-"` // WebSocket connection
	mutex     sync.RWMutex    `json:"-"`
}

// DocumentManager manages all documents
type DocumentManager struct {
	documents map[string]*Document
	mutex     sync.RWMutex
}

// NewDocumentManager creates a new document manager
func NewDocumentManager() *DocumentManager {
	return &DocumentManager{
		documents: make(map[string]*Document),
	}
}

// NewDocument creates a new document
func NewDocument(id, title string) *Document {
	return &Document{
		ID:         id,
		Title:      title,
		Content:    "",
		Version:    0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Operations: make([]Operation, 0),
		Clients:    make(map[string]*Client),
	}
}

// NewClient creates a new client
func NewClient(id, username string, conn *websocket.Conn) *Client {
	colors := []string{"#FF6B6B", "#4ECDC4", "#45B7D1", "#96CEB4", "#FFEAA7", "#DDA0DD", "#98D8C8", "#F7DC6F"}
	colorIndex := len(id) % len(colors)

	return &Client{
		ID:       id,
		Username: username,
		Color:    colors[colorIndex],
		LastSeen: time.Now(),
		Conn:     conn,
	}
}

// UpdateLastSeen updates the client's last seen time
func (c *Client) UpdateLastSeen() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.LastSeen = time.Now()
}

// GetLastSeen returns the client's last seen time
func (c *Client) GetLastSeen() time.Time {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.LastSeen
}

// AddClient adds a client to the document
func (d *Document) AddClient(client *Client) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Ensure the Clients map is initialized
	if d.Clients == nil {
		d.Clients = make(map[string]*Client)
	}

	d.Clients[client.ID] = client
}

// RemoveClient removes a client from the document
func (d *Document) RemoveClient(clientID string) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	delete(d.Clients, clientID)
}

// GetClient gets a client by ID
func (d *Document) GetClient(clientID string) (*Client, bool) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	client, exists := d.Clients[clientID]
	return client, exists
}

// GetAllClients returns all connected clients
func (d *Document) GetAllClients() []*Client {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	clients := make([]*Client, 0, len(d.Clients))
	for _, client := range d.Clients {
		clients = append(clients, client)
	}
	return clients
}

// ApplyOperation applies an operation to the document
func (d *Document) ApplyOperation(op *Operation) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Apply the operation to the content
	newContent, err := op.Apply(d.Content)
	if err != nil {
		return err
	}

	d.Content = newContent
	d.Version++
	d.UpdatedAt = time.Now()

	// Add to operation history
	op.Version = d.Version
	d.Operations = append(d.Operations, *op)

	return nil
}

// GetState returns the current document state
func (d *Document) GetState() map[string]interface{} {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	return map[string]interface{}{
		"id":      d.ID,
		"title":   d.Title,
		"content": d.Content,
		"version": d.Version,
		"clients": d.GetAllClients(),
	}
}

// Document Manager Methods

// CreateDocument creates a new document
func (dm *DocumentManager) CreateDocument(id, title string) *Document {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()

	doc := NewDocument(id, title)
	dm.documents[id] = doc
	return doc
}

// GetDocument gets a document by ID
func (dm *DocumentManager) GetDocument(id string) (*Document, bool) {
	dm.mutex.RLock()
	defer dm.mutex.RUnlock()

	doc, exists := dm.documents[id]
	return doc, exists
}

// GetOrCreateDocument gets a document or creates it if it doesn't exist
func (dm *DocumentManager) GetOrCreateDocument(id, title string) *Document {
	if doc, exists := dm.GetDocument(id); exists {
		return doc
	}
	return dm.CreateDocument(id, title)
}

// ListDocuments returns all documents
func (dm *DocumentManager) ListDocuments() []*Document {
	dm.mutex.RLock()
	defer dm.mutex.RUnlock()

	docs := make([]*Document, 0, len(dm.documents))
	for _, doc := range dm.documents {
		docs = append(docs, doc)
	}
	return docs
}

// DeleteDocument deletes a document
func (dm *DocumentManager) DeleteDocument(id string) bool {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()

	if _, exists := dm.documents[id]; exists {
		delete(dm.documents, id)
		return true
	}
	return false
}

// AddDocumentFromStorage adds a document loaded from storage
func (dm *DocumentManager) AddDocumentFromStorage(id string, doc *Document) {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()

	// Ensure the Clients map is initialized (storage might not include it)
	if doc.Clients == nil {
		doc.Clients = make(map[string]*Client)
	}

	dm.documents[id] = doc
}
