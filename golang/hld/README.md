Template:

Decision Making | Trade-off | Which to select and why to select. | Reduce code

**Step:**

1. **Functional Requirements**.

    what the software system should do?
    
    They outline the specific functions, features, and capabilities that the software must provide to meet the user's needs

    for e.g., we want to make a cache

   1. get(key)
   2. put(key, value) etc.
   3. in short features part comes in this.
      → API designing. CRUD
      → API error handling.
      → Data modeling.
      → Query.

2. **Non-Function Requirements. (Jo ki tume batai nahi jayagi , but understood hota hai)**

    These requirements answer questions like "How should the software perform?”

    for e.g. performance, scalability, availability, security, etc.

    → load balancing <br>
    → consistent hashing <br>
    → sharding <br>
    → data backup, replication, distribution <br>
    → security firewall vpn etc. authentication + authorisation + access control + rate limiting <br>
    → caching <br>
    → monitoring and alerting. <br>
    → metrics and analytics. <br>
    → logs analysis. <br>
    → continuous deployment → deployment without downtime. <br>
    → Https using <br>
    → testing before deployment  → load testing + unit test + integration test <br>
    → async messaging system, queuing system <br>
    → DNS domain name system. <br>


3. **Data Estimations.**

    The back of the envelope estimation:
    
    → This is a rough estimation
    → Don’t spend much time, because the interview says scalable. so it mostly will have scalable components like a load balancer etc.
    → Keep number simple 2^x
    
    | Zeros | Traffic | Storage |
    | --- | --- | --- |
    | 3 | thousand | KB |
    | 6 | million | MB |
    | 9 | billion | GB |
    | 12 | Trillian | TB |
    | 15 | quadrillion | PB |
    
    char →  2 bytes
    
    long/double → 8 bytes
    
    image avg → 300kb
    
    trade → cap theorem
    
    …
    5 millions user * 1kb → 5 gb storage
    
**Steps:**

**Traffic Estimations:
→** total no of requests per second on service.

→ total no of queries per second on storage.



DAU → daily active user <br>
MAU → monthly active user.

→ Write request <br>
→ Read request <br>
→ read  heavy or write heavy ? <br>



**Storage Estimations:
→** 5-year storage


**Bandwidth estimations:**

amount of data transfer per second.

**Ram Estimations(for e.g cache):**

→ total ram

**No of Machines with RAM:**



| Availability (percentage) | Downtime per year |
| --- | --- |
| 99 | 3.6 days |
| 99.99 | 52 minutes |
| 99.999 | 5 minutes |
| 99.9999 | 31 seconds |


-> **Ask the interviewer about the trade-off of CAP** <br>
→ **High-level design (API + diagraming)** <br>
→ **Deep dive into the design and its components.** <br>
→ **Identify bottlenecks and scalability.** <br>