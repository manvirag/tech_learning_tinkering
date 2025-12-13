#include <iostream>
#include <thread>
#include <mutex>
#include <shared_mutex>
#include <chrono>
#include <vector>
#include <functional>
using namespace std;

/*
There are several types of mutex.
Most commonly used are:

std::mutex
Reference

Just your plain lockable mutex
std::timed_mutex
Reference

Timed mutex
You can lock for a specified amount of time with try_lock_for() and try_lock_until()
std::recursive_mutex
Reference

Multiple locks can be acquired by the same thread
You need to call unlock the same amount of times you've called lock before the lock is released
std::recursive_timed_mutex
Reference

Same as the recursive mutex, except it also has the timed locking methods that timed mutexes have
std::shared_mutex
Reference

Read-Write mutex
Can acquire both exclusive or shared locks (just use the appropriate lock guard type!)
std::shared_timed_mutex
Reference

Read-Write mutex with timeout support
Same as shared_mutex, but also supports try_lock_for() and try_lock_until()

*/

// ============================================================================
// 1. MUTEX - Basic mutual exclusion
// ============================================================================
void example1_Mutex() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 1: Basic Mutex (std::mutex)" << endl;
    cout << string(70, '=') << endl;
    cout << "Just your plain lockable mutex\n" << endl;
    
    mutex mtx;
    int counter = 0;
    
    auto increment = [&mtx, &counter](int id) {
        mtx.lock();
        cout << "Thread " << id << ": Counter before = " << counter;
        counter++;
        cout << ", Counter after = " << counter << endl;
        mtx.unlock();
    };
    
    thread t1(increment, 1);
    thread t2(increment, 2);
    thread t3(increment, 3);
    
    t1.join();
    t2.join();
    t3.join();
    
    cout << "\nFinal counter value: " << counter << endl;
    cout << "✓ Basic mutex ensures mutual exclusion\n" << endl;
}

// ============================================================================
// 2. RECURSIVE MUTEX - Allows same thread to lock multiple times
// ============================================================================
void example2_RecursiveMutex() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 2: Recursive Mutex (std::recursive_mutex)" << endl;
    cout << string(70, '=') << endl;
    cout << "Multiple locks can be acquired by the same thread\n";
    cout << "You need to call unlock the same amount of times you've called lock\n" << endl;
    
    recursive_mutex rmtx;
    int value = 0;
    
    // Function that calls itself recursively
    function<void(int, int)> recursiveFunction = [&rmtx, &value, &recursiveFunction](int depth, int maxDepth) {
        rmtx.lock();
        value++;
        cout << "Depth " << depth << ": Value = " << value << " (lock acquired, lock count increases)" << endl;
        
        if (depth < maxDepth) {
            // Recursive call - same thread locking again
            recursiveFunction(depth + 1, maxDepth);
        }
        
        cout << "Depth " << depth << ": Releasing lock (lock count decreases)" << endl;
        rmtx.unlock();
    };
    
    // With regular mutex, this would deadlock!
    // With recursive_mutex, same thread can lock multiple times
    thread t1([&]() {
        cout << "Thread 1: Starting recursive function\n";
        recursiveFunction(1, 3);
    });
    
    thread t2([&]() {
        this_thread::sleep_for(chrono::milliseconds(100));
        rmtx.lock();
        cout << "Thread 2: Acquired lock, value = " << value << endl;
        rmtx.unlock();
    });
    
    t1.join();
    t2.join();
    
    cout << "\nFinal value: " << value << endl;
    cout << "✓ Recursive mutex allows same thread to lock multiple times\n" << endl;
}

// ============================================================================
// 3. TIMED MUTEX - Allows try_lock_for and try_lock_until
// ============================================================================
void example3_TimedMutex() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 3: Timed Mutex (std::timed_mutex)" << endl;
    cout << string(70, '=') << endl;
    cout << "You can lock for a specified amount of time with try_lock_for() and try_lock_until()\n" << endl;
    
    timed_mutex tmtx;
    int sharedData = 0;
    
    // Thread that holds lock for a long time
    thread longTask([&]() {
        tmtx.lock();
        cout << "Long task: Acquired lock, working for 2 seconds..." << endl;
        this_thread::sleep_for(chrono::seconds(2));
        sharedData = 100;
        cout << "Long task: Finished, releasing lock" << endl;
        tmtx.unlock();
    });
    
    // Give long task time to acquire lock
    this_thread::sleep_for(chrono::milliseconds(100));
    
    // Thread that tries to acquire lock with timeout using try_lock_for
    thread quickTask([&]() {
        cout << "Quick task: Trying to acquire lock using try_lock_for() (timeout: 1 second)..." << endl;
        
        // Try to lock with timeout
        if (tmtx.try_lock_for(chrono::seconds(1))) {
            cout << "Quick task: Successfully acquired lock!" << endl;
            sharedData = 200;
            tmtx.unlock();
        } else {
            cout << "Quick task: Failed to acquire lock within timeout!" << endl;
        }
    });
    
    // Another thread using try_lock_until
    thread timedTask([&]() {
        this_thread::sleep_for(chrono::milliseconds(500));
        auto timeout = chrono::steady_clock::now() + chrono::milliseconds(500);
        cout << "Timed task: Trying to acquire lock using try_lock_until()..." << endl;
        
        if (tmtx.try_lock_until(timeout)) {
            cout << "Timed task: Successfully acquired lock!" << endl;
            sharedData = 300;
            tmtx.unlock();
        } else {
            cout << "Timed task: Failed to acquire lock within timeout!" << endl;
        }
    });
    
    longTask.join();
    quickTask.join();
    timedTask.join();
    
    cout << "\nFinal shared data: " << sharedData << endl;
    cout << "✓ Timed mutex allows non-blocking lock attempts with timeout\n" << endl;
}

