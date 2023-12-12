#include <iostream>
#include <thread>
using namespace std;

// Sad not working in my window some issue in gcc will check later

// A simple function that will be executed by the thread
void threadFunction1() {
    for (int i = 0; i < 5; ++i) {
        cout << "Thread1: " << i << endl;
    }
}

void threadFunction2() {
    for (int i = 0; i < 5; ++i) {
        cout << "Thread2: " << i << endl;
    }
}


int main() {
    // Create a thread and launch it by calling the threadFunction
    thread myThread1(threadFunction1);
    thread myThread2(threadFunction2);

   
    return 0;
}
