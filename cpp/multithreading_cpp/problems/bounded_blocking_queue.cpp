#include <queue>
#include <mutex>
#include <condition_variable>
#include <iostream>
#include <thread>

using namespace std;
// https://algo.monster/liteproblems/1188

class BoundedBlockingQueue {
public:
    // Old-style constructor
    BoundedBlockingQueue(int capacity) {
        this->cap = capacity;
        // mutex and condition_variable do not need explicit init
    }

    void enqueue(int element) {
        unique_lock<mutex> lock(mx);

        // Wait until queue has space
        cv.wait(lock, [&] {
            return q.size() < cap;
        });

        q.push(element);
        cout << "Enqueued: " << element << endl;
        // Notify a waiting consumer
        cv.notify_all();
    }

    int dequeue() {
        unique_lock<mutex> lock(mx);

        // Wait until queue has an element
        cv.wait(lock, [&] {
            return !q.empty();
        });

        int val = q.front();
        cout << "Dequeued: " << val << endl;
        q.pop();

        // Notify a waiting producer
        cv.notify_all();
        return val;
    }

    int size() {
        lock_guard<mutex> lock(mx);
        return q.size();
    }

private:
    queue<int> q;
    int cap;

    mutex mx;
    condition_variable cv;
    
};

int main() {
    BoundedBlockingQueue queue(10);
    thread t1([&]() {
        for (int i = 0; i < 10; i++) {
            queue.enqueue(i);
        }
    });
    thread t2([&]() {
        for (int i = 0; i < 10; i++) {
            queue.dequeue();
        }
    });
    t1.join();
    t2.join();
    return 0;
}