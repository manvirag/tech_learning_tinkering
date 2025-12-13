/*r
reference in thread
NOTES:
0. It is used to pass the variable by reference to the thread.
1. It is used to avoid the copying of the variable to the thread.
*/

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