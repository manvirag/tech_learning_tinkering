#include <iostream>
#include <vector>
#include <thread>
#include <mutex>
#include <chrono>
#include <semaphore>  // C++20
#include <queue>      // For producer-consumer example
using namespace std;

// ============================================================================
// SEMAPHORES IN C++
// ============================================================================
/*
Semaphores were introduced in C++20 (2020)

Types:
1. std::counting_semaphore<MaxValue> - General counting semaphore
2. std::binary_semaphore - Alias for counting_semaphore<1>

Key Methods:
- acquire() / wait() - Decrements counter, blocks if counter is 0
- release() / signal() - Increments counter, wakes up waiting threads
- try_acquire() - Non-blocking attempt to acquire
- try_acquire_for(timeout) - Try to acquire with timeout
- try_acquire_until(time_point) - Try to acquire until time point

Semaphore vs Mutex:
- Mutex: Only one thread can hold lock (binary)
- Semaphore: Multiple threads can acquire (counting)
- Mutex: Must be released by same thread that acquired
- Semaphore: Can be released by different thread

Use Cases:
- Resource pool management
- Producer-Consumer with bounded buffer
- Rate limiting
- Thread coordination
- Limiting concurrent access
================================================================================
*/

// ============================================================================
// Example 1: Basic Binary Semaphore
// ============================================================================
void example1_BinarySemaphore() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 1: Basic Binary Semaphore" << endl;
    cout << string(70, '=') << endl;
    cout << "Binary semaphore allows only one thread at a time\n" << endl;
    
    binary_semaphore sem(1);  // Initial value: 1 (available)
    int sharedResource = 0;
    
    auto worker = [&sem, &sharedResource](int id) {
        for (int i = 0; i < 3; i++) {
            sem.acquire();  // Wait for semaphore (decrements from 1 to 0)
            
            // Critical section
            sharedResource++;
            cout << "Thread " << id << ": Accessing resource, value = " 
                 << sharedResource << endl;
            this_thread::sleep_for(chrono::milliseconds(100));
            
            sem.release();  // Release semaphore (increments from 0 to 1)
        }
    };
    
    thread t1(worker, 1);
    thread t2(worker, 2);
    
    t1.join();
    t2.join();
    
    cout << "\nFinal value: " << sharedResource << endl;
    cout << "✓ Binary semaphore ensures mutual exclusion\n" << endl;
}

// ============================================================================
// Example 2: Counting Semaphore - Resource Pool
// ============================================================================

// const int MAX_RESOURCES = 3;
// counting_semaphore<10> sem(MAX_RESOURCES);
// //                         ↑              ↑
// //                         │              └─ Initial value = 3
// //                         └──────────────── Max possible value = 10


// counting_semaphore<10> sem(3);  // Start with 3 permits

// sem.acquire();  // Counter: 3 → 2 (one thread can proceed)
// sem.acquire();  // Counter: 2 → 1 (another thread can proceed)
// sem.acquire();  // Counter: 1 → 0 (third thread can proceed)
// sem.acquire();  // Counter: 0 → BLOCKS! (fourth thread waits)

// sem.release();  // Counter: 0 → 1 (wakes up waiting thread)
// sem.release();  // Counter: 1 → 2
// // ... can release up to 10 total (the template limit)


// // Database connection pool: max 10 connections, but only 3 available initially
// counting_semaphore<10> dbPool(3);

// // Later, you can add more connections dynamically:
// dbPool.release();  // Now 4 available
// dbPool.release();  // Now 5 available
// // ... up to 10 total
void example2_CountingSemaphore() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 2: Counting Semaphore - Resource Pool" << endl;
    cout << string(70, '=') << endl;
    cout << "Allows multiple threads to access limited resources\n" << endl;
    
    const int MAX_RESOURCES = 3;
    counting_semaphore<10> sem(MAX_RESOURCES);  // Max 3 resources available
    int resourceCount = 0;
    mutex printMtx;
    
    auto useResource = [&sem, &resourceCount, &printMtx](int threadId) {
        sem.acquire();  // Acquire one resource (decrements counter)
        
        {
            lock_guard<mutex> lock(printMtx);
            resourceCount++;
            cout << "Thread " << threadId << ": Using resource (Total in use: " 
                 << resourceCount << "/" << MAX_RESOURCES << ")" << endl;
        }
        
        // Simulate resource usage
        this_thread::sleep_for(chrono::milliseconds(500));
        
        {
            lock_guard<mutex> lock(printMtx);
            resourceCount--;
            cout << "Thread " << threadId << ": Released resource (Total in use: " 
                 << resourceCount << "/" << MAX_RESOURCES << ")" << endl;
        }
        
        sem.release();  // Release resource (increments counter)
    };
    
    // Create 5 threads competing for 3 resources
    vector<thread> threads;
    for (int i = 1; i <= 5; i++) {
        threads.emplace_back(useResource, i);
    }
    
    for (auto& t : threads) {
        t.join();
    }
    
    cout << "\n✓ Counting semaphore limits concurrent resource access\n" << endl;
}

