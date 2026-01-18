/*

Also do wtih concurrency. 

................................................................................................................................

Problem 1 -> similar to kafka

Functional Requirements
Publishers call publish(String topic, String message) to send messages asynchronously to a topic.
Subscribers call subscribe(String topic) to register for a topic and receive messages via polling (getMessages()) or callbacks.
Support unsubscribe(String topic) to stop receiving from a topic.
Each topic maintains a bounded queue (e.g., capacity 1000) for messages; handle overflow by dropping oldest or blocking (discuss).
​
................................................................................................................................


Problem 2 ( sqs)

Design an inmemory pull based queue library
where multiple publishers and consumers can publish/read 
the messages from the shared queue. 
And each message can have optional TTL .

*/



// problem 2 sqs

/*

one queue 
-> multiple publisher -> publish with ttl
-> one message will be red by one consumer 
-> after ttl message will be removed from queue. 
-> order ? -> yes gaurantee. 

we can maitain deque for messages. 
set<{timestmapl, pointer} -> cronjob will delete messages if existin gin queue
if multiple consumer polling add exlcusive lock. 

- one read by any consumer , will be deleted. 

*/


#include<iostream> 
#include<deque>
#include<map>
#include<list>
using namespace std; 

class Message {
    public: 
        int id; 
        int timestamp; // seconds
        int ttl; // seconds
    Message(){}
    Message(int id, int timestamp, int ttl = 1800): id(id), timestamp(timestamp), ttl(ttl) {}
    bool operator<(const Message &s) const{
        if(timestamp+ttl == s.timestamp+s.ttl) {
            return id < s.id;
        }
        return timestamp+ttl < s.timestamp+s.ttl;
    }
    /*
        map take as same value 
        if a < b -> false and b < a -> false and over it. 
        also a < a -> false
    */
    
    
    
};
class Sqs { 
    public: 
        list<Message> msgs; 
        map<Message, list<Message>::iterator> msgPointer;

        void publishMessage(Message msg) {
            msgs.push_back(msg);
            list<Message>::iterator it = msgs.end();
            it--;
            msgPointer[msg] = it;
        }

        Message pollMessage() {
            if(msgs.size()  == 0) {
                return Message(-1,-1,-1);
            }
            list<Message>::iterator it = msgs.begin();
            Message msg = *msgs.begin();
            msgs.erase(it);
            msgPointer.erase(msg);
            return msg;
        }

        void backgrounJob(int currentTime) {
            vector<Message> toDelete;
            for (auto& msg : msgPointer) {
                if (msg.first.timestamp + msg.first.ttl <= currentTime) {
                    toDelete.push_back(msg.first);
                } else {
                    break;
                }
            }
            for (auto& msg : toDelete) {
                msgs.erase(msgPointer[msg]);
                msgPointer.erase(msg);
            }
        }

};

int main() {
    Sqs sqs = Sqs();
    sqs.publishMessage(Message(1,100000,30));
    sqs.publishMessage(Message(2,104000,30));
    sqs.publishMessage(Message(3,130000,30));
    sqs.backgrounJob(100030);
    cout<<sqs.pollMessage().id<<endl;
    cout<<sqs.pollMessage().id<<endl;
    cout<<sqs.pollMessage().id<<endl;
    cout<<sqs.pollMessage().id<<endl;
    return 0; 
}



/*

review: 

bool operator<(const Message &s) const {
    return timestamp + ttl < s.timestamp + s.ttl;
}

Two different messages can have same expiry time
map requires strict weak ordering

If (timestamp + ttl) is equal:

a < b == false
b < a == false

→ map treats them as same key

💥 Result:
    One message overwrites another


#include <iostream>
#include <list>
#include <unordered_map>
#include <queue>
using namespace std;

class Message {
public:
    int id;
    int timestamp;
    int ttl;

    Message(int id, int timestamp, int ttl = 1800)
        : id(id), timestamp(timestamp), ttl(ttl) {}

    int expiry() const {
        return timestamp + ttl;
    }
};

class Sqs {
private:
    // FIFO order
    list<Message> queue;

    // id -> iterator in queue
    unordered_map<int, list<Message>::iterator> idToIt;

    // min-heap by expiry time
    priority_queue<
        pair<int, int>,
        vector<pair<int, int>>,
        greater<pair<int, int>>
    > expiryHeap;

public:
    void publishMessage(int id, int timestamp, int ttl = 1800) {
        Message msg(id, timestamp, ttl);
        queue.push_back(msg);

        auto it = queue.end();
        --it;
        idToIt[id] = it;

        expiryHeap.push({msg.expiry(), id});
    }

    Message pollMessage() {
        if (queue.empty()) {
            return Message(-1, -1, -1);
        }

        auto it = queue.begin();
        Message msg = *it;

        idToIt.erase(msg.id);
        queue.erase(it);

        return msg;
    }

    void backgroundJob(int currentTime) {
        while (!expiryHeap.empty()) {
            auto top = expiryHeap.top();
            int expiryTime = top.first;
            int msgId = top.second;

            if (expiryTime > currentTime) break;

            expiryHeap.pop();

            auto it = idToIt.find(msgId);
            if (it != idToIt.end()) {
                queue.erase(it->second);
                idToIt.erase(it);
            }
        }
    }
};

int main() {
    Sqs sqs;

    sqs.publishMessage(1, 100000, 30);
    sqs.publishMessage(2, 104000, 30);
    sqs.publishMessage(3, 130000, 30);

    sqs.backgroundJob(100030);

    cout << sqs.pollMessage().id << endl; // -1 (expired)
    cout << sqs.pollMessage().id << endl; // 2
    cout << sqs.pollMessage().id << endl; // 3
    cout << sqs.pollMessage().id << endl; // -1

    return 0;
}

*/