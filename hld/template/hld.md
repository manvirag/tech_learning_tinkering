# High-Level Design (HLD) Interview Template (45 Minutes)

## 1. Requirements Gathering (5-10 minutes max)
- check time in middle , shouldn't invest more than 10 mins at max -> its a red red flag.

### 1.1 Functional requirements.
  - Clear the requirements and write it as functional requirements. 
  - clarify the doubts if have any, about requirements. 
  - write some out of scope requirements -> so that interview flag else they might cheat later on -> by saying i did say this or that. 
  - if require ask about the web page flow and ui by customer, like user come there will be a search tab it will click and write up the text etc. -> this will give idea different apis and different services.
### 1.2 Non functional requirements:
  - Discuss about non functional requirements. ( latency, availability, scalability, consistency,etc.).
  - Ask if they have any specific non functional requirements. 
  - Cap theorem in distributed system. 
### 1.3 Estimations: (ask -> let me know if you are not interested in this part, may be later on or at time of design we can come on this. Even they say not much invest time in calculations . ) 
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
  - mutex lock/unblock -> 100ns

## 2. System Design (10-15 minutes) ( sum -> 15 - 20) (20mins at max -> so total at max 25mins -> red red flag.)

### 2.1 High-Level Architecture

- Ask interview -> i usually draw boxes and do the api and db design. and then later on go into nfr like scaling, other details, let me know if any want to do have some other way. -> some can say first do api design and db design. then hld but mostly don't say any thing so continue. **Mostly prefer to have entities and  then most suitable api first then other things later on like hld db schema deep dive.** 
- Think about the solutions with past problem , Draw the boxes of different components, something new just draw boxes with some responsibility and later figure out in deep dive. 
- they can be different servers, gateway, database, kafka, elastic search, s3, graphql, client, CDN, redis, aws lambda, consumer, stream, batch ( only if multi country expensive)  etc just black boxes like database can be box as of now. b
- Draw the arrows -> to have high level idea of arrows. tell the flow one by one and make the arrow. 
- https, websocket, webrtc, sse, long polling, sse,  async, event driven, sync, grpc , graphql etc. 

- check the time.

### 2.2 API Design + database schema ( can also discuss Database type or may be later.).
- Go flow by flow and complete one flow -> then jump on the other. for example one flow is user service flow, other write flow, read flow etc. ( go by the highest priority to lowest )
- write the must have apis 
    - write its POST/GET etc. , 
    - api/v1/{domain}/{routename} , 
    - if GET write in query, POST write its json, 
    - write the request and response.
    - if require to have pagination add that.
    - if require to have cursor add that. ( always better to use cursor. )
    - talk about headers ( sone pe suhata.) -> token, jwt, saml , authid come there.
    - client -> graphql
    - server to server -> grpc
    - talk about different status codes -> 
        - 200: success , all good
        - 301: redirect to new url -> send this code and url.
        - 302: temporary redirect ->  mean we will come back on this like A/B testing etc.
        - 400: bad request, request is bad validation failed. 
        - 401: Unauthorized -> user needs to log in or authenticate to access the page.
        - 403: Forbidden -> server understood the request but won't fulfill it, due to a lack of permissions.
        - 404: Not Found -> server couldn't find the page requested ( might be incorrect url)
        - 500: internal server error , server broke.
        - 502: Bad Gateway -> This happens when one server, acting as a gateway or proxy, receives a faulty response from an upstream server.
        - 503: server unavaible -> not able to take request may be undeploy or down. 
        - 101: switching protocol.
    - **request meta , token, jwt token, saml**  -> for **authentication** , **header**  ( DIFF ), for websocket -> headers upgrade , key and version etc. -> res 101 switching protocol etc.
    ![api_design](./apidesign.png)

- write the db schema of that 
    - in mind think about the db, may be you can tell to interview thinking about this -> a/c to read heavy ( mysql , kv pair) or write heavy ( cassandra) with these questions -> what is data access pattern, what is query pattern, btree, lsm sst, columnar etc. -> more in db type deep dive. , or simply key value getting, or m:m relation graph, olap ( cassandra, snowflake datalake), vector db for embeddings, 
    - figure out the entities and write schema as per db.
    - write the structure, fields -> default -> id, createdtime, updatedtime. createdby, updatedby, globalcontextid , foreignId etc. 
    - write the type as well , but write the primary key and foreign key detail. 
    - mysql -> CHAR ( 0-255B), VARCHAR(0-64KB) -> both (can contain letters, numbers, and special characters) || TEXT -> string (0-64kB) , LONGTEXT(0-4GB) , TIMESTAMP (19731230153000 sample) , INT(4B), BIGINT(8B), BOOLEAN(1B), ENUM(variable) ... so on.
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
    ![alt_image](./dbschema.png)

## 3. Deep Dive (15-20 minutes) (20-40) -> no limit until interviewer wants. ( 40mins-45 mins)

