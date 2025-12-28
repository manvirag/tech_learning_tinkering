package stl

import "fmt"

// ============================================================================
// QUEUE in C++ vs SLICE in Go
// Go slices can be used as queues (FIFO - First In First Out)
// ============================================================================

func Queue() {
	// ========================================================================
	// C++: queue<int> q;
	// Go:   var q []int
	// ========================================================================
	var q []int

	// ========================================================================
	// C++: q.push(1);
	// Go:   q = append(q, 1)
	// ========================================================================
	q = append(q, 1)
	q = append(q, 2)
	q = append(q, 3)
	fmt.Println("After push:", q) // [1 2 3]

	// ========================================================================
	// C++: q.front()
	// Go:   q[0]
	// ========================================================================
	fmt.Println("Front:", q[0]) // 1

	// ========================================================================
	// C++: q.back()
	// Go:   q[len(q)-1]
	// ========================================================================
	fmt.Println("Back:", q[len(q)-1]) // 3

	// ========================================================================
	// C++: q.pop()
	// Go:   q = q[1:]
	// ========================================================================
	q = q[1:]
	fmt.Println("After pop:", q)    // [2 3]
	fmt.Println("Front now:", q[0]) // 2

	// ========================================================================
	// C++: q.size()
	// Go:   len(q)
	// ========================================================================
	fmt.Println("Size:", len(q)) // 2

	// ========================================================================
	// C++: q.empty()
	// Go:   len(q) == 0
	// ========================================================================
	fmt.Println("Empty?", len(q) == 0) // false

	// Pop all
	q = q[1:]
	q = q[1:]
	fmt.Println("After popping all, Empty?", len(q) == 0) // true

	// ========================================================================
	// WITH STRUCTS
	// ========================================================================
	type Task struct {
		ID   int
		Name string
	}

	var taskQueue []Task
	taskQueue = append(taskQueue, Task{1, "Task A"})
	taskQueue = append(taskQueue, Task{2, "Task B"})
	taskQueue = append(taskQueue, Task{3, "Task C"})

	fmt.Println("\nTask Queue:")
	for len(taskQueue) > 0 {
		front := taskQueue[0]
		fmt.Printf("  Processing: ID=%d, Name=%s\n", front.ID, front.Name)
		taskQueue = taskQueue[1:]
	}

	// ========================================================================
	// EXAMPLE: BFS using Queue
	// ========================================================================
	fmt.Println("\n=== Example: BFS Level Order ===")
	// Simulating level-order traversal

	var queue []int
	queue = append(queue, 1) // Start with root

	fmt.Println("BFS Order:")
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("  Visit: %d\n", node)

		// Add children (simplified - in real BFS, you'd have graph/tree structure)
		if node == 1 {
			queue = append(queue, 2, 3)
		} else if node == 2 {
			queue = append(queue, 4, 5)
		} else if node == 3 {
			queue = append(queue, 6)
		}
	}
}
