

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
	        - Sorting in linux work on external sorting, while langugae code work in inmemory, external sorting -> use disk if require would be able to sort if very large file not able fill in in memory use merge sort.

				
	#### MapReduce and Distributed Filesystems
	- MapReduce:	
		- **MapReduce** is a bit like Unix tools, but distributed across potentially thousands of machines.
		- A single **MapReduce** job is comparable to a single Unix process: it takes one or more inputs and produces one or more outputs. Don't affect input just generate output. immutable input.
		- Instead of stdin or stdout, MapReduce jobs read and write files on a **distributed filesystem**. (e.d. **HDFS** for Hadoop distributed file system, open source version of **GFS**, there are other alternative also like s3 etc.) 
		- **HDFS** is based on the **shared-nothing principle** (see the introduction to Part II). ( have data on their machine )
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
	- Above shows the dataflow in a Hadoop MapReduce job -> hadoop kind of orchestrator, we give the jar files of map and reduce
	- parallelization is based on partitioning
	- **Putting the computation near the data**: it saves copying the input file over the network, reducing network load and increasing locality.
	- The reduce side of the computation is also partitioned.  
	- The key-value pairs must be sorted, but the dataset is likely too large to be sorted with a conventional sorting algorithm on a single machine. Instead, the sorting is performed in stages.
	- The process of partitioning by reducer, sorting, and copying data partitions from mappers to reducers is known as the **shuffle**. shufftle is handled by framework itself only map and reduce given by user.
	- **MapReduce workflows**:
	    - it is very common for MapReduce jobs to be **chained** together into workflows, such that the output of one job becomes the input to the next job. 
	        - this chaining is done **implicitly** by directory name: 
	        - First job must be configured to write its output to a designated directory in HDFS, 
	        - Second job must be configured to read that same directory name as its input.
	    - Various **workflow schedulers** for Hadoop have been developed, including Oozie, Azkaban, Luigi, Airflow, and Pinball 
	        - Workflows consisting of 50 to 100 MapReduce jobs are common when building recommendation systems.
	    - Various **higher-level tools** for Hadoop, such as Pig, Hive, Cascading, Crunch, and FlumeJava.
	- **Reduce-Mapper-Side** **Joins : 
		- in map reduce we do join of multiple file to get data , that we are saying jion in below we want to know actvity with date , this join can be done on mapper side and on reduce side . here we are discuss reduce side then will discuss on map side ). 
		- As of now not going in details, in this book mentioned different ways to do this join and how can be optimised . 
	
	![](https://lh6.googleusercontent.com/sIOJ108OnK11GyxAzyb2iReZKhDFwFqkc5-4d8t5RXDJrofdnh0xx8pc6rAbSLcgJ1xYSEWSCoheh4F0Z0e8rqrfnjNfzntbbK7EdMeIR34gM62LdXsVKgt8UVJk-Hq8AGpuCWGj)
	
	- 
		![](https://lh4.googleusercontent.com/Ah4fzlrSOWx-I8wL3rjp-8ENf-AWkcvJtpTS_DC8tsS6YOf8q41aYbUbSyAvvvLFUD5HflSJrqhcNdm03Hwlqh-KEXkcg2gVD4EcGpYpKSs_47GRmQMOOJGlGlmrpifcD0UIJtOd)
	- **The Output of Batch Workflows**:
		    - Where does batch processing fit in? 
		        - It is not transaction processing, nor is it analytics. It is **closer to analytics**, in that a batch process typically scans over large portions of an input dataset.
		    - The output of a batch process is often not a report, but some other kind of structure.
		    - **Building search indexes**: (documents in, indexes out.)
		        - Google’s original use of MapReduce was to build indexes for its search engine, which was implemented as a workflow of 5 to 10 MapReduce jobs.  (e.g. still used today by **Lucene/Solr**)
		        - Recall **full-text search index**: it is a file (the term dictionary) in which you can efficiently look up a particular keyword and find the list of all the document IDs containing that keyword (inverted index). 
		    - **Philosophy of batch process outputs**:
		        - In the process, the **input is left unchanged**, any previous **output is completely replaced** with the new output, and there are no other side effects.
	- **Comparing Hadoop to Distributed Databases**:
		    - **Hadoop** is somewhat like a distributed version of Unix, where **HDFS** is the filesystem and MapReduce is a quirky implementation of a Unix process(which happens to always run the sort utility between the map phase and the reduce phase). 
		    - **MapReduce** and a **Distributed Filesystem** provides something much more like a general-purpose operating system that can run arbitrary programs.
		    - **Diversity of storage**:
		        - **Databases require you to structure** data according to a particular model (e.g., relational or documents), 
		            - whereas files in a distributed filesystem are just byte sequences, which can be written using any data model and encoding. 
		        - Collecting data in its raw form, and worrying about schema design later, allows the data collection to be speeded up (a concept sometimes known as a “**data lake**” or “**enterprise data hub**” ). 
		            - Aka. **sushi principle**: “**raw (data) is better**” 
		        - Indiscriminate data dumping shifts the burden of interpreting the data 
		        - Data modeling still happens, but it is in a separate step, decoupled from the data collection. 
		            - This decoupling is possible because a distributed filesystem supports data encoded in any format.
		    - **Designing for frequent faults**:
		        - When comparing MapReduce to MPP(massively parallel processing) databases, two more differences in design approach stand out: **the handling of faults** and **the use of memory and disk**.
		            - **MPP** databases prefer to keep as much data as possible in memory (e.g., using hash joins) to avoid the cost of reading from disk.
		            - **MapReduce** is very eager to write data to disk, partly for fault tolerance, and partly on the assumption that the dataset will be too big to fit in memory anyway.
		        - It’s not because the hardware is particularly unreliable, it’s because the freedom to arbitrarily terminate processes enables better resource utilization in a computing cluster. ( google used to use the node which could have imp task , that will stop these map reduce job, so made this fault tolerant)
		            - Among open source cluster schedulers, preemption is less widely used. (e.g. YARN’s CapacityScheduler)
		
	#### Beyond MapReduce
	
	- Implementing a complex processing job using the raw MapReduce APIs is actually quite hard and laborious—for instance, you would need to implement any join algorithms from scratch. 
	- In response to the difficulty of using MapReduce directly, various higher-level programming models (Pig, Hive, Cascading, Crunch) were created as abstractions on top of MapReduce. 
	- **Materialization of Intermediate State**:
	    - Publishing data to a well-known location in the distributed file system allows loose coupling so that jobs don’t need to know who is producing their input or consuming their output.
	    - **Intermediate state**: a means of passing data from one job to the next. (not shared between job or team) 
	        - The process of writing out this intermediate state to files is called **materialization**. 
	        - In contrast: **Pipes** do not fully materialize the intermediate state, but instead **stream** the output to the input incrementally, using only a small in-memory buffer.
	    - MapReduce’s approach of **fully materializing intermediate state** has downsides compared to Unix pipes: 
	        - A MapReduce job can only start when all tasks in the preceding jobs (that generate its inputs) have completed;
	        - Mappers are often redundant: they just read back the same file that was just written by a reducer, and prepare it for the next stage of partitioning and sorting.
	        - Storing intermediate state in a distributed file system means those files are replicated across several nodes, which is often overkill for such temporary data.
	    - **Dataflow engines**:
	        - New execution engine created to solve previou problems. E.g. **Spark**, Tez, and Flink. 
	            
	    - **Fault tolerance**:
	        - An advantage of fully materializing intermediate state to a distributed file system is that it is durable, which makes fault tolerance fairly easy.
	            - if a task fails, it can just be restarted on another machine and read the same input again from the filesystem.
	        - **Spark, Flink, and Tez** avoid writing intermediate state to HDFS, so they take a different approach to tolerating faults: 
	            - if a machine fails and the intermediate state on that machine is lost, it is **recomputed** from other data that is still available.
	        - Recovering from faults by recomputing data is not always the right answer: 
	            - if the intermediate data is much smaller than the source data, or if the computation is very CPU-intensive, it is probably cheaper to materialize the intermediate data to files than to recompute it.
	    - **Discussion of materialization**:
	        - **Flink** especially is built around the idea of pipelined execution: that is, incrementally passing the output of an operator to other operators, and not waiting for the input to be complete before starting to process it.
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
	




