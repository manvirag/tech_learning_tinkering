/*
Functional Requirements
Publishers call publish(String topic, String message) to send messages asynchronously to a topic.

Subscribers call subscribe(String topic) to register for a topic and receive messages via polling (getMessages()) or callbacks.

Support unsubscribe(String topic) to stop receiving from a topic.

Each topic maintains a bounded queue (e.g., capacity 1000) for messages; handle overflow by dropping oldest or blocking (discuss).
​
...

queue 
    -> topic
    -> multiple producer -> append message 
    -> consuerm -> topic:offset

publishser 
consumer    
observer
    -> consumer 
pubsub system 
        pubsubstore   
   publish(publisher, topic, message)
   subscrbie(consumer, topic)
   unsubscrbie(consumer, topic)

pubsubstore
- topic: []message 
- list<obserser> 
- consumer:offset 
- list<publishser> 

...



-> pub-sub 

Write code to create a publisher-consumer model where multiple producers publish messages to a queue, and multiple consumers can consume messages from the queue. System can have multiple such queues.

**Solution**: I created all required entities and had PublisherQueueSubscription and ConsumerQueueSubscription stored as a HashMap/List in a subscription repository. My design followed a synchronous approach where the message was deleted from the queue once all consumers received it by getting he front message. I also proposed an asynchronous solution, where each consumer would have its own local queue to poll messages (similar to how Amazon SQS/SNS works).

**Verdict**: Rejected

Reason:

- Managerial Round: I was told that I could not provide more complex answers.
- LLD Round: I overcomplicated my solution and was not able to write running code at the end due to some syntax issues.
- Design an inmemory pull based queue library where multiple publishers and consumers can publish/read the messages from the shared queue. And each message can have optional TTL .
*/

#include<iostream> 
#include<map>
#include<vector>
#include<string>
#include<thread>
using namespace std; 

class Publisher{
    public: 
        int id ;
};
class Observer{
    public: 
        virtual void notify() = 0;
        virtual int getId() = 0;
};
class Consumer: public Observer{
      int id; 
    public: 
       Consumer(){}
       Consumer(int i): id(i){}
       void notify() override{
         cout<< "I have been notified "<<this->id<<endl;
       }
       int getId() override{
        return this->id; 
       }
};
class PubSubStore {
    public: 
        map<string, vector<string> > topicVsMessages;
        map<string , map<int,Observer*>> subscriber; 
        map<string, map<int, int>> consumerVsOffset; 

};
class PubSub{
    PubSubStore pss; 
    public: 
        PubSub(){}
        PubSub(string topic) {
            vector<string> tempMessageArray;
            this -> pss.topicVsMessages[topic] = tempMessageArray;
        }
        void publish(string topic, string message) {
            if(this->pss.topicVsMessages.find(topic) == this->pss.topicVsMessages.end()){
                cout<<"topic not found first create it"<<endl;
                return ;
            }
            cout<<topic<<" "<<message<<endl;
            this->pss.topicVsMessages[topic].push_back(message);
            return ;
        }
        void subscribe(Observer* o, string topic) {
            if(this->pss.topicVsMessages.find(topic) == this->pss.topicVsMessages.end()){
                cout<<"topic not found first create it"<<endl;
                return ;
            }
            this->pss.subscriber[topic][o->getId()] = o; 
        }
        void unsubscribe(Observer* o, string topic) {
            if(this->pss.topicVsMessages.find(topic) == this->pss.topicVsMessages.end()){
                cout<<"topic not found first create it"<<endl;
                return ;
            }
            this->pss.subscriber[topic].erase(o->getId());
        }
        string poll(Observer *c , string topic) {
            if(this->pss.topicVsMessages.find(topic) == this->pss.topicVsMessages.end()){
                cout<<"topic not found first create it"<<endl;
                return "";
            }
            int currOffset = this->pss.consumerVsOffset[topic][c->getId()]; 
            if(currOffset >= this->pss.topicVsMessages[topic].size()) {
                // cout<<"wrong offset "<<currOffset<<endl;
                return "";
            }
            string msg = this->pss.topicVsMessages[topic][currOffset];
            this->pss.consumerVsOffset[topic][c->getId()]++;
            return msg;
        }
};


int main() {
    PubSub ps = PubSub("topic1");
    
    vector<thread> threads;
    for(int i=0;i<10;i++){
        threads.push_back(thread(&PubSub::publish, &ps, "topic1", to_string(i)));
    }

    Observer * c1 = new Consumer(1);
    ps.subscribe(c1, "topic1");

    auto consuming = [&ps, &c1]() {
        while(true) {
            
            string msg = ps.poll(c1, "topic1");
            if(msg != "") 
            cout<<"polled msg: "<<msg<<endl; 
            
        }    
    };

    thread t1(consuming);

    for(auto &t: threads) {
        t.join();
    }       
    t1.join();  

    return 0;
}