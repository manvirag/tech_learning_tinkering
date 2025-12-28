package stl

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// ============================================================================
// RUNE in Go vs CHAR in C++
// Go runes are Unicode code points (int32), C++ char is single byte (char)
// ============================================================================

func Rune() {
	// ========================================================================
	// C++: char c = 'A';
	// Go:   var r rune = 'A'  or  r := 'A'
	// ========================================================================
	var r rune = 'A'
	fmt.Printf("Rune: %c (value: %d)\n", r, r) // A (value: 65)

	// ========================================================================
	// C++: char c = 65;
	// Go:   r := rune(65)
	// ========================================================================
	r = rune(65)
	fmt.Printf("Rune from int: %c\n", r) // A

	// ========================================================================
	// C++: 'A' + 1
	// Go:   'A' + 1
	// ========================================================================
	r = 'A' + 1
	fmt.Printf("'A' + 1 = %c\n", r) // B

	// ========================================================================
	// C++: isalpha(c)
	// Go:   unicode.IsLetter(r)
	// ========================================================================
	fmt.Println("\n=== Character Classification ===")
	fmt.Printf("IsLetter('A'): %v\n", unicode.IsLetter('A')) // true
	fmt.Printf("IsLetter('5'): %v\n", unicode.IsLetter('5')) // false
	fmt.Printf("IsDigit('5'): %v\n", unicode.IsDigit('5'))   // true
	fmt.Printf("IsDigit('A'): %v\n", unicode.IsDigit('A'))   // false
	fmt.Printf("IsSpace(' '): %v\n", unicode.IsSpace(' '))   // true
	fmt.Printf("IsUpper('A'): %v\n", unicode.IsUpper('A'))   // true
	fmt.Printf("IsLower('a'): %v\n", unicode.IsLower('a'))   // true
	fmt.Printf("IsPunct('!'): %v\n", unicode.IsPunct('!'))   // true

	// ========================================================================
	// C++: toupper(c), tolower(c)
	// Go:   unicode.ToUpper(r), unicode.ToLower(r)
	// ========================================================================
	fmt.Println("\n=== Case Conversion ===")
	fmt.Printf("ToUpper('a'): %c\n", unicode.ToUpper('a')) // A
	fmt.Printf("ToLower('A'): %c\n", unicode.ToLower('A')) // a

	// ========================================================================
	// C++: string s = "hello"; char c = s[0];
	// Go:   s := "hello"; r := rune(s[0])  (for ASCII)
	//       For Unicode: use []rune(s) or range
	// ========================================================================
	fmt.Println("\n=== String to Rune ===")
	s := "hello"
	fmt.Printf("String: %s\n", s)
	fmt.Printf("s[0] as byte: %d ('%c')\n", s[0], s[0])
	fmt.Printf("s[0] as rune: %c\n", rune(s[0]))

	// For multi-byte characters
	s = "hello世界"
	fmt.Printf("\nString with Unicode: %s\n", s)
	fmt.Printf("Length in bytes: %d\n", len(s))                    // 11 (5 + 6 for 世界)
	fmt.Printf("Length in runes: %d\n", utf8.RuneCountInString(s)) // 7

	// ========================================================================
	// C++: for (char c : s) { ... }
	// Go:   for _, r := range s { ... }  (r is rune)
	// ========================================================================
	fmt.Println("\n=== Iterate String by Rune ===")
	s = "hello"
	fmt.Println("Iterate by rune (range):")
	for i, r := range s {
		fmt.Printf("  [%d] = '%c' (rune: %d)\n", i, r, r)
	}

	// ========================================================================
	// C++: string s = "hello"; s[0] = 'H';
	// Go:   Strings are IMMUTABLE - cannot modify directly
	// ========================================================================
	fmt.Println("\n=== String Immutability ===")
	s = "hello"
	fmt.Println("Original:", s)
	// s[0] = 'H'  // ❌ COMPILE ERROR - strings are immutable

	// Convert to []rune, modify, convert back
	runes := []rune(s)
	runes[0] = 'H'
	s = string(runes)
	fmt.Println("After modification:", s) // "Hello"

	// ========================================================================
	// C++: char c = 'A'; string s = string(1, c);
	// Go:   r := 'A'; s := string(r)
	// ========================================================================
	fmt.Println("\n=== Rune to String ===")
	r = 'A'
	s = string(r)
	fmt.Printf("Rune '%c' to string: \"%s\"\n", r, s)

	// Multiple runes
	runes = []rune{'h', 'e', 'l', 'l', 'o'}
	s = string(runes)
	fmt.Printf("Runes to string: \"%s\"\n", s)

	// ========================================================================
	// C++: char c = 'A'; int i = c;
	// Go:   r := 'A'; i := int(r)
	// ========================================================================
	fmt.Println("\n=== Rune to Int ===")
	r = 'A'
	i := int(r)
	fmt.Printf("Rune '%c' to int: %d\n", r, i) // 65

	// ========================================================================
	// C++: int i = 65; char c = i;
	// Go:   i := 65; r := rune(i)
	// ========================================================================
	fmt.Println("\n=== Int to Rune ===")
	i = 65
	r = rune(i)
	fmt.Printf("Int %d to rune: '%c'\n", i, r) // A

	// ========================================================================
	// Unicode Code Points
	// ========================================================================
	fmt.Println("\n=== Unicode Code Points ===")
	runes = []rune{'A', '中', '😀', 'ñ'}
	for _, r := range runes {
		fmt.Printf("  '%c' = U+%04X (decimal: %d)\n", r, r, r)
	}

	// ========================================================================
	// Multi-byte Characters
	// ========================================================================
	fmt.Println("\n=== Multi-byte Characters ===")
	s = "世界"
	fmt.Printf("String: %s\n", s)
	fmt.Printf("Byte length: %d\n", len(s))                   // 6 (3 bytes per character)
	fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(s)) // 2

	// Iterate and show byte positions
	fmt.Println("Byte positions:")
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("  Byte [%d]: '%c' (size: %d bytes)\n", i, r, size)
		i += size
	}

	// ========================================================================
	// C++: isalnum(c), isalpha(c), isdigit(c)
	// Go:   unicode.IsLetter(r), unicode.IsDigit(r), unicode.IsLetterOrDigit(r)
	// ========================================================================
	fmt.Println("\n=== Character Checks ===")
	testRunes := []rune{'A', '5', '!', 'ñ', '中'}
	for _, r := range testRunes {
		fmt.Printf("  '%c': Letter=%v, Digit=%v, LetterOrDigit=%v\n",
			r,
			unicode.IsLetter(r),
			unicode.IsDigit(r),
			unicode.IsLetter(r) || unicode.IsDigit(r))
	}

	// ========================================================================
	// C++: string s = "Hello"; transform(s.begin(), s.end(), ::tolower);
	// Go:   Convert to []rune, modify, convert back
	// ========================================================================
	fmt.Println("\n=== Transform String ===")
	s = "Hello World"
	runes = []rune(s)
	for i, r := range runes {
		runes[i] = unicode.ToLower(r)
	}
	s = string(runes)
	fmt.Println("ToLower:", s) // "hello world"

	// ========================================================================
	// C++: find_if(s.begin(), s.end(), isdigit)
	// Go:   Loop and check unicode.IsDigit
	// ========================================================================
	fmt.Println("\n=== Find First Digit ===")
	s = "hello123world"
	for i, r := range s {
		if unicode.IsDigit(r) {
			fmt.Printf("First digit at index %d: '%c'\n", i, r)
			break
		}
	}

	// ========================================================================
	// C++: count_if(s.begin(), s.end(), isalpha)
	// Go:   Loop and count
	// ========================================================================
	fmt.Println("\n=== Count Letters ===")
	s = "Hello123World!"
	count := 0
	for _, r := range s {
		if unicode.IsLetter(r) {
			count++
		}
	}
	fmt.Printf("Letter count in '%s': %d\n", s, count) // 10

	// ========================================================================
	// C++: remove_if(s.begin(), s.end(), isdigit)
	// Go:   Filter runes
	// ========================================================================
	fmt.Println("\n=== Remove Digits ===")
	s = "Hello123World"
	runes = []rune(s)
	filtered := []rune{}
	for _, r := range runes {
		if !unicode.IsDigit(r) {
			filtered = append(filtered, r)
		}
	}
	s = string(filtered)
	fmt.Printf("After removing digits: %s\n", s) // "HelloWorld"

	// ========================================================================
	// UTF-8 Encoding/Decoding
	// ========================================================================
	fmt.Println("\n=== UTF-8 Encoding ===")
	r = '中'
	buf := make([]byte, 4)
	n := utf8.EncodeRune(buf, r)
	fmt.Printf("Rune '%c' encoded to %d bytes: %v\n", r, n, buf[:n])

	// Decode
	decodedRune, size := utf8.DecodeRune(buf[:n])
	fmt.Printf("Decoded: '%c' (size: %d bytes)\n", decodedRune, size)

	// ========================================================================
	// C++: char c = 'A'; bool isUpper = (c >= 'A' && c <= 'Z');
	// Go:   r := 'A'; isUpper := unicode.IsUpper(r)
	// ========================================================================
	fmt.Println("\n=== Range Checks ===")
	fmt.Println("C++ way (ASCII only):")
	fmt.Println("  isUpper = (c >= 'A' && c <= 'Z')")
	fmt.Println("Go way (Unicode aware):")
	fmt.Println("  isUpper := unicode.IsUpper(r)")
	fmt.Println("  Works for all Unicode characters, not just ASCII")

	// Example with non-ASCII
	fmt.Printf("  IsUpper('A'): %v\n", unicode.IsUpper('A')) // true
	fmt.Printf("  IsUpper('А'): %v\n", unicode.IsUpper('А')) // true (Cyrillic)
	fmt.Printf("  IsUpper('a'): %v\n", unicode.IsUpper('a')) // false
}
