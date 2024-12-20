****1. mysql.****
- SQL is also very scalable if you don’t care about relations or foreign key constraints
- store data in row-wise.
- if we don't have index, require linear time to search the rows.
- it uses b-tree  to do indexing ( B-tree indexes are a common data structure for organizing and accessing data efficiently on disk) , each leaf node of b-tree has page of row ( 1 page have multiple rows depending upon side of page. ) [ Note sure, but can imagine like we have a file that is b-tree index , it have the address of disk where does its leaf node data exist. Now at any search it will find it and fetch it and save in ram kind of, in update it will rebalance itself  and save its file to disk and help to find the address of the block to be changed and let say other function with updated it. ]
- b-tree make the read to log(n) for that particular field on which index is made. leaf node is page of it.
- on the other hand it make write slower, since it has to rebalance the tree after insertion. So too much index is n't good for write
- though can scale with sharding -> but with limitations since join. but vertical scale and replication there alwasy.
- transactions is a plus point. ( very expensive in case of distributed )
#### More
- https://vipulpachauri12.medium.com/how-mysql-stores-data-in-disk-ee51a9e81c39#:~:text=MySQL%20stores%20each%20database%20(also,the%20table%20definition%20in%20MyTable.
- https://www.youtube.com/watch?v=DbxddGtHl70
- https://www.lucavall.in/blog/how-databases-store-and-retrieve-data-with-b-trees
-

****2. mongodb ( document ).****

