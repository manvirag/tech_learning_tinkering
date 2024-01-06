package main

/*
it's important to note that this implementation is not thread-safe.
If your program has multiple goroutines concurrently trying to access GetInstance(),
there is a possibility of race conditions leading to unexpected behavior

To make the singleton implementation thread-safe, you should use synchronization mechanisms such as mutexes.
Here's an example of how you can modify your implementation to make it thread-safe:

package main

import "sync"

	type Singleton interface {
		AddOne() int
	}

	type singleton struct {
		count int
	}

var (

	instance     *singleton
	instanceOnce sync.Once

)

	func GetInstance() Singleton {
		instanceOnce.Do(func() {
			instance = new(singleton)
		})

		return instance
	}

	func (s *singleton) AddOne() int {
		s.count++
		return s.count
	}
*/
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}

	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
