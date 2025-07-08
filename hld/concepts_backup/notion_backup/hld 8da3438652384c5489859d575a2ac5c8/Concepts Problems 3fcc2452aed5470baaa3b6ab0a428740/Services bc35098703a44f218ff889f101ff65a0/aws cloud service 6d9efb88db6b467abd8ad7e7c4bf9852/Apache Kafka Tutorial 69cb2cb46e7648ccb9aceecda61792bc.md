# Apache Kafka Tutorial

[https://www.youtube.com/playlist?list=PLa7VYi0yPIH2PelhRHoFR5iQgflg-y6JA](https://www.youtube.com/playlist?list=PLa7VYi0yPIH2PelhRHoFR5iQgflg-y6JA)

[https://www.youtube.com/watch?v=j4bqyAMMb7o&list=PLa7VYi0yPIH0KbnJQcMv5N9iW8HkZHztH](https://www.youtube.com/watch?v=j4bqyAMMb7o&list=PLa7VYi0yPIH0KbnJQcMv5N9iW8HkZHztH)

[https://medium.com/inspiredbrilliance/kafka-basics-and-core-concepts-5fd7a68c3193](https://medium.com/inspiredbrilliance/kafka-basics-and-core-concepts-5fd7a68c3193)

Motivation → New York Newspaper use it as a database :D. Since it can have infinite retention .
Is this PubSub or consumer need to poll it like in SQS  ? 🤔 (Yes it has pubsub mechanism)
can we run locally to experiment on this  ? 🤔 (Yet not aws msk but opensource apache down apache kafka, zookeeper set it up etc.)

**Distributed**  mean → we have data x we distributed it to like x/5 in 5 service to save it from server crash.
And if we are putting x in each 5 service its called **replication.

User disk storage . to save message**

## Brief:

*Kafka is a Distributed Streaming Platform or a Distributed Commit Log*

Its complex than other messaging service. Its like the one messaging queue with multiple kinesis, that can be denoted by a topic. Apache Kafka is made by Linkedin and open sourced it and donated to apache foundation. Aws Msk is managed version of this open sourced apache Kafka that is handled by amazon web service.

## Architecture:

![Untitled](Apache%20Kafka%20Tutorial%2069cb2cb46e7648ccb9aceecda61792bc/Untitled.png)

## Components/Terminology:

**Commit Log:**

When you push data to Kafka it takes and appends them to a stream of records, like appending logs in a log file. This stream of data can be “Replayed” or read from any point in time.

**Distributed:**

Kafka works as a cluster of one or more nodes that can live in different Data centres, we can distribute data/ load across different nodes in the Kafka Cluster, and it is inherently scalable, available, and fault-tolerant.

**Streaming Platform:**

Kafka stores data as a stream of continuous records which can be processed in different methods.

**Topic:**

![Untitled](Apache%20Kafka%20Tutorial%2069cb2cb46e7648ccb9aceecda61792bc/Untitled%201.png)

(logical segregation between messages, sort of like having different tables for holding different types of data.)

Analogy → Like in kinesis , we consume all messages its not like categorisation of message like in sns, after consuming message we check we require this or not.
Now assume kafka is like multiple kinesis each kinesis is denoted by topic. So consumer can decide the topic and listen the message only for those topics, Means two topic messages are separated they are not interlinked. 

![Untitled](Apache%20Kafka%20Tutorial%2069cb2cb46e7648ccb9aceecda61792bc/Untitled%202.png)

**Partition:**

(We can increase the partition but can’t decrease , since might get data lost).

the core concept behind Kafka’s scaling capabilities. Let’s say that our system becomes really popular and hence there are millions of log messages per second. So now the node on which XYZ topic is present, is unable to hold all the data that is coming in. We initially solve this by adding more storage to our node i.e. vertical scaling. But as we all know vertical scaling has its limit, once that threshold is reached we need to horizontally scale, which means we need to add more nodes and split the data between the nodes. When we split data of a topic into multiple streams, we call all of those smaller streams the “Partition” of that topic.

![Untitled](Apache%20Kafka%20Tutorial%2069cb2cb46e7648ccb9aceecda61792bc/Untitled%203.png)

This image depicts the **idea of partitions**, where a single **topic has 4 partitions**, and all of them **hold a different set of data**. The blocks you see here are the different messages in that partition. Let’s imagine the topic to be an array, now due to memory constraint we have split the single array into 4 different smaller arrays. And when we write a new message to a topic, the relevant partition is selected and then that message is added at the end of the array.

An **offset** for a message is the **index of the array** for that message. The numbers on the blocks in this picture denote the **Offset,** the first block is at the 0th offset and the last block would on the (n-1)th offset. (Please note that on Kafka it is not going to be an actual array but a symbolic one)

**Producer:**

A producer is the Kafka client that publishes messages to a Kafka topic. Also one of the core responsibilities of the Producer is to decide which partition to send the messages to.
      **Key is not a partition , its like partition key means same key will have same hash and same has will have same partition , mean if we sending data with same key then they will be in order and in a single partition**

1. **No Key specified =>** When no key is specified in the message the producer will randomly decide partition and would try to balance the total number of messages on all partitions.
2. **Key Specified =>** When a key is specified with the message, then the producer uses [Consistent Hashing](https://www.toptal.com/big-data/consistent-hashing) to map the key to a partition.(the same key same hash is generated always), and it minimises the redistribution of keys on a re-hashing scenario like a node add or a node removal to the cluster. 
3. **Partition Specified =>** You can hardcode the destination partition as well.
4. **Custom Partitioning logic =>** We can write some rules depending on which the partition can be decided.

**Consumer:**

A consumer reads messages from partitions, in an ordered fashion. So if 1, 2, 3, 4 was inserted into a topic, the consumer will read it in the same order. Since every message has an offset, every time a consumer reads a message it stores the offset value onto Kafka or Zookeeper, denoting that it is the last message that the consumer read. So in case, a consumer node goes down, it can come back and resume from the last read position. Also if at any point in time a consumer needs to go back in time and read older messages, it can do so by just resetting the offset position.

![Untitled](Apache%20Kafka%20Tutorial%2069cb2cb46e7648ccb9aceecda61792bc/Untitled%204.png)

**Consumer Group**

![Untitled](Apache%20Kafka%20Tutorial%2069cb2cb46e7648ccb9aceecda61792bc/Untitled%205.png)

If user doesn’t mention a consumer group while creating consumer. Then kafka automatically do that . Else create group which is given by user and add those consumer to that group.

So As per my understanding , kafka created this to do parallel processing. And have some rule, like every consumer will have some group. Also No duplication in consumers in same group. So it means 🤔, if we want to use this concept properly. It will be helpful if (for e.g. we have some how did the concurrency of consumer to consume data.) . Also kafka will assume that all consumer group is a consumer that he have to maintain order . So it means if any other consumer of same group reassign then it start from the previous offset not from start. Also in this case we want to maintain order as well. So Below paragraph will help how (Not sure but if we haven’t given group explicitly it might create the consumer in different group and order might not maintain  confirm this)

One partition can not be read by multiple consumers in the same consumer group. This is enabled by the consumer group only, only one consumer in the group gets to read from a single partition.Our topic has 3 partitions, and due to consistent hashing messages with the same key always go to the same partition, so all the messages with “A” as the key will get grouped and the same for B and C. Now as each partition has only one consumer, they get messages in order only. So the consumer will receive A1 before A2 and B1 before B2, and thus the order is maintained

![Untitled](Apache%20Kafka%20Tutorial%2069cb2cb46e7648ccb9aceecda61792bc/Untitled%206.png)

**Consumer in group and partition combination sample example**

So for 3 partitions, you can have a max of 3 consumers, if you had 4 consumers, one consumer will be sitting idle. But for 3 partitions you can have 2 consumers, then one consumer will read from one partition and one consumer will read from two partitions. If one consumer goes down in this case, the last surviving consumer will end up reading from all the three partitions, and when new consumers are added back, again partition would be split between consumers, this is called re-balancing.

![Untitled](Apache%20Kafka%20Tutorial%2069cb2cb46e7648ccb9aceecda61792bc/Untitled%207.png)

![Untitled](Apache%20Kafka%20Tutorial%2069cb2cb46e7648ccb9aceecda61792bc/Untitled%208.png)

**Broker:**

A broker is a single Kafka server

![Untitled](Apache%20Kafka%20Tutorial%2069cb2cb46e7648ccb9aceecda61792bc/Untitled%209.png)

**Kafka Cluster:**

A Kafka cluster is a group of broker nodes working together to provide, scalability, availability, and fault tolerance. One of the brokers in a cluster works as the Controller, which basically assigns partitions to brokers, monitors for broker failure to do certain administrative stuff. ( zookeeper have info about controller and controller is the real man to actually do the work)

Reassigning partitions → let say we want to have partition in particular broker, in this case case controller do assigning-reassigning.

![Untitled](Apache%20Kafka%20Tutorial%2069cb2cb46e7648ccb9aceecda61792bc/Untitled%2010.png)

**Zoo Keeper:**

![Untitled](Apache%20Kafka%20Tutorial%2069cb2cb46e7648ccb9aceecda61792bc/Untitled%2011.png)

Zookeeper is also a software like the apache. which we need to install and run like apache kafka. If you are not using aws msk you can search on net to download library of zookeeper and apache kafka.
Kafka does not function without zookeeper( at least for now, they have plans to deprecate zookeeper in near future). Zookeeper works as the central configuration and consensus management system for Kafka. It tracks the brokers, topics, and partition assignment, leader election, basically all the metadata about the cluster.

**Replication/Replication Factor:**

In a cluster, partitions are replicated on multiple brokers depending on the replication factor of the topic to have ***failover*** capability.What I mean is, for a topic of replication factor 3, each partition of that topic will live onto 3 different brokers. When a partition is replicated onto 3 brokers, one of the brokers will act as the leader for that partition and the rest two will be followers.Data is always written on the leader broker and then replicated to the followers. This way we do not lose data nor availability of the cluster, and if the leader goes down another leader is elected

![Untitled](Apache%20Kafka%20Tutorial%2069cb2cb46e7648ccb9aceecda61792bc/Untitled%2012.png)

**Kafka Transaction:**

![Untitled](Apache%20Kafka%20Tutorial%2069cb2cb46e7648ccb9aceecda61792bc/Untitled%2013.png)

[https://www.youtube.com/watch?v=Ki2D2o9aVl8&ab_channel=Confluent](https://www.youtube.com/watch?v=Ki2D2o9aVl8&ab_channel=Confluent)

## AWS Msk:

[https://docs.aws.amazon.com/msk/latest/developerguide/what-is-msk.html](https://docs.aws.amazon.com/msk/latest/developerguide/what-is-msk.html)

## How to implement it, what things we need to take care:

## Monitoring: