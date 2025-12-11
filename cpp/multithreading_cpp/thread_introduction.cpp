#include <iostream>
#include <thread>
using namespace std;

void printSum(int a , int b) {
    cout<<a+b<<endl;
}
void printProduct(int a , int b) {
    cout<<a*b<<endl;
}

int main() {
    thread t1(printSum, 1, 2);
    thread t2(printProduct, 3, 4);
    t1.join();
    t2.join();
    cout<<"Main thread ends"<<endl;
    return 0;
}