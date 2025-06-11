package services

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"drive-mvp/internal/models"

	"github.com/google/uuid"
)

// FileService handles file operations
type FileService struct {
	filesDir     string
	chunkService *ChunkService
}

// NewFileService creates a new file service
func NewFileService(filesDir string, chunkService *ChunkService) *FileService {
	return &FileService{
		filesDir:     filesDir,
		chunkService: chunkService,
	}
}

// UploadFile handles file upload with chunking
func (fs *FileService) UploadFile(name, contentType string, reader io.Reader) (*models.FileMetadata, error) {
	fileID := uuid.New().String()

	// Create chunks
	chunks, err := fs.chunkService.CreateChunks(reader)
	if err != nil {
		return nil, fmt.Errorf("error creating chunks: %w", err)
	}

	// Calculate total size
	var totalSize int64
	for _, chunk := range chunks {
		totalSize += chunk.Size
	}

	// Create file metadata
	fileMetadata := &models.FileMetadata{
		ID:          fileID,
		Name:        name,
		Size:        totalSize,
		ContentType: contentType,
		Chunks:      chunks,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Version:     1,
	}

	// Save metadata
	if err := fs.saveFileMetadata(fileMetadata); err != nil {
		return nil, fmt.Errorf("error saving file metadata: %w", err)
	}

	return fileMetadata, nil
}

// UpdateFile handles file update with differential chunking
func (fs *FileService) UpdateFile(fileID string, reader io.ReadSeeker) (*models.FileMetadata, error) {
	// Load existing file metadata
	existingFile, err := fs.GetFileMetadata(fileID)
	if err != nil {
		return nil, fmt.Errorf("error loading existing file: %w", err)
	}

	// Create new chunks from updated file
	newChunks, err := fs.chunkService.CreateChunks(reader)
	if err != nil {
		return nil, fmt.Errorf("error creating new chunks: %w", err)
	}

	// Find changed chunks
	changedChunks, err := fs.chunkService.CompareChunks(existingFile.Chunks, newChunks)
	if err != nil {
		return nil, fmt.Errorf("error comparing chunks: %w", err)
	}

	// Update only changed chunks
	if len(changedChunks) > 0 {
		if err := fs.chunkService.UpdateChunks(changedChunks, reader); err != nil {
			return nil, fmt.Errorf("error updating chunks: %w", err)
		}
	}

	// Calculate new total size
	var totalSize int64
	for _, chunk := range newChunks {
		totalSize += chunk.Size
	}

	// Update file metadata
	existingFile.Chunks = newChunks
	existingFile.Size = totalSize
	existingFile.UpdatedAt = time.Now()
	existingFile.Version++

	// Save updated metadata
	if err := fs.saveFileMetadata(existingFile); err != nil {
		return nil, fmt.Errorf("error saving updated metadata: %w", err)
	}

	return existingFile, nil
}

// DownloadFile reconstructs and streams a file
func (fs *FileService) DownloadFile(fileID string, writer io.Writer) (*models.FileMetadata, error) {
	// Load file metadata
	fileMetadata, err := fs.GetFileMetadata(fileID)
	if err != nil {
		return nil, fmt.Errorf("error loading file metadata: %w", err)
	}

	// Reconstruct file from chunks
	if err := fs.chunkService.ReconstructFile(fileMetadata.Chunks, writer); err != nil {
		return nil, fmt.Errorf("error reconstructing file: %w", err)
	}

	return fileMetadata, nil
}

// GetFileMetadata loads file metadata
func (fs *FileService) GetFileMetadata(fileID string) (*models.FileMetadata, error) {
	metadataPath := filepath.Join(fs.filesDir, fileID+".json")

	data, err := os.ReadFile(metadataPath)
	if err != nil {
		return nil, fmt.Errorf("error reading metadata file: %w", err)
	}

	var metadata models.FileMetadata
	if err := json.Unmarshal(data, &metadata); err != nil {
		return nil, fmt.Errorf("error unmarshaling metadata: %w", err)
	}

	return &metadata, nil
}

// ListFiles returns all file metadata
func (fs *FileService) ListFiles() ([]models.FileMetadata, error) {
	files, err := os.ReadDir(fs.filesDir)
	if err != nil {
		return nil, fmt.Errorf("error reading files directory: %w", err)
	}

	var fileList []models.FileMetadata
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			fileID := file.Name()[:len(file.Name())-5] // Remove .json extension
			metadata, err := fs.GetFileMetadata(fileID)
			if err != nil {
				continue // Skip files with invalid metadata
			}
			fileList = append(fileList, *metadata)
		}
	}

	return fileList, nil
}

// DeleteFile removes a file and its chunks
func (fs *FileService) DeleteFile(fileID string) error {
	// Load file metadata first
	metadata, err := fs.GetFileMetadata(fileID)
	if err != nil {
		return fmt.Errorf("error loading file metadata: %w", err)
	}

	// Delete metadata file
	metadataPath := filepath.Join(fs.filesDir, fileID+".json")
	if err := os.Remove(metadataPath); err != nil {
		return fmt.Errorf("error deleting metadata file: %w", err)
	}

	// Note: We don't delete chunks here because they might be shared by other files
	// In a production system, you'd want to implement reference counting for chunks
	_ = metadata // Use the metadata if needed for chunk cleanup

	return nil
}

// saveFileMetadata saves file metadata to disk
func (fs *FileService) saveFileMetadata(metadata *models.FileMetadata) error {
	metadataPath := filepath.Join(fs.filesDir, metadata.ID+".json")

	data, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling metadata: %w", err)
	}

	if err := os.WriteFile(metadataPath, data, 0644); err != nil {
		return fmt.Errorf("error writing metadata file: %w", err)
	}

	return nil
}
