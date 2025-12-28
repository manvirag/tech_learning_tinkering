package stl

import "fmt"

// ============================================================================
// MAP KEYS IN GO - What can and cannot be keys
// ============================================================================

func MapKeysExplained() {
	fmt.Println("=== MAP KEYS IN GO ===")

	// ========================================================================
	// VALID KEY TYPES (Comparable types)
	// ========================================================================
	fmt.Println("\n1. VALID KEY TYPES:")

	// ✅ Primitive types - all work
	intMap := make(map[int]string)
	intMap[1] = "one"
	fmt.Println("  ✅ int:", intMap[1])

	stringMap := make(map[string]int)
	stringMap["hello"] = 42
	fmt.Println("  ✅ string:", stringMap["hello"])

	boolMap := make(map[bool]string)
	boolMap[true] = "yes"
	fmt.Println("  ✅ bool:", boolMap[true])

	// ✅ Arrays (not slices!) - arrays are comparable
	arrayKeyMap := make(map[[3]int]string)
	arrayKeyMap[[3]int{1, 2, 3}] = "array key"
	fmt.Println("  ✅ array:", arrayKeyMap[[3]int{1, 2, 3}])

	// ✅ Structs with comparable fields
	type Point struct {
		X, Y int // both int (comparable)
	}
	pointMap := make(map[Point]string)
	pointMap[Point{1, 2}] = "point"
	fmt.Println("  ✅ struct with comparable fields:", pointMap[Point{1, 2}])

	// ✅ Pointers - pointers are comparable
	x := 10
	ptrMap := make(map[*int]string)
	ptrMap[&x] = "pointer"
	fmt.Println("  ✅ pointer:", ptrMap[&x])

	// ✅ Interfaces - if underlying type is comparable
	var i interface{} = 42
	interfaceMap := make(map[interface{}]string)
	interfaceMap[i] = "interface"
	fmt.Println("  ✅ interface:", interfaceMap[i])

	// ========================================================================
	// INVALID KEY TYPES (Non-comparable types)
	// ========================================================================
	fmt.Println("\n2. INVALID KEY TYPES:")

	// ❌ Slices - NOT comparable
	// sliceMap := make(map[[]int]string)  // COMPILE ERROR!
	fmt.Println("  ❌ slice - cannot use []int as key")
	fmt.Println("     Reason: Slices are not comparable (no == operator)")

	// ❌ Maps - NOT comparable
	// mapMap := make(map[map[string]int]string)  // COMPILE ERROR!
	fmt.Println("  ❌ map - cannot use map[string]int as key")
	fmt.Println("     Reason: Maps are not comparable (no == operator)")

	// ❌ Functions - NOT comparable
	// funcMap := make(map[func()]string)  // COMPILE ERROR!
	fmt.Println("  ❌ function - cannot use func() as key")
	fmt.Println("     Reason: Functions are not comparable")

	// ❌ Structs with non-comparable fields
	type BadStruct struct {
		X int
		Y []int // slice - not comparable!
	}
	// badMap := make(map[BadStruct]string)  // COMPILE ERROR!
	fmt.Println("  ❌ struct with slice/map fields - cannot use as key")
	fmt.Println("     Reason: Contains non-comparable fields")

	// ========================================================================
	// WORKAROUNDS for invalid key types
	// ========================================================================
	fmt.Println("\n3. WORKAROUNDS:")

	// Workaround 1: Convert slice to string (if elements are simple)
	fmt.Println("  Workaround 1: Convert to string representation")
	slice := []int{1, 2, 3}
	// Convert slice to string key (for simple cases)
	sliceAsString := fmt.Sprintf("%v", slice)
	stringKeyMap := make(map[string]int)
	stringKeyMap[sliceAsString] = 42
	fmt.Printf("     Using string key: %s -> %d\n", sliceAsString, stringKeyMap[sliceAsString])

	// Workaround 2: Use pointer to slice (pointers are comparable)
	fmt.Println("  Workaround 2: Use pointer to slice")
	ptrToSliceMap := make(map[*[]int]string)
	ptrToSliceMap[&slice] = "pointer to slice"
	fmt.Println("     Using pointer key:", ptrToSliceMap[&slice])

	// Workaround 3: Use struct with comparable fields
	fmt.Println("  Workaround 3: Use struct wrapper")
	type SliceWrapper struct {
		Data string // store as string or use hash
	}
	wrapperMap := make(map[SliceWrapper]int)
	wrapperMap[SliceWrapper{Data: "[1,2,3]"}] = 100
	fmt.Println("     Using wrapper struct:", wrapperMap[SliceWrapper{Data: "[1,2,3]"}])
}

// ============================================================================
// WHY WE CAN'T MODIFY STRUCT FIELD IN MAP DIRECTLY (vs C++)
// ============================================================================

