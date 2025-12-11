#include <iostream>
#include <thread>
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

mutex mtx;
int counter = 0;
void incrementCounter() {
    mtx.lock();
    counter++;
    mtx.unlock();
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