// ============================================================================
// Example 3: Producer-Consumer with Semaphore
// ============================================================================
void example3_ProducerConsumer() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 3: Producer-Consumer with Semaphore" << endl;
    cout << string(70, '=') << endl;
    cout << "Semaphores coordinate producers and consumers\n" << endl;
    
    const int BUFFER_SIZE = 5;
    queue<int> buffer;
    mutex bufferMtx;
    
    // Semaphores for coordination
    counting_semaphore<10> emptySlots(BUFFER_SIZE);  // Initially all slots empty
    counting_semaphore<10> fullSlots(0);             // Initially no items
    
    auto producer = [&](int id) {
        for (int i = 1; i <= 3; i++) {
            int item = id * 100 + i;
            
            emptySlots.acquire();  // Wait for empty slot
            
            {
                lock_guard<mutex> lock(bufferMtx);
                buffer.push(item);
                cout << "Producer " << id << ": Produced item " << item 
                     << " (Buffer size: " << buffer.size() << ")" << endl;
            }
            
            fullSlots.release();  // Signal that item is available
            this_thread::sleep_for(chrono::milliseconds(100));
        }
    };
    
    auto consumer = [&](int id) {
        for (int i = 1; i <= 3; i++) {
            fullSlots.acquire();  // Wait for item to be available
            
            int item;
            {
                lock_guard<mutex> lock(bufferMtx);
                item = buffer.front();
                buffer.pop();
                cout << "Consumer " << id << ": Consumed item " << item 
                     << " (Buffer size: " << buffer.size() << ")" << endl;
            }
            
            emptySlots.release();  // Signal that slot is empty
            this_thread::sleep_for(chrono::milliseconds(150));
        }
    };
    
    thread p1(producer, 1);
    thread p2(producer, 2);
    thread c1(consumer, 1);
    thread c2(consumer, 2);
    
    p1.join();
    p2.join();
    c1.join();
    c2.join();
    
    cout << "\n✓ Producer-Consumer coordination with semaphores\n" << endl;
}

// ============================================================================
// Example 4: Try Acquire (Non-blocking)
// ============================================================================
void example4_TryAcquire() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 4: Try Acquire (Non-blocking)" << endl;
    cout << string(70, '=') << endl;
    cout << "Non-blocking semaphore acquisition\n" << endl;
    
    counting_semaphore<5> sem(2);  // Only 2 permits available
    int successCount = 0;
    int failCount = 0;
    mutex mtx;
    
    auto tryAccess = [&](int id) {
        if (sem.try_acquire()) {
            lock_guard<mutex> lock(mtx);
            successCount++;
            cout << "Thread " << id << ": Successfully acquired semaphore" << endl;
            
            this_thread::sleep_for(chrono::milliseconds(200));
            sem.release();
        } else {
            lock_guard<mutex> lock(mtx);
            failCount++;
            cout << "Thread " << id << ": Failed to acquire semaphore (busy)" << endl;
        }
    };
    
    // Create 5 threads trying to acquire
    vector<thread> threads;
    for (int i = 1; i <= 5; i++) {
        threads.emplace_back(tryAccess, i);
    }
    
    for (auto& t : threads) {
        t.join();
    }
    
    cout << "\nSuccess: " << successCount << ", Failed: " << failCount << endl;
    cout << "✓ Non-blocking acquisition prevents thread blocking\n" << endl;
}

