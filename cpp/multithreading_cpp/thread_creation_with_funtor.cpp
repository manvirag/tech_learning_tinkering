#include <iostream>
#include <thread>
using namespace std;

// functor (function object) is a class that overloads the () operator

class PrintSum {
    public:
    void operator()(int a , int b) {
        cout<<a+b<<endl;
    }
};
int main() {

    thread t1(PrintSum(), 1, 2);
    t1.join();   
    cout<<"Main thread ends"<<endl;
    return 0;
}