- This part basically , solve problem, let say if any algorithm remain like order book in stock exchange, window algo in rate limitings, consistency, orderering of events. trade offs, scaling, specific nfr etc done let's summarise and cover previous part or have mroe time discuss about the moniroting, logging, support system etc. Trade is ver very important -> means tell what are the option we have and which one you are choosing by doing some trade offs, error handling is ver very important, edge case handling ver very important, like consistency then how to make sure loss etc. 
- Ask interview which part they want to cover first and go in deep dive, -> if yes go into that. else below: 
- Telling the things below in order do them one by one , deep dive one and move to other. 
- Think about any end to end flow,  which you think important , then Do the dfs. Let say you have select some flow. Then go to client add details there , go the next component of flow add details there .... to end. Details are the below order wise.
- Tell the same to the interviewer

### 3.1 Order wide details -> could be possible for some component couldn't be possible for other. 

- When deep diving into any component, follow this systematic approach:

- Remeber these words -> Scalability, bottlenecks, Availability(after crash), Reliability(no loss), SPOF, Retry , error handling, consistency, concurrency, race condition, transaction, distributed complexity, backups, reconsilation, idempotency

- Let say first talk about client. ( like what would be happening on UI.)
    - say like UI will come on this page client on this. 
    - these are static pages -> via s3 or server -> CDN if high static caching backed by s3 or having multip customer.
    - say this will call to graphql. Not need to add much details. 
    - some time client do some work that can tell like file upload to s3 and send link to server etc.
![client](./client.png)
    

- [Gate way](https://www.hellointerview.com/learn/system-design/deep-dives/api-gateway) -> let talk about this -> this is kind of statting point like app-gatekeeper. 
    - use it when you have a microservices architecture and don't use it when you have a simple client-server architecture.
    - tell particular route will get hit -> for e.g. /graphql , or may be direct api. 
    - it will send this route to particular server. 
    - say k8s can be used here -> like gate pay or load balancer high to the ingress of k8s then it will send to particular requests. 
    - tell bit about graphql -> why to use, request response, bff, dataloader etc. 
    - now this will call the backend server either via grpc or api.
    - also tell here like industry wide use k8s cluster so this need not to go to again via load balancer
    - scale graphql horizontally -> can tell -> k8s pods , replica count, hpa , service for load balancing etc can be used here
    - make secure by rate limiting, authentication (middle ware work)
    - request/response transformation -> header manipulation, payload modification
![gateway](./apigateway.png)

- For any general purpose server. ( sync called, stateless ):
    - is it https vs grpc vs websocket decide. etc. 
    - ask your self what's the responsibility of this server and wrote it
    - stateless -> scale horizontally depending upon qps, add loadbalancer ( service in k8s ) ( nginx (ingress) -> service (LB) -> pods (horizontal server)). -> available
    - what about hotspot  -> stateless shouldn't be case
    - Consistency and error handling , exponential retry, or dlq
    - or may be server level rate limiting, authentication, timeout
    - health checks endpoints, monitoring , logging and alertings
    - deployment strategy -> blue-green, canary, rolling updates
    - Ready-made solutions: AWS Lambda, Kubernetes Deployments, Google Cloud Run
![](./sync_stateless_server.png)

-  for any general purpose consumer ( consuming async events):
    - consuming events via queue/kafka. 
    - write responsibility.
    - message processing -> batch vs single, error handling ( like retry or not.)
    - some also use retry queue for retrying events. ( change order -> in retry batch duplicacy.)
    - dead letter queue -> failed message handling
    - scaling strategy -> consumer group scaling -> equal to parition.
    - message ordering -> partition key -> failure retry -> but idempotency at target.
    - idempotency -> deduplication, message replay.
    - discuss about push vs pull. 
    - monitoring -> consumer lag, processing rate, logger
    - failure handling to target -> retry policies
    - Ready-made solutions: aws lambda, k8s

-  for any general purpose server/consumer/processor( stateful ):

-  stream/batch, aggregation windowing algorithms ? in batch, issues in clock syn. ( stateful )

-  for any general purpose cronjob ( running at interval with some input event ):

-  queue for any async or decoupling purpose and its complicacy: push pull

-  for different types of databases and its challences, transactions, distribute tx: ( altogether different and vast thing will talk about gen.) 

-  for any blob storage usecase: 

-  what about the file storage ? like logs etc:

-  talk about cache for any general purpose, its data structure:

-  talk about CDN: 

-  monitoring , pagerduty, only, incident cases (case our system broke and need someone to intervene and how can be restore state):

-  what about load balancers ?

-  what to do read / write amplification.

-  what to do when need strict ordering things ? and indempotency and replay kind of thing

-  clock synchronisation complicacy in distributed system.

-  what about hashing mechanism ? 

-  what the heck about service discovery ? 

-  details about decoding/encoding/compression/trancoding/base64/ascii/utf-8(x)/storage-for-char/  ? -> for cost and storage optimization

-  advance mmap event source mappingn low latency.

-  ledger reconsilation in finance system
    - double-entry accounting
    - transaction logs
    - audit trails
    - consistency checks
    - rollback mechanisms
    - reconciliation jobs
    - financial reporting
    - Ready-made solutions: AWS QLDB, Hyperledger Fabric, Stellar

## 4. Wrap-up (5 minutes)
- Summarize the design
- Address any remaining concerns
- Discuss potential future improvements


