#include <iostream>
#include <map>
#include <string>
#include <unordered_set>
using namespace std;

class KeyValueTransaction {
    public:
      map<int, string> localStore;
      unordered_set<int> deletedKeys;
      KeyValueTransaction* parentTransaction;
      KeyValueTransaction() {
        parentTransaction = nullptr;
      }
      KeyValueTransaction(KeyValueTransaction* parentTransaction) {
        this->parentTransaction = parentTransaction;
      }
    void set(int key, string value) {
        localStore[key] = value; 
        if(deletedKeys.find(key) != deletedKeys.end()) {
            deletedKeys.erase(key);
        }
    }
    void deleteKey(int key) {
        deletedKeys.insert(key);
        if(localStore.find(key) != localStore.end()) {
            localStore.erase(key);
        }
    }
    string get(int key,  map<int, string> &originStore) {
        if(deletedKeys.find(key) != deletedKeys.end()) {
            return "";
        } else if(localStore.find(key) != localStore.end()) {
            return localStore[key];
        } else if(parentTransaction) {
            return parentTransaction->get(key, originStore);
        } else if(originStore.find(key) != originStore.end()) {
            return originStore[key];
        }
        return "";
    }
    void commit(map<int, string> &originStore) {
        if(parentTransaction) {
            for(auto &kv : localStore) {
                parentTransaction->localStore[kv.first] = kv.second;
            }
            for(auto &key : deletedKeys) {
                parentTransaction->deletedKeys.insert(key);
            }
        } else {
            for(auto &kv : localStore) {
                originStore[kv.first] = kv.second;
            }
            for(auto &key : deletedKeys) {
                originStore.erase(key);
            }
        }
        localStore.clear();
        deletedKeys.clear();
        parentTransaction = nullptr;
    }
};
class KeyValueStore {
    public: 
      map<int, string> store; 
      vector<KeyValueTransaction*> transactions;
    KeyValueStore() {
        transactions.clear();
    }
    void set(int key, string value) {
        if (transactions.size() > 0) {
            transactions.back()->set(key, value);
        } else {
            store[key] = value;
        }
    }
    
    string get(int key) {
        if (transactions.size() > 0) {
            return transactions.back()->get(key, store);
        } else {
            return store[key];
        }
    }

    void deleteKey(int key) {
        if (transactions.size() > 0) {
            transactions.back()->deleteKey(key);
        } else {
            if(store.find(key) != store.end()) {
                store.erase(key);
            }
        }
        return;
    }

    void begin() {
        if(transactions.size() > 0) {
            KeyValueTransaction* currentTransaction = transactions.back();
            transactions.push_back(new KeyValueTransaction(currentTransaction));
        } else {
            transactions.push_back(new KeyValueTransaction());
        }
    }

    void commit() {
        if(transactions.size() > 0) {
            KeyValueTransaction* currentTransaction = transactions.back();
            transactions.pop_back();
            currentTransaction->commit(store);
            delete currentTransaction;
        } 
    }

    void rollback() {
        if (transactions.size() > 0) {
            KeyValueTransaction* currentTransaction = transactions.back();
            transactions.pop_back();
            delete currentTransaction;
        }
        
        
    }
};
int main() {
    KeyValueStore kv;
    kv.set(1, "Hello");
    kv.set(2, "World");
    cout<<kv.get(1)<<endl;
    cout<<kv.get(2)<<endl;


    kv.begin();
    kv.set(1, "Hello1");

    kv.begin();
    kv.set(2, "World1");
    cout<<kv.get(1)<<endl;
    cout<<kv.get(2)<<endl;


    kv.begin();
    kv.set(2, "World2");
    cout<<kv.get(2)<<endl;
    cout<<kv.get(2)<<endl;

    kv.rollback();
    cout<<kv.get(2)<<endl;

    kv.rollback();  
    cout<<kv.get(2)<<endl;

    kv.rollback();
    cout<<kv.get(2)<<endl;

    return 0;
}