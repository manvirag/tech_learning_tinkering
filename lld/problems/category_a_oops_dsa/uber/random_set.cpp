/*

**Design Question:** Implement a simplified **Randomized Set** data structure

- Functional Requirements:
    - Insert, delete, search — all in average constant time
    - Should support random element access
    - Should handle duplicates if required



    - Proposed a hybrid structure using:
    - HashMap (for O(1) lookups)
    - Array/List (for O(1) random access)
- Implemented insert and getRandom
- Covered:
    - Tradeoffs between space and time
    - Handling edge cases (like deleting the last element)
    - Testing strategy for random behavior

  
    
vector<int>        -> index vs value
map -> int vector<int>  val -> list of index

insert -> pushback 
delete -> swap popback 
search -> map 
duplicate -> list 
*/

#include<iostream>
#include<random>
using namespace std; 


random_device rd; 
mt19937 rng(rd());
class RandomizedSet {
    public: 
        vector<int> values; 
        unordered_map<int,vector<int>> valVsIndexMap; 

    void insertV(int value) {
        values.push_back(value);
        valVsIndexMap[value].push_back(values.size()-1);
    }
    void deleteV(int value) {
        // dependent on interview , current  removing last one -> else can use deque if require
        if(valVsIndexMap.find(value) == valVsIndexMap.end()) {
            return ; 
        }
        int index = valVsIndexMap[value].back(); 
        int lastIndex = values.size()-1;
        if(index != lastIndex) {
            swap(values[index], values[lastIndex]);
            for(auto &ix: valVsIndexMap[values[index]]){
                if(ix == lastIndex){
                    ix = index;
                    break;
                }
            }
        }
        
        values.pop_back();
        valVsIndexMap[value].pop_back();
        if(valVsIndexMap[value].size() == 0) {
            valVsIndexMap.erase(value);
        }

        

        cout<<"rint"<<endl;
        for(auto x: valVsIndexMap) {
            cout<<x.first<<endl;
            for(auto y: valVsIndexMap[x.first]) {
                cout<<y<<" ";
            }
            cout<<endl;
        }
        cout<<endl;
        cout<<endl;
        
    }
    bool searchV(int value) {
        return valVsIndexMap.find(value) != valVsIndexMap.end();
    }
    int getRandomV() {
        if(values.size() == 0) {
            return -1;
        }
        uniform_int_distribution<int> ig(0, values.size()-1);
        int randomIx = ig(rng);

        return values[randomIx];
    }
};
int main() {
    RandomizedSet rs = RandomizedSet();
    rs.insertV(2);
    rs.insertV(24);
    rs.insertV(4);
    rs.insertV(3);
    rs.insertV(22);
    rs.insertV(3);
    rs.insertV(4);

    rs.deleteV(3); 

    cout<<rs.getRandomV()<<endl;
    return 0;
}



/*
deletecase issue -> both index is last one -> update now.
use -> unordered_set<int> -> in duplicate case instead of vector to improve
*/