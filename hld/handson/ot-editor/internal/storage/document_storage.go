package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// DocumentStorage interface for saving/loading documents
type DocumentStorage interface {
	SaveDocument(id string, data interface{}) error
	LoadDocument(id string, target interface{}) error
	DeleteDocument(id string) error
	ListDocuments() ([]string, error)
	DocumentExists(id string) bool
}

// FileStorage implements DocumentStorage using local files
type FileStorage struct {
	basePath string
	mutex    sync.RWMutex
}

// NewFileStorage creates a new file-based storage
func NewFileStorage(basePath string) *FileStorage {
	// Create base directory if it doesn't exist
	if err := os.MkdirAll(basePath, 0755); err != nil {
		log.Printf("Warning: Could not create storage directory %s: %v", basePath, err)
	}

	return &FileStorage{
		basePath: basePath,
	}
}

// SaveDocument saves a document to a JSON file
func (fs *FileStorage) SaveDocument(id string, data interface{}) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	// Create wrapper with metadata
	wrapper := DocumentWrapper{
		ID:      id,
		Data:    data,
		SavedAt: time.Now(),
		Version: 1,
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(wrapper, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal document %s: %v", id, err)
	}

	// Write to file
	filePath := fs.getFilePath(id)
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write document %s to file: %v", id, err)
	}

	log.Printf("Document %s saved to %s", id, filePath)
	return nil
}

// LoadDocument loads a document from a JSON file
func (fs *FileStorage) LoadDocument(id string, target interface{}) error {
	fs.mutex.RLock()
	defer fs.mutex.RUnlock()

	filePath := fs.getFilePath(id)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("document %s not found", id)
	}

	// Read file
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read document %s: %v", id, err)
	}

	// Parse wrapper
	var wrapper DocumentWrapper
	err = json.Unmarshal(jsonData, &wrapper)
	if err != nil {
		return fmt.Errorf("failed to unmarshal document %s: %v", id, err)
	}

	// Extract data to target
	dataBytes, err := json.Marshal(wrapper.Data)
	if err != nil {
		return fmt.Errorf("failed to marshal document data: %v", err)
	}

	err = json.Unmarshal(dataBytes, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to target: %v", err)
	}

	log.Printf("Document %s loaded from %s", id, filePath)
	return nil
}

// DeleteDocument deletes a document file
func (fs *FileStorage) DeleteDocument(id string) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	filePath := fs.getFilePath(id)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("document %s not found", id)
	}

	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("failed to delete document %s: %v", id, err)
	}

	log.Printf("Document %s deleted from %s", id, filePath)
	return nil
}

// ListDocuments returns all document IDs
func (fs *FileStorage) ListDocuments() ([]string, error) {
	fs.mutex.RLock()
	defer fs.mutex.RUnlock()

	files, err := ioutil.ReadDir(fs.basePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read storage directory: %v", err)
	}

	var documentIDs []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			// Remove .json extension to get document ID
			id := file.Name()[:len(file.Name())-5]
			documentIDs = append(documentIDs, id)
		}
	}

	return documentIDs, nil
}

// DocumentExists checks if a document exists
func (fs *FileStorage) DocumentExists(id string) bool {
	fs.mutex.RLock()
	defer fs.mutex.RUnlock()

	filePath := fs.getFilePath(id)
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// getFilePath returns the file path for a document ID
func (fs *FileStorage) getFilePath(id string) string {
	// Sanitize ID to be safe for filesystem
	safeID := filepath.Base(id)
	return filepath.Join(fs.basePath, safeID+".json")
}

// DocumentWrapper wraps documents with metadata
type DocumentWrapper struct {
	ID      string      `json:"id"`
	Data    interface{} `json:"data"`
	SavedAt time.Time   `json:"saved_at"`
	Version int         `json:"version"`
}

// AutoSaveManager handles automatic saving of documents
type AutoSaveManager struct {
	storage  DocumentStorage
	interval time.Duration
	quit     chan bool
	mutex    sync.Mutex
}

// NewAutoSaveManager creates a new auto-save manager
func NewAutoSaveManager(storage DocumentStorage, interval time.Duration) *AutoSaveManager {
	return &AutoSaveManager{
		storage:  storage,
		interval: interval,
		quit:     make(chan bool),
	}
}

// Start begins the auto-save routine
func (asm *AutoSaveManager) Start() {
	go func() {
		ticker := time.NewTicker(asm.interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				// Auto-save logic would be implemented by the calling service
				log.Printf("Auto-save tick (interval: %v)", asm.interval)
			case <-asm.quit:
				log.Println("Auto-save manager stopped")
				return
			}
		}
	}()
	log.Printf("Auto-save manager started with interval: %v", asm.interval)
}

// Stop stops the auto-save routine
func (asm *AutoSaveManager) Stop() {
	asm.mutex.Lock()
	defer asm.mutex.Unlock()

	select {
	case asm.quit <- true:
	default:
	}
}

// BackupManager handles document backups
type BackupManager struct {
	storage    DocumentStorage
	backupPath string
	mutex      sync.Mutex
}

// NewBackupManager creates a new backup manager
func NewBackupManager(storage DocumentStorage, backupPath string) *BackupManager {
	// Create backup directory
	if err := os.MkdirAll(backupPath, 0755); err != nil {
		log.Printf("Warning: Could not create backup directory %s: %v", backupPath, err)
	}

	return &BackupManager{
		storage:    storage,
		backupPath: backupPath,
	}
}

// BackupDocument creates a timestamped backup of a document
func (bm *BackupManager) BackupDocument(id string, data interface{}) error {
	bm.mutex.Lock()
	defer bm.mutex.Unlock()

	timestamp := time.Now().Format("20060102_150405")
	backupID := fmt.Sprintf("%s_backup_%s", id, timestamp)

	wrapper := DocumentWrapper{
		ID:      backupID,
		Data:    data,
		SavedAt: time.Now(),
		Version: 1,
	}

	jsonData, err := json.MarshalIndent(wrapper, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal backup: %v", err)
	}

	backupFile := filepath.Join(bm.backupPath, backupID+".json")
	err = ioutil.WriteFile(backupFile, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write backup: %v", err)
	}

	log.Printf("Document %s backed up to %s", id, backupFile)
	return nil
}

// CleanupOldBackups removes backups older than the specified duration
func (bm *BackupManager) CleanupOldBackups(maxAge time.Duration) error {
	bm.mutex.Lock()
	defer bm.mutex.Unlock()

	files, err := ioutil.ReadDir(bm.backupPath)
	if err != nil {
		return fmt.Errorf("failed to read backup directory: %v", err)
	}

	cutoff := time.Now().Add(-maxAge)
	var cleaned int

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			if file.ModTime().Before(cutoff) {
				backupFile := filepath.Join(bm.backupPath, file.Name())
				if err := os.Remove(backupFile); err != nil {
					log.Printf("Failed to remove old backup %s: %v", backupFile, err)
				} else {
					cleaned++
				}
			}
		}
	}

	if cleaned > 0 {
		log.Printf("Cleaned up %d old backup files", cleaned)
	}

	return nil
}
