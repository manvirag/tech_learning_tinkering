#include <iostream>
#include <thread>
using namespace std;


int main() {

    // auto printSum = [](int a , int b) {
    //     cout<<a+b<<endl;
    // };


    thread t1([](int a , int b) {
        cout<<a+b<<endl;
    }, 1, 2);
    t1.join();   
    cout<<"Main thread ends"<<endl;
    return 0;
}