// ============================================================================
// Example 5: Try Acquire with Timeout
// ============================================================================
void example5_TryAcquireTimeout() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 5: Try Acquire with Timeout" << endl;
    cout << string(70, '=') << endl;
    cout << "Semaphore acquisition with timeout\n" << endl;
    
    counting_semaphore<5> sem(1);  // Only 1 permit
    
    // Thread that holds semaphore for a while
    thread holder([&sem]() {
        sem.acquire();
        cout << "Holder: Acquired semaphore, holding for 2 seconds..." << endl;
        this_thread::sleep_for(chrono::seconds(2));
        sem.release();
        cout << "Holder: Released semaphore" << endl;
    });
    
    this_thread::sleep_for(chrono::milliseconds(100));
    
    // Thread trying with timeout
    thread waiter([&sem]() {
        cout << "Waiter: Trying to acquire with 1 second timeout..." << endl;
        
        auto timeout = chrono::milliseconds(1000);
        if (sem.try_acquire_for(timeout)) {
            cout << "Waiter: Successfully acquired!" << endl;
            sem.release();
        } else {
            cout << "Waiter: Timeout! Could not acquire within 1 second" << endl;
        }
    });
    
    holder.join();
    waiter.join();
    
    cout << "\n✓ Timeout prevents indefinite blocking\n" << endl;
}

// ============================================================================
// Example 6: Rate Limiting with Semaphore
// ============================================================================
void example6_RateLimiting() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 6: Rate Limiting with Semaphore" << endl;
    cout << string(70, '=') << endl;
    cout << "Limit number of operations per time period\n" << endl;
    
    const int MAX_REQUESTS = 3;
    counting_semaphore<10> rateLimiter(MAX_REQUESTS);
    
    auto makeRequest = [&rateLimiter](int requestId) {
        rateLimiter.acquire();
        
        cout << "Request " << requestId << ": Processing..." << endl;
        this_thread::sleep_for(chrono::milliseconds(300));
        cout << "Request " << requestId << ": Completed" << endl;
        
        rateLimiter.release();
    };
    
    // Simulate 10 requests
    vector<thread> requests;
    for (int i = 1; i <= 10; i++) {
        requests.emplace_back(makeRequest, i);
        this_thread::sleep_for(chrono::milliseconds(50));
    }
    
    for (auto& t : requests) {
        t.join();
    }
    
    cout << "\n✓ Rate limiting controls concurrent operations\n" << endl;
}

int main() {
    cout << "\n" << string(70, '=') << endl;
    cout << "SEMAPHORES IN C++ - COMPREHENSIVE EXAMPLES" << endl;
    cout << string(70, '=') << endl;
    cout << "\nSemaphores were introduced in C++20 (2020)" << endl;
    cout << "Types: counting_semaphore<Max> and binary_semaphore\n" << endl;
    
    example1_BinarySemaphore();
    example2_CountingSemaphore();
    example3_ProducerConsumer();
    example4_TryAcquire();
    example5_TryAcquireTimeout();
    example6_RateLimiting();
    
    cout << "\n" << string(70, '=') << endl;
    cout << "SEMAPHORE SUMMARY:" << endl;
    cout << string(70, '=') << endl;
    cout << "Introduced: C++20 (2020)" << endl;
    cout << "\nTypes:" << endl;
    cout << "  - counting_semaphore<MaxValue>: General counting semaphore" << endl;
    cout << "  - binary_semaphore: Alias for counting_semaphore<1>" << endl;
    cout << "\nKey Methods:" << endl;
    cout << "  - acquire() / wait(): Decrement, block if 0" << endl;
    cout << "  - release() / signal(): Increment, wake waiting threads" << endl;
    cout << "  - try_acquire(): Non-blocking attempt" << endl;
    cout << "  - try_acquire_for(timeout): Try with timeout" << endl;
    cout << "  - try_acquire_until(time_point): Try until time point" << endl;
    cout << "\nUse Cases:" << endl;
    cout << "  - Resource pool management" << endl;
    cout << "  - Producer-Consumer coordination" << endl;
    cout << "  - Rate limiting" << endl;
    cout << "  - Thread coordination" << endl;
    cout << "  - Limiting concurrent access" << endl;
    cout << string(70, '=') << endl;
    
    return 0;
}

