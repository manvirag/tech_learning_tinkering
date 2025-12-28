package stl

import (
	"fmt"

	"github.com/emirpasic/gods/maps/treemap"
)

// ============================================================================
// MAP in C++ vs TREEMAP in Go (using gods library)
// Time Complexity: O(log n) for all operations (same as C++ map)
// ============================================================================

func OrderedMap() {
	// ========================================================================
	// C++: map<string, int> m;
	// Go:   m := treemap.NewWithStringComparator()
	// ========================================================================
	m := treemap.NewWithStringComparator()

	// ========================================================================
	// C++: m["apple"] = 5;
	// Go:   m.Put("apple", 5)
	// ========================================================================
	m.Put("apple", 5)
	m.Put("banana", 3)
	m.Put("cherry", 8)
	m.Put("apple", 10) // Overwrites previous value
	fmt.Println("After insert:", m.Size(), "keys")

	// ========================================================================
	// C++: m["apple"]
	// Go:   val, found := m.Get("apple")
	// ========================================================================
	val, found := m.Get("apple")
	if found {
		fmt.Printf("apple: %d\n", val) // 10
	}

	val, found = m.Get("grape")
	if !found {
		fmt.Println("grape: not found")
	}

	// ========================================================================
	// C++: m.find("banana") != m.end()
	// Go:   m.Get("banana") - check found
	// ========================================================================
	_, found = m.Get("banana")
	fmt.Println("Contains 'banana':", found) // true

	// ========================================================================
	// C++: m.erase("banana");
	// Go:   m.Remove("banana")
	// ========================================================================
	m.Remove("banana")
	fmt.Println("After erase 'banana', size:", m.Size()) // 2

	// ========================================================================
	// C++: m.size()
	// Go:   m.Size()
	// ========================================================================
	fmt.Println("Size:", m.Size()) // 2

	// ========================================================================
	// C++: m.empty()
	// Go:   m.Empty()
	// ========================================================================
	fmt.Println("Empty?", m.Empty()) // false

	// ========================================================================
	// C++: m.clear()
	// Go:   m.Clear()
	// ========================================================================
	m.Clear()
	fmt.Println("After clear, Empty?", m.Empty()) // true

	// ========================================================================
	// C++: for (auto [key, value] : m)
	// Go:   it := m.Iterator(); for it.Next() { key, value := it.Key(), it.Value() }
	// ========================================================================
	m = treemap.NewWithStringComparator()
	m.Put("zebra", 1)
	m.Put("apple", 2)
	m.Put("banana", 3)

	fmt.Println("\nIterating map (sorted by key):")
	it := m.Iterator()
	for it.Next() {
		key := it.Key()
		value := it.Value()
		fmt.Printf("  %s: %d\n", key, value)
	}
	// Output: apple: 2, banana: 3, zebra: 1 (sorted)

	// ========================================================================
	// WITH STRUCTS as keys
	// ========================================================================
	fmt.Println("\n=== With Struct Types as Keys ===")

	type Point struct {
		X, Y int
	}

	// Create map with custom comparator for Point key
	pointMap := treemap.NewWith(func(a, b interface{}) int {
		p1 := a.(Point)
		p2 := b.(Point)
		if p1.X != p2.X {
			if p1.X < p2.X {
				return -1
			}
			return 1
		}
		if p1.Y < p2.Y {
			return -1
		} else if p1.Y > p2.Y {
			return 1
		}
		return 0
	})

	// Insert
	pointMap.Put(Point{3, 4}, "C")
	pointMap.Put(Point{1, 2}, "A")
	pointMap.Put(Point{5, 6}, "B")

	fmt.Println("Point map size:", pointMap.Size()) // 3

	// Get value
	val, found = pointMap.Get(Point{1, 2})
	if found {
		fmt.Printf("Value at (1,2): %s\n", val) // "A"
	}

	// Iterate
	fmt.Println("Points (sorted by key):")
	it = pointMap.Iterator()
	for it.Next() {
		key := it.Key().(Point)
		value := it.Value()
		fmt.Printf("  (%d, %d): %s\n", key.X, key.Y, value)
	}

	// Remove
	pointMap.Remove(Point{3, 4})
	fmt.Println("After remove (3,4), size:", pointMap.Size()) // 2

	// ========================================================================
	// WITH STRUCTS as values
	// ========================================================================
	fmt.Println("\n=== With Struct Types as Values ===")

	type Student struct {
		ID   int
		Name string
	}

	studentMap := treemap.NewWithIntComparator()
	studentMap.Put(1, Student{1, "Alice"})
	studentMap.Put(2, Student{2, "Bob"})
	studentMap.Put(3, Student{3, "Charlie"})

	fmt.Println("Student map:")
	it = studentMap.Iterator()
	for it.Next() {
		key := it.Key()
		value := it.Value().(Student)
		fmt.Printf("  ID %d: %s\n", key, value.Name)
	}
}

// ============================================================================
// LOWER_BOUND and UPPER_BOUND workarounds for TreeMap
// ============================================================================

// LowerBound finds the first element >= key (like C++ lower_bound)
// Uses Ceiling() which is equivalent to lower_bound
func MapLowerBound(m *treemap.Map, key interface{}) (foundKey interface{}, foundValue interface{}, found bool) {
	key, value := m.Ceiling(key)
	if key != nil {
		return key, value, true
	}
	return nil, nil, false
}

