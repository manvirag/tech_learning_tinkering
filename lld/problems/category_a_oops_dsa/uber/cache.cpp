/*

https://leetcode.com/discuss/post/2508414/sde-2-uber-august-2022-rejection-by-anon-wmhv/

You are designing a cache system for uber with following requirements

1. Adding a key, value pair to cache.
2. Removing key, value pair from cache.
3. Getting value from cache provided a key
4. Getting random key value pair from cache
Optimizied TC and SC expected for each operation.



map -> key -> (value,index)

array -> list of keys 

add -> pushback of arry -> value, array- 1
remove -> swap with last end and pop back from array
get -> simple from map 
random -> index -> key -> value 

- LLD - design a data structure like hashmap which has get, set and delete functions for key, values like a hashmap. additionally it has a getRandom function on it which returns a random key,value paid from the hashmap. I kept the hashmap's keys in an array and generated a random number less than the length of the array and returned the key at that index and the corresponding value. Extension was also to have concurrent operations on this data structure to be used by multiple threads at the same time.


capacity ? -> adding -> yes 

FR: 

You are designing a cache system for uber with following requirements

1. Adding a key, value pair to cache.
2. Removing key, value pair from cache.
3. Getting value from cache provided a key
4. Getting random key value pair from cache

*/
/*

#include<iostream> 
#include<mutex>
#include<shared_mutex>
#include<random>
#include<thread>
using namespace std; 


random_device dev;
mt19937 rng(dev());

class Cache {
  public:
    int capacity;
    mutex mu; 
    shared_mutex readmu;
    unordered_map<int, pair<int,int>> keyV; 
    vector<int> keyList; 
    Cache(int n) {
        this -> capacity = n;
    }

    void add(int key, int value) {
        lock_guard<mutex> lock(this->mu);
        if(this->keyV.find(key) != this->keyV.end()) {
            pair<int , int > valIndex = this -> keyV[key];
            valIndex.first = value;
            this -> keyV[key] = valIndex;
        } else if(keyList.size() == capacity) {
            cout<<key<<" filled"<<endl;
            return ; 
        } else {
            this -> keyList.push_back(key);
            this -> keyV[key] = {value, keyList.size()-1};
        }
    }
    void remove(int key) {
        lock_guard<mutex> lock(this->mu);
        if(this->keyV.find(key) == this->keyV.end()) {
            return ; 
        }
        int index = this -> keyV[key].second; 
        swap(this->keyList[index], this->keyList[this->keyList.size()-1]);
        this->keyList.pop_back();
        pair<int,int> value = this -> keyV[this->keyList[index]];
        value.second = index;
        this->keyV[this->keyList[index]]  = value;
        this->keyV.erase(key);
    }

    int getValue(int key) {
        shared_lock<shared_mutex> lock(this->readmu);
        if(this->keyV.find(key) == this->keyV.end()) {
            cout<<"doesn't exist"<<endl;
            return -1; 
        }
        return this -> keyV[key].first; 
    }

    int getRandomKey() {
        shared_lock<shared_mutex> lock(this->readmu);
        int sz = this->keyList.size();
        uniform_int_distribution<int> dis(0, this->keyList.size()-1);  
        int ix = dis(rng);
        return this->keyList[ix];
    }
    
};

int main() {
    Cache cc(3); 
    vector<thread> ths; 
    
    ths.push_back(thread([&](int key, int val){
        cc.add(key,val);    
    }, 1,2));
    
    ths.push_back(thread([&](int key, int val){
        cc.add(key,val);    
    }, 2,3));

    
    ths.push_back(thread([&](int key, int val){
        cc.add(key,val);    
    }, 1,5));
    
    
    ths.push_back(thread([&](int key, int val){
        cc.add(key,val);    
    }, 4,5));
    // cc.add(2,3);
    // cc.add(3,4);
    // cc.add(1,5);
    // cc.add(4,5);
    // cout<<cc.getValue(1)<<endl;
    // cout<<cc.getValue(4)<<endl;

    
    ths.push_back(thread([&](){
        cout<<cc.getValue(4)<<endl;;    
    }));

    // cout<<cc.getValue(11)<<endl;

    // cout<<cc.getRandomKey()<<endl;

    
    ths.push_back(thread([&](){
        cout<<cc.getRandomKey()<<endl;
    }));


    
    ths.push_back(thread([&](){
        cc.remove(2);
    }));

    

    // cc.remove(2);
    // cout<<cc.getValue(2)<<endl;

    
    ths.push_back(thread([&](){
        cout<<cc.getValue(2)<<endl;
    }));

    for(thread &t: ths) {
        if(t.joinable()) t.join();
    }
    return 0;
}



*/

