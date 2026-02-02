// Package golang_library provides time and random utilities for LLD, matching, competitive, and general use.
package golang_library

import (
	"crypto/rand"
	"fmt"
	"math/big"
	mathrand "math/rand"
)

// ============================================================================
// RANDOM - Quick reference for LLD, matching, competitive coding
// ============================================================================

// RandomDemo runs random-related examples.
func RandomDemo() {
	// ------------------------------------------------------------------------
	// 1. math/rand - fast, not cryptographically secure (use for games, shuffles)
	// ------------------------------------------------------------------------
	// Seed (Go 1.20+ auto-seeds; for reproducibility use Seed)
	// mathrand.Seed(42)

	// Intn(n) - [0, n) i.e. 0 to n-1
	n := mathrand.Intn(10) // 0..9
	fmt.Println("Intn(10):", n)

	// Int() - [0, MaxInt)
	_ = mathrand.Int()

	// Float64() - [0.0, 1.0)
	f := mathrand.Float64()
	fmt.Println("Float64:", f)

	// ------------------------------------------------------------------------
	// 2. Random in range [a, b] inclusive (integers)
	// ------------------------------------------------------------------------
	// [0, n) -> Intn(n)
	// [min, max] inclusive: min + Intn(max-min+1)
	min, max := 1, 6
	dice := min + mathrand.Intn(max-min+1)
	fmt.Println("Dice [1,6]:", dice)

	// [min, max) exclusive of max: min + Intn(max-min)
	// [0, max): Intn(max)

	// ------------------------------------------------------------------------
	// 3. Random float in range [min, max]
	// ------------------------------------------------------------------------
	minF, maxF := 1.0, 10.0
	r := minF + mathrand.Float64()*(maxF-minF)
	fmt.Println("Float [1,10]:", r)

	// ------------------------------------------------------------------------
	// 4. Shuffle slice (Fisher-Yates)
	// ------------------------------------------------------------------------
	slice := []int{1, 2, 3, 4, 5}
	mathrand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	fmt.Println("Shuffled:", slice)

	// ------------------------------------------------------------------------
	// 5. Perm(n) - random permutation of 0..n-1
	// ------------------------------------------------------------------------
	perm := mathrand.Perm(5)
	fmt.Println("Perm(5):", perm)

	// ------------------------------------------------------------------------
	// 6. Random element from slice
	// ------------------------------------------------------------------------
	items := []string{"a", "b", "c"}
	idx := mathrand.Intn(len(items))
	pick := items[idx]
	fmt.Println("Random element:", pick)

	// ------------------------------------------------------------------------
	// 7. crypto/rand - cryptographically secure (tokens, keys, security)
	// ------------------------------------------------------------------------
	// Random bytes
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	fmt.Println("Crypto random bytes (first 4):", b[:4])

	// Random int in [0, n) using crypto/rand
	maxBig := big.NewInt(100)
	nBig, _ := rand.Int(rand.Reader, maxBig)
	fmt.Println("Crypto Int [0,100):", nBig)

	// Random in [min, max] inclusive with crypto/rand
	// min + Int(rand.Reader, big.NewInt(int64(max-min+1)))
	minC, maxC := 1, 6
	maxRange := big.NewInt(int64(maxC - minC + 1))
	x, _ := rand.Int(rand.Reader, maxRange)
	diceSecure := minC + int(x.Int64())
	fmt.Println("Crypto dice [1,6]:", diceSecure)

	// ------------------------------------------------------------------------
	// 8. Reproducible random (same seed = same sequence)
	// ------------------------------------------------------------------------
	// mathrand.Seed(42)
	// a := mathrand.Intn(10)
	// b := mathrand.Intn(10)
	// Same seed -> same a, b every run (useful for testing)
	fmt.Println("Reproducible: mathrand.Seed(42) then use mathrand.*")
}
