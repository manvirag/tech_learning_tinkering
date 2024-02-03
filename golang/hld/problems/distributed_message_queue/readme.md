Design Distributed Message queue:

#### Message queue ( RabbitMQ ) vs Event streaming platforms ( Kafka ):

Event streaming platform are distributed message queue with extra features like have high data retention, repeated consumption of message.

***We will be discussing the message queue which have all these features. Since its not like designing question. Its a learning so will share how these works***


### Content

1. Messaging Models. ✅
   1. Point-to-Point. ✅
   2. Publish-Subscribe. ✅
2. Topic, Partitions and Brokers. ✅
3. Consumer Groups. ✅
4. High level design. ✅
5. Design deep dive. 
   1. Data storage
   2. Message data structure.
   3. Batching 
   4. Flow. Push/Pull Method. 
   5. Consumer Re-balancing [ Skip ]
   6. State storage, metadata storage and zookeeper. 
   7. Replication and in-sync replica. 
   8. Data delivery semantics.

#### Note we guarantee the order in partition level of a topic.

#### Message Models

![alt_text](./images/img.png)

1. Point-to-Point: One message consumed by only one consumer. [ SQS ].
2. Publish-Subscribe: Consumer subscribe for a particular topic and consume message from it. But point-to-point can also be simulated by ***consumer group***.

#### Topic , Partition, Brokers

![alt_text](./images/img_1.png)

1. Topic: You can assume its like one queue, and we have separate queue according to topic name. So that consumer no need to listen events from other domain.
2. Brokers: These are the service which maintaining and processing your data.
3. Partition: For scaling topic and fault tolerance, we distribute the message in multiple brokers in multiple partition. So if we have lots of events in particular topic we can scale it. Partitions have different messages for a particular topic.

#### Consumer Groups

![alt_text](./images/img_2.png)

As said messaging queue allow one message to only one consumer. [ we have both feature this one as well as multiple ].
Multiple is easy , many consumer can subscribe same topic that's it.

But How to implement this ? 

That's where consumer groups come into picture.

***As name , consumer groups are the collection of consumers, one consumer can only be in one group. Messages in a partition will be consumed by only one consumer in a group but consumer can consume from multiple partition of different topic or same topic. This is like point to point model.***


#### High level design

![alt_text](./images/img_3.png)

Let's discuss about storage and coordination service in brief.

1. Data Storage
   1. State storage: keep the configuration or state of consumers
   2. Metadata storage: keep the configuration and properties of topic.
   3. Data storage: use to retain the messages

2. Coordination service
   1. Service discovery: which broker are alive.
   2. Leader Selection: Select one active controller among the brokers which will handle the partition things.
   3. Apache Zookeeper: commonly use to elect a controller.



#### Deep-dive High level design










Reference:
1. System design alex xu.
