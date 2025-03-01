

- CHAPTER 10 - BATCH PROCESSING :


	- Three different types of systems:
	    - **Services (online systems)**: Response time is usually the primary measure of performance of a service and availability is often very important. (e.g. API) 
	    - **Batch processing systems (offline systems)**: primary performance measure of a batch job is usually throughput (the time it takes to crunch through an input dataset of a certain size). (e.g. MapReduce, a batch processing algo. although the imp. of this reducing now a days .) 
	    - **Stream processing systems (near-real-time systems)**: As stream processing builds upon batch processing. (e.g. Kafka) 
	- As we shall see in this chapter, **batch processing** is an important building block in our quest to build reliable, scalable, and maintainable applications. (e.g. MapReduce → Hadoop, CouchDB, and MongoDB)
	
	#### Batch Processing with Unix Tools 
	
	- **Simple Log Analysis**
	    - Surprisingly many data analyses can be done in a few minutes using some combination of **awk, sed, grep, sort, uniq, and xargs**, and they perform surprisingly well.
	      ![[Pasted image 20250301183429.png]]
	    - **Chain of commands ( above one -> ex. of batch processing ) vs. Custom program**:
	        - Instead of the chain of Unix commands, you could write a simple program to do the same thing.  (e.g. Ruby, Python) , command one is very performance , that you'll notice once you do this for large files.
	        - limitation work on single machine.

				
	#### MapReduce and Distributed Filesystems
	- MapReduce:	
		- **MapReduce** is a bit like Unix tools, but distributed across potentially thousands of machines.
		- A single **MapReduce** job is comparable to a single Unix process: it takes one or more inputs and produces one or more outputs. Don't affect input just generate output.
		- Instead of stdin or stdout, MapReduce jobs read and write files on a **distributed filesystem**. (e.d. **HDFS** for Hadoop distributed file system, open source version of **GFS**, there are other alternative also like s3 etc.) 
		- **HDFS** is based on the **shared-nothing principle** (see the introduction to Part II).
		- **HDFS** consists of a **daemon process** running on each machine, exposing a network service that allows other nodes to access files stored on that machine (assuming that every general-purpose machine in a datacenter has some disks attached to it).
		    - A central server called the **NameNode** keeps track of which file blocks are stored on which machine. 
		    - Thus, **HDFS** conceptually creates one big filesystem that can use the space on the disks of all machines running the daemon.
		- In order to tolerate machine and disk failures, **file blocks are replicated on multiple machines**. 
		    - ensure coding scheme such as Reed–Solomon codes, kind like RAID. 
		- **HDFS has scaled well**: at the time of writing, the biggest HDFS deployments run on tens of thousands of machines, with combined storage capacity of hundreds of petabytes.
	- **MapReduce Job Execution**:
	    - **MapReduce is a programming framework** with which you can write code to process large datasets in a distributed filesystem like HDFS.
	    - To create a MapReduce job, you need to implement two callback functions:
	        - **Mapper**: The mapper is called once for every input record, and its job is to extract the key and value from the input record.
	        - **Reducer**: The MapReduce framework takes the key-value pairs produced by the mappers, collects all the values belonging to the same key, and calls the reducer with an iterator over that collection of values. 
	    - Viewed like this;
	        - The role of the **mapper** is to prepare the data by putting it into a form that is suitable for sorting.
	        - The role of the **reducer** is to process the data that has been sorted.
	    - **Distributed execution of MapReduce**:
	
	![](https://lh5.googleusercontent.com/MUHPW9rOI7U8kwsoTpX9mwt_OC2S9imphpV9c5vazwQCIwkerKwMTYiqhhsClY8NBDAPFcikbkILhKtWUeVTJIrjUqU9aHdUuBZDv_Ab6TKYeYvWh6i6goNLP8RIKRB5N-MY51iB)
	
	- In Hadoop MapReduce, the mapper and reducer are each a Java class that implements a particular interface.
	- Above shows the dataflow in a Hadoop MapReduce job
	- parallelization is based on partitioning
	- **Putting the computation near the data**: it saves copying the input file over the network, reducing network load and increasing locality.
	- The reduce side of the computation is also partitioned.  
	- The key-value pairs must be sorted, but the dataset is likely too large to be sorted with a conventional sorting algorithm on a single machine. Instead, the sorting is performed in stages.
	- The process of partitioning by reducer, sorting, and copying data partitions from mappers to reducers is known as the **shuffle**. 
	- **MapReduce workflows**:
	    - it is very common for MapReduce jobs to be **chained** together into workflows, such that the output of one job becomes the input to the next job. 
	        - this chaining is done **implicitly** by directory name: 
	        - First job must be configured to write its output to a designated directory in HDFS, 
	        - Second job must be configured to read that same directory name as its input.
	    - Various **workflow schedulers** for Hadoop have been developed, including Oozie, Azkaban, Luigi, Airflow, and Pinball 
	        - Workflows consisting of 50 to 100 MapReduce jobs are common when building recommendation systems.
	    - Various **higher-level tools** for Hadoop, such as Pig, Hive, Cascading, Crunch, and FlumeJava.
	- **Reduce-Side** **Joins and Grouping**:
	    - A **foreign key** in a relational model, a **document reference** in a document model, or **an edge** in a graph model. 
	    - MapReduce has no concept of indexes—at least not in the usual sense.
	    - **Example: analysis of user activity events**:
	
	![](https://lh6.googleusercontent.com/sIOJ108OnK11GyxAzyb2iReZKhDFwFqkc5-4d8t5RXDJrofdnh0xx8pc6rAbSLcgJ1xYSEWSCoheh4F0Z0e8rqrfnjNfzntbbK7EdMeIR34gM62LdXsVKgt8UVJk-Hq8AGpuCWGj)
	
	- **Star schema**: the log of events is the fact table, and the user database is one of the dimensions.
	- In order to achieve good throughput in a batch process, the computation must be (as much as possible) local to one machine.
	- **Sort-merge joins**:
	
	![](https://lh4.googleusercontent.com/Ah4fzlrSOWx-I8wL3rjp-8ENf-AWkcvJtpTS_DC8tsS6YOf8q41aYbUbSyAvvvLFUD5HflSJrqhcNdm03Hwlqh-KEXkcg2gVD4EcGpYpKSs_47GRmQMOOJGlGlmrpifcD0UIJtOd)
	
	- The effect is that all the activity events and the user record with the same user ID become adjacent to each other in the reducer input. (along with secondary sort) 
	- **sort-merge join**: Since the reducer processes all of the records for a particular user ID in one go, it only needs to keep one user record in memory at any one time, and it never needs to make any requests over the network. (C: this explained why ETL tool Pentaho/Kettle need to always sort the value before “Merge Join Row”) 
	- **Bringing related data together in the same place**:
	    - One way of looking at this architecture is that mappers “send messages” to the reducers. 
	        - When a mapper emits a key-value pair, the **key acts like the destination address** to which the value should be delivered. 
	    - Using the **MapReduce programming model** has **separated** the physical network communication aspects of the computation (getting the data to the right machine) **from** the application logic (processing the data once you have it).
	- **GROUP BY**:
	    - **The simplest way** is to set up the mappers so that the key-value pairs they produce use the desired grouping key. 
	    - **Another common use** for grouping is collating all the activity events for a particular user session, in order to find out the sequence of actions that the user took—a process called **sessionization**. (e.g. for A/B testing) 
	- **Handling skew**:
	    - The pattern of “bringing all records with the same key to the same place” breaks down if there is a very large amount of data related to a single key. 
	        - e.g. celebrities in SNS, Such disproportionately active database records are known as linchpin objects or hot keys.
	    - If a join input has hot keys, there are a few algorithms you can use to compensate. (e.g. skewed join method in Pig, sharded join method in Crunch) 
	    - Hive’s skewed join optimization takes an alternative approach.
	    - When grouping records by a hot key and aggregating them, you can perform the grouping in two stages. 
	- **Map-Side** **Joins**:
	    - The reduce-side approach has the advantage that you do not need to make any assumptions about the input data.
	    - if you can **make certain assumptions** about your input data, it is possible to **make joins faster** by using a so-called **map-side join**.
	        - This approach uses a cut-down MapReduce job in which there are no reducers and no sorting.
	    - **Broadcast hash joins**:
	        - The simplest way of performing a map-side join applies in the case where a **large dataset is joined with a small dataset**.
	            - the small dataset needs to be small enough that it can be loaded entirely into memory in each of the mappers
	        - This simple but effective algorithm is called a **broadcast hash join**: 
	            - The word **broadcast** reflects the fact that each mapper for a partition of the large input reads the entirety of the small input (so the small input is effectively “**broadcast**” to all partitions of the large input), and the word **hash** reflects its use of a hash table.
	            - E.g. Pig (under the name “replicated join”), Hive (“MapJoin”), Cascading, and Crunch, Data-warehouse engine Impala.
	        - Instead of loading the small join input into an in-memory hash table, an alternative is to store the small join input in a read-only index on the local disk. (fit in OS’ page cache, almost as fast as memory) 
	    - **Partitioned hash joins**: (e.g. bucketed map joins in Hive) 
	        - If the inputs to the map-side join are partitioned in the same way, then the hash join approach can be applied to each partition independently.
	        - This approach only works if both of the join’s inputs have the same number of partitions, with records assigned to partitions based on the same key and the same hash function.
	    - **Map-side merge joins**:
	        - not only partitioned in the same way, but also sorted based on the same key.
	    - **MapReduce workflows with map-side joins**:
	        - When the output of a MapReduce join is consumed by downstream jobs, the choice of map-side or reduce-side join affects the structure of the output.
	        - Knowing about the physical layout of datasets in the distributed filesystem becomes important when optimizing join strategies. 
	            - In the Hadoop ecosystem, this kind of **metadata** about the partitioning of datasets is often maintained in **HCatalog** and the **Hive metastore**.
	- **The Output of Batch Workflows**:
	    - Where does batch processing fit in? 
	        - It is not transaction processing, nor is it analytics. It is **closer to analytics**, in that a batch process typically scans over large portions of an input dataset.
	    - The output of a batch process is often not a report, but some other kind of structure.
	    - **Building search indexes**: (documents in, indexes out.)
	        - Google’s original use of MapReduce was to build indexes for its search engine, which was implemented as a workflow of 5 to 10 MapReduce jobs.  (e.g. still used today by **Lucene/Solr**)
	        - Recall **full-text search index**: it is a file (the term dictionary) in which you can efficiently look up a particular keyword and find the list of all the document IDs containing that keyword (the postings list). 
	    - **Key-value stores as batch process output**: (database files in, database out)
	        - Another common use for batch processing is to build **machine learning systems** such as **classifiers** (e.g., spam filters, anomaly detection, image recognition) and **recommendation systems** (e.g., people you may know, products you may be interested in, or related searches)
	        - Build a brand-new database inside the batch job and write it as files to the job’s output directory in the distributed filesystem, just like the search indexes in the last section.
	            - Various key-value stores support building database files in MapReduce jobs, including Voldemort, Terrapin, ElephantDB, and HBase bulk loading.
	    - **Philosophy of batch process outputs**:
	        - In the process, the **input is left unchanged**, any previous **output is completely replaced** with the new output, and there are no other side effects.
	        - By treating **inputs as immutable** and avoiding side effects (such as writing to external databases), batch jobs not only achieve good performance but also become much easier to maintain. 
	        - On Hadoop, some of those low-value syntactic conversions are eliminated by using more structured file formats: e.g. Avro, Parquet 
	- **Comparing Hadoop to Distributed Databases**:
	    - **Hadoop** is somewhat like a distributed version of Unix, where **HDFS** is the filesystem and MapReduce is a quirky implementation of a Unix process(which happens to always run the sort utility between the map phase and the reduce phase). 
	    - **MapReduce** and a **Distributed Filesystem** provides something much more like a general-purpose operating system that can run arbitrary programs.
	    - **Diversity of storage**:
	        - Databases require you to structure data according to a particular model (e.g., relational or documents), 
	            - whereas files in a distributed filesystem are just byte sequences, which can be written using any data model and encoding. 
	        - Collecting data in its raw form, and worrying about schema design later, allows the data collection to be speeded up (a concept sometimes known as a “**data lake**” or “**enterprise data hub**” ). 
	            - Aka. **sushi principle**: “**raw (data) is better**” 
	        - Indiscriminate data dumping shifts the burden of interpreting the data from producer to consumer’s problem (schema-on-read approach). 
	        - There may not even be one ideal data model, but rather different views onto the data that are suitable for different purposes.
	        - Data modeling still happens, but it is in a separate step, decoupled from the data collection. 
	            - This decoupling is possible because a distributed filesystem supports data encoded in any format.
	    - **Diversity of processing models**:
	        - MapReduce gave engineers the ability to easily run their own code over large datasets.
	        - Sometimes having two processing models, SQL and MapReduce, was not enough. 
	        - The system is flexible enough to support a diverse set of workloads within the same cluster. 
	        - Not having to move data around makes it a lot easier to derive value from the data, and a lot easier to experiment with new processing models.
	    - **Designing for frequent faults**:
	        - When comparing MapReduce to MPP databases, two more differences in design approach stand out: **the handling of faults** and **the use of memory and disk**.
	            - **MPP** databases prefer to keep as much data as possible in memory (e.g., using hash joins) to avoid the cost of reading from disk.
	            - **MapReduce** is very eager to write data to disk, partly for fault tolerance, and partly on the assumption that the dataset will be too big to fit in memory anyway.
	        - Overcommitting resources in turn allows better utilization of machines and greater efficiency compared to systems that segregate production and non-production tasks.
	        - It’s not because the hardware is particularly unreliable, it’s because the freedom to arbitrarily terminate processes enables better resource utilization in a computing cluster.
	            - Among open source cluster schedulers, preemption is less widely used. (e.g. YARN’s CapacityScheduler)
	
	#### Beyond MapReduce
	
	- Depending on the volume of data, the structure of the data, and the type of processing being done with it, other tools may be more appropriate for expressing a computation.
	- Implementing a complex processing job using the raw MapReduce APIs is actually quite hard and laborious—for instance, you would need to implement any join algorithms from scratch. 
	- In response to the difficulty of using MapReduce directly, various higher-level programming models (Pig, Hive, Cascading, Crunch) were created as abstractions on top of MapReduce. 
	- **Materialization of Intermediate State**:
	    - Publishing data to a well-known location in the distributed file system allows loose coupling so that jobs don’t need to know who is producing their input or consuming their output.
	    - **Intermediate state**: a means of passing data from one job to the next. (not shared between job or team) 
	        - The process of writing out this intermediate state to files is called **materialization**. (C: recall “materialized views” from previously) 
	        - In contrast: **Pipes** do not fully materialize the intermediate state, but instead **stream** the output to the input incrementally, using only a small in-memory buffer.
	    - MapReduce’s approach of **fully materializing intermediate state** has downsides compared to Unix pipes: (C: which derived from its advantages) 
	        - A MapReduce job can only start when all tasks in the preceding jobs (that generate its inputs) have completed;
	        - Mappers are often redundant: they just read back the same file that was just written by a reducer, and prepare it for the next stage of partitioning and sorting.
	        - Storing intermediate state in a distributed file system means those files are replicated across several nodes, which is often overkill for such temporary data.
	    - **Dataflow engines**:
	        - New execution engine created to solve previou problems. E.g. **Spark**, Tez, and Flink. 
	            - They handle an entire workflow as one job, rather than breaking it up into independent subjobs.
	        - Since they explicitly model the flow of data through several processing stages, these systems are known as **dataflow engines**.
	        - Offers several advantages compared to the MapReduce model:
	            - Expensive work such as sorting need only be performed in places where it is actually required.
	            - There are no unnecessary map tasks
	            - Can make locality optimizations.
	            - It is usually sufficient for intermediate state between operators to be kept in memory or written to local disk.
	            - Operators can start executing as soon as their input is ready;
	            - Existing Java Virtual Machine (JVM) processes can be reused to run new operators, reducing startup overheads compared to MapReduce (which launches a new JVM for each task).
	    - **Fault tolerance**:
	        - An advantage of fully materializing intermediate state to a distributed file system is that it is durable, which makes fault tolerance fairly easy.
	            - if a task fails, it can just be restarted on another machine and read the same input again from the filesystem.
	        - **Spark, Flink, and Tez** avoid writing intermediate state to HDFS, so they take a different approach to tolerating faults: 
	            - if a machine fails and the intermediate state on that machine is lost, it is **recomputed** from other data that is still available.
	        - To enable this **recomputation**, the framework must keep track of how a given piece of data was computed—which input partitions it used, and which operators were applied to it. 
	            - When recomputing data, it is important to know whether the computation is **deterministic**.
	                - The solution in the case of non-deterministic operators is normally to kill the downstream operators as well, and run them again on the new data.
	            - In order to avoid such cascading faults, **it is better to make operators deterministic**.
	        - Recovering from faults by recomputing data is not always the right answer: 
	            - if the intermediate data is much smaller than the source data, or if the computation is very CPU-intensive, it is probably cheaper to materialize the intermediate data to files than to recompute it.
	    - **Discussion of materialization**:
	        - **Flink** especially is built around the idea of pipelined execution: that is, incrementally passing the output of an operator to other operators, and not waiting for the input to be complete before starting to process it.
	- **Graphs and Iterative Processing**:
	    - In graph processing, the data itself has the form of a graph. 
	    - This need often arises in machine learning applications such as recommendation engines, or in ranking systems.  (e.g. PageRank)
	    - Iterative style: (works, but very inefficient with MapReduce) 
	        - 1. An external scheduler runs a batch process to calculate one step of the algorithm.
	        - 2. When the batch process completes, the scheduler checks whether it has finished.
	        - 3. If it has not yet finished, the scheduler goes back to step 1 and runs another round of the batch process.
	    - **The Pregel processing model**:
	        - As an optimization for batch processing graphs, the **bulk synchronous parallel (BSP) model** of computation has become popular. Aka. “**Pregel model**”. 
	            - (implemented by Apache Giraph, Spark’s GraphX API, and Flink’s Gelly API.)
	        - **Idea behind Pregel**: one vertex can “send a message” to another vertex, and typically those messages are sent along the edges in a graph.
	        - In each iteration, a function is called for each vertex, passing it all the messages that were sent to it—much like a call to the reducer.  (It’s a bit similar to the actor model) 
	    - **Fault tolerance**:
	        - Pregel implementations **guarantee that messages are processed exactly once** at their destination vertex in the following iteration.
	        - This fault tolerance is achieved by periodically **check-pointing** the state of all vertices at the end of an iteration. (i.e., writing their full state to durable storage.) 
	    - **Parallel execution**:
	        - A vertex does not need to know on which physical machine it is executing; when it sends messages to other vertices, it simply sends them to a vertex ID.
	- **High-Level APIs and Languages**:
	    - **Spark** and **Flink** also include their own high-level dataflow APIs, often taking inspiration from **FlumeJava**.
	    - These dataflow APIs generally use relational-style building blocks to express a computation: 
	        - joining datasets on the value of some field; 
	        - grouping tuples by key; 
	        - filtering by some condition; 
	        - and aggregating tuples by counting, summing, or other functions.
	    - **The move toward declarative query languages**:
	        - The choice of join algorithm can make a big difference to the performance of a batch job; 
	        - This is possible if joins are specified in a declarative way: the application simply states which joins are required, and the query optimizer decides how they can best be executed. 
	        - Hive, Spark DataFrames, and Impala also use **vectorized execution**: iterating over data in a tight inner loop that is friendly to CPU caches, and avoiding function calls. 
	        - batch processing frameworks begin to look more like **MPP databases** (and can achieve comparable performance) they retain their flexibility advantage.
	    - **Specialization for different domains**:
	        - Another domain of increasing importance is **statistical** and **numerical** algorithms, which are needed for machine learning applications such as classification and recommendation systems. 
	
	#### Summary
	
	- In this chapter we explored the topic of batch processing.
	- In the Unix world, the uniform interface that allows one program to be composed with another in **files** and **pipes**; 
	    - In **MapReduce**, that interface is a distributed file system. 
	- Dataflow engines add their own pipe-like data transport mechanisms to avoid materializing intermediate state to the distributed file system, but the initial input and final output of a job is still usually HDFS.
	- The two main problems that distributed batch processing frameworks need to solve are:
	    - **Partitioning**: In MapReduce, mappers are partitioned according to input file blocks.
	    - **Fault tolerance**: MapReduce frequently writes to disk, which makes it easy to recover from an individual failed task. 
	- **Join algorithms for MapReduce**:
	    - **Sort-merge joins**: Each of the inputs being joined goes through a mapper that extracts the join key.
	    - **Broadcast hash joins**: One of the two join inputs is small, so it is not partitioned and it can be entirely loaded into a hash table. 
	    - **Partitioned hash joins**: If the two join inputs are partitioned in the same way (using the same key, same hash function, and same number of partitions), then the hash table approach can be used independently for each partition.
	- Distributed batch processing engines have a deliberately restricted programming model: **callback functions** (such as mappers and reducers) are assumed to be **stateless** and to have **no externally visible side effects** besides their designated output.
	- Does **not need to worry about implementing fault-tolerance mechanisms**: the framework can guarantee that the final output of a job is the same as if no faults had occurred, even though in reality various tasks perhaps had to be retried.
	- The output is derived from the input.  And the input data is **bounded**: it has a known, fixed size (for example, it consists of a set of log files at some point in time, or a snapshot of a database’s contents).
	



CHAPTER 11: STREAM PROCESSING:


A complex system that works is invariably found to have evolved from a simple system that works. The inverse proposition also appears to be true: A complex system designed from scratch never works and cannot be made to work.

—John Gall, Systemantics (1975)

- Batch process is under the assumption all the data is **Bounded**, which means we know the finite size of the data we are dealing with, so it is known when the job is finished. 
- In reality, a lot of the data is unbounded. (because data keeps generated every second). This issue will force the batch process to divide data into “chunks”.  
    - Which means the result/derived data is **delayed** based on the interval you are chosen.  (this is too slow for impatient users) 
- Continuously data processing without interruptions(break into chunks) is the idea behind **streaming processing**.
    - Think of “Stream” as a never stop flow of water/river that keep feeding data in;
    - E.g. (stdin/stdout, file inputstream, TCP stream etc.) 
    - “Event Stream” as a data management mechanism 

#### Transmitting Event Streams

- Batch processing: Input are **Files**  vs. Streaming Processing: Input are **Events**;
- What is **Event**: a small, self-contained, immutable object containing the details of something that happened at some point in time. 
- Event generated by a producer(Publisher/Sender) and then processed by multiple consumers(subscribers/recipients). 
    - Related “events” are grouped together by **Topic/Stream**.
- Batch vs. Stream processing is kind like the difference between Pull and Push. 
    - Batch: is the consumer keep Pull event from DB. 
    - Stream: is the DB keep push Event to Consumer.
- Traditional DB/(RDMS) is not designed for “Stream/Event” processing 
- **Message Systems**:
    - A producer sends a message containing the event, which is then pushed to consumers.
    - MQ vs. Unix Pipe or TCP 
        - MQ allows many-to-many relationships (Producer vs. Consumer)
        - As Unix Pipe & TCP is usually one-to-one 
    - What happens if the producers send messages faster than the consumers can process them?  Three options
        - Drop message;  Queue ; backpressure (flow control) 
        - What if Queue is full ? 
    - What happens if nodes crash or temporarily go offline—are any messages lost? 
        - Still a trade off between C & A (Consistency and Availability) 
    - **Direct messaging from producers to consumers**: (prone to loss data) 
        - **UDP multicast**: used in financial industry for streams such as stock market feeds; 
        - **Brokerless message library**: ZeroMQ. 
        - **StatsD & Brubeck**: UDP messaging. 
    - **Message brokers**: (aka. Message Queue, e.g. ActiveMQ, RabbitMQ)
        - A special type of DB that optimized for Message Streams; 
        - Producer → Broker → Consumer 
        - Better fault tolerance; Message could be persist to disk; 
        - It all asynchronously; 
    - **Message brokers compared to databases**: (JMS, AMQP) 
        - In MQ Some even support 2PC (two-phase commit); 
        - In MQ, data is deleted right-after message is been consumed; 
        - In MQ, usually assume the Queue is short. 
        - DB Secondary Indexes vs. MQ subset of topics
        - MQ has no support for queries, but notify clients when data change
        - E.g. RabbitMQ, ActiveMQ, HornetQ, Qpid, TIBCO Enterprise Message Service, IBM MQ, Azure Service Bus, and Google Cloud Pub/Sub. 
    - **Multiple consumers**:
        - **Load balancing**:  arbitrarily assigned to worker/consumer; Good for parallel processing expensive work-load. 
        - **Fan-out**: Each message is sent to all consumers/worker; (topic subscription in JMS, exchange bindings in AMQP) 

![](https://lh5.googleusercontent.com/RiEGJXtyVVLGFU_VGZGPn5LQ1vcZaygSRk7Pl3pIF7YxtDvmbqMZda9ECli1RdIwWC_lOw-3zV0OKW9W-xhEbSpeood34t1O7j292eTTeYAj2-pXPc3wskE02Yiup7ptF8L-Rhcn)

- Two patterns could be combined. 
- **Acknowledgments and redelivery**:
    - **acknowledgments**: a confirmation from a client that it has finished processing a message so that the broker can remove it from the queue.
        - Note: due to network issues, the ack. Could be lost, then cause the ordering of the message change (C: In general, when you using Queue, you shouldn’t have cared about the order at first place) 
- **Partitioned Logs**:
    - **Transient messaging mindset**: transient operation that leaves no permanent trace.  (Which is totally opposite than DB or FileSystem) 
    - **MQ is NOT idempotent**: receiving a message is destructive if the acknowledgment causes it to be deleted from the broker;
    - Why can we not have a hybrid, combining the durable storage approach of databases with the low-latency notification facilities of messaging?
        - Yes, **log-based message brokers**.
    - **Using logs for message storage**:
        - Log: is simply an append-only sequence of records on disk.
        - **Log-based Message broker**: A producer sends a message by appending it to the end of the log, and a consumer receives messages by reading the log sequentially. 
        - **This log can be partitioned**. 
        - Each message is offset by a sequence number. (totally ordered) 
            - Note: no ordering guarantee across partitions tho. 

![](https://lh5.googleusercontent.com/vcEA6f2Hreq8PEMT-rxaf5ScgujSbRYGF2BIN8oHYA9D7UVVm66weFW3AXxREYjKJfM5CbpSJ1wUgbHhzGTVZyp-6sToR_RUTD4GN0t8fvsiGfZ2IbkhzkrUUwGlxQU6f8fkAaYa)

- E.g. Apache **Kafka**, Amazon Kinesis Streams, and Twitter’s DistributedLog. (millions of MPS by partitioning) 
- **Logs compared to traditional messaging***: 
    - the broker can assign entire partitions to nodes in the consumer group, Then each client consumes all the messages in the partitions it has been assigned.
    - **JMS/AMQP style of message broker is preferable**: Message is expensive, parallel processing, order doesn’t matter. 
    - **Log-based approach**: high message throughput, where each message is fast to process and where message ordering is important. 
- **Consumer offsets**:
    - Similar to “log sequence number” in single-leader DB replication; 
    - The message broker behaves like a leader database, and the consumer like a follower.
- **Disk space usage**:
    - To prevent run out of storage, log is divided into segments, old segments are deleted or archived. 
    - **“Log” is kind of like a bounded-size buffer**, if the consumer can’t keep up, the old message will be discarded. (aka. Circular buffer, ring buffer) 
        - E.g. 6T HDD with 150MB/s write speed can buffer up to 11 hrs of messages. 
- **When consumers cannot keep up with producers**:
    - Three choices: **dropping**, **buffering** or **backpressure**(flow-control) 
    - Consumers are independent from each other; 
- **Replaying old messages**:
    - it is a read-only operation that does not change the log.
    - Offset is under the consumer’s control. So, it has the freedom to go back to previous data/offset. 
    - This made it easier for integration with other dataflows. 

#### Databases and Streams

- DB and Stream are correlated deeper; (e.g. write is an event)  
- The replication log is a stream of database write events, produced by the leader as it processes transactions. 
- **Keeping Systems in Sync**:
    - Often need to combine several different technologies in order to satisfy their requirements. (Write/Read/Search/Analytics etc.) 
    - It is essential to keep all the data in-sync. if an item is updated in the database, it also needs to be updated in the cache, search indexes, and data warehouse. (usually through ETL processes.) 
        - Or “Dual writes”, which prone to issue like race-conditiotion; 

![](https://lh5.googleusercontent.com/V8jDH8Txozb5aHwFjw7oqyJnJXSr8sFXVKcurQnbT-J1XN8KYsey7O6l6HSNiQPe2FH35Wv9L9j9tE1xL2JO2GyhWYCLsyoe8QJq-aeCkOUQt2fwnhpHZ_ckA66AcdpRx0Dk3iUo)

- The key is to determine if we can have only one “Source of Truth” (aka. Leader) 
- **Change Data Capture**:
    - CDC(Change Data Capture): is a process that extracts DB changes and puts it into other systems. 
        - E.g. in the form of “Stream” 

![](https://lh6.googleusercontent.com/17JtXcnyV-Kp9NUlh6RHp99KHz4qIgLf-qQBlMgYMZ4RVH53JQMq0gBXNG0wdA9YqL19TAlReB9p5INuOoZObWmtt25sxJU6wRKKttnwMenjaZ2dt0S8UF3XurdW60mUreplCByi)

- **Implementing change data capture**:
    - We call any “log consumer” a “derived data system”. The idea behind CDC is to ensure all those “derived data systems” got the up-to-date changes. 
    - Essentially, CDC makes one database the **leader** (the one from which the changes are captured), and turns the others into followers.
    - Potential Mechanism: DB Triggers(Con:performance overheads), Parsing replication log(Con:schema changes); 
    - Usually done asynchronously → replication lag 
- **Initial Snapshot**:
- **Log compaction**: (e.g. Apache Kafka) 
    - This is used when need add a new “derived data system” 
    - This allows the message broker to be used for durable storage, not just for transient messaging. 
- **API support for change streams**: (e.g. RethinkDB, FIreBase, CouchDB, Meteor, VoltDB) 
    - DBs engine started to support change streams; 
    - A table will hold transactions but can’t be queried. 
- **Event Sourcing**:
    - **Event Sourcing**: a technique that was developed in the domain-driven design (DDD) community.
    - **CDC vs. ES**(Event Sourcing):  different level of abstraction.
        - **CDC**: application isn’t aware of CDC occurring, so it happens at a lower level. 
        - **ES**: reflect things that happened at the application level. 
    - **ES is a powerful technique for data modeling**, because it makes more sense to record a user’s action as immutable events, rather than the effect of the actions on a mutable DB. 
    - Event sourcing is similar to the chronicle data model. (or “fact table”).
    - **Deriving current state from the event log**:
        - Applications need to take the events and transform it to an application state that is suitable for the user to view. (deterministic) 
    - **Commands and events**:
        - First it comes as “Command” and then after it successfully executed, it becomes “Event” which is durable and immutable. 
            - when the event is generated, it becomes a fact.
        - A consumer of the event stream is not allowed to reject an event;
        - Any validation of a command needs to happen synchronously, before it becomes an event. 
- **State, Streams, and Immutability**:
    - **Immutability** is also what makes event sourcing and change data capture powerful.
    - Whenever you have a state that changes, that state is the result of the events that mutated it over time.
    - **mutable state** and an **append-only log** of immutable events do not contradict each other: they are two sides of the same coin. 
        - the **changelog**, represents the evolution of state over time.
    - In terms of mathematical: 
        - **application state** is what you get when you **integrate** an event stream over time;
        - **a change stream** is what you get when you **differentiate** the state by time; 
        - ![](https://lh4.googleusercontent.com/Ny-nL5DbdWJNFuMqm35nGzRw8HvQfqPaBtptcpBHnhsbNkE7qUPzZsnKprORFWbQog4NZCF14V42hlWP3CI1qYMXaGWwOwJljSKXkajy3a9Aoeomn-_HWZDvC6YpM7uxFSRqO38x)
    - **Quote from Pat Helland**:
        - _Transaction logs record all the changes made to the database. High-speed appends are the only way to change the log. From this perspective, the contents of the database hold a caching of the latest record values in the logs._ **_The truth is the log_**_. The database is a cache of a subset of the log. That cached subset happens to be the latest value of each record and index value from the log_.
    - **Log compaction**: bridging the distinction between log and DB state.  it retains only the latest version of each record, and discards overwritten versions.
    - **Advantages of immutable events**: (e.g. Account ledger) 
        - Particularly important in financial systems, it is also beneficial for many other systems.
        - Capture more information than just the current state.
        - E.g. Shopping cart history with append-only event could help analytic in the future; 
    - **Deriving several views from the same event log**:
        - You can derive several different read-oriented representations from the same log of events. (C: This idea is similar to the talk about Apache Kafka,built application around the Kafka stream) 
        - Having an explicit translation step from an event log to a database makes it easier to evolve your application over time.  (C: enable old & new system running side by side) 
        - **Command Query Responsibility Segregation (CQRS)**: you gain a lot of flexibility by separating the form in which data is written from the form it is read. 
    - **Concurrency control**:
        - The biggest downside of event sourcing and change data capture is asynchronous. → cause delay. 
        - Potential Solutions: 
            - “Reading your own writes” 
            - “Implementing linearizable storage using total order broadcast” 
    - **Limitations of immutability**:
        - Truly deleting data could be difficult because data live in many places. 
        - Deletion is more a matter of “making it harder to retrieve the data” than actually “making it impossible to retrieve the data.” 

#### Processing Streams

- Where streams come from (user activity events, sensors, and writes to databases).
- How streams are transported (through direct messaging, via message brokers, and in event logs).
- Last question is **What can you do with Stream** ? three major options 
    - 1, write it to a database, cache, search index, or similar storage system so it can be used by other applications/clients/systems. 
        - Kind like **maintaining materialized views**.
    - 2, push the events to users directly; 
    - 3, process one or more input streams to produce one or more output streams. (pipelining); **processing streams to produce other, derived streams**. 
- A block of Code that processes streams so called “Operator” or “Job”. (kind like Unique processes or MapReduce job) 
    - Since the Stream never ends(unbounded), so sorting doesn’t make sense here, neither does sort-merge joins will be used. 
    - Fault-tolerance mechanisms also need to be revised. 
- **Use of Stream Processing**
    - **Monitoring System**: Fraud detection, Trading System, Manufacturing System, Military and Intelligence systems. 
    - **Complex event processing(CEP)**: emerged from the 90s
        - CEP allows you to specify rules to search for certain patterns of events in a stream.
    - **Stream analytics**: (e.g. Apache Storm, Spark Streaming, Flink, Concord, Samza, and Kafka Streams, Google Cloud Dataflow and Azure Stream Analytics) 
        - More oriented toward aggregations and statistical metrics over a large number of events;
        - Stream analytics systems sometimes use **probabilistic algorithms**, such as Bloom filters. 
    - **Maintaining materialized views**:
        - Derived data systems can be treated as maintaining materialized views. 
    - **Search on streams**:
        - The percolator feature of Elasticsearch is one option for implementing this kind of stream search.
    - **Message passing and RPC**:
- **Reasoning About Time**
    - Time “window” 
    - Using the timestamps in the events allows the processing to be **deterministic**. 
    - **Event time versus processing time**: (e.g. Star War movies) 
        - Processing may be delayed. 
        - Confusing event time and processing time leads to bad data.
        - ![](https://lh4.googleusercontent.com/ZhGu2_uf1KuNRRGTeh2LjALbRhhqgIvpYvPb3JjlrdT_qAR058dV6hGy_nPZHU3YUaaIgiSQXyD8MMhpu4jqKFY3FA3Av5yMRa6x93s9qnxmfO_8NBq1HmSRuCSW-59s9SKc99AB)
    - **Knowing when you’re ready**:
        - need to be able to handle such **straggler** events that arrive after the window has already been declared complete.
            - 1, Ignore the straggler events;
            - 2, Publish a correction; 
    - **Whose clock are you using, anyway**?
        - Need address Incorrect device clocks, log three timestamps:
            - The time at which the event occurred, according to the device clock
            - The time at which the event was sent to the server, according to the device clock
            - The time at which the event was received by the server, according to the server clock
    - **Types of windows**:
        - Tumbling window: fixed length, and every event belongs to exactly one window.
        - Hopping window: fixed length, but allows windows to overlap in order to provide some smoothing.
        - Sliding window: contains all the events that occur within some interval of each other.
        - Session window: has no fixed duration. But, grouping together all events relative to the same user that occur closely together in time. (e.g. website analytics) 
- **Stream Joins**
    - Similar to batch jobs; However, since new events can appear anytime on a stream makes joins on streams more challenging than in batch jobs.
    - three different types of joins: **stream-stream joins**, **stream-table joins**, and **table-table joins**.
    - **Stream-stream join (window join)**:
        - a stream processor needs to maintain state. 
    - **Stream-table join (stream enrichment)**:
        - Enriching the activity events with information from the database.
        - Instead of performing remote SQL queries, we can cache up a copy of DB. (In Memory hashtable or local disk index) 
            - Need CDC to ensure the stream data is up-to-date; 
        - A stream-table join is actually very similar to a stream-stream join, but in this case we have “table changelog stream” involved. 
    - **Table-table join (materialized view maintenance)**: (e.g. Tweets)
        - it maintains a materialized view for a query that joins two tables. 
    - **Time-dependence of joins**:
        - **Common**: they all require the stream processor to maintain some state based on one join input, and query that state on messages from the other join input.
        - If the state changes over time, and you join with some state, what point in time do you use for the join ? (e.g. sales Tax calculation) 
        - If the ordering of events across streams is undetermined, the join becomes nondeterministic;
        - **slowly changing dimension (SCD)**:  addressed by using a unique identifier for a particular version of the joined record. (but this approach made log compaction impossible, because we need retain all version of the records) 
- **Fault Tolerance** 
    - You can’t wait until a stream is finished to validate its output/result, since all the stream is unbounded and will never really finish/complete. 
    - **Microbatching and checkpointing**:
        - **Microbatching**: break the stream into small blocks, and treat each block like a miniature batch process.  (e.g. **Spark** Streaming)  usually one second interval.
            - Smaller the batches size the greater overhead. 
            - Larger batches size means longer delay of results. 
            - implicitly provides a tumbling window equal to the batch size
        - **Checkpointing**: triggered by barriers in the message stream, similar to the boundaries between microbatches, but without forcing a particular window size.  (e.g. Apache **Flink**) 
        - Both approaches won’t prevent external side effects after the results have been written into External Systems. 
    - **Atomic commit revisited**:
        - Achieve “Exactly-Once” processing without transactions across heterogeneous technologies. 
        - **Idempotence**:
            - Distributed transactions are one way of achieving that goal, but another way is to rely on **idempotence**.
            - if an operation is not naturally idempotent, it can often be made idempotent with a bit of extra metadata. (e.g. Kafka with some offset value) 
        - **Rebuilding state after a failure**:
            - keep state local to the stream processor, and replicate it periodically.
            - sometimes the state can be rebuilt from the input streams. 

#### Summary

- Discussed **event streams**, what purposes they serve, and how to process them. 
    - Similar to “batch processing” but unbounded. 
    - **message brokers** and **event logs** serve as the streaming equivalent of a filesystem.
- Two types of Message brokers:
    - **AMQP/JMS-style message broker**: exact order is not important 
    - **Log-based message broker**: order is kept. 
        - Similar to log-structured storage engines
- Where streams come from ? 
    - user activity events, sensors providing periodic readings, and data feeds (e.g., market data in finance)
    - writes to a database as a stream: capture the changelog
        - Change Data Capture 
        - Event Sourcing 
- DB as streams is very useful for integrating different systems
    - E.g.   search indexes, caches, and analytics systems. 
- Stream joins and fault tolerance:  By maintaining state as streams and replaying messages
    - searching for event patterns (complex event processing), 
    - computing windowed aggregations (stream analytics), 
    - keeping derived data systems up to date (materialized views).
- three types of joins:
    - Stream-stream joins
    - Stream-table joins
    - Table-table joins
- fault tolerance and exactly-once semantics
    - microbatching, 
    - checkpointing, 
    - transactions, 
    - idempotent writes.