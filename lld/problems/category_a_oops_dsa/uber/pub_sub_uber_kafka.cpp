/*

Also do wtih concurrency. 

................................................................................................................................

Problem 1 -> similar to kafka

Functional Requirements
Publishers call publish(String topic, String message) to send messages asynchronously to a topic.
Subscribers call subscribe(String topic) to register for a topic and receive messages via polling (getMessages()) or callbacks.
Support unsubscribe(String topic) to stop receiving from a topic.
Each topic maintains a bounded queue (e.g., capacity 1000) for messages; handle overflow by dropping oldest or blocking (discuss).




topic -> list of mesage
topic -> consumer
consumer -> offset 
oldestOffset 
​

*/

/*

#include<iostream>
#include<unordered_map>
#include<mutex>
using namespace std ;
class Message {
    public: 
        int id; 
    Message(int id): id(id){}
};
class Consumer{
    public: 
        int id;
    Consumer(){}
    Consumer(int id): id(id) {}
}; 
class TopicStore{ // as of now asume infinite retention
    public: 
        string name;
        vector<Message> messages;
        unordered_map<int,Consumer> consumers ;
        unordered_map<int, int> consumerVsOffset;
        TopicStore(){}
        TopicStore(string name): name(name){}
};
class PubSub{
    public: 
        unordered_map<string,TopicStore*> datastore; 
    ~PubSub() {
        for(auto &topic: datastore){
            delete topic.second;
        }
    }
    PubSub(){}

    bool isTopicExist(string topic) {
        if(this -> datastore.count(topic)) {
            return true;
        }
        return false;
    }
    bool isTopicConsumerExist(string topic, int consumerId) {   
        if(!isTopicExist(topic)){
            return false;
        }
        if(datastore[topic]->consumerVsOffset.find(consumerId) != datastore[topic]->consumerVsOffset.end()){
            return true;
        }
        return false;
    }

    void registerTopic(string topic) {
        if(isTopicExist(topic)) return ;
        datastore[topic] = new TopicStore(topic);
        datastore[topic]->messages = {};
        cout<<"registration done"<<endl;
    }
    void publishMessage(string topic, Message msg) {
        if(!isTopicExist(topic)) {
            cout<<"register topic first"<<endl;
            return ;
        } 
        
        datastore[topic]->messages.push_back(msg);
        cout<<"published done"<<endl;
    }

    void subscribeTopic(string topic, Consumer consumer){
        if(!isTopicExist(topic)) {
            cout<<"register topic first"<<endl;
            return ;
        }
        datastore[topic]->consumers[consumer.id]=consumer;
        datastore[topic]->consumerVsOffset[consumer.id] = -1;

    }

    vector<Message> pollMessage(string topic, Consumer consumer){
        vector<Message> messages;
        if(!isTopicConsumerExist(topic, consumer.id)){
            cout<<"Either topic or consumer not exist"<<endl;
            return messages;
        }

        int offset = datastore[topic]->consumerVsOffset[consumer.id];
        cout<<offset<<endl;
        if(offset + 1 == datastore[topic]->messages.size()) {
            cout<<"wait for new message"<<endl;
            return messages;
        }
        offset++;
        datastore[topic]->consumerVsOffset[consumer.id] = offset;
        messages.push_back(datastore[topic]->messages[offset]);
        return messages;
    }

};
int main() {
    PubSub* ps = new PubSub();
    
    
    // The PubSub class contains a mutex member (line 54). mutex is non-copyable, so the copy constructor is implicitly deleted. On line 134, PubSub ps = PubSub(); triggers copy initialization, which fails.
    // Fix: Use direct initialization instead:
    // int main() {    PubSub ps;  // Direct initialization, not copy    // ... rest of code}
    // Or use a pointer:
    // int main() {    PubSub* ps = new PubSub();    // ... use ps->    delete ps;}
    // Explanation:
    // mutex cannot be copied, so classes containing mutex cannot be copy-constructed
    // PubSub ps = PubSub(); tries to copy-construct ps from a temporary
    // PubSub ps; directly constructs ps without copying
    // Change line 134 from PubSub ps = PubSub(); to PubSub ps;.
    
    ps->registerTopic("updated_user");
    ps->publishMessage("updated_user", Message(1));
    ps->publishMessage("updated_user", Message(2));
    ps->publishMessage("updated_user", Message(3));
    ps->publishMessage("updated_user", Message(4));
    Consumer c = Consumer(1);
    ps->subscribeTopic("updated_user", c);
    vector<Message> msgs = ps->pollMessage("updated_user", c );
    msgs.push_back(ps->pollMessage("updated_user", c)[0]);
    msgs.push_back(ps->pollMessage("updated_user", c)[0]);
    msgs.push_back(ps->pollMessage("updated_user", c)[0]);
    cout<<endl;
    for(auto msg: msgs){
        cout<<msg.id<<endl;
    }
    msgs = ps->pollMessage("update_user_2", c );
    for(auto msg: msgs){
        cout<<msg.id<<endl;
    }


}

*/

