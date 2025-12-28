package stl

import (
	"container/heap"
	"fmt"
)

// ============================================================================
// PRIORITY QUEUE in C++ vs HEAP in Go
// Time Complexity: O(log n) for push/pop, O(1) for top
// ============================================================================

// IntHeap is a min-heap of ints
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// ============================================================================
// STRUCT HEAP IMPLEMENTATIONS (Package level - can be used anywhere)
// ============================================================================

// Task struct for examples
type Task struct {
	Priority int
	Name     string
}

// TaskHeap is a min-heap of Tasks
type TaskHeap []Task

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { return h[i].Priority < h[j].Priority }
func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TaskHeap) Push(x interface{}) {
	*h = append(*h, x.(Task))
}

func (h *TaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MaxTaskHeap is a max-heap of Tasks
type MaxTaskHeap []Task

func (h MaxTaskHeap) Len() int           { return len(h) }
func (h MaxTaskHeap) Less(i, j int) bool { return h[i].Priority > h[j].Priority } // Reverse for max
func (h MaxTaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxTaskHeap) Push(x interface{}) {
	*h = append(*h, x.(Task))
}

func (h *MaxTaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func PriorityQueue() {
	// ========================================================================
	// C++: priority_queue<int> pq;  (max-heap by default)
	// Go:   h := &IntHeap{}; heap.Init(h)  (min-heap by default)
	// ========================================================================
	h := &IntHeap{2, 1, 5, 3, 4}
	heap.Init(h)
	fmt.Println("Initial heap:", *h) // [1 2 5 3 4] (min-heap order)

	// ========================================================================
	// C++: pq.push(0);
	// Go:   heap.Push(h, 0)
	// ========================================================================
	heap.Push(h, 0)
	fmt.Println("After push(0):", *h) // [0 1 5 3 4 2]

	// ========================================================================
	// C++: pq.top()
	// Go:   (*h)[0]  (access first element without removing)
	// ========================================================================
	fmt.Println("Top (min):", (*h)[0])       // 0
	fmt.Println("Size before top:", len(*h)) // Still 6 elements

	// Top doesn't remove the element
	topValue := (*h)[0]
	fmt.Printf("Top value: %d (heap unchanged)\n", topValue)
	fmt.Println("Size after accessing top:", len(*h)) // Still 6

	// ========================================================================
	// C++: pq.pop();
	// Go:   heap.Pop(h)
	// ========================================================================
	top := heap.Pop(h).(int)
	fmt.Printf("Pop: %d, remaining: %v\n", top, *h) // Pop: 0

	// ========================================================================
	// C++: pq.size()
	// Go:   len(*h)
	// ========================================================================
	fmt.Println("Size:", len(*h)) // 5

	// ========================================================================
	// C++: pq.empty()
	// Go:   len(*h) == 0
	// ========================================================================
	fmt.Println("Empty?", len(*h) == 0) // false

	// ========================================================================
	// MAX-HEAP (C++ default behavior)
	// ========================================================================
	fmt.Println("\n=== MAX-HEAP (C++ priority_queue default) ===")
	fmt.Println("For max-heap, use generic PriorityQueue with reversed comparison")
	fmt.Println("See Max-Heap with Structs example below")

	// ========================================================================
	// WITH STRUCTS (Simple Example)
	// ========================================================================
	fmt.Println("\n=== With Struct Types ===")

	// Create and initialize min-heap
	th := &TaskHeap{
		{3, "Low"},
		{1, "High"},
		{2, "Medium"},
	}
	heap.Init(th)

	fmt.Println("Task heap (min by priority):")
	for th.Len() > 0 {
		task := heap.Pop(th).(Task)
		fmt.Printf("  Priority %d: %s\n", task.Priority, task.Name)
	}
	// Output: High (1), Medium (2), Low (3)

	// ========================================================================
	// MAX-HEAP WITH STRUCTS (Simple Example)
	// ========================================================================
	fmt.Println("\n=== Max-Heap with Structs ===")

	// Create and initialize max-heap
	mth := &MaxTaskHeap{
		{3, "Low"},
		{1, "High"},
		{2, "Medium"},
	}
	heap.Init(mth)

	fmt.Println("Task max-heap (max by priority):")
	for mth.Len() > 0 {
		task := heap.Pop(mth).(Task)
		fmt.Printf("  Priority %d: %s\n", task.Priority, task.Name)
	}
	// Output: Low (3), Medium (2), High (1)

	// ========================================================================
	// GETTING TOP WITHOUT POPPING
	// ========================================================================
	fmt.Println("\n=== Getting Top Without Popping ===")

	mth2 := &MaxTaskHeap{
		{5, "Task A"},
		{1, "Task B"},
		{3, "Task C"},
	}
	heap.Init(mth2)

	// Get top without popping - just access first element
	topTask := (*mth2)[0]
	fmt.Printf("Top task: Priority %d, Name %s\n", topTask.Priority, topTask.Name)
	fmt.Println("Size before pop:", len(*mth2)) // 3

	// Now pop it
	poppedTask := heap.Pop(mth2).(Task)
	fmt.Printf("Popped: Priority %d, Name %s\n", poppedTask.Priority, poppedTask.Name)
	fmt.Println("Size after pop:", len(*mth2)) // 2
}