// ============================================================================
// 4. RECURSIVE TIMED MUTEX - Recursive + Timed
// ============================================================================
void example4_RecursiveTimedMutex() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 4: Recursive Timed Mutex (std::recursive_timed_mutex)" << endl;
    cout << string(70, '=') << endl;
    cout << "Same as the recursive mutex, except it also has the timed locking methods\n" << endl;
    
    recursive_timed_mutex rtmtx;
    int value = 0;
    
    // Recursive function with timeout capability
    function<void(int, int)> recursiveFunctionWithTimeout = [&rtmtx, &value, &recursiveFunctionWithTimeout](int depth, int maxDepth) {
        // Try to lock with timeout
        if (rtmtx.try_lock_for(chrono::milliseconds(100))) {
            value++;
            cout << "Depth " << depth << ": Value = " << value << " (lock acquired)" << endl;
            
            if (depth < maxDepth) {
                // Recursive call - same thread locking again
                recursiveFunctionWithTimeout(depth + 1, maxDepth);
            }
            
            cout << "Depth " << depth << ": Releasing lock" << endl;
            rtmtx.unlock();
        } else {
            cout << "Depth " << depth << ": Failed to acquire lock within timeout!" << endl;
        }
    };
    
    thread t1([&]() {
        cout << "Thread 1: Starting recursive function with timeout support\n";
        recursiveFunctionWithTimeout(1, 3);
    });
    
    thread t2([&]() {
        this_thread::sleep_for(chrono::milliseconds(50));
        cout << "Thread 2: Trying to acquire lock with timeout..." << endl;
        if (rtmtx.try_lock_for(chrono::milliseconds(500))) {
            cout << "Thread 2: Acquired lock, value = " << value << endl;
            rtmtx.unlock();
        } else {
            cout << "Thread 2: Failed to acquire lock within timeout!" << endl;
        }
    });
    
    t1.join();
    t2.join();
    
    cout << "\nFinal value: " << value << endl;
    cout << "✓ Recursive timed mutex: Recursive + timeout support\n" << endl;
}

// ============================================================================
// 5. SHARED MUTEX - Read-Write mutex
// ============================================================================
void example5_SharedMutex() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 5: Shared Mutex (std::shared_mutex)" << endl;
    cout << string(70, '=') << endl;
    cout << "Read-Write mutex - Can acquire both exclusive or shared locks\n";
    cout << "Use shared_lock for readers, unique_lock for writers\n";
    cout << "Multiple readers can read simultaneously, but only one writer at a time\n" << endl;
    
    shared_mutex smtx;
    int data = 0;
    
    // Reader function - uses shared_lock (multiple readers can read simultaneously)
    auto reader = [&smtx, &data](int id) {
        shared_lock<shared_mutex> lock(smtx);
        cout << "Reader " << id << ": Reading data = " << data << " (shared lock acquired)" << endl;
        this_thread::sleep_for(chrono::milliseconds(200));
        cout << "Reader " << id << ": Finished reading" << endl;
    };
    
    // Writer function - uses unique_lock (exclusive access)
    auto writer = [&smtx, &data](int id, int newValue) {
        unique_lock<shared_mutex> lock(smtx);
        cout << "Writer " << id << ": Writing data from " << data << " to " << newValue 
             << " (exclusive lock acquired)" << endl;
        this_thread::sleep_for(chrono::milliseconds(300));
        data = newValue;
        cout << "Writer " << id << ": Finished writing. New value = " << data << endl;
    };
    
    // Multiple readers can read simultaneously
    cout << "[Starting multiple readers - they can read simultaneously]\n" << endl;
    vector<thread> readers;
    for (int i = 1; i <= 3; i++) {
        readers.emplace_back(reader, i);
    }
    
    // Wait a bit, then start a writer (will wait for readers)
    this_thread::sleep_for(chrono::milliseconds(50));
    cout << "\n[Starting writer - it will wait for all readers to finish]\n" << endl;
    thread w1([&]() { writer(1, 100); });
    
    // Wait for readers
    for (auto& t : readers) {
        t.join();
    }
    
    // Writer can now proceed
    w1.join();
    
    // Start more readers after writer finishes
    cout << "\n[Writer finished - starting more readers]\n" << endl;
    thread r4([&]() { reader(4); });
    thread r5([&]() { reader(5); });
    
    r4.join();
    r5.join();
    
    // Demonstrate that writers block each other
    cout << "\n[Starting multiple writers - they will execute one at a time]\n" << endl;
    vector<thread> writers;
    for (int i = 2; i <= 3; i++) {
        writers.emplace_back([&, i]() { writer(i, i * 10); });
    }
    
    for (auto& t : writers) {
        t.join();
    }
    
    cout << "\nFinal data: " << data << endl;
    cout << "✓ Shared mutex: Multiple readers OR single writer\n" << endl;
}

