# High-Level Design (HLD) Interview Template (45 Minutes)

## 1. Requirements Gathering (5-10 minutes max)
- check time in middle , shouldn't invest more than 10 mins at max -> its a red red flag.

### 1.1 Functional requirements.
  - Clear the requirements and write it as functional requirements. 
  - clarify the doubts if have any, about requirements. 
### 1.2 Non functional requirements:
  - Discuss about non functional requirements. ( latency, availability, scalability, consistency,etc.).
  - Ask if they have any specific non functional requirements. 
### 1.3 Estimations: (ask -> let me know if you are not interested in this part, may be later on or at time of design we can come on this. ) 
  - Come on to esitmations. 
  - ask about the DAU. 
  - find the read qps.
  - ask about read: write ratio. 
  - find the write qps. 
  - ( 1 year = 400, 1 day = 10^5, 10^3 (K)(KB)(milli second (-ve)) , 10^6(M)(MB)(micro second (-ve)), .. B(GB)(nano), 10^3*B , (TB), ps)
  - Storage -> each row -> 500B, image -> 1KB -> find the 1 year.
  - usually storage ( MT rds) -> ~5OTB -> greater than this -> talk about partition , sharding.
  - find read heavy or write heavy.
  - find peak.
  - notice any specific constraints. ( like some system very less write, some have concurrent etc. )
  - if single digit latency , 5 thread -> 1k qps -> 1k8
  - kafka -> aws msk -> avg -> 50k-200k qps per broker -> as per instance. ( since 50MB / s -> 1 kb assume.)
  - redis -> micro second, sql -> single digit read, double digit write , cassandra -> read sql/s , write -> sql read.
  - transaction limit assume 1-10k at max.

## 2. System Design (10-15 minutes) ( sum -> 15 - 20) (20mins at max -> so total at max 25mins -> red red flag.)

### 2.1 High-Level Architecture

- Ask interview -> i usually draw boxes and do the api and db design. and then later on go into nfr like scaling, other details, let me know if any want to do have some other way. -> some can say first do api design and db design. then hld but mostly don't say any thing so continue. 
- Think about the solutions with past problem , Draw the boxes of different components, something new just draw boxes with some responsibility and later figure out in deep dive. 
- they can be different servers, gateway, database, kafka, elastic search, s3, graphql, client, CDN, redis, aws lambda, consumer, stream, batch ( only if multi country expensive)  etc just black boxes like database can be box as of now. b
- Draw the arrows -> to have high level idea of arrows. tell the flow one by one and make the arrow. 
- https, websocket, webrtc, sse, long polling, async, event driven, sync, grpc , graphql etc. 

- check the time.

### 2.2 API Design + database schema ( can also discuss Database type or may be later.).
- Go flow by flow and complete one flow -> then jump on the other. for example one flow is user service flow, other write flow, read flow etc. ( go by the highest priority to lowest )
- write the must have apis 
    - write its POST/GET etc. , 
    - api/v1/{domain}/{routename} , 
    - if GET write in query, POST write its json, 
    - write the request and response.
    - if require to have pagination add that.
    - if require to have cursor add that. 
    - client -> graphql
    - server to server -> grpc

- write the db schema of that
    - in mind think about the db, may be you can tell to interview thinking about this -> a/c to read heavy ( mysql , kv pair) or write heavy ( cassandra) with these questions -> what is data access pattern, what is query pattern , or simply key value getting, or m:m relation graph, olap ( cassandra, snowflake datalake), vector db for embeddings, 
    - figure out the entities and write schema as per db.
    - write the structure, fields -> default -> id, createdtime, updatedtime. createdby, updatedby, globalcontextid , foreignId etc. 
    - write the type as well , but write the primary key and foreign key detail. 
    - any specific contraint of field.
    - do as much normalised as possible. 
    - create as much as table as possible or as much as entities.
    - create 1:m , m:1, m:m relation drawing. 
    - though below one don't specificly showed primary, foreignkey or type of field. but this is for breaking.
    ![alt_image](./image.png)
    - this is mostly require for mysql , or document db, but generally work for all. 
    - Now ask interview do i have to invest more times on api and db design for other flows, or should i move on to details of flows and non functional requirements and scalabiltity. since i have less time. ? 
        - if yes -> do same for other flows. 
        - if yes-no -> just tell about the apis on high level or via text and db type and why. 
        - else move to next part. 
    - check the time don't invest more than ( 20mins at max -> so total at max 25mins -> red red flag.)

## 3. Deep Dive (15-20 minutes) (20-40) -> no limit until interviewer wants. ( 40mins-45 mins)

- This part basically , solve problem, let say if any algorithm remain like order book in stock exchange, window algo in rate limitings, consistency, orderering of events. trade offs, scaling, specific nfr etc done let's summarise and cover previous part or have mroe time discuss about the moniroting, logging, support system etc. Trade is ver very important -> means tell what are the option we have and which one you are choosing by doing some trade offs, error handling is ver very important, edge case handling ver very important, like consistency then how to make sure loss etc. 
- Ask interview which part they want to cover first and go in deep dive, -> if yes go into that. else below: 
- Telling the things below in order do them one by one , deep dive one and move to other. 
- Think about any end to end flow,  which you think important , then Do the dfs. Let say you have select some flow. Then go to client add details there , go the next component of flow add details there .... to end. Details are the below order wise.

### 3.1 Order wide details -> could be possible for some component couldn't be possible for other. 

- What is this component. 

### 3.1 Scalability
- Discuss horizontal/vertical scaling strategies
- Identify potential bottlenecks
- Propose solutions for scaling

### 3.2 Availability & Reliability
- Discuss redundancy and failover
- Identify single points of failure
- Propose solutions for high availability

### 3.3 Data Consistency
- Discuss consistency requirements
- Choose appropriate consistency model
- Handle edge cases and race conditions


## 3.4 Trade-offs & Optimization (5 minutes)
- Discuss design decisions and their trade-offs
- Identify potential optimizations
- Consider cost implications


## 4. Wrap-up (5 minutes)
- Summarize the design
- Address any remaining concerns
- Discuss potential future improvements


