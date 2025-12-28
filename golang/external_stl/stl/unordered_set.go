package stl

import "fmt"

// ============================================================================
// UNORDERED_SET in C++ vs MAP in Go
// Go maps can be used as unordered sets (hash set)
// Time Complexity: O(1) average for all operations
// ============================================================================

func UnorderedSet() {
	// ========================================================================
	// C++: unordered_set<int> s;
	// Go:   s := make(map[int]bool)
	// ========================================================================
	s := make(map[int]bool)

	// ========================================================================
	// C++: s.insert(1);
	// Go:   s[1] = true
	// ========================================================================
	s[1] = true
	s[2] = true
	s[3] = true
	s[1] = true                     // Duplicate - no effect
	fmt.Println("After insert:", s) // map[1:true 2:true 3:true]

	// ========================================================================
	// C++: s.find(2) != s.end()  or  s.count(2) > 0
	// Go:   _, exists := s[2]
	// ========================================================================
	if _, exists := s[2]; exists {
		fmt.Println("2 exists in set")
	}
	if _, exists := s[5]; !exists {
		fmt.Println("5 does NOT exist in set")
	}

	// ========================================================================
	// C++: s.erase(2);
	// Go:   delete(s, 2)
	// ========================================================================
	delete(s, 2)
	fmt.Println("After erase 2:", s) // map[1:true 3:true]

	// ========================================================================
	// C++: s.size()
	// Go:   len(s)
	// ========================================================================
	fmt.Println("Size:", len(s)) // 2

	// ========================================================================
	// C++: s.empty()
	// Go:   len(s) == 0
	// ========================================================================
	fmt.Println("Empty?", len(s) == 0) // false

	// ========================================================================
	// C++: s.clear()
	// Go:   s = make(map[int]bool)  or loop and delete
	// ========================================================================
	s = make(map[int]bool)
	fmt.Println("After clear, Empty?", len(s) == 0) // true

	// ========================================================================
	// C++: for (auto x : s)
	// Go:   for x := range s
	// ========================================================================
	s = map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
	fmt.Println("\nIterating set:")
	for x := range s {
		fmt.Printf("  %d\n", x)
	}

	// ========================================================================
	// WITH STRUCTS (need comparable types)
	// ========================================================================
	type Point struct {
		X, Y int
	}

	pointSet := make(map[Point]bool)
	pointSet[Point{1, 2}] = true
	pointSet[Point{3, 4}] = true
	pointSet[Point{1, 2}] = true // Duplicate

	fmt.Println("\nPoint Set:")
	for p := range pointSet {
		fmt.Printf("  (%d, %d)\n", p.X, p.Y)
	}

	// Check if point exists
	if pointSet[Point{3, 4}] {
		fmt.Println("Point (3,4) exists")
	}

	// ========================================================================
	// EXAMPLE: Find unique elements
	// ========================================================================
	fmt.Println("\n=== Example: Find Unique Elements ===")
	arr := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
	unique := make(map[int]bool)

	for _, val := range arr {
		unique[val] = true
	}

	fmt.Println("Original array:", arr)
	fmt.Print("Unique elements: ")
	for val := range unique {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// ========================================================================
	// EXAMPLE: Check if array has duplicates
	// ========================================================================
	fmt.Println("\n=== Example: Check for Duplicates ===")
	arr2 := []int{1, 2, 3, 4, 5}
	seen := make(map[int]bool)
	hasDuplicate := false

	for _, val := range arr2 {
		if seen[val] {
			hasDuplicate = true
			break
		}
		seen[val] = true
	}

	fmt.Println("Array:", arr2)
	fmt.Println("Has duplicates?", hasDuplicate)
}
