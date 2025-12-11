#include <iostream>
#include <thread>
using namespace std;


void printSum(int a , int b) {
    cout<<a+b<<endl;
}
int main() {
    vector<thread> threads;
    for(int i = 0; i < 10; i++) {
        threads.push_back(thread(printSum, 1, 2));
    }
    for(auto &t : threads) {
        t.join();
    }
    cout<<"Main thread ends"<<endl;
    return 0;
}