// UpperBound finds the first element > key (like C++ upper_bound)
// Uses Ceiling and then checks if it equals key, if so gets next element
func MapUpperBound(m *treemap.Map, key interface{}) (foundKey interface{}, foundValue interface{}, found bool) {
	// First find ceiling (lower_bound)
	ceilingKey, ceilingValue := m.Ceiling(key)
	if ceilingKey == nil {
		return nil, nil, false
	}

	// If ceiling equals key, we need the next element
	// Compare using the map's comparator
	if compareMapKeys(ceilingKey, key) == 0 {
		// Key exists, get next element
		it := m.Iterator()
		for it.Next() {
			if compareMapKeys(it.Key(), key) > 0 {
				return it.Key(), it.Value(), true
			}
		}
		return nil, nil, false
	}

	// Ceiling is already > key
	return ceilingKey, ceilingValue, true
}

// Helper function to compare map keys
func compareMapKeys(a, b interface{}) int {
	switch aVal := a.(type) {
	case int:
		if bVal, ok := b.(int); ok {
			if aVal < bVal {
				return -1
			} else if aVal > bVal {
				return 1
			}
			return 0
		}
	case string:
		if bVal, ok := b.(string); ok {
			if aVal < bVal {
				return -1
			} else if aVal > bVal {
				return 1
			}
			return 0
		}
	}
	return 0
}

func MapBoundsExample() {
	fmt.Println("\n=== LOWER_BOUND and UPPER_BOUND in TreeMap ===")

	m := treemap.NewWithIntComparator()
	m.Put(1, "one")
	m.Put(3, "three")
	m.Put(5, "five")
	m.Put(7, "seven")
	m.Put(9, "nine")

	fmt.Println("Map keys:", m.Keys()) // [1 3 5 7 9]

	// Lower bound: first element >= 4
	key, value, found := MapLowerBound(m, 4)
	if found {
		fmt.Printf("Lower bound of 4: key=%d, value=%s\n", key, value) // 5, "five"
	}

	// Lower bound: first element >= 3 (exact match)
	key, value, found = MapLowerBound(m, 3)
	if found {
		fmt.Printf("Lower bound of 3: key=%d, value=%s\n", key, value) // 3, "three"
	}

	// Upper bound: first element > 5
	key, value, found = MapUpperBound(m, 5)
	if found {
		fmt.Printf("Upper bound of 5: key=%d, value=%s\n", key, value) // 7, "seven"
	}

	// Upper bound: first element > 9
	key, value, found = MapUpperBound(m, 9)
	if found {
		fmt.Printf("Upper bound of 9: key=%d, value=%s\n", key, value)
	} else {
		fmt.Println("Upper bound of 9: not found (no element > 9)")
	}

	// Using built-in Ceiling (equivalent to lower_bound)
	fmt.Println("\nUsing built-in Ceiling() (same as lower_bound):")
	key, value = m.Ceiling(4)
	if key != nil {
		fmt.Printf("Ceiling of 4: key=%d, value=%s\n", key, value) // 5, "five"
	}

	// Using built-in Floor (largest key <= given key)
	fmt.Println("\nUsing built-in Floor() (largest <= key):")
	key, value = m.Floor(6)
	if key != nil {
		fmt.Printf("Floor of 6: key=%d, value=%s\n", key, value) // 5, "five"
	}

	// ========================================================================
	// BENEFITS OF LOWER_BOUND/UPPER_BOUND (even with same O(log n))
	// ========================================================================
	fmt.Println("\n=== BENEFITS OF LOWER_BOUND/UPPER_BOUND ===")
	fmt.Println("1. RANGE QUERIES:")
	fmt.Println("   Find all key-value pairs in range [a, b):")
	fmt.Println("   - lower_bound(a) to upper_bound(b)")
	fmt.Println("   - Example: Get all students with ID 10-20")

	fmt.Println("\n2. EFFICIENT RANGE ITERATION:")
	fmt.Println("   - Start from lower_bound, iterate to upper_bound")
	fmt.Println("   - O(k) where k = elements in range, not O(n)")

	fmt.Println("\n3. BINARY SEARCH OPERATIONS:")
	fmt.Println("   - Check if key exists: Get(lower_bound) == key")
	fmt.Println("   - Count keys in range")
	fmt.Println("   - Find closest key")

	fmt.Println("\n4. COMPETITIVE PROGRAMMING:")
	fmt.Println("   - Standard operations for sorted maps")
	fmt.Println("   - Clear semantic intent")
	fmt.Println("   - Essential for many algorithms")

	fmt.Println("\n5. TIME COMPLEXITY:")
	fmt.Println("   - Both O(log n), but lower_bound is MORE EFFICIENT:")
	fmt.Println("     * Direct tree navigation (O(log n))")
	fmt.Println("     * vs Iterating from start (O(n) worst case)")
	fmt.Println("     * vs Contains + iteration (O(log n) + O(n))")
	fmt.Println("   - For range queries: O(log n + k) vs O(n)")

	// Example: Range query
	fmt.Println("\n=== EXAMPLE: Range Query [3, 7) ===")
	lowerKey, _, _ := MapLowerBound(m, 3)
	upperKey, _, upperFound := MapUpperBound(m, 7)

	fmt.Printf("Range [3, 7):\n")
	it := m.Iterator()
	started := false
	for it.Next() {
		k := it.Key().(int)
		v := it.Value().(string)

		// Start from lower_bound
		if !started && k == lowerKey.(int) {
			started = true
		}

		if started {
			fmt.Printf("  %d: %s\n", k, v)
			// Stop at upper_bound
			if upperFound && k == upperKey.(int) {
				break
			}
			if !upperFound || k >= 7 {
				break
			}
		}
	}
}
