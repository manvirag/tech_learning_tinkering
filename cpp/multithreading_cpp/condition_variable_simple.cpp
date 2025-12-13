#include <iostream>
#include <thread>
#include <mutex>
#include <condition_variable>
using namespace std;


void addMoney(int money, condition_variable& cv, mutex& m, long& balance) {
    std::lock_guard<mutex> lg(m);
    balance += money;
    cout << "Amount Added Current Balance: " << balance << endl;
    cv.notify_all();
}

void withdrawMoney(int money, condition_variable& cv, mutex& m, long& balance) {
    unique_lock<mutex> ul(m);
    cv.wait(ul, [&]() { return (balance != 0) ? true : false; }); // The lambda function checks if balance is non-zero, false -> relase lock and wait
    if (balance >= money) {
        balance -= money;
        cout << "Amount Deducted: " << money << " Current Balance: " << balance << endl;
    }
    else {
        cout << "Amount Can't Be Deducted, Current Balance Is Less Than " << money << endl;
    }
    cout << "Current Balance Is: " << balance << endl;
}

int main() {
    std::condition_variable cv;
    std::mutex m;
    long balance = 0;
    thread t1(withdrawMoney, 500, ref(cv), ref(m), ref(balance));
    this_thread::sleep_for(chrono::seconds(10));
    thread t2(addMoney, 500, ref(cv), ref(m), ref(balance));
    t1.join();
    t2.join();
    return 0;
}