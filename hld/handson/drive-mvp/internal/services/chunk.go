package services

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"drive-mvp/internal/models"
)

const (
	ChunkSize = 1024 * 1024 // 1MB chunks
)

// ChunkService handles chunking operations
type ChunkService struct {
	chunksDir string
}

// NewChunkService creates a new chunk service
func NewChunkService(chunksDir string) *ChunkService {
	return &ChunkService{
		chunksDir: chunksDir,
	}
}

// CreateChunks splits a file into chunks and returns metadata
func (cs *ChunkService) CreateChunks(reader io.Reader) ([]models.Chunk, error) {
	var chunks []models.Chunk
	buffer := make([]byte, ChunkSize)
	chunkIndex := 0
	var totalOffset int64

	for {
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("error reading data: %w", err)
		}

		if n == 0 {
			break
		}

		chunkData := buffer[:n]

		// Calculate hash
		hash := fmt.Sprintf("%x", sha256.Sum256(chunkData))
		chunkID := fmt.Sprintf("chunk_%s", hash)

		chunk := models.Chunk{
			ID:     chunkID,
			Hash:   hash,
			Size:   int64(n),
			Index:  chunkIndex,
			Offset: totalOffset,
		}

		// Save chunk to disk
		if err := cs.saveChunk(chunkID, chunkData); err != nil {
			return nil, fmt.Errorf("error saving chunk %s: %w", chunkID, err)
		}

		chunks = append(chunks, chunk)
		chunkIndex++
		totalOffset += int64(n)
	}

	return chunks, nil
}

// saveChunk saves a chunk to disk
func (cs *ChunkService) saveChunk(chunkID string, data []byte) error {
	chunkPath := filepath.Join(cs.chunksDir, chunkID)

	// Check if chunk already exists (deduplication)
	if _, err := os.Stat(chunkPath); err == nil {
		return nil // Chunk already exists, skip saving
	}

	file, err := os.Create(chunkPath)
	if err != nil {
		return fmt.Errorf("error creating chunk file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

// ReadChunk reads a chunk from disk
func (cs *ChunkService) ReadChunk(chunkID string) ([]byte, error) {
	chunkPath := filepath.Join(cs.chunksDir, chunkID)
	return os.ReadFile(chunkPath)
}

// CompareChunks compares two sets of chunks and returns which chunks need to be updated
func (cs *ChunkService) CompareChunks(oldChunks, newChunks []models.Chunk) ([]models.Chunk, error) {
	oldChunkMap := make(map[int]string) // index -> hash
	for _, chunk := range oldChunks {
		oldChunkMap[chunk.Index] = chunk.Hash
	}

	var changedChunks []models.Chunk
	for _, newChunk := range newChunks {
		if oldHash, exists := oldChunkMap[newChunk.Index]; !exists || oldHash != newChunk.Hash {
			changedChunks = append(changedChunks, newChunk)
		}
	}

	return changedChunks, nil
}

// UpdateChunks updates only the changed chunks
func (cs *ChunkService) UpdateChunks(changedChunks []models.Chunk, reader io.ReadSeeker) error {
	for _, chunk := range changedChunks {
		// Seek to the chunk's position in the file
		_, err := reader.Seek(chunk.Offset, io.SeekStart)
		if err != nil {
			return fmt.Errorf("error seeking to chunk offset %d: %w", chunk.Offset, err)
		}

		// Read the chunk data
		chunkData := make([]byte, chunk.Size)
		n, err := reader.Read(chunkData)
		if err != nil && err != io.EOF {
			return fmt.Errorf("error reading chunk data: %w", err)
		}

		chunkData = chunkData[:n]

		// Save the updated chunk
		if err := cs.saveChunk(chunk.ID, chunkData); err != nil {
			return fmt.Errorf("error saving updated chunk %s: %w", chunk.ID, err)
		}
	}

	return nil
}

// ReconstructFile reconstructs a file from its chunks
func (cs *ChunkService) ReconstructFile(chunks []models.Chunk, writer io.Writer) error {
	for _, chunk := range chunks {
		chunkData, err := cs.ReadChunk(chunk.ID)
		if err != nil {
			return fmt.Errorf("error reading chunk %s: %w", chunk.ID, err)
		}

		_, err = writer.Write(chunkData)
		if err != nil {
			return fmt.Errorf("error writing chunk data: %w", err)
		}
	}

	return nil
}
