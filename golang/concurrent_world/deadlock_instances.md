### Deadlock Instances:

1. Suppose we have a function, in which we have acquired RLock and goroutine 2 trying to get Lock ( i.e. write lock ). And gouroutine 3 trying to get RLock. Then goroutine 3 will be deadlock. Because first is waiting for gorouting 3 to get execute.

**As per documentation**: If any goroutine calls Lock while the lock is already held by one or more readers, concurrent calls to RLock will block until the writer has acquired (and released) the lock, to ensure that the lock eventually becomes available to the writer. Note that this prohibits recursive read-locking.

More detail: https://jarv.org/posts/go-deadlock/#:~:text=Connect%20your%20browser%20to%20http,RLock()%20and%20Lock()%20.
Documentation: https://pkg.go.dev/sync#RWMutex
