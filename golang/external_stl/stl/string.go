package stl

import (
	"fmt"
	"strings"
)

// ============================================================================
// STRING in C++ vs STRING in Go
// Go strings are immutable byte slices
// ============================================================================

func String() {
	// ========================================================================
	// C++: string s = "hello";
	// Go:   s := "hello"
	// ========================================================================
	s := "hello"
	fmt.Println("String:", s)

	// ========================================================================
	// C++: s.length()  or  s.size()
	// Go:   len(s)
	// ========================================================================
	fmt.Println("Length:", len(s)) // 5

	// ========================================================================
	// C++: s[i]
	// Go:   s[i]  (returns byte, not rune for multi-byte)
	// ========================================================================
	fmt.Printf("s[0]: %c (byte: %d)\n", s[0], s[0]) // 'h'

	// ========================================================================
	// C++: s += " world";
	// Go:   s += " world"  or  s = s + " world"
	// ========================================================================
	s = s + " world"
	fmt.Println("After concatenation:", s) // "hello world"

	// ========================================================================
	// C++: s1 == s2
	// Go:   s1 == s2
	// ========================================================================
	s1 := "hello"
	s2 := "hello"
	s3 := "world"
	fmt.Println("s1 == s2:", s1 == s2) // true
	fmt.Println("s1 == s3:", s1 == s3) // false

	// ========================================================================
	// C++: s.substr(start, length)
	// Go:   s[start:end]  (end is exclusive)
	// ========================================================================
	s = "hello world"
	sub := s[0:5] // "hello" (indices 0 to 4)
	fmt.Println("Substring [0:5]:", sub)

	sub = s[6:] // "world" (from index 6 to end)
	fmt.Println("Substring [6:]:", sub)

	sub = s[:5] // "hello" (from start to index 4)
	fmt.Println("Substring [:5]:", sub)

	// ========================================================================
	// C++: s.find("world")
	// Go:   strings.Index(s, "world")
	// ========================================================================
	idx := strings.Index(s, "world")
	fmt.Printf("Index of 'world': %d\n", idx) // 6

	idx = strings.Index(s, "xyz")
	fmt.Printf("Index of 'xyz': %d (not found = -1)\n", idx) // -1

	// ========================================================================
	// C++: s.find_last_of("o")
	// Go:   strings.LastIndex(s, "o")
	// ========================================================================
	lastIdx := strings.LastIndex(s, "o")
	fmt.Printf("Last index of 'o': %d\n", lastIdx) // 7

	// ========================================================================
	// C++: s.empty()
	// Go:   len(s) == 0  or  s == ""
	// ========================================================================
	empty := ""
	fmt.Println("Is empty?", len(empty) == 0) // true
	fmt.Println("Is empty?", empty == "")     // true

	// ========================================================================
	// C++: s.clear()
	// Go:   s = ""
	// ========================================================================
	s = "hello"
	s = ""
	fmt.Println("After clear:", s, "Empty?", s == "")

	// ========================================================================
	// C++: s.replace(pos, len, "new")
	// Go:   strings.Replace(s, old, new, n)  or manual slice
	// ========================================================================
	s = "hello world"
	replaced := strings.Replace(s, "world", "golang", 1)
	fmt.Println("After replace:", replaced) // "hello golang"

	// Replace all occurrences
	s = "hello hello hello"
	replaced = strings.ReplaceAll(s, "hello", "hi")
	fmt.Println("Replace all:", replaced) // "hi hi hi"

	// ========================================================================
	// C++: std::transform(s.begin(), s.end(), ::toupper)
	// Go:   strings.ToUpper(s)
	// ========================================================================
	s = "hello"
	upper := strings.ToUpper(s)
	lower := strings.ToLower("HELLO")
	fmt.Println("ToUpper:", upper) // "HELLO"
	fmt.Println("ToLower:", lower) // "hello"

	// ========================================================================
	// C++: s.starts_with("he")
	// Go:   strings.HasPrefix(s, "he")
	// ========================================================================
	s = "hello"
	fmt.Println("Starts with 'he':", strings.HasPrefix(s, "he")) // true
	fmt.Println("Starts with 'wo':", strings.HasPrefix(s, "wo")) // false

	// ========================================================================
	// C++: s.ends_with("lo")
	// Go:   strings.HasSuffix(s, "lo")
	// ========================================================================
	fmt.Println("Ends with 'lo':", strings.HasSuffix(s, "lo")) // true
	fmt.Println("Ends with 'he':", strings.HasSuffix(s, "he")) // false

	// ========================================================================
	// C++: s.find_first_of("aeiou")
	// Go:   strings.IndexAny(s, "aeiou")
	// ========================================================================
	s = "hello"
	idx = strings.IndexAny(s, "aeiou")
	fmt.Printf("First vowel at index: %d ('%c')\n", idx, s[idx]) // 1 ('e')

	// ========================================================================
	// C++: std::stoi(s)  or  atoi(s.c_str())
	// Go:   strconv.Atoi(s)
	// ========================================================================
	// Note: Need to import "strconv" for this
	// For now, just show the concept
	fmt.Println("\nString to int:")
	fmt.Println("  C++: std::stoi(\"123\")")
	fmt.Println("  Go:   strconv.Atoi(\"123\")")

	// ========================================================================
	// C++: std::to_string(123)
	// Go:   strconv.Itoa(123)  or  fmt.Sprintf("%d", 123)
	// ========================================================================
	fmt.Println("\nInt to string:")
	fmt.Println("  C++: std::to_string(123)")
	fmt.Println("  Go:   strconv.Itoa(123)  or  fmt.Sprintf(\"%d\", 123)")

	// ========================================================================
	// C++: std::getline(cin, s)
	// Go:   reader.ReadLine()  or  scanner.Scan()
	// ========================================================================
	fmt.Println("\nReading line:")
	fmt.Println("  C++: std::getline(cin, s)")
	fmt.Println("  Go:   scanner.Scan()  or  reader.ReadLine()")

	// ========================================================================
	// C++: std::reverse(s.begin(), s.end())
	// Go:   convert to []rune, reverse, convert back
	// ========================================================================
	s = "hello"
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reversed := string(runes)
	fmt.Println("\nReverse 'hello':", reversed) // "olleh"

	// ========================================================================
	// C++: std::sort(s.begin(), s.end())
	// Go:   convert to []rune, sort, convert back
	// ========================================================================
	// Note: Need to import "sort" for this
	fmt.Println("\nSort string:")
	fmt.Println("  C++: std::sort(s.begin(), s.end())")
	fmt.Println("  Go:   Convert to []rune, use sort.Slice(), convert back")

	// ========================================================================
	// C++: s.erase(pos, len)
	// Go:   s[:pos] + s[pos+len:]
	// ========================================================================
	s = "hello world"
	// Remove "world" (from index 6)
	erased := s[:6] + s[11:]
	fmt.Println("\nAfter erase 'world':", erased) // "hello "

	// ========================================================================
	// C++: s.insert(pos, "new")
	// Go:   s[:pos] + "new" + s[pos:]
	// ========================================================================
	s = "hello world"
	// Insert "beautiful " at position 6
	inserted := s[:6] + "beautiful " + s[6:]
	fmt.Println("After insert 'beautiful ':", inserted) // "hello beautiful world"

	// ========================================================================
	// C++: std::count(s.begin(), s.end(), 'l')
	// Go:   strings.Count(s, "l")
	// ========================================================================
	s = "hello"
	count := strings.Count(s, "l")
	fmt.Printf("\nCount of 'l' in '%s': %d\n", s, count) // 2

	// ========================================================================
	// C++: std::trim(s)
	// Go:   strings.TrimSpace(s)  or  strings.Trim(s, " ")
	// ========================================================================
	s = "  hello world  "
	trimmed := strings.TrimSpace(s)
	fmt.Printf("Trimmed: '%s'\n", trimmed) // "hello world"

	// ========================================================================
	// C++: std::split(s, delimiter)
	// Go:   strings.Split(s, delimiter)
	// ========================================================================
	s = "apple,banana,cherry"
	parts := strings.Split(s, ",")
	fmt.Println("\nSplit by ',':", parts) // ["apple" "banana" "cherry"]

	// ========================================================================
	// C++: std::join(vector, delimiter)
	// Go:   strings.Join(slice, delimiter)
	// ========================================================================
	joined := strings.Join(parts, " | ")
	fmt.Println("Join with ' | ':", joined) // "apple | banana | cherry"

	// ========================================================================
	// ITERATION
	// ========================================================================
	fmt.Println("\n=== Iteration ===")
	s = "hello"

	// C++: for (char c : s)
	// Go:   for _, c := range s  (c is rune, not byte)
	fmt.Println("Iterate by rune:")
	for i, c := range s {
		fmt.Printf("  [%d] = '%c'\n", i, c)
	}

	// Iterate by byte
	fmt.Println("Iterate by byte:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("  [%d] = '%c' (byte: %d)\n", i, s[i], s[i])
	}
}
