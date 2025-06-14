package handlers

import (
	"encoding/json"
	"net/http"

	"ot-editor/internal/services"

	"github.com/gorilla/mux"
)

// HTTPHandler handles HTTP requests
type HTTPHandler struct {
	documentService *services.DocumentService
}

// NewHTTPHandler creates a new HTTP handler
func NewHTTPHandler(documentService *services.DocumentService) *HTTPHandler {
	return &HTTPHandler{
		documentService: documentService,
	}
}

// HomeHandler serves the main editor page
func (h *HTTPHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/static/index.html")
}

// DocumentHandler serves the document editor for a specific document
func (h *HTTPHandler) DocumentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	documentID := vars["documentId"]

	if documentID == "" {
		http.Error(w, "Document ID is required", http.StatusBadRequest)
		return
	}

	// Serve the same editor page (JavaScript will handle the document ID)
	http.ServeFile(w, r, "web/static/editor.html")
}

// ListDocumentsHandler returns a list of all documents
func (h *HTTPHandler) ListDocumentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	documents := h.documentService.ListDocuments()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"documents": documents,
		"count":     len(documents),
	})
}

// CreateDocumentHandler creates a new document
func (h *HTTPHandler) CreateDocumentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if request.ID == "" {
		http.Error(w, "Document ID is required", http.StatusBadRequest)
		return
	}

	if request.Title == "" {
		request.Title = "Untitled Document"
	}

	document := h.documentService.CreateDocument(request.ID, request.Title)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(document)
}

// GetDocumentHandler returns a specific document
func (h *HTTPHandler) GetDocumentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	documentID := vars["documentId"]

	if documentID == "" {
		http.Error(w, "Document ID is required", http.StatusBadRequest)
		return
	}

	document, err := h.documentService.GetDocumentState(documentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(document)
}

// DeleteDocumentHandler deletes a document
func (h *HTTPHandler) DeleteDocumentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	documentID := vars["documentId"]

	if documentID == "" {
		http.Error(w, "Document ID is required", http.StatusBadRequest)
		return
	}

	err := h.documentService.DeleteDocument(documentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Document deleted successfully"})
}

// GetDocumentHistoryHandler returns the operation history for a document
func (h *HTTPHandler) GetDocumentHistoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	documentID := vars["documentId"]

	if documentID == "" {
		http.Error(w, "Document ID is required", http.StatusBadRequest)
		return
	}

	operations, err := h.documentService.GetOperationHistory(documentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"operations": operations,
		"count":      len(operations),
	})
}

// GetDocumentClientsHandler returns all clients currently connected to a document
func (h *HTTPHandler) GetDocumentClientsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	documentID := vars["documentId"]

	if documentID == "" {
		http.Error(w, "Document ID is required", http.StatusBadRequest)
		return
	}

	clients, err := h.documentService.GetDocumentClients(documentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"clients": clients,
		"count":   len(clients),
	})
}

// HealthHandler returns the health status of the service
func (h *HTTPHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	clientManager := h.documentService.GetClientManager()
	clientCount := clientManager.GetClientCount()
	documents := h.documentService.ListDocuments()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":            "healthy",
		"connected_clients": clientCount,
		"document_count":    len(documents),
	})
}

// StatsHandler returns detailed statistics about the service
func (h *HTTPHandler) StatsHandler(w http.ResponseWriter, r *http.Request) {
	clientManager := h.documentService.GetClientManager()
	documents := h.documentService.ListDocuments()

	// Calculate statistics
	totalOperations := 0
	for _, doc := range documents {
		totalOperations += len(doc.Operations)
	}

	stats := map[string]interface{}{
		"total_documents":   len(documents),
		"connected_clients": clientManager.GetClientCount(),
		"total_operations":  totalOperations,
		"documents":         documents,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
