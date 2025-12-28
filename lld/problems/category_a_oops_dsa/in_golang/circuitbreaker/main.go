package circuitbreaker

import (
	"errors"
	"sync"
	"time"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	mu            sync.Mutex
	state         State
	failureCount  int
	successCount  int
	failureThresh int
	timeout       time.Duration
	lastFailTime  time.Time
}

// NewCircuitBreaker creates a new CircuitBreaker with failure threshold and timeout duration.
func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:         Closed,
		failureThresh: threshold,
		timeout:       timeout,
	}
}

// Call executes the function and manages the state of the circuit breaker.
func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	// Handle Open state: if timeout has passed, move to Half-Open state
	if cb.state == Open {
		if time.Since(cb.lastFailTime) >= cb.timeout {
			cb.state = HalfOpen
		} else {
			return errors.New("circuit breaker is open")
		}
	}

	// Handle Half-Open: allow one request
	if cb.state == HalfOpen {
		err := fn()
		if err != nil {
			cb.state = Open
			cb.lastFailTime = time.Now()
			return err
		}
		cb.state = Closed
		cb.resetCounts()
		return nil
	}

	// Handle Closed state
	err := fn()
	if err != nil {
		cb.failureCount++
		if cb.failureCount >= cb.failureThresh {
			cb.state = Open
			cb.lastFailTime = time.Now()
		}
		return err
	}

	cb.resetCounts()
	return nil
}

// Reset sets the state back to Closed and resets the failure count.
func (cb *CircuitBreaker) Reset() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.state = Closed
	cb.resetCounts()
}

// resetCounts resets the failure and success counts.
func (cb *CircuitBreaker) resetCounts() {
	cb.failureCount = 0
	cb.successCount = 0
}