- also store data row-wise, but not rigid structural , its store as json.
- means instead of table , it does replication in the json.
- means it can be scaled horizontally more freely. since no joins. ( mysql you can't put table in different node if they have join with some table, the ony way is to keep in such way that no need node wise join situation)
- it also uses b-tree same like the mysql
- mongodb keep a global unique id for each document , that's how here the joins work by keeping that unique document id in first document.
- it also has acidic transactions.
- real time analytics.
- single leader and multiple read replica approach.
- aws documentdb, (aws dynamodb also provide json support), mongodb, couchbase etc.

#### More
- https://medium.com/nerd-for-tech/all-basics-of-mongodb-in-10-minutes-baddaf6b6625
- https://www.youtube.com/watch?v=ONzdr4SmOng


- Confused ? which to use between these -> when your priority is only scalability and flexibility. strong acidic and schema and secure go for sql.

- Or in short you can also skip this db in interview , no having very high advantage. So always prefix sql here ( saying only with compare to mongodb not other nosql ). So now we will compare other db with sql directly.
- or it is used when we have not structure data but want quality kind of similar to sql.

![img.png](img.png)

****3. cassandra, cosmosdb ( column ).****
- index based on LSM tree ( Log-Structured Merge-Tree [Link](https://www.youtube.com/watch?v=MbwmMCu9ltg&list=PLwrbo0b_XxA8BaxKRHuGHAQsBrmhYBsh1&index=4) ) and SSTables ( Sorted String Tables ). ( optimised for fast write )
- This is not real colomn oriented ( DDIA ). Yes its true. It is wide-colum store and colmn store is different form this. https://discuss.educative.io/t/is-cassandra-really-a-column-based-database/39536
- Quote: "Wide-column stores such as Bigtable and Apache Cassandra are not column stores in the original sense of the term, since their two-level structures do not use a columnar data layout. In genuine column stores, a columnar data layout is adopted such that each column is stored separately on disk. Wide-column stores do often support the notion of column families that are stored separately. However, each such column family typically contains multiple columns that are used together, similar to traditional relational database tables. Within a given column family, all data is stored in a row-by-row fashion, such that the columns for a given row are stored together, rather than each column being stored separately."
  - On high level it save data in tree like structure which is called memtable ( balanced binary search tree ) in memory, then flash it to immutable( no update it will be created with new value) table which is called sstable. internally it uses merge sort etc at each level , that complete architecture is called LSM tree.
  - so write is simply write in tree then it will be flushed.
  - but in reading it will first check in memory lsm tree, then most recent lsm tree then second most recent, that's why reading is having little more complexity.
  - reading is optimised with bloom filter.
  - b-tree is random access, and in disk its not considered efficient , sequential access considered efficent, lsm tree follow this. while reading it has in memory indexing. [link](https://medium.com/@qiaojialinwolf/conceptions-you-should-know-about-hard-disk-997545b316e7)
  - compaction run in background
- Cassandra works really well if you want to write and store a large amount of data in a distributed system, don’t care much about ACID with good performance.
- Cassandra is an early NoSQL database with a hybrid design between a tabular and key-value store.
- Cassandra stores data as key-value stores ( then why no use key store ? -> key store is more simpler mostly use in caching. In this it shows data in tabular form allow to have more complex query as well based on the coloum value. ). It allows you to define tables with rows and columns, but the tabular structure isn’t used in actual storage. Instead, it uses the wide column-oriented database model ( still confusion whether its key value store or coloum oriented. As of now assume its assigned towards coloum oriented ) , so each row in the table can have a different set of columns.
- multi master apporach
- for e.g. query: 
```cassandraql
SELECT * FROM Orders
WHERE customerId = 'customer123'
AND orderDate >= '2024-01-01' AND orderDate <= '2024-01-31';
```
- this db is used in system design mostly as name of nosql. ( other are also used , but as compare to mongodb ).
- used when high write volume and high write , high availability , distributed, but no consistency.  Cassandra’s true use case, in a single word, is scale.

![img_6.png](img_6.png)
![img_1.png](img_1.png)
![img_7.png](img_7.png)
![img_8.png](img_8.png)
#### More details
- https://www.youtube.com/watch?v=I6jB0nM9SKU
- https://medium.com/@qiaojialinwolf/lsm-tree-the-underlying-design-of-nosql-database-cf30218e82f3
- https://medium.com/@qiaojialinwolf/conceptions-you-should-know-about-hard-disk-997545b316e7
- https://aws.amazon.com/compare/the-difference-between-cassandra-and-mongodb/#:~:text=Cassandra%20stores%20data%20as%20key%2Dvalue%20stores.
- actual coloum oriented, although it also use the lsm tree -> https://www.youtube.com/watch?v=Zt7rqtJ3uWA  ( no sure cassandra is 100% like this or not )
- ![img_2.png](img_2.png)




****7. cosmosdb, Redis, dynamodb ( key store ).****

- Only key-value store like hashing.
- We can query with key only to get data , not any parameter of value.
- value can have different data types.
- mostly use for caching or config store with key kind of thing where no need query on basis of value field etc.
- simple as compare to other no sql.
- it also used multi master replication. ( specifically for dynamodb not sure for other) [ alex xu key value -> cassandra and dynamo] 

****4. graph ( graph ).****

- Niche
- Useful when thing is represented by graph i.e node and edges.
```json
(:User {id: "user1", name: "Alice"}) -[:FRIENDS_WITH]-> (:User {id: "user2", name: "Bob"})
(:User {id: "user1", name: "Alice"}) -[:FRIENDS_WITH]-> (:User {id: "user3", name: "Charlie"})
(:User {id: "user1", name: "Alice"}) -[:FRIENDS_WITH]-> (:User {id: "user4", name: "David"})
(:User {id: "user2", name: "Bob"}) -[:FRIENDS_WITH]-> (:User {id: "user3", name: "Charlie"})
(:User {id: "user3", name: "Charlie"}) -[:FRIENDS_WITH]-> (:User {id: "user4", name: "David"})

```
- can be used in social network.
- why not relational database ? , because of time complexity and scaling. let consider below example ( log(E) + Log(N) => find all friends + find name of those friends) 
![img_3.png](img_3.png)
- why not existing nosql ? , ( Log(N) + duplicate data )
![img_4.png](img_4.png)
- In neo4j it store the address directly to can access to disk directly
![img_5.png](img_5.png)

****- More details****
- https://www.youtube.com/watch?v=Sdw_D-Gllac&t=94s


****5. time series.[Don't know internal]****

- Niche
- Time series databases are specifically designed to handle time-stamped or time-series data efficiently, optimizing for high write throughput and unique query patterns based on time.
- Use column oriented storage.
- query, indexing on basis of timestamp.
- LSM tree is used in this, ( though TSM tree is used which internally use LSM ).
- it has chunks internally and chunk have its own indexing .
- existing nosql and rds is not much optimised for handling timeseries data workload. ( why not sure -> as of now assume that tsd are made in every step in consideration of time series data )
```json
| Timestamp           | Temperature (°C) |
|---------------------|------------------|
| 2024-03-01 00:00:00 | 15               |
| 2024-03-01 01:00:00 | 14               |
| 2024-03-01 02:00:00 | 13               |
| ...                 | ...              |
| 2024-03-02 00:00:00 | 14               |
| 2024-03-02 01:00:00 | 13               |
| ...                 | ...              |
| 2024-03-03 00:00:00 | 16               |
| ...                 | ...              |

```
![img_9.png](img_9.png)
- https://www.timescale.com/blog/time-series-data-why-and-how-to-use-a-relational-database-instead-of-nosql-d0cd6975e87c/  [ ek no. ]
- https://www.youtube.com/watch?v=QVa8k36w0Ig&list=PLwrbo0b_XxA8BaxKRHuGHAQsBrmhYBsh1&index=7

****6. search engine.****

- Niche - searching on the text. ( Performing a full-text search would mean that any user can search for something like “java” or “learn programming,” and you need to figure out all the blog posts where these words appear within a few milliseconds )
- Elasticsearch isn't a database, not in the same way that, for example, MySQL is. (but data can be read ,write, update in this for searching . )
- Elasticsearch is a #JSON document repository that is based on the #Apache #Lucene search engine
- Or in simpler language, Lucene is libraray which is used to do all these operations, but that is single node functionality , so elastic search is only the orchestrator , which advantage it to distributed and multiple nodes, but internally use the lucene itself. 
```json
{
  "_id": "9a91473c-522e-4174-bf7f-f55293b8e526",
  "post_title": "Learning about Elasticsearch",
  "author_name": "Sanil Khurana",
  .....
}
```
- Lucene works its magic by indexing documents.
- distributed on multiple node, parallel searching, inverted index.  The core of high speed is derived from parallel computing and inverted index.
- inverted index: An inverted index lists every unique word that appears in any document and identifies all of the documents each word occurs in.
![img_11.png](img_11.png)
![img_10.png](img_10.png)



- Deep Dive: 
![alt text](image.png)
  - The important concepts of Elasticsearch from a client perspective are documents, indices, mappings, and fields.
  - Documents are the individual units of data that you're searching over.
  Book Doc:
  ```
    {
      "id": "XYZ123",
      "title": "The Great Gatsby",
      "author": "F. Scott Fitzgerald",
      "price": 10.99,
      "createdAt": "2024-01-01T00:00:00.000Z"
    }
  ```
  - Index -> An index is a collection of documents. Each document is associated with a unique ID and a set of fields.
  -  a mapping is the schema of the index. It defines the fields that the index will have, the data type of each field, and any other properties like how the field is processed and indexed. You can put whatever data you want in the document, but the mapping determines which fields are searchable and what type of data they contain. So in sort document can have a lot field you want to search on some field you add them only with their type (so that can do some operation according to them , later).
  - Below is sample index of book
  ```
      // PUT /books/_mapping
            {
              "properties": {
                "title": { "type": "text" },
                "author": { "type": "keyword" },
                "description": { "type": "text" },
                "price": { "type": "float" },
                "publish_date": { "type": "date" },
                "categories": { "type": "keyword" },
                "reviews": {
                  "type": "nested",
                  "properties": {
                    "user": { "type": "keyword" },
                    "rating": { "type": "integer" },
                    "comment": { "type": "text" }
                  }
                }
              }
            }
  ```
  - We can now add the documents in these index and update, delete etc.
  ```
      // POST /books/_doc
          {
            "title": "The Great Gatsby",
            "author": "F. Scott Fitzgerald",
            "description": "A novel about the American Dream in the Jazz Age",
            "price": 9.99,
            "publish_date": "1925-04-10",
            "categories": ["Classic", "Fiction"],
            "reviews": [
              {
                "user": "reader1",
                "rating": 5,
                "comment": "A masterpiece!"
              },
              {
                "user": "reader2",
                "rating": 4,
                "comment": "Beautifully written, but a bit sad."
              }
            ]
          }
  ```
  - Now query time.
  - ES provide a lots of different type of query on that index. 
  - Like Range, Match , Not match ,fussy, semantic, sort etc.
  - For our usecase (search_filter_sorting problem). 
  - We can use match query in filter fields, fussy search in name , sorting with rating etc.
  - This can return data in pagination form and can fetch any page see doc for syntax.


- Internals

  - On high level lucene works on inverted index, and it provide a lot function with overlapping on it.
  - We specially interested in elastic search internal working. ( or how it scale lucene or make it distributed system )
  - now that you have a basic understanding of how you might use Elasticsearch as a client. Elasticsearch can be thought of as a high-level orchestration framework for Apache Lucene.
  - Elasticsearch handles the distributed systems aspects: cluster coordination, APIs, aggregations, and real-time capabilities while the "heart" of the search functionality is handled by Lucene.
  - Elasticsearch is a distributed search engine. When you spin up an Elasticsearch cluster, you're actually spinning up multiple nodes. Nodes can be of 5 types which are specified when the instance is started.
  - Master Node is responsible for coordinating the cluster.It's the only node that can perform cluster-level operations like adding or removing nodes, and creating or deleting indices. Think of it like the "admin".
  - Data Node is responsible for storing the data. It's where your data is actually stored. You'll have lots of these in a big cluster.
  - Coordinating Node is responsible for coordinating the search requests across the cluster. It's the node that receives the search request from the client and sends it to the appropriate nodes. This is the frontend for your cluster.
  - Ingest Node is responsible for data ingestion. It's where your data is transformed and prepared for indexing.
  - Machine Learning Node is responsible for machine learning tasks.
  ![alt text](image-1.png)
  

- https://betterprogramming.pub/system-design-series-elasticsearch-architecting-for-search-5d5e61360463
- https://www.reddit.com/r/rails/comments/66q413/why_is_elastic_search_faster_at_querying_compared/
- https://medium.com/analytics-vidhya/how-elasticsearch-search-so-fast-248630b70ba4
- https://www.hellointerview.com/learn/system-design/deep-dives/elasticsearch

****7. Vector database****

- Niche
- for ml embedding, not going in deep of this currently



#### Reference:

- https://www.youtube.com/watch?v=cODCpXtPHbQ&t=38s
- https://www.codekarle.com/system-design/Database-system-design.html
- if its structure and not need acid can use any.
- when data is not structure its json also need query , then can use document db like mongodb , couchbase.
- and if you don't have much complex query 
![alt_text](./img_12.png)
