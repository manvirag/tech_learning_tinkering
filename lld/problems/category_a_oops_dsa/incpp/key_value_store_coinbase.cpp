#include<iostream> 
#include<map> 
#include<vector> 
#include<string> 
#include<set> 
using namespace std;
/*

Interview Problem: Two-Key Time-Based Key-Value Store

You are asked to design an in-memory key-value store that supports two levels of keys and time-based queries.
The problem is divided into four parts, where each part builds on top of the previous one.

You should assume all data fits in memory.

Part 1 – Basic Storage and Queries

You need to support storing values using two keys and a timestamp.

Each value is identified by:

(primaryKey, secondaryKey, timestamp) → value

Required operations

set(primaryKey, secondaryKey, value, timestamp)
Stores the given value.

getLatest(primaryKey, secondaryKey)
Returns the most recently stored value for the given keys.
If no value exists, return an empty result.

getAll(primaryKey)
Returns all latest values associated with the given primary key
(one value per secondary key).

Part 2 – Time-Based Retrieval

Extend the system to support historical lookups.

New operation

getAt(primaryKey, secondaryKey, timestamp)
Returns the value stored at the largest timestamp less than or equal to the given timestamp.
If no such value exists, return an empty result.

Part 3 – Delete Operations

Add support for deleting data.

New operations

deleteSecondaryKey(primaryKey, secondaryKey)
Removes all values associated with the given secondary key under the primary key.

deletePrimaryKey(primaryKey)
Removes all data associated with the primary key.

Notes

Deletions affect all future queries.

Queries for deleted keys should return empty results.

Part 4 – Consistency After Deletes

Ensure the system behaves correctly under mixed operations.

Your implementation should correctly handle:

Setting values after a key or subkey has been deleted

Time-based queries after deletes

Multiple deletes and re-insertions

Interleaved operations across different keys

No new APIs are added in this part — correctness is the focus.

General Constraints & Notes

Timestamps are integers

You may assume timestamps for a given (primaryKey, secondaryKey) are monotonically increasing

Return empty results (not errors) when data is missing

The system does not need to be thread-safe

Focus on correctness and clean behavior over performance optimizations

*/

class KeyValue {
    map<string, map<string, set<pair<int, string>>>> kvStore; 
    public: 
    KeyValue() {}
    // assumption all argument are valid;
    void setKey(string pk, string sk, int t, string val) {
        this -> kvStore[pk][sk].insert({-t, val}); 
    }
    string getLatest(string pk, string sk) {
        string res = "";
        if(this -> kvStore[pk][sk].size() > 0) {
            res = this -> kvStore[pk][sk].begin() -> second; 
        }
        return res; 
    }
    vector<string> getAll(string pk) {
        vector<string> sks; 
        for(auto sk: this -> kvStore[pk]) {
                sks.push_back(sk.first);
        }
        return sks; 
    }
    string getAt(string pk, string sk, int t) {
        set<pair<int, string>> res = this -> kvStore[pk][sk];
        for(auto x: res){
            int ts = abs(x.first);
            if(ts <= t) {
                return x.second;
            }
        }
        return "";
    }

    void deleteSecondaryKey(string pk, string sk) {
        this -> kvStore[pk].erase(sk);
    }
    void deletePrimaryKey(string pk) {
        this -> kvStore.erase(pk);
    }
    

};
int main() {
    KeyValue* kv = new KeyValue();
    kv -> setKey("a", "b", 1, "c");
    kv -> setKey("a", "b", 3, "d");
    kv -> setKey("a", "c", 3, "e");
    cout<<kv->getLatest("a", "b")<<endl;
    cout<<kv->getAt("a", "b", 2)<<endl;
    vector<string> rr = kv -> getAll("a");
    for(auto x: rr) {
        cout<<x<<" ";
    }
    cout<<endl;
    return 0;
}