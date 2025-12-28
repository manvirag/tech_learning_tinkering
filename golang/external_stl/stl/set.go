package stl

import (
	"fmt"

	"github.com/emirpasic/gods/sets/treeset"
)

// ============================================================================
// SET in C++ vs TREESET in Go (using gods library)
// Time Complexity: O(log n) for all operations (same as C++ set)
// ============================================================================

func Set() {
	// ========================================================================
	// C++: set<int> s;
	// Go:   s := treeset.NewWithIntComparator()
	// ========================================================================
	s := treeset.NewWithIntComparator()

	// ========================================================================
	// C++: s.insert(3);
	// Go:   s.Add(3)
	// ========================================================================
	s.Add(3)
	s.Add(1)
	s.Add(2)
	s.Add(3)                                 // Duplicate - won't be added
	fmt.Println("After insert:", s.Values()) // [1 2 3]

	// ========================================================================
	// C++: s.find(2) != s.end()  or  s.count(2) > 0
	// Go:   s.Contains(2)
	// ========================================================================
	fmt.Println("Contains 2:", s.Contains(2)) // true
	fmt.Println("Contains 5:", s.Contains(5)) // false

	// ========================================================================
	// C++: s.erase(2);
	// Go:   s.Remove(2)
	// ========================================================================
	s.Remove(2)
	fmt.Println("After erase 2:", s.Values()) // [1 3]

	// ========================================================================
	// C++: s.size()
	// Go:   s.Size()
	// ========================================================================
	fmt.Println("Size:", s.Size()) // 2

	// ========================================================================
	// C++: s.empty()
	// Go:   s.Empty()
	// ========================================================================
	fmt.Println("Empty?", s.Empty()) // false

	// ========================================================================
	// C++: s.clear()
	// Go:   s.Clear()
	// ========================================================================
	s.Clear()
	fmt.Println("After clear, Empty?", s.Empty()) // true

	// ========================================================================
	// C++: for (auto x : s)
	// Go:   for _, v := range s.Values()
	// ========================================================================
	s = treeset.NewWithIntComparator()
	s.Add(5)
	s.Add(1)
	s.Add(3)
	fmt.Println("\nIterating set:")
	for _, v := range s.Values() {
		fmt.Printf("  %d\n", v)
	}

	// ========================================================================
	// WITH STRUCTS
	// ========================================================================
	fmt.Println("\n=== With Struct Types ===")

	type Point struct {
		X, Y int
	}

	// Create set with custom comparator
	pointSet := treeset.NewWith(func(a, b interface{}) int {
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

	// Insert points
	pointSet.Add(Point{3, 4})
	pointSet.Add(Point{1, 2})
	pointSet.Add(Point{5, 6})
	pointSet.Add(Point{1, 2}) // Duplicate

	fmt.Println("Point set size:", pointSet.Size())                // 3
	fmt.Println("Contains (1,2):", pointSet.Contains(Point{1, 2})) // true

	// Iterate
	fmt.Println("Points (sorted):")
	for _, v := range pointSet.Values() {
		p := v.(Point)
		fmt.Printf("  (%d, %d)\n", p.X, p.Y)
	}

	// Remove
	pointSet.Remove(Point{3, 4})
	fmt.Println("After remove (3,4), size:", pointSet.Size()) // 2
}

// ============================================================================
// WHAT CAN BE VALUES IN TREESET
// ============================================================================

func SetValuesExplained() {
	fmt.Println("\n=== WHAT CAN BE VALUES IN TREESET ===")

	// ========================================================================
	// VALID VALUE TYPES
	// ========================================================================
	fmt.Println("\n1. VALID VALUE TYPES:")

	// ✅ Primitive types - all work with built-in comparators
	fmt.Println("  ✅ Primitive types:")

	// Int
	intSet := treeset.NewWithIntComparator()
	intSet.Add(1)
	intSet.Add(2)
	fmt.Println("    int:", intSet.Values())

	// String
	stringSet := treeset.NewWithStringComparator()
	stringSet.Add("apple")
	stringSet.Add("banana")
	fmt.Println("    string:", stringSet.Values())

	// Float64 (need custom comparator)
	floatSet := treeset.NewWith(func(a, b interface{}) int {
		f1 := a.(float64)
		f2 := b.(float64)
		if f1 < f2 {
			return -1
		} else if f1 > f2 {
			return 1
		}
		return 0
	})
	floatSet.Add(3.14)
	floatSet.Add(2.71)
	fmt.Println("    float64:", floatSet.Values())

	// ✅ Structs - need custom comparator
	fmt.Println("\n  ✅ Structs (with custom comparator):")

	type Student struct {
		ID   int
		Name string
	}

	studentSet := treeset.NewWith(func(a, b interface{}) int {
		s1 := a.(Student)
		s2 := b.(Student)
		if s1.ID < s2.ID {
			return -1
		} else if s1.ID > s2.ID {
			return 1
		}
		return 0
	})

	studentSet.Add(Student{2, "Bob"})
	studentSet.Add(Student{1, "Alice"})
	fmt.Println("    struct:", studentSet.Values())

	// ✅ Arrays (not slices!) - arrays are comparable
	fmt.Println("\n  ✅ Arrays (not slices):")

	arraySet := treeset.NewWith(func(a, b interface{}) int {
		arr1 := a.([3]int)
		arr2 := b.([3]int)
		for i := 0; i < 3; i++ {
			if arr1[i] < arr2[i] {
				return -1
			} else if arr1[i] > arr2[i] {
				return 1
			}
		}
		return 0
	})

	arraySet.Add([3]int{1, 2, 3})
	arraySet.Add([3]int{1, 2, 4})
	fmt.Println("    array:", arraySet.Values())

	// ========================================================================
	// INVALID VALUE TYPES (need workarounds)
	// ========================================================================
	fmt.Println("\n2. INVALID VALUE TYPES (need workarounds):")

	// ❌ Slices - NOT directly comparable
	fmt.Println("  ❌ Slices - cannot use []int directly")
	fmt.Println("     Reason: Slices are not comparable (no == operator)")
	fmt.Println("     ")
	fmt.Println("     Workaround 1: Convert to string")
	sliceSet := treeset.NewWithStringComparator()
	sliceSet.Add(fmt.Sprintf("%v", []int{1, 2, 3}))
	sliceSet.Add(fmt.Sprintf("%v", []int{4, 5, 6}))
	fmt.Println("       Using string representation:", sliceSet.Values())

	fmt.Println("     ")
	fmt.Println("     Workaround 2: Use pointer to slice")
	ptrSliceSet := treeset.NewWith(func(a, b interface{}) int {
		ptr1 := a.(*[]int)
		ptr2 := b.(*[]int)
		// Compare by pointer address (or implement custom logic)
		if fmt.Sprintf("%p", ptr1) < fmt.Sprintf("%p", ptr2) {
			return -1
		} else if fmt.Sprintf("%p", ptr1) > fmt.Sprintf("%p", ptr2) {
			return 1
		}
		return 0
	})
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	ptrSliceSet.Add(&slice1)
	ptrSliceSet.Add(&slice2)
	fmt.Println("       Using pointer:", ptrSliceSet.Size(), "elements")

	// ❌ Maps - NOT comparable
	fmt.Println("\n  ❌ Maps - cannot use map[string]int directly")
	fmt.Println("     Reason: Maps are not comparable")
	fmt.Println("     Workaround: Convert to string or use wrapper struct")

	// ❌ Functions - NOT comparable
	fmt.Println("\n  ❌ Functions - cannot use func() directly")
	fmt.Println("     Reason: Functions are not comparable")

	// ========================================================================
	// KEY REQUIREMENT: COMPARATOR FUNCTION
	// ========================================================================
	fmt.Println("\n3. KEY REQUIREMENT:")
	fmt.Println("   Treeset needs a COMPARATOR function that returns:")
	fmt.Println("     -1 if a < b")
	fmt.Println("      0 if a == b")
	fmt.Println("      1 if a > b")
	fmt.Println("   ")
	fmt.Println("   Built-in comparators available:")
	fmt.Println("     - treeset.NewWithIntComparator()")
	fmt.Println("     - treeset.NewWithStringComparator()")
	fmt.Println("     - treeset.NewWithFloat64Comparator()")
	fmt.Println("   ")
	fmt.Println("   For custom types, provide your own comparator:")
	fmt.Println("     treeset.NewWith(func(a, b interface{}) int { ... })")

	// ========================================================================
	// EXAMPLE: Custom Comparator for Complex Struct
	// ========================================================================
	fmt.Println("\n4. EXAMPLE: Complex Struct with Multiple Fields")

	type Person struct {
		Name string
		Age  int
		City string
	}

	personSet := treeset.NewWith(func(a, b interface{}) int {
		p1 := a.(Person)
		p2 := b.(Person)
		// Compare by Name first, then Age, then City
		if p1.Name < p2.Name {
			return -1
		} else if p1.Name > p2.Name {
			return 1
		}
		if p1.Age < p2.Age {
			return -1
		} else if p1.Age > p2.Age {
			return 1
		}
		if p1.City < p2.City {
			return -1
		} else if p1.City > p2.City {
			return 1
		}
		return 0
	})

	personSet.Add(Person{"Alice", 25, "NYC"})
	personSet.Add(Person{"Bob", 30, "LA"})
	personSet.Add(Person{"Alice", 20, "NYC"})

	fmt.Println("   Person set (sorted by Name, Age, City):")
	for _, v := range personSet.Values() {
		p := v.(Person)
		fmt.Printf("     %s, %d, %s\n", p.Name, p.Age, p.City)
	}
}

// ============================================================================
// LOWER_BOUND and UPPER_BOUND workarounds for TreeSet
// ============================================================================

// LowerBound finds the first element >= key (like C++ lower_bound)
// Time Complexity: O(n) worst case - TreeSet doesn't expose Floor/Ceiling
// Returns the element and true if found, nil and false otherwise
// NOTE: This is O(n) because we must iterate from the beginning.
//
//	For O(log n) operations, consider using TreeMap instead.
//
// WHY IT'S NOT POSSIBLE IN GO:
// 1. TreeSet.tree field is unexported (lowercase) - can't access from outside package
// 2. TreeSet doesn't expose Floor()/Ceiling() methods
// 3. Go's package system prevents accessing unexported fields (even with reflection is limited)
// 4. The underlying red-black tree has Floor/Ceiling, but TreeSet wrapper doesn't expose them
func SetLowerBound(s *treeset.Set, key interface{}) (interface{}, bool) {
	// TreeSet doesn't expose Floor/Ceiling directly (tree field is unexported)
	// Workaround: Iterate from beginning until we find element >= key
	it := s.Iterator()

	// If key exists, return it (O(log n) check)
	if s.Contains(key) {
		return key, true
	}

	// Otherwise, iterate to find first element >= key (O(n) worst case)
	for it.Next() {
		val := it.Value()
		if compareValues(val, key) >= 0 {
			return val, true
		}
	}
	return nil, false
}

// UpperBound finds the first element > key (like C++ upper_bound)
// Time Complexity: O(n) worst case - TreeSet doesn't expose Floor/Ceiling
// Returns the element and true if found, nil and false otherwise
// NOTE: This is O(n) because we must iterate from the beginning.
//
//	If upper_bound doesn't exist (all elements <= key), we iterate through ALL elements.
//	For O(log n) operations, consider using TreeMap instead.
func SetUpperBound(s *treeset.Set, key interface{}) (interface{}, bool) {
	it := s.Iterator()

	// Iterate from beginning to find first element > key
	// Worst case: If no upper_bound exists (all elements <= key),
	// we iterate through ALL n elements - O(n)
	for it.Next() {
		val := it.Value()
		if compareValues(val, key) > 0 {
			return val, true
		}
	}

	// No element > key found (upper_bound doesn't exist)
	return nil, false
}

// Helper function to compare values (simple implementation)
func compareValues(a, b interface{}) int {
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

func SetBoundsExample() {
	fmt.Println("\n=== LOWER_BOUND and UPPER_BOUND in TreeSet ===")

	s := treeset.NewWithIntComparator()
	s.Add(1)
	s.Add(3)
	s.Add(5)
	s.Add(7)
	s.Add(9)

	fmt.Println("Set:", s.Values()) // [1 3 5 7 9]

	// Lower bound: first element >= 4
	lower, found := SetLowerBound(s, 4)
	if found {
		fmt.Printf("Lower bound of 4: %d\n", lower) // 5
	}

	// Lower bound: first element >= 3
	lower, found = SetLowerBound(s, 3)
	if found {
		fmt.Printf("Lower bound of 3: %d\n", lower) // 3 (exact match)
	}

	// Upper bound: first element > 5
	upper, found := SetUpperBound(s, 5)
	if found {
		fmt.Printf("Upper bound of 5: %d\n", upper) // 7
	}

	// Upper bound: first element > 9
	upper, found = SetUpperBound(s, 9)
	if found {
		fmt.Printf("Upper bound of 9: %d\n", upper)
	} else {
		fmt.Println("Upper bound of 9: not found (no element > 9)")
	}

	// ========================================================================
	// BENEFITS OF LOWER_BOUND/UPPER_BOUND (even with same O(log n))
	// ========================================================================
	fmt.Println("\n=== BENEFITS OF LOWER_BOUND/UPPER_BOUND ===")
	fmt.Println("1. RANGE QUERIES:")
	fmt.Println("   Find all elements in range [a, b):")
	fmt.Println("   - lower_bound(a) to upper_bound(b)")
	fmt.Println("   - Essential for range-based operations")

	fmt.Println("\n2. BINARY SEARCH PATTERNS:")
	fmt.Println("   - Check if element exists: lower_bound == upper_bound")
	fmt.Println("   - Count elements in range: distance(lower, upper)")

	fmt.Println("\n3. ITERATOR POSITIONING:")
	fmt.Println("   - Start iteration from specific position")
	fmt.Println("   - Efficient range traversal")

	fmt.Println("\n4. COMPETITIVE PROGRAMMING:")
	fmt.Println("   - Very common operations")
	fmt.Println("   - Clear semantic intent")
	fmt.Println("   - Standard pattern for sorted containers")

	fmt.Println("\n5. TIME COMPLEXITY COMPARISON:")
	fmt.Println("   TreeMap (with Ceiling/Floor):")
	fmt.Println("     - O(log n) - direct tree navigation")
	fmt.Println("     - No need to iterate from beginning")
	fmt.Println("   ")
	fmt.Println("   TreeSet (current implementation):")
	fmt.Println("     - O(n) worst case - must iterate from start")
	fmt.Println("     - If upper_bound doesn't exist: O(n) to check all elements")
	fmt.Println("     - TreeSet doesn't expose Floor/Ceiling (tree field unexported)")
	fmt.Println("   ")
	fmt.Println("   WHY IT'S NOT POSSIBLE IN GO:")
	fmt.Println("     1. Package encapsulation: TreeSet.tree is unexported (lowercase)")
	fmt.Println("     2. No public API: TreeSet doesn't expose Floor()/Ceiling() methods")
	fmt.Println("     3. Go's visibility rules: Can't access unexported fields from outside package")
	fmt.Println("     4. Library design: gods library chose not to expose these methods for TreeSet")
	fmt.Println("   ")
	fmt.Println("   RECOMMENDATION:")
	fmt.Println("     - For O(log n) bounds operations, use TreeMap instead")
	fmt.Println("     - TreeSet bounds are O(n) due to library limitations")
	fmt.Println("     - Or fork/modify the library to add Floor/Ceiling to TreeSet")
}
