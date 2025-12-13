#include <iostream>
#include <thread>
#include <mutex>  // Required for scoped_lock, unique_lock and mutex
#include <condition_variable>
#include <chrono>
using namespace std;

/*
================================================================================
std::scoped_lock<> - C++17 Feature
================================================================================

std::scoped_lock is the modern replacement for std::lock_guard (C++17+).
It provides RAII (Resource Acquisition Is Initialization) benefits with
the ability to lock MULTIPLE mutexes simultaneously.

KEY FEATURES:
1. Can lock a SINGLE mutex (like lock_guard)
2. Can lock MULTIPLE mutexes simultaneously (unique to scoped_lock)
3. Automatically prevents deadlocks when locking multiple mutexes
4. Automatically unlocks when going out of scope (exception-safe)

SYNTAX:
    scoped_lock<mutex> lock(mtx);              // Single mutex
    scoped_lock<mutex, mutex> lock(mtx1, mtx2); // Multiple mutexes

ADVANTAGES over lock_guard:
- Can handle multiple mutexes at once
- Uses std::lock() internally to prevent deadlocks
- More flexible and modern (C++17 standard)
================================================================================

changed from 
g++ -std=c++11 -pthread practice.cpp -o practice && ./practice

to 
g++ -std=c++17 -pthread practice.cpp -o practice && ./practice

multile mutex use relative less as compared to single mutex.
they are mostly using in multiple data updates.
*/

// Real-world example: Bank Account Transfer System
// Each account has its own mutex to protect its balance

class BankAccount {
private:
    int balance;
    mutex mtx;  // Each account has its own mutex
    
    // Internal methods that don't lock (assume caller already has lock)
    int getBalanceUnsafe() {
        return balance;
    }
    
    void depositUnsafe(int amount) {
        balance += amount;
    }
    
    void withdrawUnsafe(int amount) {
        balance -= amount;
    }
    
public:
    BankAccount(int initialBalance) : balance(initialBalance) {}
    
    // Public methods with locking for single-threaded operations
    int getBalance() {
        lock_guard<mutex> lock(mtx);
        return balance;
    }
    
    mutex& getMutex() {
        return mtx;  // Expose mutex for scoped_lock
    }
    
    // Friend function can access private members when mutexes are already locked
    friend void transfer(BankAccount& from, BankAccount& to, int amount);
};

// Transfer money from one account to another
// This is where scoped_lock shines - we need to lock BOTH accounts simultaneously
void transfer(BankAccount& from, BankAccount& to, int amount) {
    // Lock BOTH account mutexes at once - prevents deadlock!
    // scoped_lock uses std::lock() internally which uses deadlock avoidance algorithm
    // This ensures both accounts are locked atomically, preventing race conditions
    scoped_lock<mutex, mutex> lock(from.getMutex(), to.getMutex());
    
    // Now we can safely update both accounts without double-locking
    int fromBalance = from.getBalanceUnsafe();
    if (fromBalance >= amount) {
        from.withdrawUnsafe(amount);
        to.depositUnsafe(amount);
        cout << "Transferred $" << amount 
             << " (From balance: $" << fromBalance 
             << " -> $" << from.getBalanceUnsafe() << ")" << endl;
    } else {
        cout << "Transfer failed: Insufficient funds (Balance: $" 
             << fromBalance << ", Required: $" << amount << ")" << endl;
    }
    
    // Both mutexes automatically unlocked when function exits
    // Even if an exception occurs, the locks are released safely
}

int main() {
    BankAccount account1(1000);  // Account 1 starts with $1000
    BankAccount account2(500);   // Account 2 starts with $500
    
    cout << "Initial balances:" << endl;
    cout << "Account1: $" << account1.getBalance() << endl;
    cout << "Account2: $" << account2.getBalance() << endl;
    cout << endl;
    
    // Simulate multiple concurrent transfers
    thread t1(transfer, ref(account1), ref(account2), 200);
    thread t2(transfer, ref(account2), ref(account1), 100);
    thread t3(transfer, ref(account1), ref(account2), 150);
    
    t1.join();
    t2.join();
    t3.join();
    
    cout << endl;
    cout << "Final balances:" << endl;
    cout << "Account1: $" << account1.getBalance() << endl;
    cout << "Account2: $" << account2.getBalance() << endl;
    
    // Demonstrate unique_lock features
    demonstrateUniqueLock();
    
    return 0;
}

/*
================================================================================
std::unique_lock<> - Flexible Lock Management
================================================================================

std::unique_lock is more flexible than lock_guard and scoped_lock.
It provides additional features while maintaining RAII benefits.

KEY FEATURES:
1. Can unlock manually before going out of scope
2. Can be moved/transferred between functions
3. Supports deferred locking (lock later, not immediately)
4. Can be used with condition variables (required for wait operations)
5. Supports try_lock operations (non-blocking)
6. Can check if lock is owned

SYNTAX:
    unique_lock<mutex> lock(mtx);              // Lock immediately
    unique_lock<mutex> lock(mtx, defer_lock);  // Defer locking
    unique_lock<mutex> lock(mtx, try_to_lock); // Try to lock (non-blocking)

ADVANTAGES over lock_guard:
- More flexible: can unlock early
- Can be used with condition variables
- Supports deferred and try locking
- Can be moved (transfer ownership)

WHEN TO USE:
- When you need to unlock before scope ends
- With condition variables (required for cv.wait())
- When you need try_lock functionality
- When you need to transfer lock ownership
================================================================================
*/