#include <iostream>
#include <mutex>
#include <shared_mutex>
#include <random>
#include <thread>
#include <unordered_map>
#include <vector>

using namespace std;

// Thread-safe RNG
// after chat gpt revsiion
thread_local mt19937 rng(random_device{}());

class Cache {
public:
    int capacity;
    unordered_map<int, pair<int,int>> keyV; // key -> {value, index}
    vector<int> keyList;

private:
    shared_mutex mu;

public:
    Cache(int n) : capacity(n) {}

    void add(int key, int value) {
        unique_lock<shared_mutex> lock(mu);

        auto it = keyV.find(key);
        if (it != keyV.end()) {
            it->second.first = value;
            return;
        }

        if (keyList.size() == capacity) {
            cout << key << " filled" << endl;
            return;
        }

        keyList.push_back(key);
        keyV[key] = {value, (int)keyList.size() - 1};
    }

    void remove(int key) {
        unique_lock<shared_mutex> lock(mu);

        auto it = keyV.find(key);
        if (it == keyV.end()) return;

        int index = it->second.second;
        int lastKey = keyList.back();

        swap(keyList[index], keyList.back());
        keyV[lastKey].second = index;

        keyList.pop_back();
        keyV.erase(it);
    }

    int getValue(int key) {
        shared_lock<shared_mutex> lock(mu);

        auto it = keyV.find(key);
        if (it == keyV.end()) {
            cout << "doesn't exist" << endl;
            return -1;
        }

        return it->second.first;
    }

    int getRandomKey() {
        shared_lock<shared_mutex> lock(mu);

        if (keyList.empty()) {
            cout << "cache empty" << endl;
            return -1;
        }

        uniform_int_distribution<int> dist(0, keyList.size() - 1);
        return keyList[dist(rng)];
    }
};


int main() {
    Cache cc(3); 
    vector<thread> ths; 
    
    ths.push_back(thread([&](int key, int val){
        cc.add(key,val);    
    }, 1,2));
    
    ths.push_back(thread([&](int key, int val){
        cc.add(key,val);    
    }, 2,3));

    
    ths.push_back(thread([&](int key, int val){
        cc.add(key,val);    
    }, 1,5));
    
    
    ths.push_back(thread([&](int key, int val){
        cc.add(key,val);    
    }, 4,5));
    // cc.add(2,3);
    // cc.add(3,4);
    // cc.add(1,5);
    // cc.add(4,5);
    // cout<<cc.getValue(1)<<endl;
    // cout<<cc.getValue(4)<<endl;

    
    ths.push_back(thread([&](){
        cout<<cc.getValue(4)<<endl;;    
    }));

    // cout<<cc.getValue(11)<<endl;

    // cout<<cc.getRandomKey()<<endl;

    
    ths.push_back(thread([&](){
        cout<<cc.getRandomKey()<<endl;
    }));


    
    ths.push_back(thread([&](){
        cc.remove(2);
    }));

    

    // cc.remove(2);
    // cout<<cc.getValue(2)<<endl;

    
    ths.push_back(thread([&](){
        cout<<cc.getValue(2)<<endl;
    }));

    for(thread &t: ths) {
        if(t.joinable()) t.join();
    }
    return 0;
}
