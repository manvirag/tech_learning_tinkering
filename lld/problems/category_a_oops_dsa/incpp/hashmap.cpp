#include <iostream>
#include <map>
#include <vector>
#include <set>
#include <list>
using namespace std;
// assuming key -> not exist -> put 
class HashMap {

    public:
        int bucketSize;
        vector<list<pair<int,int>>> buckets;
        HashMap(int bucketSize) {
            this->bucketSize = bucketSize;
            buckets.resize(bucketSize);
        }
        void put(int key, int value) {
            int bucketIndex = key % bucketSize;
            buckets[bucketIndex].push_back(make_pair(key, value));
        }
        int get(int key) {
            int bucketIndex = key % bucketSize;
            for(auto &pair : buckets[bucketIndex]) {
                if(pair.first == key) {
                    return pair.second;
                }
            }
            return -1;
        }
        void remove(int key) {
            int bucketIndex = key % bucketSize;
            list<pair<int,int>>::iterator pointer = buckets[bucketIndex].end();
            for(auto it = buckets[bucketIndex].begin(); it != buckets[bucketIndex].end(); it++) {
                if(it->first == key) {
                    pointer = it;
                    break;
                }
            }
            buckets[bucketIndex].erase(pointer);
        }
        
};
int main() {
    HashMap hashMap(10);
    hashMap.put(1, 1);
    hashMap.put(11, 11);
    cout<<hashMap.get(1)<<endl;
    hashMap.remove(1);
    cout<<hashMap.get(1)<<endl;
    cout<<hashMap.get(11)<<endl;
    return 0;
}