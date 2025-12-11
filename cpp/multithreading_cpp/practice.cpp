// TOPIC: lock_guard In C++ (std::lock_guard<mutex> lock(m1))
// NOTES:
// 0. It is very light weight wrapper for owning mutex on scoped basis.
// 1. It aquires mutex lock the moment you create the object of lock_guard.
// 2. It automatically removes the lock while goes out of scope.
// 3. You can not explicitly unlock the lock_guard.
// 4. You can not copy lock_guard. .

#include <iostream>
#include <thread>
#include <mutex>
using namespace std;

void updateCounter(int &counter){
    counter++;
    
}
int main(){
    int counter = 0;
    thread t1(updateCounter, ref(counter));
    t1.join();
    cout<<"Counter: "<<counter<<endl;
    cout<<"Main thread ends"<<endl;
    return 0;
}