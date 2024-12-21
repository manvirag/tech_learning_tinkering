# Master-Slave DD

Databases are usually read-heavy, having a read-to-write ratio of approximately 90:10.

![Untitled](Master-Slave%20DD%20e22a3ba5ccd5477192b05bcdab051673/Untitled.png)

Or it can be like, we have a single load balancer(consistent hashing) that’s handling about write and reads by automatically writing to master and reading from any of slave.

Master → main/latest data.

Challenges:

1. How are we scaling master? 
    1. Write requests to master can hardly be scaled. One of the only options to scale the writing requests is to increase the compute capacity(CPU and ROM) of the master database.
    2. Master-Master replication. (multi-master replication)
    
2. What if the master is down? 
    
    In this case, we promote one of the slave databases to the master and it starts working as the master.    
    
    How do we do this? 
    
    Using the leader election algorithm, a **voting algorithm**  which is called as **Consensus Algorithms**
    
    ### **Further Reading (Refer to Distributed consensus notion)**
    
    - **The split-brain problem in Consensus Algorithm.**
    - CORS [[Link](https://martinfowler.com/bliki/CQRS.html)]
    
3. How we get to know if some node is failure or not ? 
    1. Heartbeat → Know more about this.

1. What if the slave is down?
    1. So in the scenario of follower failure, we can use a strategy called **Catchup recovery**. In this strategy, the follower (who got disconnected) can connect to the leader again and request all data changes that occur when the follower is disconnected.Every follower maintains a **Log** of data changes received from the Leader on its **Local Disk**

There are 3 types of Replication Strategies, you can choose one of them depending on your use case: **[ Note: In these cases, I haven’t considered any slave-slave communication for read replication ]**

(How does this happen at a low-level implementation? )

**Synchronous Replication: (banking)**

Synchronous replication ensures data consistency by updating replicas simultaneously and waiting for confirmations before responding to the client. (**Strong Consistency)**

Challenges → Write latency (Don’t have any other option, if possible go with semi)

**Asynchronous Replication: (Instagram post, means we use this when read data need to be latest)**

In Asynchronous Replication, once the Master node updates its copy of the data, it immediately completes the operation by responding to the Client. It does not wait for the changes to be propagated to the Replicas, thus minimizing the block for the Client and maximizing the throughput.

Challenges → Failure handling. How do we do this and consistency? 

**A/c to me** → Infinite retry with ordering maintenance ( means block subsequent requests 🤔 ) + alerting (pager duty) + fix live on production.

**Consistency** → yes have the possibility, that we immediately do not get the data from the slave, (CAP). But it is what it is. 

 

![Untitled](Master-Slave%20DD%20e22a3ba5ccd5477192b05bcdab051673/Untitled%201.png)

**Semi-synchronous Replication:** 

In Semi-synchronous Replication, which sits right between the Synchronous and Asynchronous Replication strategies, once the Master node updates its copy of the data, it synchronously replicates the data to a subset of Replicas and asynchronously to others.

![Untitled](Master-Slave%20DD%20e22a3ba5ccd5477192b05bcdab051673/Untitled%202.png)

Challenges → 

How do we do this what algorithm? → maybe no algorithm not sure.

Failure handling → Same as discussed in async

Consistency → Same as discussed in async → It can more less probability of having

**How to do eventual consistency here and conflict resolution? :**