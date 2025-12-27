# C++ Threading - Machine Coding Reference

Concise reference for concurrent programming in machine coding interviews.

---

## 0. Simple Mental Model

**Think of it like this:**
- **Thread** = Worker doing a job (like Go's `go function()`)
- **Mutex** = Lock on a door (only one person can enter at a time)
- **Condition Variable** = Waiting room (wait until something happens, then proceed)
- **Semaphore** = Limited parking spots (N cars can park, others wait)

**When to use what:**
- **Need to protect shared data?** → Use `mutex` + `lock_guard`
- **Need to wait for something?** → Use `condition_variable`
- **Need to limit concurrent access?** → Use `semaphore`
- **Multiple readers, one writer?** → Use `shared_mutex`

**The Golden Rule:**
- Always use `lock_guard` or `unique_lock` (never manual `lock()/unlock()`)
- Always `join()` threads (or explicitly `detach()`)
- Always use predicate with `cv.wait()` (prevents spurious wakeups)

---

## 1. Thread Creation

### Basic Thread Creation
```cpp
#include <thread>
#include <iostream>
using namespace std;

// Method 1: Function pointer
void printMessage(string msg) {
    cout << msg << endl;
}
thread t1(printMessage, "Hello");

// Method 2: Lambda function
thread t2([](int a, int b) {
    cout << a + b << endl;
}, 10, 20);

// Method 3: Member function
class MyClass {
public:
    void memberFunc(int x) {
        cout << x << endl;
    }
};
MyClass obj;
thread t3(&MyClass::memberFunc, &obj, 100);

// Always join threads
t1.join();
t2.join();
t3.join();
```

### Thread Management
```cpp
thread t(function, args...);

t.join();      // Wait for thread to finish
t.detach();    // Let thread run independently (use carefully!)
t.joinable();  // Check if thread can be joined

// Get thread ID
thread::id id = t.get_id();
this_thread::get_id();  // Current thread ID

// Sleep
this_thread::sleep_for(chrono::seconds(1));
this_thread::sleep_until(time_point);
```

**Key Points:**
- Always `join()` threads (or `detach()` if intentional)
- Pass arguments by value or use `ref()` for references
- Use `std::ref()` when passing mutex/condition_variable to threads

---

## 2. Mutex and Locks

**Simple Analogy:** Mutex is like a bathroom lock. Only one person can enter, others wait outside.

### Mutex Types (99% of time, just use `mutex`)
```cpp
#include <mutex>

mutex mtx;                    // ✅ Use this 99% of the time
recursive_mutex rmtx;         // Same thread can lock multiple times (rare)
timed_mutex tmtx;             // Can try_lock_for/until (rare)
shared_mutex smtx;            // Multiple readers OR one writer (C++17, for reader-writer pattern)
```

### ❌ Manual Locking (Don't Do This)
```cpp
mutex mtx;
mtx.lock();
// Critical section
mtx.unlock();  // ❌ Easy to forget! Deadlock risk!
```

### ✅ lock_guard (Use This - Simple & Safe)
```cpp
mutex mtx;

void safeFunction() {
    lock_guard<mutex> lock(mtx);  // Lock automatically
    // Critical section - protected!
    // Unlocks automatically when function ends (even if exception!)
}

// That's it! No need to unlock manually.
```

**Why `lock_guard`?**
- **Automatic**: Locks when created, unlocks when destroyed
- **Exception-safe**: Unlocks even if exception is thrown
- **Simple**: Just one line, no manual management

### unique_lock (When You Need Flexibility)
```cpp
mutex mtx;

// Use unique_lock when you need condition_variable
void withConditionVariable() {
    unique_lock<mutex> lock(mtx);  // Must use unique_lock (not lock_guard)
    cv.wait(lock, [&]() { return ready; });
}

// Can also unlock early if needed
void flexibleFunction() {
    unique_lock<mutex> lock(mtx);
    // Do critical work
    lock.unlock();  // Release early
    // Do non-critical work (don't hold lock)
}
```

**When to Use `unique_lock`:**
- **Required** for `condition_variable` (won't work with `lock_guard`)
- Need to unlock before scope ends (rare)
- Otherwise, just use `lock_guard` (simpler!)

### shared_lock (Multiple Readers)
```cpp
#include <shared_mutex>

shared_mutex smtx;

// Multiple threads can read simultaneously
void reader() {
    shared_lock<shared_mutex> lock(smtx);  // Shared lock
    // Read data
}

// Only one thread can write
void writer() {
    unique_lock<shared_mutex> lock(smtx);  // Exclusive lock
    // Write data
}
```

**Key Points:**
- `lock_guard`: Simple, automatic, exception-safe
- `unique_lock`: Flexible, needed for condition variables
- `shared_lock`: For reader-writer scenarios
- Always pass mutex by reference: `ref(mtx)`

---

## 3. Condition Variables

**Simple Analogy:** Like a waiting room. Thread waits until condition is met, then proceeds.

### Basic Pattern (3 Steps)
```cpp
#include <condition_variable>
#include <mutex>

mutex mtx;
condition_variable cv;
bool ready = false;

// Step 1: Waiter thread (waits for condition)
void waiter() {
    unique_lock<mutex> lock(mtx);  // Must use unique_lock!
    
    // Wait until condition is true
    cv.wait(lock, [&]() {
        return ready;  // Keep waiting until this is true
    });
    
    // Now ready is true, proceed!
    cout << "Ready!" << endl;
}

// Step 2: Notifier thread (signals condition)
void notifier() {
    {
        lock_guard<mutex> lock(mtx);
        ready = true;  // Change condition
    }
    cv.notify_all();  // Wake all waiting threads
}
```

**The Pattern:**
1. **Wait**: `cv.wait(lock, predicate)` - waits until predicate is true
2. **Change**: Modify condition (with lock held)
3. **Notify**: `cv.notify_all()` - wake waiting threads

### Producer-Consumer Pattern
```cpp
class BoundedBlockingQueue {
private:
    queue<int> q;
    int capacity;
    mutex mtx;
    condition_variable cv;
    
public:
    BoundedBlockingQueue(int cap) : capacity(cap) {}
    
    void enqueue(int item) {
        unique_lock<mutex> lock(mtx);
        
        // Wait until queue has space
        cv.wait(lock, [&]() {
            return q.size() < capacity;
        });
        
        q.push(item);
        cv.notify_all();  // Notify consumers
    }
    
    int dequeue() {
        unique_lock<mutex> lock(mtx);
        
        // Wait until queue has items
        cv.wait(lock, [&]() {
            return !q.empty();
        });
        
        int item = q.front();
        q.pop();
        cv.notify_all();  // Notify producers
        return item;
    }
};
```

### Condition Variable Methods
```cpp
cv.wait(lock, predicate);           // Wait until predicate is true
cv.wait_for(lock, duration, pred);   // Wait with timeout
cv.wait_until(lock, time, pred);    // Wait until time point

cv.notify_one();   // Wake one waiting thread
cv.notify_all();   // Wake all waiting threads
```

**Key Points:**
- Must use `unique_lock` (not `lock_guard`) with condition variables
- Always use predicate in `wait()` to avoid spurious wakeups
- `notify_one()` vs `notify_all()`: Use `notify_all()` when multiple threads might be waiting

---

## 4. Semaphores (C++20)

**Simple Analogy:** Like parking spots. N spots available, others wait.

### Basic Usage
```cpp
#include <semaphore>

// Binary semaphore (like mutex, but can be released by different thread)
binary_semaphore sem(1);  // 1 spot available

sem.acquire();  // Take spot (blocks if 0)
// Do work
sem.release();  // Free spot (wakes waiting threads)
```

### Counting Semaphore (Most Useful)
```cpp
// Allow N threads to access resource simultaneously
counting_semaphore<10> sem(3);  // Max 10, currently 3 available

sem.acquire();  // 3 → 2 (one thread enters)
sem.acquire();  // 2 → 1 (another thread enters)
sem.acquire();  // 1 → 0 (third thread enters)
sem.acquire();  // Blocks! (no spots left)

sem.release();  // 0 → 1 (one thread leaves, wakes waiting thread)
```

### Semaphore Methods
```cpp
sem.acquire();                    // Blocking acquire
sem.release();                    // Release (increment)
sem.try_acquire();                // Non-blocking
sem.try_acquire_for(duration);    // With timeout
sem.try_acquire_until(time_point); // Until time point
```

### Use Cases
```cpp
// 1. Resource Pool (limit concurrent connections)
counting_semaphore<10> dbPool(5);  // Max 5 connections

void useDatabase() {
    dbPool.acquire();
    // Use database connection
    dbPool.release();
}

// 2. Producer-Consumer
counting_semaphore<10> emptySlots(5);  // 5 empty slots
counting_semaphore<10> fullSlots(0);  // 0 items initially

void producer() {
    emptySlots.acquire();  // Wait for empty slot
    // Add item to buffer
    fullSlots.release();   // Signal item available
}

void consumer() {
    fullSlots.acquire();   // Wait for item
    // Remove item from buffer
    emptySlots.release();  // Signal slot empty
}
```

**Key Points:**
- C++20 feature (may not be available in older compilers)
- Semaphore vs Mutex: Semaphore can be released by different thread
- Use for resource pools, rate limiting, producer-consumer coordination

---

## 5. Quick Decision Guide

**"What should I use?"**

| Scenario | Solution | Code |
|----------|----------|------|
| Protect shared variable | `mutex` + `lock_guard` | `lock_guard<mutex> lock(mtx);` |
| Wait for condition | `condition_variable` | `cv.wait(lock, [&]() { return condition; });` |
| Limit concurrent access (N threads) | `semaphore` | `sem.acquire(); ... sem.release();` |
| Multiple readers, one writer | `shared_mutex` | `shared_lock` for read, `unique_lock` for write |
| Thread-safe counter | `mutex` + `lock_guard` | Simple mutex protection |
| Producer-Consumer queue | `mutex` + `condition_variable` | Wait for empty/full |
| Wait for all threads | `barrier` or `condition_variable` | Count threads, notify when all ready |

**Simple Rules:**
- **Just protecting data?** → `lock_guard<mutex>`
- **Need to wait?** → `condition_variable` with `unique_lock`
- **Need to limit access?** → `semaphore`
- **Multiple readers?** → `shared_mutex`

---

## 6. Common Patterns

### Pattern 1: Thread-Safe Counter (Simplest)
**Use Case:** Multiple threads incrementing a counter
```cpp
class ThreadSafeCounter {
private:
    int count = 0;
    mutex mtx;
    
public:
    void increment() {
        lock_guard<mutex> lock(mtx);
        count++;
    }
    
    int get() {
        lock_guard<mutex> lock(mtx);
        return count;
    }
};
```

### Pattern 2: Reader-Writer Lock
**Use Case:** Many readers, few writers (e.g., cache, configuration)
```cpp
class ThreadSafeData {
private:
    int data = 0;
    shared_mutex smtx;
    
public:
    int read() {
        shared_lock<shared_mutex> lock(smtx);  // Multiple readers can read together
        return data;
    }
    
    void write(int value) {
        unique_lock<shared_mutex> lock(smtx);  // Only one writer at a time
        data = value;
    }
};
```

### Pattern 3: Producer-Consumer Queue (Most Common in Interviews)
**Use Case:** One thread produces, another consumes (e.g., task queue, message queue)
```cpp
class BlockingQueue {
private:
    queue<int> q;
    mutex mtx;
    condition_variable cv;
    
public:
    void push(int item) {
        lock_guard<mutex> lock(mtx);
        q.push(item);
        cv.notify_one();
    }
    
    int pop() {
        unique_lock<mutex> lock(mtx);
        cv.wait(lock, [&]() { return !q.empty(); });
        int item = q.front();
        q.pop();
        return item;
    }
};
```

### Pattern 4: Barrier (Wait for All Threads)
```cpp
class Barrier {
private:
    int count;
    int waiting = 0;
    mutex mtx;
    condition_variable cv;
    
public:
    Barrier(int n) : count(n) {}
    
    void wait() {
        unique_lock<mutex> lock(mtx);
        waiting++;
        if (waiting == count) {
            waiting = 0;
            cv.notify_all();
        } else {
            cv.wait(lock, [&]() { return waiting == 0; });
        }
    }
};
```

---

## 7. Common Pitfalls

### ❌ Deadlock
```cpp
// Problem: Locking multiple mutexes in different order
mutex mtx1, mtx2;

void thread1() {
    lock_guard<mutex> lock1(mtx1);
    lock_guard<mutex> lock2(mtx2);  // Deadlock if thread2 has reverse order
}

void thread2() {
    lock_guard<mutex> lock2(mtx2);
    lock_guard<mutex> lock1(mtx1);  // Different order = deadlock!
}
```

**Solution:**
```cpp
// Always lock in same order
void thread1() {
    lock_guard<mutex> lock1(mtx1);
    lock_guard<mutex> lock2(mtx2);
}

void thread2() {
    lock_guard<mutex> lock1(mtx1);  // Same order
    lock_guard<mutex> lock2(mtx2);
}

// OR use std::lock (C++11)
void safeLock() {
    lock(mtx1, mtx2);  // Locks both atomically
    lock_guard<mutex> lock1(mtx1, adopt_lock);
    lock_guard<mutex> lock2(mtx2, adopt_lock);
}
```

### ❌ Forgetting to Join
```cpp
void badExample() {
    thread t([]() { /* work */ });
    // Forgot to join! Program may terminate before thread finishes
}

// ✅ Always join
void goodExample() {
    thread t([]() { /* work */ });
    t.join();  // Wait for completion
}
```

### ❌ Race Condition
```cpp
int counter = 0;  // Shared variable

void increment() {
    counter++;  // ❌ Not thread-safe! Race condition!
}

// ✅ Use mutex
mutex mtx;
void safeIncrement() {
    lock_guard<mutex> lock(mtx);
    counter++;
}
```

### ❌ Spurious Wakeups (Condition Variable)
```cpp
// ❌ Wrong: No predicate
cv.wait(lock);  // May wake up even when condition is false!

// ✅ Correct: Use predicate
cv.wait(lock, [&]() { return ready; });
```

### ❌ Using lock_guard with condition_variable
```cpp
// ❌ Wrong: condition_variable needs unique_lock
lock_guard<mutex> lock(mtx);
cv.wait(lock);  // Compile error!

// ✅ Correct: Use unique_lock
unique_lock<mutex> lock(mtx);
cv.wait(lock, [&]() { return ready; });
```

### ❌ Passing Mutex by Value
```cpp
// ❌ Wrong: Mutex is not copyable
void badFunc(mutex m) {  // Compile error!
    lock_guard<mutex> lock(m);
}

// ✅ Correct: Pass by reference
void goodFunc(mutex& m) {
    lock_guard<mutex> lock(m);
}

// When creating thread:
thread t(goodFunc, ref(mtx));  // Use ref() wrapper
```

---

## 8. Quick Reference

### Thread Creation
```cpp
thread t(function, args...);
t.join();
t.detach();
```

### Locks
```cpp
lock_guard<mutex> lock(mtx);           // Simple, auto-unlock
unique_lock<mutex> lock(mtx);          // Flexible, for condition_variable
shared_lock<shared_mutex> lock(smtx);  // Multiple readers
```

### Condition Variable
```cpp
unique_lock<mutex> lock(mtx);
cv.wait(lock, predicate);
cv.notify_one();
cv.notify_all();
```

### Semaphore (C++20)
```cpp
counting_semaphore<Max> sem(initial);
sem.acquire();
sem.release();
```

### Common Mutex Types
```cpp
mutex              // Basic
recursive_mutex    // Same thread can lock multiple times
shared_mutex       // Multiple readers OR one writer
timed_mutex        // Can try_lock_for/until
```

---

## 9. Machine Coding Tips

1. **Always use RAII locks** (`lock_guard`, `unique_lock`) - never manual `lock()/unlock()`
2. **Use `unique_lock` with condition variables** - `lock_guard` won't work
3. **Always use predicate in `cv.wait()`** - prevents spurious wakeups
4. **Lock mutexes in same order** - prevents deadlocks
5. **Always `join()` threads** - or explicitly `detach()` if intentional
6. **Pass mutex by reference** - use `ref(mtx)` when passing to threads
7. **Use `notify_all()` when multiple threads might wait** - safer than `notify_one()`
8. **Keep critical sections small** - minimize time holding locks

---

**Remember**: For machine coding, focus on:
- Thread creation and joining
- Mutex with `lock_guard` or `unique_lock`
- Condition variables for coordination
- Common patterns (producer-consumer, reader-writer)

---

## TL;DR - The Essentials

**3 Things You Need:**
1. **Thread**: `thread t(func, args); t.join();`
2. **Mutex**: `lock_guard<mutex> lock(mtx);` (protect shared data)
3. **Condition Variable**: `cv.wait(lock, predicate); cv.notify_all();` (wait for condition)

**The 3 Rules:**
1. Always use `lock_guard` or `unique_lock` (never manual lock/unlock)
2. Always `join()` threads
3. Always use predicate with `cv.wait()` (prevents spurious wakeups)

**Most Common Pattern (Producer-Consumer):**
```cpp
// Producer
void producer() {
    lock_guard<mutex> lock(mtx);
    queue.push(item);
    cv.notify_all();
}

// Consumer
void consumer() {
    unique_lock<mutex> lock(mtx);
    cv.wait(lock, [&]() { return !queue.empty(); });
    int item = queue.front();
    queue.pop();
}
```

That's it! You're ready for 90% of machine coding concurrency problems.

