#include <iostream>
#include <thread>
#include <mutex>
#include <shared_mutex>  // Required for shared_mutex and shared_lock
#include <chrono>
#include <vector>
using namespace std;

/*

A shared lock is just like a unique lock, except the lock is a shared lock as opposed to an exclusive one.


Just like the normal lock guard, except...
It initialises a shared lock
It can be returned from the function without releasing the lock (via move semantics)
It can be released before it is destroyed
You can also use nifty lock methods!



std::shared_lock my_mutex;
std::shared_lock<std::shared_mutex> guard(my_mutex);

// Check if guard owns lock (either works)
guard.owns_lock();
bool(guard);

// Return function without releasing the lock
return std::move(guard);

// Release lock before destruction
guard.unlock();




If you defer the locks, you can use the nifty lock methods!

// Initialise the lock guard, but don't actually lock yet
std::shared_lock<std::shared_mutex> guard(mutex_1, std::defer_lock);

// Now you can do some of the following!
guard.lock(); // Lock now!
guard.try_lock(); // Won't block if it can't acquire
guard.try_lock_for(); // Only for timed_mutexes
guard.try_lock_until(); // Only for timed_mutexes

require c++17



Exclusive lock mode prevents the associated resource from being shared. This lock mode is obtained to modify data. The first transaction to lock a resource exclusively is the only transaction that can alter the resource until the exclusive lock is released.

Share lock mode allows the associated resource to be shared, depending on the operations involved. Multiple users reading data can share the data, holding share locks to prevent concurrent access by a writer (who needs an exclusive lock). Several transactions can acquire share locks on the same resource.

Notice this means that if an object is shared locked, you can acquire shared locks, but not exclusive locks.

Basically:

If there are multiple readers, no writers can bind, but readers can bind.
If there is one writer, no one can bind.

*/

// Real-world example: Shared Data Store
// Demonstrates exclusive locks (writers) vs shared locks (readers)
class SharedDataStore {
private:
    int data;
    shared_mutex mtx;  // Can be locked in shared or exclusive mode
    
public:
    SharedDataStore(int initialValue) : data(initialValue) {}
    
    // READ operation - uses SHARED LOCK (multiple readers can read simultaneously)
    int read(int readerId) {
        shared_lock<shared_mutex> lock(mtx);  // Shared lock - allows multiple readers
        
        cout << "Reader " << readerId << ": Reading data = " << data 
             << " (shared lock acquired)" << endl;
        
        // Simulate reading time
        this_thread::sleep_for(chrono::milliseconds(200));
        
        cout << "Reader " << readerId << ": Finished reading" << endl;
        
        return data;
        // Shared lock automatically released
    }
    
    // WRITE operation - uses EXCLUSIVE LOCK (only one writer at a time)
    void write(int newValue, int writerId) {
        unique_lock<shared_mutex> lock(mtx);  // Exclusive lock - blocks all others
        
        cout << "Writer " << writerId << ": Writing data from " << data 
             << " to " << newValue << " (exclusive lock acquired)" << endl;
        
        // Simulate writing time
        this_thread::sleep_for(chrono::milliseconds(300));
        
        data = newValue;
        
        cout << "Writer " << writerId << ": Finished writing. New value = " << data << endl;
        
        // Exclusive lock automatically released
    }
};

// Example 1: Multiple readers can read simultaneously (shared locks)
void example1_MultipleReaders() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 1: Multiple Readers (Shared Locks)" << endl;
    cout << string(70, '=') << endl;
    cout << "Multiple readers can read simultaneously - they don't block each other!\n" << endl;
    
    SharedDataStore store(100);
    vector<thread> readers;
    
    // Create 5 readers that all read at the same time
    for (int i = 1; i <= 5; i++) {
        readers.emplace_back([&store, i]() {
            store.read(i);
        });
    }
    
    // All readers should execute concurrently
    for (auto& t : readers) {
        t.join();
    }
    
    cout << "\n✓ All readers completed - they ran simultaneously!" << endl;
}

// Example 2: Only one writer at a time (exclusive locks)
void example2_MultipleWriters() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 2: Multiple Writers (Exclusive Locks)" << endl;
    cout << string(70, '=') << endl;
    cout << "Writers block each other - only one can write at a time!\n" << endl;
    
    SharedDataStore store(100);
    vector<thread> writers;
    
    // Create 3 writers
    for (int i = 1; i <= 3; i++) {
        writers.emplace_back([&store, i]() {
            store.write(i * 10, i);
        });
    }
    
    // Writers will execute one at a time (sequentially)
    for (auto& t : writers) {
        t.join();
    }
    
    cout << "\n✓ All writers completed - they ran one at a time!" << endl;
    cout << "Final value: " << store.read(0) << endl;
}