func StructFieldModificationExplained() {
	fmt.Println("\n=== WHY WE CAN'T MODIFY STRUCT FIELD IN MAP DIRECTLY ===")

	type Person struct {
		Name string
		Age  int
	}

	// ========================================================================
	// C++ Behavior
	// ========================================================================
	fmt.Println("\n1. C++ unordered_map:")
	fmt.Println("   unordered_map<string, Person> m;")
	fmt.Println("   m[\"john\"] = Person{\"John\", 25};")
	fmt.Println("   m[\"john\"].Age = 26;  // ✅ WORKS - modifies directly")
	fmt.Println("   ")
	fmt.Println("   Why it works:")
	fmt.Println("   - m[\"john\"] returns a REFERENCE to the value (not a copy)")
	fmt.Println("   - You can modify fields directly through the reference")
	fmt.Println("   - If key doesn't exist, operator[] creates it with default value")
	fmt.Println("   ")
	fmt.Println("   Example:")
	fmt.Println("     unordered_map<string, Person> m;")
	fmt.Println("     m[\"john\"].Age = 26;  // ✅ Works!")
	fmt.Println("     // If \"john\" doesn't exist, creates Person{} first")
	fmt.Println("     // Then modifies Age to 26")
	fmt.Println("     // This works for both std::map and std::unordered_map")

	// ========================================================================
	// Go Behavior - DOESN'T WORK
	// ========================================================================
	fmt.Println("\n2. Go map - DOESN'T WORK:")
	fmt.Println("   m := make(map[string]Person)")
	fmt.Println("   m[\"john\"] = Person{\"John\", 25}")
	fmt.Println("   m[\"john\"].Age = 26  // ❌ COMPILE ERROR!")
	fmt.Println("   ")
	fmt.Println("   Error: cannot assign to m[\"john\"].Age")
	fmt.Println("   Reason: map values are NOT addressable")

	// ========================================================================
	// Why It Doesn't Work in Go
	// ========================================================================
	fmt.Println("\n3. WHY IT DOESN'T WORK:")

	fmt.Println("   ❌ Map values are NOT addressable")
	fmt.Println("      - When you access m[\"john\"], you get a COPY, not a reference")
	fmt.Println("      - Go doesn't allow taking address of map values")
	fmt.Println("      - This is by design for safety and simplicity")

	fmt.Println("\n   ❌ Cannot take address of map value")
	fmt.Println("      - &m[\"john\"] is NOT allowed")
	fmt.Println("      - This prevents dangling pointers if map is resized")

	// ========================================================================
	// CORRECT WAY IN GO
	// ========================================================================
	fmt.Println("\n4. CORRECT WAY IN GO:")

	m := make(map[string]Person)
	m["john"] = Person{"John", 25}

	// Method 1: Get, modify, put back
	fmt.Println("   Method 1: Get, modify, put back")
	person := m["john"]
	person.Age = 26
	m["john"] = person
	fmt.Printf("   After: %+v\n", m["john"])

	// Method 2: Direct assignment of entire struct
	fmt.Println("\n   Method 2: Direct assignment of entire struct")
	m["john"] = Person{"John", 27}
	fmt.Printf("   After: %+v\n", m["john"])

	// Method 3: Use pointer as value (like C++ reference)
	fmt.Println("\n   Method 3: Use pointer as value (similar to C++ reference)")
	pointerMap := make(map[string]*Person)
	pointerMap["john"] = &Person{"John", 25}
	pointerMap["john"].Age = 28 // ✅ WORKS - because value is a pointer
	fmt.Printf("   After: %+v\n", *pointerMap["john"])

	// ========================================================================
	// Comparison
	// ========================================================================
	fmt.Println("\n5. COMPARISON:")

	fmt.Println("   C++:")
	fmt.Println("     map<string, Person> m;")
	fmt.Println("     m[\"john\"] = Person{\"John\", 25};")
	fmt.Println("     m[\"john\"].Age = 26;  // ✅ Direct modification")

	fmt.Println("\n   Go (struct value):")
	fmt.Println("     m := make(map[string]Person)")
	fmt.Println("     m[\"john\"] = Person{\"John\", 25}")
	fmt.Println("     m[\"john\"].Age = 26  // ❌ Error")
	fmt.Println("     ")
	fmt.Println("     // Must do:")
	fmt.Println("     p := m[\"john\"]")
	fmt.Println("     p.Age = 26")
	fmt.Println("     m[\"john\"] = p")

	fmt.Println("\n   Go (pointer value):")
	fmt.Println("     m := make(map[string]*Person)")
	fmt.Println("     m[\"john\"] = &Person{\"John\", 25}")
	fmt.Println("     m[\"john\"].Age = 26  // ✅ Works!")

	// ========================================================================
	// Example: Why This Matters
	// ========================================================================
	fmt.Println("\n6. EXAMPLE:")

	// Wrong way (doesn't work)
	fmt.Println("   Wrong way (compile error):")
	fmt.Println("     m := make(map[string]Person)")
	fmt.Println("     m[\"alice\"] = Person{\"Alice\", 20}")
	fmt.Println("     m[\"alice\"].Age++  // ❌ Error!")

	// Right way 1
	fmt.Println("\n   Right way 1 (get-modify-put):")
	m2 := make(map[string]Person)
	m2["alice"] = Person{"Alice", 20}
	alice := m2["alice"]
	alice.Age++
	m2["alice"] = alice
	fmt.Printf("     Alice's age: %d\n", m2["alice"].Age)

	// Right way 2 (pointer)
	fmt.Println("\n   Right way 2 (use pointer):")
	m3 := make(map[string]*Person)
	m3["bob"] = &Person{"Bob", 30}
	m3["bob"].Age++ // ✅ Works directly
	fmt.Printf("     Bob's age: %d\n", m3["bob"].Age)

	// ========================================================================
	// Summary
	// ========================================================================
	fmt.Println("\n7. SUMMARY:")
	fmt.Println("   ✅ In C++: Map returns reference, can modify directly")
	fmt.Println("   ❌ In Go: Map returns copy, cannot modify directly")
	fmt.Println("   ")
	fmt.Println("   Solutions in Go:")
	fmt.Println("     1. Get value, modify, put back")
	fmt.Println("     2. Use pointer as map value (m[string]*Struct)")
	fmt.Println("     3. Assign entire struct")
}

// Call MapKeysExplained() and StructFieldModificationExplained() from main.go
