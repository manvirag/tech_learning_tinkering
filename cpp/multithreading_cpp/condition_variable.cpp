#include <iostream>
#include <thread>
#include <mutex>
#include <condition_variable>
#include <queue>
#include <vector>
#include <chrono>
#include <random>
using namespace std;

/*

Condition Variables:


Sometimes you need to do some nice signal/event handling.

It's possible to do it using a global variable that you constantly lock threads for to check, but it's far more efficient to use condition variables.


IMPORTANT:
A condition variable allows you to wait for some condition to be true before continuing thread execution. During this time, any locks that were passed to the waiting function are released until the condition is fulfilled. Following which, the lock is reacquired.



Example Flow

Thread acquires lock
Check if condition is false
If false, call wait(), which releases the lock and blocks the thread until the condition is fulfilled
If a condition is fulfilled, the condition variable must be notified before it can check
Once the condition check succeeds, thread reacquires lock and continues execution


Let's try it out!

Condition variables use unique_locks, so we'll use that.



#include <condition_variable>

// Init
std::condition_variable condition_var;
std::mutex mutex;
bool condition(false);

// Acquire lock
std::unique_lock<std::mutex> guard(mutex);

// Avoid spurious wakeups and 
// ensure wait is only called when the condition has not been fulfilled
while (!condition)
{
  condition_var.wait(guard);
}

// Now in some other thread
{
  // Acquire lock
  std::unique_lock<std::mutex> guard(mutex);

  // We can set the condition to true
  condition = true;

  // And notify one blocked thread by the condition variable that it's ok to wake up, randomly one of the blocked threads will be woken up
  // (In this case we only have one)
  condition_var.notify_one();

  // If we want to notify all of them instead...
  condition_var.notify_all();
    
  // If we didn't surround the threads with the while (!condition) loop,
  // Notifying the threads will cause the wait to return. So there's no condition check.
  // But this is dangerous since random wakeups can occur without notifications!
}
*/



// ============================================================================
// REAL-WORLD EXAMPLE: Producer-Consumer Pattern with Bounded Buffer
// ============================================================================
// This is a classic real-world scenario where condition variables are essential:
// - Producers create items and add them to a shared buffer
// - Consumers remove items from the buffer
// - Buffer has limited capacity (bounded)
// - Producers wait when buffer is full
// - Consumers wait when buffer is empty
// ============================================================================

class BoundedBuffer {
private:
    queue<int> buffer;
    size_t maxSize;
    mutex mtx;
    condition_variable notFull;   // Signaled when buffer is not full
    condition_variable notEmpty;  // Signaled when buffer is not empty
    
public:
    BoundedBuffer(size_t size) : maxSize(size) {}
    
    // Producer: Add item to buffer
    void produce(int item, int producerId) {
        unique_lock<mutex> lock(mtx);
        
        // Wait while buffer is full (predicate prevents spurious wakeups)
        notFull.wait(lock, [this]() { 
            return buffer.size() < maxSize; 
        });
        
        // Buffer has space, add item
        buffer.push(item);
        cout << "Producer " << producerId << ": Produced item " << item 
             << " (Buffer size: " << buffer.size() << "/" << maxSize << ")" << endl;
        
        // Notify one waiting consumer that buffer is not empty
        notEmpty.notify_one();
    }
    
    // Consumer: Remove item from buffer
    int consume(int consumerId) {
        unique_lock<mutex> lock(mtx);
        
        // Wait while buffer is empty (predicate prevents spurious wakeups)
        notEmpty.wait(lock, [this]() { 
            return !buffer.empty(); 
        });
        
        // Buffer has items, remove one
        int item = buffer.front();
        buffer.pop();
        cout << "Consumer " << consumerId << ": Consumed item " << item 
             << " (Buffer size: " << buffer.size() << "/" << maxSize << ")" << endl;
        
        // Notify one waiting producer that buffer is not full
        notFull.notify_one();
        
        return item;
    }
    
    size_t getSize() {
        lock_guard<mutex> lock(mtx);
        return buffer.size();
    }
};

