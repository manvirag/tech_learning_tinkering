package ot

import (
	"fmt"
	"ot-editor/internal/models"
)

// TransformResult holds the result of transforming two operations
type TransformResult struct {
	Op1Prime *models.Operation // Transformed version of op1
	Op2Prime *models.Operation // Transformed version of op2
}

// Transform transforms two concurrent operations against each other
// This is the core of Operational Transform
func Transform(op1, op2 *models.Operation) (*TransformResult, error) {
	if op1 == nil || op2 == nil {
		return nil, fmt.Errorf("operations cannot be nil")
	}

	// Clone operations to avoid modifying originals
	op1Prime := op1.Clone()
	op2Prime := op2.Clone()

	// Transform based on operation types
	switch {
	case op1.Type == models.OpInsert && op2.Type == models.OpInsert:
		return transformInsertInsert(op1Prime, op2Prime)
	case op1.Type == models.OpInsert && op2.Type == models.OpDelete:
		return transformInsertDelete(op1Prime, op2Prime)
	case op1.Type == models.OpDelete && op2.Type == models.OpInsert:
		result, err := transformInsertDelete(op2Prime, op1Prime)
		if err != nil {
			return nil, err
		}
		// Swap the results since we swapped the order
		return &TransformResult{Op1Prime: result.Op2Prime, Op2Prime: result.Op1Prime}, nil
	case op1.Type == models.OpDelete && op2.Type == models.OpDelete:
		return transformDeleteDelete(op1Prime, op2Prime)
	default:
		return &TransformResult{Op1Prime: op1Prime, Op2Prime: op2Prime}, nil
	}
}

// transformInsertInsert handles transformation of two insert operations
func transformInsertInsert(op1, op2 *models.Operation) (*TransformResult, error) {
	if op1.Position <= op2.Position {
		// op1 is before or at the same position as op2
		// op2's position needs to be adjusted by the length of op1's insert
		op2.Position += len([]rune(op1.Content))
	} else {
		// op2 is before op1
		// op1's position needs to be adjusted by the length of op2's insert
		op1.Position += len([]rune(op2.Content))
	}

	return &TransformResult{Op1Prime: op1, Op2Prime: op2}, nil
}

// transformInsertDelete handles transformation of insert and delete operations
func transformInsertDelete(insert, delete *models.Operation) (*TransformResult, error) {
	if insert.Position <= delete.Position {
		// Insert is before the delete position
		// Delete position needs to be adjusted by the insert length
		delete.Position += len([]rune(insert.Content))
	} else if insert.Position < delete.Position+delete.Length {
		// Insert is within the delete range
		// This is a complex case - the insert position is within what's being deleted
		// We need to adjust the insert position to be at the delete position
		insert.Position = delete.Position
	} else {
		// Insert is after the delete range
		// Insert position needs to be adjusted by the delete length
		insert.Position -= delete.Length
	}

	return &TransformResult{Op1Prime: insert, Op2Prime: delete}, nil
}

// transformDeleteDelete handles transformation of two delete operations
func transformDeleteDelete(op1, op2 *models.Operation) (*TransformResult, error) {
	// Case 1: op1 is completely before op2
	if op1.Position+op1.Length <= op2.Position {
		// op1 is completely before op2
		op2.Position -= op1.Length
		return &TransformResult{Op1Prime: op1, Op2Prime: op2}, nil
	}

	// Case 2: op2 is completely before op1
	if op2.Position+op2.Length <= op1.Position {
		// op2 is completely before op1
		op1.Position -= op2.Length
		return &TransformResult{Op1Prime: op1, Op2Prime: op2}, nil
	}

	// Case 3: Overlapping deletes - this is complex
	// We need to handle the overlapping regions

	// Calculate the overlap
	overlapStart := max(op1.Position, op2.Position)
	overlapEnd := min(op1.Position+op1.Length, op2.Position+op2.Length)
	overlapLength := overlapEnd - overlapStart

	if overlapLength > 0 {
		// There is an overlap
		if op1.Position <= op2.Position {
			// op1 starts first
			op1.Length -= overlapLength
			op2.Position = op1.Position + op1.Length
			op2.Length -= overlapLength
		} else {
			// op2 starts first
			op2.Length -= overlapLength
			op1.Position = op2.Position + op2.Length
			op1.Length -= overlapLength
		}
	}

	return &TransformResult{Op1Prime: op1, Op2Prime: op2}, nil
}

// TransformAgainstOperationList transforms an operation against a list of operations
func TransformAgainstOperationList(op *models.Operation, opList []models.Operation) (*models.Operation, error) {
	result := op.Clone()

	for _, existingOp := range opList {
		transformResult, err := Transform(result, &existingOp)
		if err != nil {
			return nil, fmt.Errorf("failed to transform operation: %w", err)
		}
		result = transformResult.Op1Prime
	}

	return result, nil
}

// ComposeOperations composes a sequence of operations into a minimal set
// This is useful for optimization and reducing the number of operations
func ComposeOperations(ops []models.Operation) ([]models.Operation, error) {
	if len(ops) == 0 {
		return ops, nil
	}

	var result []models.Operation

	for _, op := range ops {
		if len(result) == 0 {
			result = append(result, op)
			continue
		}

		lastOp := &result[len(result)-1]

		// Try to compose with the last operation
		if canCompose(lastOp, &op) {
			composed, err := compose(lastOp, &op)
			if err != nil {
				return nil, err
			}
			result[len(result)-1] = *composed
		} else {
			result = append(result, op)
		}
	}

	return result, nil
}

// canCompose checks if two operations can be composed together
func canCompose(op1, op2 *models.Operation) bool {
	// Same author and consecutive operations
	if op1.Author != op2.Author {
		return false
	}

	// Two consecutive inserts at adjacent positions
	if op1.Type == models.OpInsert && op2.Type == models.OpInsert {
		return op1.Position+len([]rune(op1.Content)) == op2.Position
	}

	// Two consecutive deletes at the same position
	if op1.Type == models.OpDelete && op2.Type == models.OpDelete {
		return op1.Position == op2.Position
	}

	return false
}

// compose combines two operations into one
func compose(op1, op2 *models.Operation) (*models.Operation, error) {
	result := op1.Clone()

	if op1.Type == models.OpInsert && op2.Type == models.OpInsert {
		// Combine the content
		result.Content = op1.Content + op2.Content
		return result, nil
	}

	if op1.Type == models.OpDelete && op2.Type == models.OpDelete {
		// Combine the lengths
		result.Length = op1.Length + op2.Length
		return result, nil
	}

	return nil, fmt.Errorf("cannot compose operations of types %s and %s", op1.Type, op2.Type)
}

// Helper functions
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
