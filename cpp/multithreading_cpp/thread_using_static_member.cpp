#include <iostream>
#include <thread>
using namespace std;

// Using static member function as a thread function

class MyClass {
    public:   
    static void printSum(int a , int b) {
        cout<<a+b<<endl;
    }
};

int main() {

    // &MyClass::printSum is the address of the printSum function
    thread t1(&MyClass::printSum, 1, 2); 
    t1.join();   
    cout<<"Main thread ends"<<endl;
    return 0;
}