/*

review 
- unbounded queue
- not threads
- no batch consumer 
- better to have topic level data store. -> updated


*/


// version 2 with concurrent



#include<iostream>
#include<unordered_map>
#include<mutex>
using namespace std ;
class Message {
    public: 
        int id; 
    Message(int id): id(id){}
};
class Consumer{
    public: 
        int id;
    Consumer(){}
    Consumer(int id): id(id) {}
}; 
class TopicStore{ // as of now asume infinite retention
    public: 
        mutex topicMu;
        string name;
        vector<Message> messages;
        unordered_map<int,Consumer> consumers ;
        unordered_map<int, int> consumerVsOffset;
        TopicStore(){}
        TopicStore(string name): name(name){}
    
};
class PubSub{
    public: 
        mutex mu;
        unordered_map<string,TopicStore*> datastore; 
    ~PubSub() {
        for(auto &topic: datastore){
            delete topic.second;
        }
    }
    PubSub(){}

    bool isTopicExist(string topic) {
        lock_guard<mutex> lock(mu);
        if(this -> datastore.count(topic)) {
            return true;
        }
        return false;
    }
    bool isTopicConsumerExist(string topic, int consumerId) {   
        TopicStore* topicP;
        {
            lock_guard<mutex> lock(mu);
            if(!this -> datastore.count(topic))
                return false;

                topicP = datastore[topic];
        }

        {
          lock_guard<mutex> lock(topicP->topicMu);
          if(topicP->consumerVsOffset.find(consumerId) != topicP->consumerVsOffset.end())
            return true;
        }
        return false;
    }

    void registerTopic(string topic) {
        if(isTopicExist(topic)) return ;
        TopicStore *topicP;
        {
            lock_guard<mutex> lock(mu);
            datastore[topic] = new TopicStore(topic);
            topicP = datastore[topic];
        }
        lock_guard<mutex> lock(topicP->topicMu);
        topicP->messages = {};
        cout<<"registration done"<<endl;
    }
    void publishMessage(string topic, Message msg) {
        if(!isTopicExist(topic)) {
            cout<<"register topic first"<<endl;
            return ;
        } 
        
        datastore[topic]->messages.push_back(msg);
        cout<<"published done"<<endl;
    }

    void subscribeTopic(string topic, Consumer consumer){
        if(!isTopicExist(topic)) {
            cout<<"register topic first"<<endl;
            return ;
        }
        datastore[topic]->consumers[consumer.id]=consumer;
        datastore[topic]->consumerVsOffset[consumer.id] = -1;

    }

    vector<Message> pollMessage(string topic, Consumer consumer){
        vector<Message> messages;
        if(!isTopicConsumerExist(topic, consumer.id)){
            cout<<"Either topic or consumer not exist"<<endl;
            return messages;
        }

        int offset = datastore[topic]->consumerVsOffset[consumer.id];
        cout<<offset<<endl;
        if(offset + 1 == datastore[topic]->messages.size()) {
            cout<<"wait for new message"<<endl;
            return messages;
        }
        offset++;
        datastore[topic]->consumerVsOffset[consumer.id] = offset;
        messages.push_back(datastore[topic]->messages[offset]);
        return messages;
    }

};
int main() {
    PubSub* ps = new PubSub();
    /*
    
    The PubSub class contains a mutex member (line 54). mutex is non-copyable, so the copy constructor is implicitly deleted. On line 134, PubSub ps = PubSub(); triggers copy initialization, which fails.
    Fix: Use direct initialization instead:
    int main() {    PubSub ps;  // Direct initialization, not copy    // ... rest of code}
    Or use a pointer:
    int main() {    PubSub* ps = new PubSub();    // ... use ps->    delete ps;}
    Explanation:
    mutex cannot be copied, so classes containing mutex cannot be copy-constructed
    PubSub ps = PubSub(); tries to copy-construct ps from a temporary
    PubSub ps; directly constructs ps without copying
    Change line 134 from PubSub ps = PubSub(); to PubSub ps;.
    */
    ps->registerTopic("updated_user");
    ps->publishMessage("updated_user", Message(1));
    ps->publishMessage("updated_user", Message(2));
    ps->publishMessage("updated_user", Message(3));
    ps->publishMessage("updated_user", Message(4));
    Consumer c = Consumer(1);
    ps->subscribeTopic("updated_user", c);
    cout<<"fds"<<endl;
    vector<Message> msgs = ps->pollMessage("updated_user", c );
    // msgs.push_back(ps->pollMessage("updated_user", c)[0]);
    // msgs.push_back(ps->pollMessage("updated_user", c)[0]);
    // msgs.push_back(ps->pollMessage("updated_user", c)[0]);
    cout<<endl;
    for(auto msg: msgs){
        cout<<msg.id<<endl;
    }
    msgs = ps->pollMessage("update_user_2", c );
    for(auto msg: msgs){
        cout<<msg.id<<endl;
    }


}