// ============================================================================
// Example 1: Basic Producer-Consumer with Single Producer and Consumer
// ============================================================================
void example1_BasicProducerConsumer() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 1: Basic Producer-Consumer (1 Producer, 1 Consumer)" << endl;
    cout << string(70, '=') << endl;
    cout << "Demonstrates basic condition variable usage with wait() and notify_one()\n" << endl;
    
    BoundedBuffer buffer(5);  // Buffer can hold 5 items
    
    thread producer([&buffer]() {
        for (int i = 1; i <= 10; i++) {
            buffer.produce(i, 1);
            this_thread::sleep_for(chrono::milliseconds(100));
        }
        cout << "\nProducer 1: Finished producing all items\n" << endl;
    });
    
    thread consumer([&buffer]() {
        for (int i = 1; i <= 10; i++) {
            buffer.consume(1);
            this_thread::sleep_for(chrono::milliseconds(150));
        }
        cout << "\nConsumer 1: Finished consuming all items\n" << endl;
    });
    
    producer.join();
    consumer.join();
    
    cout << "✓ Basic producer-consumer completed successfully\n" << endl;
}

// ============================================================================
// Example 2: Multiple Producers and Consumers
// ============================================================================
void example2_MultipleProducersConsumers() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 2: Multiple Producers and Consumers" << endl;
    cout << string(70, '=') << endl;
    cout << "Demonstrates condition variables with multiple threads\n";
    cout << "Multiple producers can produce simultaneously when buffer has space\n";
    cout << "Multiple consumers can consume simultaneously when buffer has items\n" << endl;
    
    BoundedBuffer buffer(3);  // Small buffer to see blocking behavior
    
    vector<thread> producers;
    vector<thread> consumers;
    
    // Create 3 producers
    for (int i = 1; i <= 3; i++) {
        producers.emplace_back([&buffer, i]() {
            for (int j = 1; j <= 5; j++) {
                int item = i * 100 + j;
                buffer.produce(item, i);
                this_thread::sleep_for(chrono::milliseconds(50));
            }
            cout << "Producer " << i << ": Finished\n" << endl;
        });
    }
    
    // Create 2 consumers
    for (int i = 1; i <= 2; i++) {
        consumers.emplace_back([&buffer, i]() {
            for (int j = 1; j <= 7; j++) {  // Each consumer consumes 7 items
                buffer.consume(i);
                this_thread::sleep_for(chrono::milliseconds(80));
            }
            cout << "Consumer " << i << ": Finished\n" << endl;
        });
    }
    
    // Wait for all producers
    for (auto& t : producers) {
        t.join();
    }
    
    // Wait for all consumers
    for (auto& t : consumers) {
        t.join();
    }
    
    cout << "✓ Multiple producers-consumers completed successfully\n" << endl;
}

// ============================================================================
// Example 3: Task Queue (Real-world application)
// ============================================================================
// This simulates a real-world task queue where workers process tasks
// ============================================================================
class TaskQueue {
private:
    queue<string> tasks;
    mutex mtx;
    condition_variable hasTasks;
    bool shutdown = false;
    
public:
    void addTask(const string& task) {
        lock_guard<mutex> lock(mtx);
        tasks.push(task);
        cout << "[TaskQueue] Added task: " << task << " (Queue size: " << tasks.size() << ")" << endl;
        hasTasks.notify_one();  // Notify one waiting worker
    }
    
    string getTask(int workerId) {
        unique_lock<mutex> lock(mtx);
        
        // Wait until there's a task OR shutdown is requested
        hasTasks.wait(lock, [this]() { 
            return !tasks.empty() || shutdown; 
        });
        
        if (shutdown && tasks.empty()) {
            return "";  // Signal to stop
        }
        
        string task = tasks.front();
        tasks.pop();
        cout << "[Worker " << workerId << "] Processing task: " << task 
             << " (Remaining: " << tasks.size() << ")" << endl;
        
        return task;
    }
    
    void shutdownQueue() {
        lock_guard<mutex> lock(mtx);
        shutdown = true;
        hasTasks.notify_all();  // Notify all waiting workers
    }
    
    size_t getSize() {
        lock_guard<mutex> lock(mtx);
        return tasks.size();
    }
};