// ============================================================================
// 6. SHARED TIMED MUTEX - Read-Write mutex with timeout
// ============================================================================
void example6_SharedTimedMutex() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 6: Shared Timed Mutex (std::shared_timed_mutex)" << endl;
    cout << string(70, '=') << endl;
    cout << "Read-Write mutex with timeout support\n";
    cout << "Same as shared_mutex, but also supports try_lock_for() and try_lock_until()\n";
    cout << "Use shared_lock for readers, unique_lock for writers\n" << endl;
    
    shared_timed_mutex stmtx;
    int data = 0;
    
    // Reader with timeout - uses shared_lock
    auto readerWithTimeout = [&stmtx, &data](int id) {
        cout << "Reader " << id << ": Trying to acquire shared lock (timeout: 500ms)..." << endl;
        
        shared_lock<shared_timed_mutex> lock(stmtx, defer_lock);
        if (lock.try_lock_for(chrono::milliseconds(500))) {
            cout << "Reader " << id << ": Acquired shared lock, reading data = " << data << endl;
            this_thread::sleep_for(chrono::milliseconds(200));
            cout << "Reader " << id << ": Finished reading" << endl;
        } else {
            cout << "Reader " << id << ": Failed to acquire lock within timeout!" << endl;
        }
    };
    
    // Writer with timeout - uses unique_lock
    auto writerWithTimeout = [&stmtx, &data](int id, int newValue) {
        cout << "Writer " << id << ": Trying to acquire exclusive lock (timeout: 1s)..." << endl;
        
        unique_lock<shared_timed_mutex> lock(stmtx, defer_lock);
        if (lock.try_lock_for(chrono::seconds(1))) {
            cout << "Writer " << id << ": Acquired exclusive lock, writing " << newValue << endl;
            this_thread::sleep_for(chrono::milliseconds(300));
            data = newValue;
            cout << "Writer " << id << ": Finished writing" << endl;
        } else {
            cout << "Writer " << id << ": Failed to acquire lock within timeout!" << endl;
        }
    };
    
    // Long-running writer
    thread longWriter([&]() {
        unique_lock<shared_timed_mutex> lock(stmtx);
        cout << "Long writer: Acquired exclusive lock, working for 1.5 seconds..." << endl;
        this_thread::sleep_for(chrono::milliseconds(1500));
        data = 999;
        cout << "Long writer: Finished, releasing lock" << endl;
    });
    
    // Give long writer time to acquire lock
    this_thread::sleep_for(chrono::milliseconds(100));
    
    // Multiple readers trying with timeout (will fail because writer has exclusive lock)
    thread r1([&]() { readerWithTimeout(1); });
    thread r2([&]() { readerWithTimeout(2); });
    
    // Writer trying with timeout (will fail because another writer has exclusive lock)
    thread w1([&]() { writerWithTimeout(1, 100); });
    
    longWriter.join();
    r1.join();
    r2.join();
    w1.join();
    
    // Now readers should succeed (no writer blocking)
    cout << "\n[Writer released lock - readers can now proceed]\n" << endl;
    thread r3([&]() { readerWithTimeout(3); });
    thread r4([&]() { readerWithTimeout(4); });
    
    r3.join();
    r4.join();
    
    cout << "\nFinal data: " << data << endl;
    cout << "✓ Shared timed mutex: Read-Write locks + timeout support\n" << endl;
}

int main(){
    cout << "\n" << string(70, '=') << endl;
    cout << "MUTEX TYPES IN C++ - COMPREHENSIVE EXAMPLES" << endl;
    cout << string(70, '=') << endl;
    
    example1_Mutex();
    example2_RecursiveMutex();
    example3_TimedMutex();
    example4_RecursiveTimedMutex();
    example5_SharedMutex();
    example6_SharedTimedMutex();
    
    cout << "\n" << string(70, '=') << endl;
    cout << "SUMMARY:" << endl;
    cout << string(70, '=') << endl;
    cout << "1. std::mutex: Basic mutual exclusion" << endl;
    cout << "2. std::recursive_mutex: Same thread can lock multiple times" << endl;
    cout << "3. std::timed_mutex: Supports try_lock_for() and try_lock_until()" << endl;
    cout << "4. std::recursive_timed_mutex: Recursive + timed locking methods" << endl;
    cout << "5. std::shared_mutex: Read-Write mutex" << endl;
    cout << "   - Use shared_lock for readers (multiple can read simultaneously)" << endl;
    cout << "   - Use unique_lock for writers (exclusive access)" << endl;
    cout << "6. std::shared_timed_mutex: Read-Write mutex with timeout support" << endl;
    cout << "   - Same as shared_mutex + timeout methods" << endl;
    cout << string(70, '=') << endl;
    
    return 0;
}
