/*

Also do wtih concurrency. 

................................................................................................................................

Problem 1 -> similar to kafka

Functional Requirements
Publishers call publish(String topic, String message) to send messages asynchronously to a topic.
Subscribers call subscribe(String topic) to register for a topic and receive messages via polling (getMessages()) or callbacks.
Support unsubscribe(String topic) to stop receiving from a topic.
Each topic maintains a bounded queue (e.g., capacity 1000) for messages; handle overflow by dropping oldest or blocking (discuss).
​
................................................................................................................................


Problem 2 ( sqs)

Design an inmemory pull based queue library
where multiple publishers and consumers can publish/read 
the messages from the shared queue. 
And each message can have optional TTL .

*/



// problem 1 kafka