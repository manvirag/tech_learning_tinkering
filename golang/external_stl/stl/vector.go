package stl

import (
	"fmt"
	"sort"
)

// ============================================================================
// VECTOR in C++ vs SLICE in Go
// Go slices are like C++ vectors - dynamic arrays
// ============================================================================

func Vector() {
	// ========================================================================
	// C++: vector<int> vec;
	// Go:   var vec []int
	// ========================================================================
	var vec []int

	// ========================================================================
	// C++: vec.push_back(1);
	// Go:   vec = append(vec, 1)
	// ========================================================================
	vec = append(vec, 1)
	vec = append(vec, 2)
	vec = append(vec, 3)
	fmt.Println("After push_back:", vec) // [1 2 3]

	// ========================================================================
	// C++: vec.size()
	// Go:   len(vec)
	// ========================================================================
	fmt.Println("Size:", len(vec)) // 3

	// ========================================================================
	// C++: vec[i]
	// Go:   vec[i]
	// ========================================================================
	fmt.Println("vec[0]:", vec[0]) // 1
	fmt.Println("vec[1]:", vec[1]) // 2

	// ========================================================================
	// C++: vec.pop_back()
	// Go:   vec = vec[:len(vec)-1]
	// ========================================================================
	vec = vec[:len(vec)-1]
	fmt.Println("After pop_back:", vec) // [1 2]

	// ========================================================================
	// C++: vec.front()
	// Go:   vec[0]
	// ========================================================================
	fmt.Println("Front:", vec[0]) // 1

	// ========================================================================
	// C++: vec.back()
	// Go:   vec[len(vec)-1]
	// ========================================================================
	fmt.Println("Back:", vec[len(vec)-1]) // 2

	// ========================================================================
	// C++: vec.empty()
	// Go:   len(vec) == 0
	// ========================================================================
	fmt.Println("Empty?", len(vec) == 0) // false

	// ========================================================================
	// C++: vec.clear()
	// Go:   vec = vec[:0]  or  vec = nil
	// ========================================================================
	vec = vec[:0]
	fmt.Println("After clear, Empty?", len(vec) == 0) // true

	// ========================================================================
	// C++: for (int i = 0; i < vec.size(); i++)
	// Go:   for i := 0; i < len(vec); i++
	// ========================================================================
	vec = []int{10, 20, 30}
	for i := 0; i < len(vec); i++ {
		fmt.Printf("vec[%d] = %d\n", i, vec[i])
	}

	// ========================================================================
	// C++: for (auto x : vec)
	// Go:   for _, x := range vec
	// ========================================================================
	for _, x := range vec {
		fmt.Println("Value:", x)
	}

	// ========================================================================
	// WITH STRUCTS
	// ========================================================================
	type Point struct {
		X, Y int
	}

	var points []Point
	points = append(points, Point{1, 2})
	points = append(points, Point{3, 4})

	fmt.Println("\nPoints:")
	for _, p := range points {
		fmt.Printf("  (%d, %d)\n", p.X, p.Y)
	}

	// ========================================================================
	// SORTING & ALGORITHMS
	// ========================================================================
	fmt.Println("\n=== SORTING & ALGORITHMS ===")

	// ========================================================================
	// C++: sort(vec.begin(), vec.end())
	// Go:   sort.Ints(vec)  or  sort.Slice(vec, func(i, j int) bool { ... })
	// ========================================================================
	vec = []int{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Println("Before sort:", vec)
	sort.Ints(vec)
	fmt.Println("After sort:", vec) // [1 1 2 3 4 5 6 9]

	// Sort descending
	vec = []int{3, 1, 4, 1, 5, 9, 2, 6}
	sort.Slice(vec, func(i, j int) bool {
		return vec[i] > vec[j] // descending
	})
	fmt.Println("Sort descending:", vec) // [9 6 5 4 3 2 1 1]

	// Sort structs
	points = []Point{{3, 4}, {1, 2}, {5, 6}, {2, 3}}
	sort.Slice(points, func(i, j int) bool {
		if points[i].X != points[j].X {
			return points[i].X < points[j].X
		}
		return points[i].Y < points[j].Y
	})
	fmt.Println("Sorted points:", points)

	// ========================================================================
	// C++: reverse(vec.begin(), vec.end())
	// Go:   for i, j := 0, len(vec)-1; i < j; i, j = i+1, j-1 { swap }
	// ========================================================================
	vec = []int{1, 2, 3, 4, 5}
	fmt.Println("\nBefore reverse:", vec)
	for i, j := 0, len(vec)-1; i < j; i, j = i+1, j-1 {
		vec[i], vec[j] = vec[j], vec[i]
	}
	fmt.Println("After reverse:", vec) // [5 4 3 2 1]

	// ========================================================================
	// C++: find(vec.begin(), vec.end(), value)
	// Go:   linear search or use sort.Search for sorted
	// ========================================================================
	vec = []int{3, 1, 4, 1, 5, 9, 2, 6}
	target := 5
	found := false
	index := -1
	for i, v := range vec {
		if v == target {
			found = true
			index = i
			break
		}
	}
	fmt.Printf("\nFind %d: found=%v, index=%d\n", target, found, index)

	// Binary search (for sorted array)
	vec = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target = 5
	idx := sort.SearchInts(vec, target)
	if idx < len(vec) && vec[idx] == target {
		fmt.Printf("Binary search %d: found at index %d\n", target, idx)
	} else {
		fmt.Printf("Binary search %d: not found\n", target)
	}

	// ========================================================================
	// C++: min_element, max_element
	// Go:   manual loop or use sort
	// ========================================================================
	vec = []int{3, 1, 4, 1, 5, 9, 2, 6}
	min := vec[0]
	max := vec[0]
	for _, v := range vec {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	fmt.Printf("\nMin: %d, Max: %d\n", min, max)

	// ========================================================================
	// C++: count(vec.begin(), vec.end(), value)
	// Go:   loop and count
	// ========================================================================
	vec = []int{1, 2, 2, 3, 2, 4, 2}
	count := 0
	for _, v := range vec {
		if v == 2 {
			count++
		}
	}
	fmt.Printf("Count of 2: %d\n", count)

	// ========================================================================
	// C++: fill(vec.begin(), vec.end(), value)
	// Go:   loop and assign
	// ========================================================================
	vec = make([]int, 5)
	for i := range vec {
		vec[i] = 42
	}
	fmt.Println("After fill(42):", vec) // [42 42 42 42 42]

	// ========================================================================
	// C++: accumulate(vec.begin(), vec.end(), 0)
	// Go:   loop and sum
	// ========================================================================
	vec = []int{1, 2, 3, 4, 5}
	sum := 0
	for _, v := range vec {
		sum += v
	}
	fmt.Printf("Sum: %d\n", sum)

	// ========================================================================
	// C++: unique(vec.begin(), vec.end())
	// Go:   manual or use map
	// ========================================================================
	vec = []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
	seen := make(map[int]bool)
	unique := []int{}
	for _, v := range vec {
		if !seen[v] {
			seen[v] = true
			unique = append(unique, v)
		}
	}
	fmt.Println("Original:", vec)
	fmt.Println("Unique:", unique)

	// ========================================================================
	// C++: lower_bound, upper_bound (for sorted)
	// Go:   sort.Search
	// ========================================================================
	vec = []int{1, 2, 2, 2, 3, 3, 4, 5}
	target = 2
	lower := sort.Search(len(vec), func(i int) bool {
		return vec[i] >= target
	})
	upper := sort.Search(len(vec), func(i int) bool {
		return vec[i] > target
	})
	fmt.Printf("\nLower bound of %d: index %d\n", target, lower)
	fmt.Printf("Upper bound of %d: index %d\n", target, upper)
}
