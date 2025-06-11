package models

import (
	"time"
)

// FileMetadata represents the metadata for a stored file
type FileMetadata struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Size        int64     `json:"size"`
	ContentType string    `json:"content_type"`
	Chunks      []Chunk   `json:"chunks"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Version     int       `json:"version"`
}

// Chunk represents a chunk of a file
type Chunk struct {
	ID     string `json:"id"`
	Hash   string `json:"hash"`   // SHA256 hash for integrity and deduplication
	Size   int64  `json:"size"`   // Size in bytes
	Index  int    `json:"index"`  // Position in the file
	Offset int64  `json:"offset"` // Byte offset in the original file
}

// UploadRequest represents a file upload request
type UploadRequest struct {
	Name        string `json:"name"`
	ContentType string `json:"content_type"`
	Size        int64  `json:"size"`
}

// UploadResponse represents a file upload response
type UploadResponse struct {
	FileID  string `json:"file_id"`
	Message string `json:"message"`
}

// UpdateRequest represents a file update request
type UpdateRequest struct {
	FileID string `json:"file_id"`
	Name   string `json:"name,omitempty"`
}

// FileListResponse represents the response for listing files
type FileListResponse struct {
	Files []FileMetadata `json:"files"`
	Total int            `json:"total"`
}
