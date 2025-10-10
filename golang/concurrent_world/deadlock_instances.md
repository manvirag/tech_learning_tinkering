### Deadlock Instances:

Problem setup (in plain English)
Goroutine 1 holds an RLock (read lock) on an RWMutex.

Goroutine 2 tries to acquire a Lock (write lock) on the same RWMutex, so it waits for all readers to release.

While Goroutine 2 is waiting, Goroutine 3 tries to acquire an RLock. According to Go's docs, once a writer is waiting, new readers must wait to ensure the writer isn't starved—so Goroutine 3 now waits too.

If Goroutine 1 is waiting for Goroutine 3 to proceed, you hit a classic deadlock: everyone is waiting for someone else.

```

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.RWMutex
	value int
}

func (c *Counter) ReadAndUpdate() {
	c.mu.RLock()
	fmt.Println("Acquired RLock for reading")

	// Simulate some read logic
	if c.value < 5 {
		fmt.Println("Value is low, trying to update...")

		// ⚠️ DEADLOCK occurs here:
		// ReadAndUpdate() calls Update(), which tries to acquire a normal Lock()
		// while RLock() is still held by the same goroutine.
		c.Update()
	}

	c.mu.RUnlock()
	fmt.Println("Released RLock")
}

func (c *Counter) Update() {
	c.mu.Lock() // will block forever — cannot upgrade RLock to Lock
	defer c.mu.Unlock()

	fmt.Println("Acquired Lock for writing")
	c.value++
	fmt.Println("Value updated to:", c.value)
}

func main() {
	counter := &Counter{value: 0}
	counter.ReadAndUpdate()
}

```


**As per documentation**: If any goroutine calls Lock while the lock is already held by one or more readers, concurrent calls to RLock will block until the writer has acquired (and released) the lock, to ensure that the lock eventually becomes available to the writer. Note that this prohibits recursive read-locking.

- More detail: https://jarv.org/posts/go-deadlock/#:~:text=Connect%20your%20browser%20to%20http,RLock()%20and%20Lock()%20.

- Documentation: https://pkg.go.dev/sync#RWMutex
- Conference: https://www.youtube.com/watch?v=9j0oQkqzhAE&ab_channel=GopherConUK




