# Distributed consensus

Note: Not Sure, but here we don’t discuss conflict resolution. For this check each replication technique’s conflict resolution. We will discuss, a solution for nodes to agree on the order of operations or decisions, reducing the likelihood of conflicts

Learn Pending:
1.  Leader selection algorithm or voting algorithms in distributed.
2.  Heartbeat detection algorithm in distributed.
3.  Paxos and raft algorithm in consensus

![Untitled](Distributed%20consensus/Untitled.png)

### **What is Consensus in a Distributed System?**

![Untitled](Distributed%20consensus/Untitled%201.png)

In a distributed system, multiple computers (known as **nodes**) are mutually connected with each other and collaborate with each other through message passing. Now, during computation, they need to agree upon a common value to coordinate among multiple processes. This phenomenon is known as **Distributed Consensus**.

Some end user use cases:

- Several people grabbing for that last seat on the airplane
- Multiple users trying to grab the same username — Note that this is subtly different problem to checking if *a* username is taken which is usually solved using [*bloom filters*](https://en.wikipedia.org/wiki/Bloom_filter)

**Why is Consensus Important in Distributed Systems?**

In distributed systems, nodes must work together to achieve a common goal, such as storing data, processing transactions, or running applications. Achieving consensus ensures that all nodes in the system agree on a particular value or state, such as the contents of a database or the result of a computation. This ensures that the system operates correctly and reliably, even in the presence of faults or failures.

Consensus is also crucial in applications such as blockchain.

### **Challenges in Distributed Consensus**

A distributed system can face mainly two types of failure

1. Crash failure → node isn’t responding. (This is a very common issue in distributed systems and it can be handled easily by simply ignoring the node’s existence.)
2. Byzantine failure → responding but malfunction/ abnormal (Handling this kind of situation is complicated in the distributed system.)

![Untitled](Distributed%20consensus/Untitled%202.png)

[ Now we know what consensus is and what are the challenges, Now how can we solve this? ]

*Non-faulty node means, a node which is not crashed attacked, or malfunctioned.*

**How to Achieve Consensus in a Distributed System**

1. `Agreement` — All non-faulty nodes agree on the same value.

 2.   `Validity` — The decided value must have been proposed by some non-faulty node in the group.  

 3.   `Termination`— All non-faulty nodes eventually decide on some value.

**The Need for Consensus Algorithms:** In distributed systems, multiple nodes work together to maintain data consistency and provide fault tolerance. However, nodes can experience failures, network delays, and partitions, which can lead to data inconsistencies. Consensus algorithms are designed to address these challenges and enable a group of nodes to agree on a single value or sequence of operations.

**Use Cases**: Consensus algorithms find applications in distributed databases, distributed file systems, distributed key-value stores, consensus services, blockchain technology, and other distributed systems requiring coordination and agreement among nodes.

Paxos and Raft are two prominent consensus algorithms used to achieve data consistency and fault tolerance in distributed systems. Both algorithms aim to ensure that a group of nodes in a distributed network agree on a single value or sequence of operations, even in the presence of failures and network partitions. Despite their similar goals, Paxos and Raft have different approaches and trade-offs.

Consensus algorithms in distributed systems are designed to achieve agreement among nodes on a shared value or sequence of operations. Key features and considerations include:

1. **Agreement:** The primary goal is to ensure that all nodes eventually agree on the same decision or value.
2. **Fault Tolerance:** Consensus algorithms are built to withstand node failures, communication delays, and network partitions, ensuring system resilience.
3. **Safety:** They guarantee safety by preventing conflicting decisions, and maintaining data consistency.
4. **Liveness:** Ensures the system can make progress and agree on new values even in the absence of failures.
5. **Quorum:** Many algorithms require a quorum (majority of nodes) to agree on a value, enhancing the robustness of decision-making.
6. **Two-Phase Approach:** Typically involves a two-phase approach – preparation and acceptance – reducing the likelihood of conflicts.
7. **Leader-Based or Leaderless:** Consensus algorithms may be leader-based, with a designated leader, or leaderless, where nodes collaborate for consensus.
8. **Message Exchange:** Nodes communicate through messages to propose values, vote, and share their states.
9. **Log Replication:** Many algorithms incorporate log replication mechanisms to maintain a consistent view of data across nodes.
10. **Consistency and Replication:** Ensures data consistency by replicating data and ensuring agreement on the same data across replicas.
11. **Membership Changes:** Some algorithms support dynamic changes in the cluster's membership, allowing nodes to join or leave gracefully.
12. **Performance Trade-offs:** Consensus algorithms vary in their trade-offs between fault tolerance, performance, and simplicity. Some prioritize simplicity, while others prioritize performance, depending on system requirements and scale.

…………………..…………………..…………………..…………………..…………………..…………………..…………………..

For reference for algorithms [Out of scope as of now]:

[https://www.linkedin.com/pulse/building-consensus-distributed-systems-power-paxos-raft-salik-tariq](https://www.linkedin.com/pulse/building-consensus-distributed-systems-power-paxos-raft-salik-tariq)

**Paxos**: Family of distributed algorithms for consensus, ensuring agreement among nodes on a single value despite failures.

[**Will check later on**]

[https://www.mydistributed.systems/2021/04/paxos.html](https://www.mydistributed.systems/2021/04/paxos.html)

[https://www.linkedin.com/pulse/raft-log-replication-leader-election-saurav-prateek?trk=portfolio_article-card_title](https://www.linkedin.com/pulse/raft-log-replication-leader-election-saurav-prateek?trk=portfolio_article-card_title)