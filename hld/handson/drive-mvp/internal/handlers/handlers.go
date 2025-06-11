package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"drive-mvp/internal/models"
	"drive-mvp/internal/services"

	"github.com/gorilla/mux"
)

// Handler holds the file service
type Handler struct {
	fileService *services.FileService
}

// NewHandler creates a new handler
func NewHandler(fileService *services.FileService) *Handler {
	return &Handler{
		fileService: fileService,
	}
}

// UploadHandler handles file uploads
func (h *Handler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form
	err := r.ParseMultipartForm(32 << 20) // 32MB max memory
	if err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Upload the file
	metadata, err := h.fileService.UploadFile(header.Filename, header.Header.Get("Content-Type"), file)
	if err != nil {
		log.Printf("Error uploading file: %v", err)
		http.Error(w, "Error uploading file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.UploadResponse{
		FileID:  metadata.ID,
		Message: "File uploaded successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateHandler handles file updates
func (h *Handler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	fileID := vars["fileId"]

	if fileID == "" {
		http.Error(w, "File ID is required", http.StatusBadRequest)
		return
	}

	// Parse multipart form
	err := r.ParseMultipartForm(32 << 20) // 32MB max memory
	if err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Convert to ReadSeeker
	seeker, ok := file.(io.ReadSeeker)
	if !ok {
		http.Error(w, "File does not support seeking", http.StatusBadRequest)
		return
	}

	// Update the file
	metadata, err := h.fileService.UpdateFile(fileID, seeker)
	if err != nil {
		log.Printf("Error updating file: %v", err)
		http.Error(w, "Error updating file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metadata)
}

// DownloadHandler handles file downloads
func (h *Handler) DownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	fileID := vars["fileId"]

	if fileID == "" {
		http.Error(w, "File ID is required", http.StatusBadRequest)
		return
	}

	// Download the file
	metadata, err := h.fileService.DownloadFile(fileID, w)
	if err != nil {
		log.Printf("Error downloading file: %v", err)
		http.Error(w, "Error downloading file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set appropriate headers
	w.Header().Set("Content-Type", metadata.ContentType)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", metadata.Name))
	w.Header().Set("Content-Length", strconv.FormatInt(metadata.Size, 10))
}

// ListFilesHandler lists all files
func (h *Handler) ListFilesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	files, err := h.fileService.ListFiles()
	if err != nil {
		log.Printf("Error listing files: %v", err)
		http.Error(w, "Error listing files: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.FileListResponse{
		Files: files,
		Total: len(files),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteHandler handles file deletion
func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	fileID := vars["fileId"]

	if fileID == "" {
		http.Error(w, "File ID is required", http.StatusBadRequest)
		return
	}

	err := h.fileService.DeleteFile(fileID)
	if err != nil {
		log.Printf("Error deleting file: %v", err)
		http.Error(w, "Error deleting file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "File deleted successfully"})
}

// HomeHandler serves the web interface
func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/static/index.html")
}
