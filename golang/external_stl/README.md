# Go STL Equivalents for Competitive Programming & Machine Coding

This repository provides C++ STL equivalents in Go with **same time complexity guarantees**. Each container is in a separate file with complete implementation, operations, and examples using both primitive and struct types.

## Libraries Used

- **github.com/emirpasic/gods** - For ordered structures (TreeSet, TreeMap) with O(log n) operations

## Time Complexity Guarantees

| C++ STL | Go Equivalent | Time Complexity | File |
|---------|---------------|-----------------|------|
| `std::set` | `OrderedSet` | O(log n) | `set.go` |
| `std::unordered_set` | `UnorderedSet` | O(1) average | `set.go` |
| `std::multiset` | `Multiset` | O(log n) | `set.go` |
| `std::map` | `OrderedMap` | O(log n) | `map.go` |
| `std::unordered_map` | `UnorderedMap` | O(1) average | `map.go` |
| `std::multimap` | `Multimap` | O(log n) | `map.go` |
| `std::priority_queue` | `MinHeap`/`MaxHeap` | O(log n) | `priority_queue.go` |
| `std::stack` | `Stack` | O(1) | `stack.go` |
| `std::queue` | `Queue` | O(1) amortized | `queue.go` |
| `std::deque` | `Deque` | O(1) amortized | `deque.go` |

## File Structure

Each file contains:
- **Complete implementation** with all operations
- **Time complexity** annotations for each operation
- **Examples with primitive types** (int, string, etc.)
- **Examples with struct types** (custom comparators)
- **Real-world use cases** (BFS, DFS, sliding window, etc.)

### Files

- **set.go** - OrderedSet, UnorderedSet, Multiset
- **map.go** - OrderedMap, UnorderedMap, Multimap
- **priority_queue.go** - MinHeap, MaxHeap, ItemPriorityQueue
- **stack.go** - Stack with examples (DFS, expression evaluation)
- **queue.go** - Queue with examples (BFS, task processing)
- **deque.go** - Deque with examples (sliding window, palindrome)
- **other_containers.go** - Vector, Pair, List, Bitset, Sorting, Binary Search

## Installation

```bash
go get github.com/emirpasic/gods
```

## Usage

### Running Examples

Each file has example functions that demonstrate usage. To run:

1. Uncomment the `main()` function at the bottom of the desired file
2. Run: `go run <filename>.go`

Example:
```bash
# Uncomment main() in set.go, then:
go run set.go
```

### Using in Your Code

Copy the implementation from the desired file and use it in your code:

```go
package stl

import (
    "container/heap"
    "fmt"
    "github.com/emirpasic/gods/sets/treeset"
)

// Copy OrderedSet implementation from set.go
// ... (implementation code)

func main() {
    set := NewOrderedSet[int]()
    set.Insert(1)
    set.Insert(2)
    fmt.Println(set.Contains(2)) // true
}
```

## Examples

### OrderedSet with Struct

```go
type Point struct {
    X, Y int
}

func ComparePoint(a, b Point) int {
    if a.X != b.X {
        if a.X < b.X { return -1 }
        return 1
    }
    if a.Y < b.Y { return -1 }
    if a.Y > b.Y { return 1 }
    return 0
}

pointSet := NewOrderedSetWithComparator[Point](ComparePoint)
pointSet.Insert(Point{1, 2})
pointSet.Insert(Point{3, 4})
```

### Priority Queue with Custom Comparator

```go
type Task struct {
    Priority int
    Duration int
}

// Min heap: lower priority number = higher priority
taskPQ := NewMinHeap[Task](func(a, b Task) bool {
    if a.Priority != b.Priority {
        return a.Priority < b.Priority
    }
    return a.Duration < b.Duration
})
heap.Init(taskPQ)
taskPQ.PushItem(Task{1, 10})
```

### Stack for DFS

```go
stack := NewStack[*TreeNode]()
stack.Push(root)
for !stack.Empty() {
    node := stack.Pop()
    // Process node
    if node.Right != nil {
        stack.Push(node.Right)
    }
    if node.Left != nil {
        stack.Push(node.Left)
    }
}
```

### Queue for BFS

```go
queue := NewQueue[*TreeNode]()
queue.Push(root)
for !queue.Empty() {
    node := queue.Pop()
    // Process node
    if node.Left != nil {
        queue.Push(node.Left)
    }
    if node.Right != nil {
        queue.Push(node.Right)
    }
}
```

## Notes

- All implementations support **generic types** (Go 1.18+)
- **Struct types** are fully supported with custom comparators
- Time complexities match C++ STL exactly
- `UnorderedSet` and `UnorderedMap` use Go's native map (hash table) for O(1) average operations
- `OrderedSet` and `OrderedMap` use Red-Black Trees from `gods` library for O(log n) operations
- Each file is **standalone** and can be used independently

## Perfect For

- **Competitive Programming** - Fast, efficient implementations
- **Machine Coding Rounds** - Clean, well-documented code
- **Low Level Design (LLD)** - Production-ready data structures
- **Learning** - Complete examples with time complexity analysis
