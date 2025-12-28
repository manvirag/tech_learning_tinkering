package stl

import "fmt"

// ============================================================================
// DEQUE in C++ vs SLICE in Go
// Go slices can be used as deques (double-ended queue)
// Time Complexity: O(1) amortized for all operations
// ============================================================================

func Deque() {
	// ========================================================================
	// C++: deque<int> dq;
	// Go:   var dq []int
	// ========================================================================
	var dq []int

	// ========================================================================
	// C++: dq.push_back(1);
	// Go:   dq = append(dq, 1)
	// ========================================================================
	dq = append(dq, 1)
	dq = append(dq, 2)
	dq = append(dq, 3)
	fmt.Println("After push_back:", dq) // [1 2 3]

	// ========================================================================
	// C++: dq.push_front(0);
	// Go:   dq = append([]int{0}, dq...)
	// ========================================================================
	dq = append([]int{0}, dq...)
	fmt.Println("After push_front(0):", dq) // [0 1 2 3]

	// ========================================================================
	// C++: dq.front()
	// Go:   dq[0]
	// ========================================================================
	fmt.Println("Front:", dq[0]) // 0

	// ========================================================================
	// C++: dq.back()
	// Go:   dq[len(dq)-1]
	// ========================================================================
	fmt.Println("Back:", dq[len(dq)-1]) // 3

	// ========================================================================
	// C++: dq.pop_back()
	// Go:   dq = dq[:len(dq)-1]
	// ========================================================================
	dq = dq[:len(dq)-1]
	fmt.Println("After pop_back:", dq) // [0 1 2]

	// ========================================================================
	// C++: dq.pop_front()
	// Go:   dq = dq[1:]
	// ========================================================================
	dq = dq[1:]
	fmt.Println("After pop_front:", dq) // [1 2]

	// ========================================================================
	// C++: dq.size()
	// Go:   len(dq)
	// ========================================================================
	fmt.Println("Size:", len(dq)) // 2

	// ========================================================================
	// C++: dq.empty()
	// Go:   len(dq) == 0
	// ========================================================================
	fmt.Println("Empty?", len(dq) == 0) // false

	// ========================================================================
	// C++: dq[i]
	// Go:   dq[i]
	// ========================================================================
	fmt.Println("dq[0]:", dq[0]) // 1
	fmt.Println("dq[1]:", dq[1]) // 2

	// ========================================================================
	// C++: dq.clear()
	// Go:   dq = dq[:0]  or  dq = nil
	// ========================================================================
	dq = dq[:0]
	fmt.Println("After clear, Empty?", len(dq) == 0) // true

	// ========================================================================
	// WITH STRUCTS
	// ========================================================================
	fmt.Println("\n=== With Struct Types ===")

	type Point struct {
		X, Y int
	}

	var pointDeque []Point

	// Push back
	pointDeque = append(pointDeque, Point{1, 2})
	pointDeque = append(pointDeque, Point{3, 4})
	pointDeque = append(pointDeque, Point{5, 6})

	// Push front
	pointDeque = append([]Point{{0, 0}}, pointDeque...)

	fmt.Println("Point deque:")
	for i, p := range pointDeque {
		fmt.Printf("  [%d] = (%d, %d)\n", i, p.X, p.Y)
	}

	// Access front and back
	fmt.Printf("Front: (%d, %d)\n", pointDeque[0].X, pointDeque[0].Y)
	fmt.Printf("Back: (%d, %d)\n", pointDeque[len(pointDeque)-1].X, pointDeque[len(pointDeque)-1].Y)

	// Pop back
	pointDeque = pointDeque[:len(pointDeque)-1]
	fmt.Println("After pop_back, size:", len(pointDeque))

	// Pop front
	pointDeque = pointDeque[1:]
	fmt.Println("After pop_front, size:", len(pointDeque))

	// ========================================================================
	// EXAMPLE: Sliding Window Maximum
	// ========================================================================
	fmt.Println("\n=== Example: Basic Deque Operations ===")

	dq2 := []int{}
	
	// Push operations
	dq2 = append(dq2, 10)           // push_back
	dq2 = append([]int{5}, dq2...)  // push_front
	dq2 = append(dq2, 15)           // push_back
	dq2 = append([]int{0}, dq2...)  // push_front

	fmt.Println("Deque:", dq2) // [0 5 10 15]

	// Access operations
	fmt.Println("Front:", dq2[0])                    // 0
	fmt.Println("Back:", dq2[len(dq2)-1])           // 15
	fmt.Println("Middle element:", dq2[2])          // 10

	// Pop operations
	dq2 = dq2[1:]                    // pop_front
	fmt.Println("After pop_front:", dq2) // [5 10 15]

	dq2 = dq2[:len(dq2)-1]           // pop_back
	fmt.Println("After pop_back:", dq2)  // [5 10]
}

