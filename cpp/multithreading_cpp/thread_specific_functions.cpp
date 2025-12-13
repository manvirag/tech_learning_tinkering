#include <iostream>
#include <thread>
#include <chrono>
using namespace std;

/*

Mutex: Mutual Exclusion

RACE CONDITION:
0. Race condition is a situation where two or more threads/process happend to change a common data at the same time.
1. If there is a race condition then we have to protect it and the protected setion is called critical section/region.

MUTEX:
0. Mutex is used to avoid race condition.
1. We use lock(), unlock() on mutex to avoid race condition.

*/

int counter = 0;
void incrementCounter() {
    // These can be used within a thread

    // Get thread ID of thread
    cout<<"Thread ID: "<<this_thread::get_id()<<endl;

    // Give priority to other threads, pause execution
    // std::this_thread::yield();

    // Sleep for some amount of time
    // this_thread::sleep_for(chrono::seconds(1));

    // Sleep until some time
    // chrono::system_clock::time_point time_point = chrono::system_clock::now()
    //                                                 + chrono::seconds(10);
    // this_thread::sleep_until(time_point);
    counter++;
}
int main() {
    thread t1(incrementCounter); 
    thread t2(incrementCounter); 
    t1.join();
    t2.join();
    cout<<"Counter: "<<counter<<endl;
    cout<<"Main thread ends"<<endl;
    return 0;
}