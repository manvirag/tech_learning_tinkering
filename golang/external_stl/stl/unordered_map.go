package stl

import "fmt"

// ============================================================================
// UNORDERED_MAP in C++ vs MAP in Go
// Go maps are like unordered_map (hash map)
// Time Complexity: O(1) average for all operations
// ============================================================================

func UnorderedMap() {
	// ========================================================================
	// C++: unordered_map<string, int> m;
	// Go:   m := make(map[string]int)
	// ========================================================================
	m := make(map[string]int)

	// ========================================================================
	// C++: m["apple"] = 5;
	// Go:   m["apple"] = 5
	// ========================================================================
	m["apple"] = 5
	m["banana"] = 3
	m["cherry"] = 8
	fmt.Println("After insert:", m) // map[apple:5 banana:3 cherry:8]

	// ========================================================================
	// C++: m["apple"]
	// Go:   m["apple"]
	// ========================================================================
	fmt.Println("apple count:", m["apple"]) // 5

	// ========================================================================
	// C++: m.find("banana") != m.end()  or  m.count("banana") > 0
	// Go:   val, exists := m["banana"]
	// ========================================================================
	val, exists := m["banana"]
	if exists {
		fmt.Printf("banana exists with value: %d\n", val)
	}

	_, exists = m["grape"]
	if !exists {
		fmt.Println("grape does NOT exist")
	}

	// ========================================================================
	// C++: m.erase("banana");
	// Go:   delete(m, "banana")
	// ========================================================================
	delete(m, "banana")
	fmt.Println("After erase banana:", m) // map[apple:5 cherry:8]

	// ========================================================================
	// C++: m.size()
	// Go:   len(m)
	// ========================================================================
	fmt.Println("Size:", len(m)) // 2

	// ========================================================================
	// C++: m.empty()
	// Go:   len(m) == 0
	// ========================================================================
	fmt.Println("Empty?", len(m) == 0) // false

	// ========================================================================
	// C++: m.clear()
	// Go:   m = make(map[string]int)  or loop and delete
	// ========================================================================
	m = make(map[string]int)
	fmt.Println("After clear, Empty?", len(m) == 0) // true

	// ========================================================================
	// C++: for (auto [key, value] : m)
	// Go:   for key, value := range m
	// ========================================================================
	m = map[string]int{
		"apple":  5,
		"banana": 3,
		"cherry": 8,
		"date":   2,
	}
	fmt.Println("\nIterating map:")
	for key, value := range m {
		fmt.Printf("  %s: %d\n", key, value)
	}

	// ========================================================================
	// C++: for (auto it = m.begin(); it != m.end(); it++)
	// Go:   for key := range m  (keys only)
	// ========================================================================
	fmt.Println("\nKeys only:")
	for key := range m {
		fmt.Printf("  %s\n", key)
	}

	// ========================================================================
	// WITH STRUCTS as keys (need comparable types)
	// ========================================================================
	type Point struct {
		X, Y int
	}

	pointMap := make(map[Point]string)
	pointMap[Point{1, 2}] = "A"
	pointMap[Point{3, 4}] = "B"
	pointMap[Point{5, 6}] = "C"

	fmt.Println("\nPoint Map:")
	for p, label := range pointMap {
		fmt.Printf("  (%d, %d): %s\n", p.X, p.Y, label)
	}

	// ========================================================================
	// WITH STRUCTS as values
	// ========================================================================
	type Student struct {
		ID   int
		Name string
	}

	studentMap := make(map[int]Student)
	studentMap[1] = Student{1, "Alice"}
	studentMap[2] = Student{2, "Bob"}
	studentMap[3] = Student{3, "Charlie"}

	fmt.Println("\nStudent Map:")
	for id, student := range studentMap {
		fmt.Printf("  ID %d: %s\n", id, student.Name)
	}

	// ========================================================================
	// EXAMPLE: Frequency counting
	// ========================================================================
	fmt.Println("\n=== Example: Frequency Counting ===")
	arr := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
	freq := make(map[int]int)

	for _, val := range arr {
		freq[val]++
	}

	fmt.Println("Array:", arr)
	fmt.Println("Frequencies:")
	for val, count := range freq {
		fmt.Printf("  %d: %d times\n", val, count)
	}

	// ========================================================================
	// EXAMPLE: Group by category
	// ========================================================================
	fmt.Println("\n=== Example: Group by Category ===")
	type Item struct {
		Name     string
		Category string
	}

	items := []Item{
		{"apple", "fruit"},
		{"banana", "fruit"},
		{"carrot", "vegetable"},
		{"potato", "vegetable"},
		{"milk", "dairy"},
	}

	grouped := make(map[string][]string)
	for _, item := range items {
		grouped[item.Category] = append(grouped[item.Category], item.Name)
	}

	fmt.Println("Grouped by category:")
	for category, names := range grouped {
		fmt.Printf("  %s: %v\n", category, names)
	}
}
