package utils

import "sync"

type AtomicCounter struct {
	mu      *sync.Mutex
	counter int64
}

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{
		mu:      &sync.Mutex{},
		counter: 0,
	}
}

func (ac *AtomicCounter) NextId() int64 {
	ac.mu.Lock()
	defer ac.mu.Unlock()
	ac.counter = ac.counter + 1
	return ac.counter
}
