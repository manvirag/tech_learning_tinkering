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




........Aboout Go Server qps , or k8s pods. 

#### HTTP Protocol Versions

##### HTTP/1.0

* **Connection per request**: Each request opens and closes a new TCP connection.
* **No keep-alive** by default.
* Result: High latency and poor scalability.

##### HTTP/1.1

* Introduced **keep-alive** (persistent connections) by default.
* A single TCP connection is **reused** for multiple requests.
* ❗ Only **one in-flight request** per connection at a time.
* Browsers work around this by opening **multiple connections per origin** (commonly 6).
* Pipelining was designed for better concurrency, but it's rarely used due to poor support.

##### HTTP/2

* Uses **a single TCP connection** per origin.
* Supports **true multiplexing**: multiple requests can be sent and responded to **concurrently over one connection**.
* Greatly improves:

  * Latency
  * Resource utilization
  * TLS performance (single handshake)

---

#### TCP Connections and Port Basics

##### How a TCP Connection Works

Each TCP connection is uniquely identified by:

```
<client IP>:<client port> → <server IP>:<server port>
```
in above cases, in case of old http browser used to send multiple client port to have custom concurrency. 

##### Server Port

* Usually fixed (e.g., `localhost:8080`).
* Listens for incoming TCP connections.

##### Client Ports

* Automatically chosen by OS from **ephemeral port range** (typically `49152–65535`).
* Allow thousands of concurrent outbound connections.
* When you make 100 concurrent requests from a client (like `curl` or `hey`), it uses 100 **different client ports**.

---

#### Goroutines in Go’s `net/http` Server

##### Server Behavior

* For each TCP connection:

  * The server spawns **one goroutine** via the `serve()` method.
* For HTTP/1.1:

  * One request is processed at a time per connection.
  * Hence, **concurrency == number of open connections**.
* For HTTP/2:

  * The connection handler demuxes and serves **many requests concurrently**, using multiple goroutines.

---

#### Keep-Alive Explained

#####  What is Keep-Alive?

* Reuses a TCP connection for multiple requests/responses.
* Enabled by default in HTTP/1.1.
* Reduces TCP connection overhead (handshakes, TLS, latency).

#####  Limitations in HTTP/1.1

* Even with keep-alive, **requests are handled serially** over a connection.
* To achieve concurrency, the client must open **multiple keep-alive connections**.

---


####  Client Examples

| Client  | TCP Behavior                                                    |
| ------- | --------------------------------------------------------------- |
| Browser | Opens \~6 TCP connections per origin                            |
| Curl    | Opens one TCP connection per request (unless keep-alive reused) |
| Postman | Uses HTTP clients under the hood, often with connection reuse   |

---

####  System-Level Constraints

##### How Many Ports?

* \~16,000+ ephemeral ports available per IP.
* Actual limit is OS-configured:

  * Linux: check with `cat /proc/sys/net/ipv4/ip_local_port_range`
  * Also limited by `ulimit -n` (max open file descriptors).

##### Finding an Available Port

* The OS dynamically allocates ephemeral ports for outbound connections.

##### Restrictions?

* Long-lived or leaked connections can exhaust ephemeral ports.
* Short-lived TCP connections reduce port reuse efficiency (common with HTTP/1.0).

---

####  Browser: Multiple TCP Connections (Pseudocode in Go)

```go
for i := 0; i < 6; i++ {
    go func(i int) {
        conn, err := net.Dial("tcp", "localhost:8080")
        if err != nil {
            log.Fatal(err)
        }
        fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: localhost\r\nConnection: keep-alive\r\n\r\n")
        io.Copy(os.Stdout, conn)
        conn.Close()
    }(i)
}
```

---

####  Key Takeaways

| Feature            | HTTP/1.1                      | HTTP/2                            |
| ------------------ | ----------------------------- | --------------------------------- |
| TCP Connections    | Multiple (e.g., 6 per origin) | Single                            |
| Concurrency        | 1 request per connection      | Many requests over one connection |
| Go server behavior | Goroutine per connection      | Multiple goroutines over one conn |
| Keep-Alive         | Enabled by default            | Required                          |
| Multiplexing       | ❌                             | ✅                                 |
| Efficiency         | Medium                        | High                              |

---

So qps ( assume explicit goroutine is one and core is 1 )
-> (1000/latency in ms ) * 10 ( let say on avg 10 either explicity connection or internal https )
