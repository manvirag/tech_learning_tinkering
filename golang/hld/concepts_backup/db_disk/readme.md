****1. mysql.****
- store data in row-wise.
- if we don't have index, require linear time to search the rows.
- it uses b-tree  to do indexing ( B-tree indexes are a common data structure for organizing and accessing data efficiently on disk) , each leaf node of b-tree has page of row ( 1 page have multiple rows depending upon side of page. ) [ Note sure, but can imagine like we have a file that is b-tree index , it have the address of disk where does its leaf node data exist. Now at any search it will find it and fetch it and save in ram kind of, in update it will rebalance itself  and save its file to disk and help to find the address of the block to be changed and let say other function with updated it. ]
- b-tree make the read to log(n) for that particular field on which index is made.
- on the other hand it make write slower, since it has to rebalance the tree after insertion. So too much index is n't good for write
- though can scale with sharding -> but with limitations since join. but vertical scale and replication there alwasy.
- transactions is a plus point. ( very expensive in case of distributed )
#### More
- https://vipulpachauri12.medium.com/how-mysql-stores-data-in-disk-ee51a9e81c39#:~:text=MySQL%20stores%20each%20database%20(also,the%20table%20definition%20in%20MyTable.
- https://www.youtube.com/watch?v=DbxddGtHl70
- https://www.lucavall.in/blog/how-databases-store-and-retrieve-data-with-b-trees

****2. mongodb ( document ).****

- also store data row-wise, but not rigid structural , its store as json.
- means instead of table , it does replication in the json.
- means it can be scaled horizontally more freely. since no joins. ( mysql you can't put table in different node if they have join with some table, the ony way is to keep in such way that no need node wise join situation)
- it also uses b-tree same like the mysql
- mongodb keep a global unique id for each document , that's how here the joins work by keeping that unique document id in first document.
- it also has acidic transactions.
- real time analytics.
- aws documentdb, (aws dynamodb also provide json support), mongodb, couchbase etc.

#### More
- https://medium.com/nerd-for-tech/all-basics-of-mongodb-in-10-minutes-baddaf6b6625
- https://www.youtube.com/watch?v=ONzdr4SmOng


- Confused ? which to use between these -> when your priority is only scalability and flexibility. strong acidic and schema and secure go for sql.

- Or in short you can also skip this db in interview , no having very high advantage. So always prefix sql here ( saying only with compare to mongodb not other nosql ). So now we will compare other db with sql directly.

![img.png](img.png)

****3. cassandra, cosmosdb ( column ).****
- index based on LSM tree and SSTables.
- Cassandra works really well if you want to write and store a large amount of data in a distributed system, don’t care much about ACID with good performance.
- 
  
- has to see, why write heavy , this index, how does it store a particular partition in a node.
- still have confusion its whether key value store or column store.
#### More details
- https://kunal14053.medium.com/about-cassandra-1423223945b3
- https://www.youtube.com/watch?v=xynXjChKkJc

****7. dynamodb, cosmosdb, Redis ( key store ), cassandra.****

- Only key-value store like hashing.
- We can query with key only to get data , not any parameter of value.
- value can have different data types.
- mostly use for caching.

****4. graph ( graph ).****

- Niche

****5. time series.****

- Niche

****6. search engine.****

- Niche - searching
- Full-text Search Complexity: Searching vast amounts of text involves considering factors like relevance, word frequency, user preferences, and handling typos, which poses significant challenges.
- Elasticsearch Solution: Elasticsearch is a specialized database designed to efficiently handle these challenges, offering tools for indexing, querying, and scoring text data for rapid and accurate searches.
- https://betterprogramming.pub/system-design-series-elasticsearch-architecting-for-search-5d5e61360463
