#include <iostream>
#include <thread>
#include <mutex>
using namespace std;

/*

unique_lock is a more flexible version of lock_guard.

Just like the normal lock guard, except...
It initialises an exclusive lock
It can be returned from the function without releasing the lock (via move semantics)
It can be released before it is destroyed
You can also use nifty lock methods!


if you defer the lock , you can use nifty lock methods!

// Initialise the lock guard, but don't actually lock yet
std::unique_lock<std::mutex> guard(mutex_1, std::defer_lock);

// Now you can do some of the following!
guard.lock(); // Lock now!
guard.try_lock(); // Won't block if it can't acquire
guard.try_lock_for(); // Only for timed_mutexes
guard.try_lock_until(); // Only for timed_mutexes

*/

/* EXAMPLE 1

void printMessage(string message, mutex &mtx) {
    unique_lock<mutex> lock(mtx);
    cout<<message<<endl;
}
int main() {
    mutex mtx;
    thread t1(printMessage, "Hello", ref(mtx));
    thread t2(printMessage, "World", ref(mtx));
    t1.join();
    t2.join();
    return 0;
}

*/

/*

EXAMPLE 2
 lock initialization is deferred
*/


void printMessage(string message, mutex &mtx) {
    unique_lock<mutex> lock(mtx, defer_lock);
    cout<<"Other work is done"<<endl;
    lock.lock(); // now the lock is acquired
    cout<<"Message is printed"<<endl;
    lock.unlock(); // can also unlock early
    cout<<message<<endl;
    // lock.unlock(); -> no need to unlock, it is automatically unlocked when the lock goes out of scope, can also use if you want to unlock early
}
int main() {
    mutex mtx;
    thread t1(printMessage, "Hello", ref(mtx));
    thread t2(printMessage, "World", ref(mtx));
    t1.join();
    t2.join();
    return 0;
}



