package util

import "sync"

var (
	counter int
	mu      sync.Mutex
)

func GetCount() int {
	mu.Lock()
	defer mu.Unlock()
	counter++
	return counter
}

