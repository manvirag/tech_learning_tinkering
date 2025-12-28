package stl

import "fmt"

// ============================================================================
// STACK in C++ vs SLICE in Go
// Go slices can be used as stacks (LIFO - Last In First Out)
// ============================================================================

func Stack() {
	// ========================================================================
	// C++: stack<int> st;
	// Go:   var st []int
	// ========================================================================
	var st []int

	// ========================================================================
	// C++: st.push(1);
	// Go:   st = append(st, 1)
	// ========================================================================
	st = append(st, 1)
	st = append(st, 2)
	st = append(st, 3)
	fmt.Println("After push:", st) // [1 2 3]

	// ========================================================================
	// C++: st.top()
	// Go:   st[len(st)-1]
	// ========================================================================
	fmt.Println("Top:", st[len(st)-1]) // 3

	// ========================================================================
	// C++: st.pop()
	// Go:   st = st[:len(st)-1]
	// ========================================================================
	st = st[:len(st)-1]
	fmt.Println("After pop:", st)          // [1 2]
	fmt.Println("Top now:", st[len(st)-1]) // 2

	// ========================================================================
	// C++: st.size()
	// Go:   len(st)
	// ========================================================================
	fmt.Println("Size:", len(st)) // 2

	// ========================================================================
	// C++: st.empty()
	// Go:   len(st) == 0
	// ========================================================================
	fmt.Println("Empty?", len(st) == 0) // false

	// Pop all
	st = st[:len(st)-1]
	st = st[:len(st)-1]
	fmt.Println("After popping all, Empty?", len(st) == 0) // true

	// ========================================================================
	// WITH STRUCTS
	// ========================================================================
	type Point struct {
		X, Y int
	}

	var pointStack []Point
	pointStack = append(pointStack, Point{1, 2})
	pointStack = append(pointStack, Point{3, 4})
	pointStack = append(pointStack, Point{5, 6})

	fmt.Println("\nPoint Stack:")
	for len(pointStack) > 0 {
		top := pointStack[len(pointStack)-1]
		fmt.Printf("  Top: (%d, %d)\n", top.X, top.Y)
		pointStack = pointStack[:len(pointStack)-1]
	}

	// ========================================================================
	// EXAMPLE: Valid Parentheses using Stack
	// ========================================================================
	fmt.Println("\n=== Example: Valid Parentheses ===")
	s := "()[]{}"
	fmt.Println("String:", s)

	var stack []rune
	valid := true
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			// Push opening bracket
			stack = append(stack, char)
		} else {
			// Check if stack is empty or top doesn't match
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				valid = false
				break
			}
			// Pop matching bracket
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) != 0 {
		valid = false
	}

	fmt.Println("Valid?", valid)
}