// Example 3: Readers block writers, but not other readers
void example3_ReadersVsWriter() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 3: Readers vs Writer" << endl;
    cout << string(70, '=') << endl;
    cout << "When readers are active, writer must wait!" << endl;
    cout << "But readers don't block each other.\n" << endl;
    
    SharedDataStore store(100);
    
    // Start multiple readers first
    vector<thread> readers;
    for (int i = 1; i <= 3; i++) {
        readers.emplace_back([&store, i]() {
            store.read(i);
        });
    }
    
    // Start a writer - it will wait for all readers to finish
    this_thread::sleep_for(chrono::milliseconds(50));
    thread writer([&store]() {
        cout << "\n[Writer trying to write while readers are reading...]" << endl;
        store.write(999, 1);
    });
    
    // Wait for all readers
    for (auto& t : readers) {
        t.join();
    }
    
    // Now writer can proceed
    writer.join();
    
    cout << "\n✓ Writer waited for all readers to finish!" << endl;
}

// Example 4: Writer blocks everyone (readers and other writers)
void example4_WriterBlocksEveryone() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 4: Writer Blocks Everyone" << endl;
    cout << string(70, '=') << endl;
    cout << "When a writer is active, NO ONE can access (readers or writers)!\n" << endl;
    
    SharedDataStore store(100);
    
    // Start a writer first
    thread writer([&store]() {
        store.write(500, 1);
    });
    
    // Give writer time to acquire lock
    this_thread::sleep_for(chrono::milliseconds(50));
    
    // Try to start readers and another writer - they all wait
    cout << "[Starting readers and another writer while first writer is active...]" << endl;
    
    vector<thread> readers;
    for (int i = 1; i <= 2; i++) {
        readers.emplace_back([&store, i]() {
            store.read(i);
        });
    }
    
    thread writer2([&store]() {
        store.write(999, 2);
    });
    
    // Wait for first writer to finish
    writer.join();
    cout << "\n[First writer finished - now others can proceed]\n" << endl;
    
    // Now readers and second writer can proceed
    for (auto& t : readers) {
        t.join();
    }
    writer2.join();
    
    cout << "\n✓ All operations completed - writer blocked everyone!" << endl;
}

// Example 5: Mixed scenario - readers and writers
void example5_MixedScenario() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 5: Mixed Scenario" << endl;
    cout << string(70, '=') << endl;
    cout << "Demonstrating the complete behavior:\n" << endl;
    
    SharedDataStore store(0);
    
    // Start multiple readers
    vector<thread> threads;
    
    // 2 readers start first
    threads.emplace_back([&store]() { store.read(1); });
    threads.emplace_back([&store]() { store.read(2); });
    
    this_thread::sleep_for(chrono::milliseconds(100));
    
    // Writer tries to write (waits for readers)
    threads.emplace_back([&store]() { store.write(100, 1); });
    
    this_thread::sleep_for(chrono::milliseconds(50));
    
    // More readers try to read (they can join the first readers)
    threads.emplace_back([&store]() { store.read(3); });
    threads.emplace_back([&store]() { store.read(4); });
    
    // Wait for all
    for (auto& t : threads) {
        t.join();
    }
    
    cout << "\n✓ Mixed scenario completed!" << endl;
}

int main() {
    cout << "\n" << string(70, '=') << endl;
    cout << "SHARED MUTEX: Exclusive vs Shared Lock Examples" << endl;
    cout << string(70, '=') << endl;
    cout << "\nKey Concepts:" << endl;
    cout << "- SHARED LOCK (readers): Multiple threads can hold simultaneously" << endl;
    cout << "- EXCLUSIVE LOCK (writers): Only one thread can hold at a time" << endl;
    cout << "- Readers block writers, but not other readers" << endl;
    cout << "- Writers block everyone (readers and other writers)" << endl;
    
    example1_MultipleReaders();
    example2_MultipleWriters();
    example3_ReadersVsWriter();
    example4_WriterBlocksEveryone();
    example5_MixedScenario();
    
    cout << "\n" << string(70, '=') << endl;
    cout << "All examples completed!" << endl;
    cout << string(70, '=') << endl;
    
    return 0;
}