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
	// STRUCTS IN VECTOR - COMPREHENSIVE GUIDE
	// ========================================================================
	fmt.Println("\n=== STRUCTS IN VECTOR ===")

	// Define a more complex struct
	type Person struct {
		ID    int
		Name  string
		Age   int
		Score float64
	}

	// Create slice of structs
	var people []Person
	people = append(people, Person{1, "Alice", 25, 85.5})
	people = append(people, Person{2, "Bob", 30, 92.0})
	people = append(people, Person{3, "Charlie", 22, 78.5})
	people = append(people, Person{4, "David", 28, 90.0})

	// Access struct fields
	fmt.Println("First person:", people[0].Name, people[0].Age)

	// Modify struct in slice
	people[0].Score = 88.0
	fmt.Println("Updated score:", people[0].Score)

	// Iterate over structs
	for i, p := range people {
		fmt.Printf("  [%d] %s (Age: %d, Score: %.1f)\n", i, p.Name, p.Age, p.Score)
	}

	// ========================================================================
	// SORTING STRUCTS - Multiple Criteria
	// ========================================================================
	fmt.Println("\n=== SORTING STRUCTS ===")

	// Sort by Age (ascending)
	people = []Person{
		{1, "Alice", 25, 85.5},
		{2, "Bob", 30, 92.0},
		{3, "Charlie", 22, 78.5},
		{4, "David", 28, 90.0},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sorted by Age:")
	for _, p := range people {
		fmt.Printf("  %s: %d\n", p.Name, p.Age)
	}

	// Sort by Score (descending)
	people = []Person{
		{1, "Alice", 25, 85.5},
		{2, "Bob", 30, 92.0},
		{3, "Charlie", 22, 78.5},
		{4, "David", 28, 90.0},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Score > people[j].Score
	})
	fmt.Println("Sorted by Score (desc):")
	for _, p := range people {
		fmt.Printf("  %s: %.1f\n", p.Name, p.Score)
	}

	// Sort by multiple fields: Age first, then Score
	people = []Person{
		{1, "Alice", 25, 85.5},
		{2, "Bob", 30, 92.0},
		{3, "Charlie", 25, 78.5}, // Same age as Alice
		{4, "David", 28, 90.0},
	}
	sort.Slice(people, func(i, j int) bool {
		if people[i].Age != people[j].Age {
			return people[i].Age < people[j].Age
		}
		return people[i].Score > people[j].Score // If same age, higher score first
	})
	fmt.Println("Sorted by Age, then Score:")
	for _, p := range people {
		fmt.Printf("  %s: Age=%d, Score=%.1f\n", p.Name, p.Age, p.Score)
	}

	// ========================================================================
	// SEARCHING STRUCTS
	// ========================================================================
	fmt.Println("\n=== SEARCHING STRUCTS ===")

	// Linear search by field
	targetName := "Bob"
	foundPersonBool := false
	var foundPersonData Person
	for _, p := range people {
		if p.Name == targetName {
			foundPersonBool = true
			foundPersonData = p
			break
		}
	}
	if foundPersonBool {
		fmt.Printf("Found: %s (Age: %d)\n", foundPersonData.Name, foundPersonData.Age)
	}

	// Search by ID
	targetID := 3
	personIndex := -1
	for i, p := range people {
		if p.ID == targetID {
			personIndex = i
			break
		}
	}
	if personIndex != -1 {
		fmt.Printf("Person with ID %d found at index %d: %s\n", targetID, personIndex, people[personIndex].Name)
	}

	// Binary search on sorted structs (by Age)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	targetAge := 25
	personIdx := sort.Search(len(people), func(i int) bool {
		return people[i].Age >= targetAge
	})
	if personIdx < len(people) && people[personIdx].Age == targetAge {
		fmt.Printf("Binary search: Found age %d at index %d: %s\n", targetAge, personIdx, people[personIdx].Name)
	}

	// ========================================================================
	// MIN/MAX WITH STRUCTS
	// ========================================================================
	fmt.Println("\n=== MIN/MAX WITH STRUCTS ===")

	// Find person with minimum age
	if len(people) > 0 {
		minAgePerson := people[0]
		for _, p := range people {
			if p.Age < minAgePerson.Age {
				minAgePerson = p
			}
		}
		fmt.Printf("Youngest: %s (Age: %d)\n", minAgePerson.Name, minAgePerson.Age)
	}

	// Find person with maximum score
	if len(people) > 0 {
		maxScorePerson := people[0]
		for _, p := range people {
			if p.Score > maxScorePerson.Score {
				maxScorePerson = p
			}
		}
		fmt.Printf("Highest score: %s (Score: %.1f)\n", maxScorePerson.Name, maxScorePerson.Score)
	}

	// ========================================================================
	// COUNT/FILTER WITH STRUCTS
	// ========================================================================
	fmt.Println("\n=== COUNT/FILTER WITH STRUCTS ===")

	// Count people above certain age
	thresholdAge := 25
	personCount := 0
	for _, p := range people {
		if p.Age > thresholdAge {
			personCount++
		}
	}
	fmt.Printf("People above age %d: %d\n", thresholdAge, personCount)

	// Filter structs (create new slice with condition)
	highScorers := []Person{}
	for _, p := range people {
		if p.Score >= 85.0 {
			highScorers = append(highScorers, p)
		}
	}
	fmt.Println("High scorers (>=85):")
	for _, p := range highScorers {
		fmt.Printf("  %s: %.1f\n", p.Name, p.Score)
	}

	// ========================================================================
	// USING sort.Sort WITH STRUCTS (Alternative to sort.Slice)
	// ========================================================================
	fmt.Println("\n=== sort.Sort INTERFACE ===")

	// Define custom type and implement sort.Interface
	type ByAge []Person

	// Note: These methods need to be defined outside the function or use a helper
	// For demonstration, we'll use sort.Slice which is more common
	peopleByAge := []Person{
		{1, "Alice", 25, 85.5},
		{2, "Bob", 30, 92.0},
		{3, "Charlie", 22, 78.5},
	}
	sort.Slice(peopleByAge, func(i, j int) bool {
		return peopleByAge[i].Age < peopleByAge[j].Age
	})
	fmt.Println("Sorted using sort.Slice (by Age):")
	for _, p := range peopleByAge {
		fmt.Printf("  %s: %d\n", p.Name, p.Age)
	}

	// ========================================================================
	// COMPARING STRUCTS
	// ========================================================================
	fmt.Println("\n=== COMPARING STRUCTS ===")

	p1 := Person{1, "Alice", 25, 85.5}
	p2 := Person{1, "Alice", 25, 85.5}
	p3 := Person{2, "Bob", 30, 92.0}

	// Compare by value (all fields)
	equal := p1.ID == p2.ID && p1.Name == p2.Name && p1.Age == p2.Age && p1.Score == p2.Score
	fmt.Printf("p1 == p2 (by value): %v\n", equal)

	// Compare specific fields
	sameAge := p1.Age == p3.Age
	fmt.Printf("p1 and p3 same age: %v\n", sameAge)

	// ========================================================================
	// REVERSE STRUCT SLICE
	// ========================================================================
	fmt.Println("\n=== REVERSE STRUCT SLICE ===")

	people = []Person{
		{1, "Alice", 25, 85.5},
		{2, "Bob", 30, 92.0},
		{3, "Charlie", 22, 78.5},
	}
	fmt.Println("Before reverse:")
	for _, p := range people {
		fmt.Printf("  %s\n", p.Name)
	}
	for i, j := 0, len(people)-1; i < j; i, j = i+1, j-1 {
		people[i], people[j] = people[j], people[i]
	}
	fmt.Println("After reverse:")
	for _, p := range people {
		fmt.Printf("  %s\n", p.Name)
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
	foundInt := false
	indexInt := -1
	for i, v := range vec {
		if v == target {
			foundInt = true
			indexInt = i
			break
		}
	}
	fmt.Printf("\nFind %d: found=%v, index=%d\n", target, foundInt, indexInt)

	// Binary search (for sorted array)
	vec = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target = 5
	idxInt := sort.SearchInts(vec, target)
	if idxInt < len(vec) && vec[idxInt] == target {
		fmt.Printf("Binary search %d: found at index %d\n", target, idxInt)
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
	countInt := 0
	for _, v := range vec {
		if v == 2 {
			countInt++
		}
	}
	fmt.Printf("Count of 2: %d\n", countInt)

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