// Real-world example: Shared Resource with Conditional Access
class SharedResource {
private:
    int data;
    mutex mtx;
    condition_variable cv;
    bool ready;
    
public:
    SharedResource() : data(0), ready(false) {}
    
    // Producer: Updates data and notifies consumers
    void updateData(int newData) {
        unique_lock<mutex> lock(mtx);  // Lock the mutex
        
        data = newData;
        ready = true;
        
        cout << "Producer: Updated data to " << data << endl;
        
        // Unlock before notifying (good practice for condition variables)
        lock.unlock();
        cv.notify_all();  // Notify all waiting threads
        
        // Lock automatically released when lock goes out of scope
    }
    
    // Consumer: Waits for data to be ready
    void waitForData() {
        unique_lock<mutex> lock(mtx);  // Lock the mutex
        
        // Wait for condition (unique_lock REQUIRED for condition_variable)
        // This will unlock the mutex while waiting and re-lock when notified
        cv.wait(lock, [this] { return ready; });
        
        cout << "Consumer: Received data = " << data << endl;
        ready = false;
        
        // Lock automatically released
    }
    
    // Example: Manual unlock before expensive operation
    void processWithEarlyUnlock() {
        unique_lock<mutex> lock(mtx);
        
        // Critical section - need lock
        int value = data;
        cout << "Processing: Got value " << value << endl;
        
        // Unlock early - don't hold lock during expensive operation
        lock.unlock();
        
        // Expensive operation (simulated)
        this_thread::sleep_for(chrono::milliseconds(100));
        cout << "Processing: Completed expensive operation" << endl;
        
        // No need to unlock - already unlocked
    }
    
    // Example: Try lock (non-blocking)
    bool tryUpdate(int newData) {
        unique_lock<mutex> lock(mtx, try_to_lock);
        
        if (lock.owns_lock()) {
            // Successfully acquired lock
            data = newData;
            cout << "TryUpdate: Successfully updated to " << data << endl;
            return true;
        } else {
            // Could not acquire lock (non-blocking)
            cout << "TryUpdate: Failed - resource busy" << endl;
            return false;
        }
    }
    
    // Method that holds lock for a while (for try_lock demonstration)
    void longOperation() {
        unique_lock<mutex> lock(mtx);
        cout << "LongOperation: Started (holding lock)" << endl;
        this_thread::sleep_for(chrono::milliseconds(200));
        cout << "LongOperation: Completed (releasing lock)" << endl;
    }
    
    int getData() {
        lock_guard<mutex> lock(mtx);
        return data;
    }
};

// Producer thread function
void producer(SharedResource& resource) {
    for (int i = 1; i <= 3; i++) {
        this_thread::sleep_for(chrono::milliseconds(500));
        resource.updateData(i * 10);
    }
}

// Consumer thread function
void consumer(SharedResource& resource, int id) {
    for (int i = 0; i < 3; i++) {
        resource.waitForData();
    }
}

// Example demonstrating unique_lock features
void demonstrateUniqueLock() {
    cout << "\n" << string(70, '=') << endl;
    cout << "UNIQUE_LOCK EXAMPLES" << endl;
    cout << string(70, '=') << "\n" << endl;
    
    SharedResource resource;
    
    // Example 1: Producer-Consumer with condition variables
    cout << "Example 1: Producer-Consumer Pattern" << endl;
    cout << "------------------------------------" << endl;
    thread prod(producer, ref(resource));
    thread cons1(consumer, ref(resource), 1);
    
    prod.join();
    cons1.join();
    
    cout << "\nExample 2: Early Unlock" << endl;
    cout << "-----------------------" << endl;
    thread t1([&resource]() { resource.processWithEarlyUnlock(); });
    thread t2([&resource]() { resource.processWithEarlyUnlock(); });
    t1.join();
    t2.join();
    
    cout << "\nExample 3: Try Lock (Non-blocking)" << endl;
    cout << "-----------------------------------" << endl;
    // Start a thread that holds the lock for a while
    thread t3([&resource]() {
        resource.longOperation();
    });
    
    // Try to update while longOperation holds the lock (will fail)
    this_thread::sleep_for(chrono::milliseconds(50));
    bool success1 = resource.tryUpdate(999);
    
    // Wait for longOperation to finish
    t3.join();
    
    // Now try again (should succeed - no contention)
    bool success2 = resource.tryUpdate(888);
    
    cout << "\nFinal data value: " << resource.getData() << endl;
}
