# Database distribution (Sharding)

In case of replication, we used to copy the database. That has its advantages. Here we are going to discuss the distribution of databases to different. so that as the database grows, it doesn’t cause any problems.

Each type of database can have the same replication. But each might have a different strategy in distribution ( How to put a particular database in different nodes/servers)

**Database sharding(concept)** is the process of splitting up a database across multiple machines to improve the scalability of an application. It involves breaking up one’s data into two or smaller chunks, called logical shards. The logical shards are then distributed across separate database nodes, referred to as physical shards, which can hold multiple logical shards.

## **Types of Sharding Architectures**

**Range Based Sharding**

- In this method, we partition the data based on the ranges of the key.
- It is very easy to implement
- Data may not be evenly distributed across shards.
- In the example below there are 5 tuples and 3 shards.

[https://miro.medium.com/v2/resize:fit:875/0*onSQGqvYMkwiye4P](https://miro.medium.com/v2/resize:fit:875/0*onSQGqvYMkwiye4P)

**Key-Based Sharding / Hash-Based Sharding**

- In this method, we generate a hash value of the key (Here key is one of the attributes of the data). This hash value determines the shard we will use to store the data.
- Using the simple hash function to distribute data can cause skewed distribution. To overcome this we can use Consistent Hashing.
- In the example below there are 6 tuples and 3 shards. We have used a simple hash function `h(x) = x%3`

[https://miro.medium.com/v2/resize:fit:875/0*YBNoolHCdtiQCVlW](https://miro.medium.com/v2/resize:fit:875/0*YBNoolHCdtiQCVlW)

**Directory-Based Sharding (Mindtickle)**

- In this method, we create a lookup table that uses a shared key to check which shard holds which data. The lookup maps each key to the shard.
- It is more flexible than range and key-based sharding.
- The lookup table is a single point of failure.

[https://miro.medium.com/v2/resize:fit:875/0*onmgdDpycjJCZ0NV](https://miro.medium.com/v2/resize:fit:875/0*onmgdDpycjJCZ0NV)

Hierarchal Sharding

![Untitled](Database%20distribution%20(Sharding)/Untitled.png)

## **Difference between Horizontal Partitioning and Sharding**

- In Horizontal Partitioning, we split the table into multiple tables in the same database instance whereas in sharding we split the table into multiple tables across multiple database instances.
- In Horizontal partitioning, we use the same database instance so the names of the partitioned tables have to be different. In sharding, since the tables are stored in different database instances, table names can be the same.

Next step → Check each database and see how we do partition there.