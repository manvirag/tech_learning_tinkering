package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// PosID is a list of integers that uniquely identify the position of a character
type PosID []int

// CharItem represents a character and its position ID + site info
type CharItem struct {
	Char   rune
	PosID  PosID
	SiteID string
	Seq    int // operation counter for tie-breaking
}

// Document holds the ordered chars and concurrency control
type Document struct {
	chars []CharItem
	mu    sync.Mutex
}

// Compare two PosIDs lexicographically
func comparePosID(a, b PosID) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}
	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}
	return 0
}

// Insert a new CharItem into the document in the right order
func (d *Document) Insert(item CharItem) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.chars = append(d.chars, item)
	// Sort by PosID, then SiteID, then Seq
	sort.Slice(d.chars, func(i, j int) bool {
		cmp := comparePosID(d.chars[i].PosID, d.chars[j].PosID)
		if cmp == 0 {
			if d.chars[i].SiteID < d.chars[j].SiteID {
				return true
			} else if d.chars[i].SiteID > d.chars[j].SiteID {
				return false
			}
			return d.chars[i].Seq < d.chars[j].Seq
		}
		return cmp < 0
	})
}

// Generate a PosID between two PosIDs (l and r) recursively
func generatePosIDBetween(l, r PosID, depth int) PosID {
	// Max integer range to choose from
	const maxDigit = 10

	// Get digit at current depth or default boundaries
	var leftDigit int = 0
	if depth < len(l) {
		leftDigit = l[depth]
	}

	var rightDigit int = maxDigit
	if depth < len(r) {
		rightDigit = r[depth]
	}

	// If space between digits, pick random digit between leftDigit+1 and rightDigit-1
	if rightDigit-leftDigit > 1 {
		randDigit := leftDigit + 1 + rand.Intn(rightDigit-leftDigit-1)
		posID := append(l[:depth], randDigit)
		return posID
	}

	// No room at this level, go deeper
	posID := append([]int{}, l[:depth]...)
	posID = append(posID, leftDigit)
	return generatePosIDBetween(l, r, depth+1)
}

// Helper to print document content as string
func (d *Document) String() string {
	d.mu.Lock()
	defer d.mu.Unlock()
	runes := []rune{}
	for _, c := range d.chars {
		runes = append(runes, c.Char)
	}
	return string(runes)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Initial document: "hello" with posIDs 1,2,3,4,5
	doc := &Document{
		chars: []CharItem{
			{'h', PosID{1}, "init", 0},
			{'e', PosID{2}, "init", 0},
			{'l', PosID{3}, "init", 0},
			{'l', PosID{4}, "init", 0},
			{'o', PosID{5}, "init", 0},
		},
	}

	// Simulate two clients concurrently inserting between 'e' and first 'l' (between posID 2 and 3)
	l := PosID{2}
	r := PosID{3}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		posID := generatePosIDBetween(l, r, 0)
		item := CharItem{'X', posID, "clientA", 1}
		doc.Insert(item)
		fmt.Println("ClientA inserted:", string(item.Char), "posID:", item.PosID)
	}()

	go func() {
		defer wg.Done()
		posID := generatePosIDBetween(l, r, 0)
		item := CharItem{'Y', posID, "clientB", 1}
		doc.Insert(item)
		fmt.Println("ClientB inserted:", string(item.Char), "posID:", item.PosID)
	}()

	wg.Wait()

	// Print final document state
	fmt.Println("Final document content:", doc.String())
	fmt.Println("Full chars with posIDs:")
	for _, c := range doc.chars {
		fmt.Printf("Char: %c, PosID: %v, SiteID: %s\n", c.Char, c.PosID, c.SiteID)
	}
}
