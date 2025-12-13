#include <iostream>
#include <thread>
#include <mutex>  // Required for lock_guard and mutex
using namespace std;

/*

It's actually better to just use a lock guard, which manages the lifecycle of a mutex for you.

It's kind of like the with: operator in Python.

Notably, a lock guard releases the lock automatically once the function that it is called in goes out of scope!



Syntax:

lock_guard<mutex> lock(mtx);
│         │       │    │
│         │       │    └─> The mutex object to lock (MUST be passed by reference)
│         │       └──────> Variable name for the lock_guard instance
│         └──────────────> Template parameter: specifies the type of mutex (could be mutex, recursive_mutex, etc.)
└────────────────────────> Template class that implements RAII (Resource Acquisition Is Initialization)

IMPORTANT: lock_guard constructor takes a REFERENCE to the mutex, not a copy!
- Mutexes are NOT copyable (copy constructor is deleted)
- You must lock the SAME mutex object, not a copy
- The constructor signature is: lock_guard(mutex_type& m)

How it works:
1. When lock_guard is created, it AUTOMATICALLY calls mtx.lock() in its constructor
2. The mutex remains locked for the entire scope where 'lock' exists
3. When 'lock' goes out of scope (function ends, block ends, exception thrown), 
   the destructor AUTOMATICALLY calls mtx.unlock()
4. This ensures the mutex is ALWAYS unlocked, even if an exception occurs!

Example comparison:

BAD (manual locking - error-prone):
    mtx.lock();
    // ... code ...
    if (error) return;  // OOPS! Forgot to unlock! Deadlock risk!
    mtx.unlock();

GOOD (lock_guard - exception-safe):
    lock_guard<mutex> lock(mtx);
    // ... code ...
    if (error) return;  // Safe! Automatically unlocks when function exits
    // No need to manually unlock - happens automatically!



 Sample Code: 
template<typename Mutex>
class lock_guard {
private:
    Mutex& mtx;  // Reference to the mutex (NOT a copy!)
    
public:
    // Constructor: Acquires the lock immediately
    explicit lock_guard(Mutex& m) : mtx(m) {
        mtx.lock();  // Lock the mutex when created
    }
    
    // Destructor: Releases the lock automatically
    ~lock_guard() {
        mtx.unlock();  // Unlock when destroyed
    }
    
    // Delete copy constructor - can't copy lock_guard
    lock_guard(const lock_guard&) = delete;
    
    // Delete assignment operator - can't assign lock_guard
    lock_guard& operator=(const lock_guard&) = delete;
};  
*/

void printMessage(string message, mutex &mtx) {
    lock_guard<mutex> lock_name(mtx);
    cout<<message<<endl;
}
int main() {
    mutex mtx;
    
    thread t1(printMessage, "Hello", ref(mtx));
    thread t2(printMessage, "World", ref(mtx));
    t1.join();
    t2.join();
    cout<<"Main thread ends"<<endl;
    return 0;
}