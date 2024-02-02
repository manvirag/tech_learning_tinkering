Design Distributed Message queue:

#### Message queue ( RabbitMQ ) vs Event streaming platforms ( Kafka ):

Event streaming platform are distributed message queue with extra features like have high data retention, repeated consumption of message.

***We will be discussing the message queue which have all these features. Since its not like designing question. Its a learning so will share how these works***


### Content

1. Messaging Models.
   1. Point-to-Point.
   2. Publish-Subscribe.
2. Topic, Partitions and Brokers.
3. Consumer Groups.
4. High level design.
5. Design deep dive.
   1. Data storage
   2. Message data structure.
   3. Batching
6. Flow. Push/Pull Method.
7. Consumer Re-balancing [ Skip ]
8. State storage, metadata storage and zookeeper.
9. Replication and in-sync replica.
10. Data delivery semantics.










Reference:
1. System design alex xu.
