package models

import (
	"encoding/json"
	"fmt"
)

// OperationType defines the type of operation
type OperationType string

const (
	OpInsert OperationType = "insert"
	OpDelete OperationType = "delete"
	OpRetain OperationType = "retain"
)

// Operation represents a single operation in the document
type Operation struct {
	Type     OperationType `json:"type"`
	Position int           `json:"position,omitempty"` // Position in the document
	Content  string        `json:"content,omitempty"`  // Content for insert operations
	Length   int           `json:"length,omitempty"`   // Length for delete/retain operations
	Author   string        `json:"author"`             // User who performed the operation
	Version  int           `json:"version"`            // Document version when operation was created
}

// OperationList represents a sequence of operations
type OperationList struct {
	Operations []Operation `json:"operations"`
	Version    int         `json:"version"`
	Author     string      `json:"author"`
	Timestamp  int64       `json:"timestamp"`
}

// Apply applies the operation to a document string and returns the new string
func (op *Operation) Apply(document string) (string, error) {
	docRunes := []rune(document)
	docLen := len(docRunes)

	switch op.Type {
	case OpInsert:
		if op.Position < 0 || op.Position > docLen {
			return "", fmt.Errorf("insert position %d out of bounds for document of length %d", op.Position, docLen)
		}

		// Insert content at position
		result := make([]rune, 0, docLen+len([]rune(op.Content)))
		result = append(result, docRunes[:op.Position]...)
		result = append(result, []rune(op.Content)...)
		result = append(result, docRunes[op.Position:]...)
		return string(result), nil

	case OpDelete:
		if op.Position < 0 || op.Position >= docLen {
			return "", fmt.Errorf("delete position %d out of bounds for document of length %d", op.Position, docLen)
		}
		if op.Position+op.Length > docLen {
			return "", fmt.Errorf("delete length %d from position %d exceeds document length %d", op.Length, op.Position, docLen)
		}

		// Delete characters from position
		result := make([]rune, 0, docLen-op.Length)
		result = append(result, docRunes[:op.Position]...)
		result = append(result, docRunes[op.Position+op.Length:]...)
		return string(result), nil

	case OpRetain:
		// Retain doesn't change the document, just moves the cursor
		return document, nil

	default:
		return "", fmt.Errorf("unknown operation type: %s", op.Type)
	}
}

// IsValid checks if the operation is valid
func (op *Operation) IsValid() bool {
	switch op.Type {
	case OpInsert:
		return op.Position >= 0 && op.Content != ""
	case OpDelete:
		return op.Position >= 0 && op.Length > 0
	case OpRetain:
		return op.Position >= 0 && op.Length > 0
	default:
		return false
	}
}

// Clone creates a deep copy of the operation
func (op *Operation) Clone() *Operation {
	return &Operation{
		Type:     op.Type,
		Position: op.Position,
		Content:  op.Content,
		Length:   op.Length,
		Author:   op.Author,
		Version:  op.Version,
	}
}

// String returns a string representation of the operation
func (op *Operation) String() string {
	switch op.Type {
	case OpInsert:
		return fmt.Sprintf("Insert('%s' at %d)", op.Content, op.Position)
	case OpDelete:
		return fmt.Sprintf("Delete(%d chars from %d)", op.Length, op.Position)
	case OpRetain:
		return fmt.Sprintf("Retain(%d chars from %d)", op.Length, op.Position)
	default:
		return fmt.Sprintf("Unknown operation: %+v", op)
	}
}

// ToJSON converts the operation to JSON
func (op *Operation) ToJSON() ([]byte, error) {
	return json.Marshal(op)
}

// FromJSON creates an operation from JSON
func FromJSON(data []byte) (*Operation, error) {
	var op Operation
	err := json.Unmarshal(data, &op)
	return &op, err
}