void example3_TaskQueue() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 3: Task Queue (Real-world Application)" << endl;
    cout << string(70, '=') << endl;
    cout << "Simulates a task queue where workers process tasks\n";
    cout << "Workers wait when queue is empty, wake up when tasks arrive\n" << endl;
    
    TaskQueue taskQueue;
    
    // Create 3 worker threads
    vector<thread> workers;
    for (int i = 1; i <= 3; i++) {
        workers.emplace_back([&taskQueue, i]() {
            int taskCount = 0;
            while (true) {
                string task = taskQueue.getTask(i);
                if (task.empty()) {
                    cout << "[Worker " << i << "] Shutting down (processed " 
                         << taskCount << " tasks)" << endl;
                    break;
                }
                
                // Simulate task processing
                this_thread::sleep_for(chrono::milliseconds(200));
                taskCount++;
            }
        });
    }
    
    // Main thread adds tasks
    this_thread::sleep_for(chrono::milliseconds(100));
    
    vector<string> taskList = {
        "Process payment",
        "Send email",
        "Generate report",
        "Update database",
        "Backup data",
        "Clean cache",
        "Validate input",
        "Compress file",
        "Deploy service",
        "Monitor system"
    };
    
    for (const auto& task : taskList) {
        taskQueue.addTask(task);
        this_thread::sleep_for(chrono::milliseconds(150));
    }
    
    // Wait a bit for tasks to be processed
    this_thread::sleep_for(chrono::seconds(1));
    
    // Shutdown the queue
    cout << "\n[Main] Shutting down task queue...\n" << endl;
    taskQueue.shutdownQueue();
    
    // Wait for all workers to finish
    for (auto& t : workers) {
        t.join();
    }
    
    cout << "\n✓ Task queue example completed successfully\n" << endl;
}

// ============================================================================
// Example 4: Condition Variable with notify_all()
// ============================================================================
void example4_NotifyAll() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 4: Using notify_all() - Event Signaling" << endl;
    cout << string(70, '=') << endl;
    cout << "Demonstrates notify_all() to wake up multiple waiting threads\n" << endl;
    
    mutex mtx;
    condition_variable cv;
    bool ready = false;
    int eventCount = 0;
    
    // Multiple threads waiting for an event
    vector<thread> waiters;
    for (int i = 1; i <= 5; i++) {
        waiters.emplace_back([&mtx, &cv, &ready, &eventCount, i]() {
            unique_lock<mutex> lock(mtx);
            
            // Wait for the event
            cv.wait(lock, [&ready]() { return ready; });
            
            cout << "Thread " << i << ": Event received! (Event count: " 
                 << eventCount << ")" << endl;
        });
    }
    
    // Give threads time to start and wait
    this_thread::sleep_for(chrono::milliseconds(200));
    
    // Signal the event
    {
        lock_guard<mutex> lock(mtx);
        ready = true;
        eventCount = 1;
        cout << "\n[Main] Signaling event to all waiting threads...\n" << endl;
        cv.notify_all();  // Wake up ALL waiting threads
    }
    
    // Wait for all threads
    for (auto& t : waiters) {
        t.join();
    }
    
    cout << "\n✓ notify_all() example completed - all threads were woken up\n" << endl;
}

int main(){
    cout << "\n" << string(70, '=') << endl;
    cout << "CONDITION VARIABLES - REAL-WORLD EXAMPLES" << endl;
    cout << string(70, '=') << endl;
    cout << "\nKey Concepts:" << endl;
    cout << "- wait(lock, predicate): Waits until condition is true (prevents spurious wakeups)" << endl;
    cout << "- notify_one(): Wakes up one waiting thread" << endl;
    cout << "- notify_all(): Wakes up all waiting threads" << endl;
    cout << "- Always use predicate with wait() to avoid spurious wakeups" << endl;
    cout << "- Condition variables work with unique_lock (not lock_guard)" << endl;
    
    example1_BasicProducerConsumer();
    example2_MultipleProducersConsumers();
    example3_TaskQueue();
    example4_NotifyAll();
    
    cout << "\n" << string(70, '=') << endl;
    cout << "SUMMARY:" << endl;
    cout << string(70, '=') << endl;
    cout << "Condition variables are essential for:" << endl;
    cout << "1. Producer-Consumer patterns (bounded buffers)" << endl;
    cout << "2. Task queues and worker pools" << endl;
    cout << "3. Event signaling between threads" << endl;
    cout << "4. Any scenario where threads need to wait for conditions" << endl;
    cout << "\nBest Practices:" << endl;
    cout << "- Always use wait(lock, predicate) to prevent spurious wakeups" << endl;
    cout << "- Use notify_one() when only one thread needs to wake up" << endl;
    cout << "- Use notify_all() when multiple threads need to wake up" << endl;
    cout << "- Lock must be held when calling notify_*() (though not required)" << endl;
    cout << string(70, '=') << endl;
    
    return 0;
}