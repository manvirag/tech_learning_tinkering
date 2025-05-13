| Tool | Read | Write |  |
| --- | --- | --- | --- |
| Redis | ≤1 ms | ≤1 ms (no AOF) | Caching, real-time analytics |
| MySQL | 1–10 ms | 10–100+ ms | ACID-compliant transactions |
| DynamoDB | 1–10 ms | 5–15 ms | Scalable NoSQL workloads |

1. 1k8s →  server → ~ latency 5ms → 200 / s → 5 threads → 1000 qps → 50 pods → 50k qps.
2. 1 mysql node storage → 
    1. **Typical QPS**:
        - 1,000 to 10,000 QPS for read-heavy workloads on optimized schemas.
        - Write-heavy workloads may drop to around 100 to 1,000 QPS depending on complexity.
3. 1 cassandra
    1. Can handle approximately 1,000 to 5,000 writes per second per node with proper tuning and partitioning.
4. Avg storage by a db node: 
    1. rds → our integration db → 64tb → on avg take like 50TB
5. Network latency: 
6. nginxi → 10-100 micro second.
7. size of each row normally → 10*50 → 500B.
8. aws msk (https://aws.amazon.com/blogs/big-data/amazon-msk-now-provides-up-to-29-more-throughput-and-up-to-24-lower-costs-with-aws-graviton3-support/)
    1. ~50MB/s → let say event is 1kb → ~50K QPS → single broker.
    2. in MT i can even see 90MB/s in a broker → ~ 90k qps
    3. for →  we also have 4x of this → 200k qps per broker → 5 broker → 1M qps.
    4. latency ? → didn’t find any direct way.
    
    | M7g.4xlarge | 16 | 64 GiB :| up to 15 Gbps | up to 10 Gbps |
    | --- | --- | --- | --- | --- |
9. RabbitMQ:
    1. less qps as compare to kafka, but low latency. ( afaik ) 
    2. scaled by multiple broker.
10. mutex lock/unlock → 100ns.
11. main memory ( RAM ) → 100ns.
