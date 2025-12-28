#pragma once
using namespace std; 

class Counter {
    private:
        static int count; 
    public: 
        static int getCount() {
            count++;
            return count; 
        }
};

int Counter::count = 0;