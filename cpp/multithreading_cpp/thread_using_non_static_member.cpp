#include <iostream>
#include <thread>
using namespace std;

// Using Non static member function as a thread function

class MyClass {
    public:   
    void printSum(int a , int b) {
        cout<<a+b<<endl;
    }
};

int main() {

    MyClass myClass;
    
    // &MyClass::printSum is the address of the printSum function
    // &myClass is the address of the object myClass
    thread t1(&MyClass::printSum, &myClass, 1, 2); 
    t1.join();   
    cout<<"Main thread ends"<<endl;
    return 0;
}