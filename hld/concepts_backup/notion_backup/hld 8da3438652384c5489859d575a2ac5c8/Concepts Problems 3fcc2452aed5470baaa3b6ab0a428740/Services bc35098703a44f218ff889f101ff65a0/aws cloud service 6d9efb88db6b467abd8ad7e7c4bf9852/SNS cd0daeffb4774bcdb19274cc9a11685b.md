# SNS

Sqs is like a db with styling quality, it has data, single data for a single person. Come and take the data and delete it (take it and go to your home). here we are guaranteeing for data up to the db after that we don't control it. 

Sns is like a magazine office with a delivery boy, we take care of the message/magazine to be on the door of a person after that we don't care. we will send this to the person who has paid a subscription or wants a magazine.

But sns is more scalable since it has topic , so its like multiple sqs with push consumer notification

[https://www.youtube.com/playlist?list=PL9nWRykSBSFg-CziAHKjr0XnvghEVkpFi](https://www.youtube.com/playlist?list=PL9nWRykSBSFg-CziAHKjr0XnvghEVkpFi)

**What is PubSub Service, and how does it work?**  

In the PubSub service, the publisher publishes events or messages to the subscriber, and the subscriber listens to those events. In the case of application-to-person communication, we typically have the recipient's email or phone number to send the response. However, for application-to-application communication, we need to provide our server endpoint for the recipient application to hit and receive the response.

Unlike other messaging services like SQS, where the service has to poll and know the SQS, in PubSub, the service does not need to know the specifics of the subscriber. It only needs to publish the message, and the subscribers who have subscribed to the topic will receive the message.
Clients can subscribe to the SNS topic and receive published messages using a supported endpoint type, such as Amazon Kinesis Data Firehose, Amazon SQS, AWS Lambda, **HTTP**, email, mobile push notifications, and mobile text messages (SMS).

Yep → [https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html](https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html)

[https://docs.aws.amazon.com/sns/latest/dg/sns-http-https-endpoint-as-subscriber.html](https://docs.aws.amazon.com/sns/latest/dg/sns-subscribe-https-s-endpoints-to-topic.html)

→ one to many  | publisher subscriber service

→ Fully Managed.

→ Topics and subscriptions.

→ Application to application and Application to person

![Untitled](SNS%20cd0daeffb4774bcdb19274cc9a11685b/Untitled.png)

**Message Ordering:**

[https://docs.aws.amazon.com/sns/latest/dg/fifo-topic-message-ordering.html](https://docs.aws.amazon.com/sns/latest/dg/fifo-topic-message-ordering.html)

**Internal Architecture**

Didn’t find much about this 💔

May be like the multiple FIFO SQS denoted as different topics, and in it we have the mechanism to send messages to all subscribers.

**Pricing**

[https://aws.amazon.com/sns/pricing/](https://aws.amazon.com/sns/pricing/)

……………………………..……………………………..……………………………..……………………………..……………………………..

Similar to a message queue, publish-subscribe is also a form of service-to-service communication that facilitates asynchronous communication. In a pub/sub model, any message published to a topic is pushed immediately to all the subscribers of the topic.

![https://raw.githubusercontent.com/karanpratapsingh/portfolio/master/public/static/courses/system-design/chapter-III/publish-subscribe/publish-subscribe.png](https://raw.githubusercontent.com/karanpratapsingh/portfolio/master/public/static/courses/system-design/chapter-III/publish-subscribe/publish-subscribe.png)

## Advantages

Let's discuss some advantages of using publish-subscribe:

- **Eliminate Polling**: Message topics allow instantaneous, push-based delivery, eliminating the need for message consumers to periodically check or *"poll"* for new information and updates. This promotes faster response time and reduces the delivery latency which can be particularly problematic in systems where delays cannot be tolerated.
- **Decoupled and Independent Scaling**: Publishers and subscribers are decoupled and work independently from each other, which allows us to develop and scale them independently.
- **Simplify Communication**: The Publish-Subscribe model reduces complexity by removing all the point-to-point connections with a single connection to a message topic, which will manage subscriptions and decide what messages should be delivered to which endpoints.

## Features

Now, let's discuss some desired features of publish-subscribe:

### Push Delivery

Pub/Sub messaging instantly pushes asynchronous event notifications when messages are published to the message topic. Subscribers are notified when a message is available.

### Multiple Delivery Protocols

In the Publish-Subscribe model, topics can typically connect to multiple types of endpoints, such as message queues, serverless functions, HTTP servers, etc.

### Fanout

This scenario happens when a message is sent to a topic and then replicated and pushed to multiple endpoints. Fanout provides asynchronous event notifications which in turn allows for parallel processing.

### Filtering

This feature empowers the subscriber to create a message filtering policy so that it will only get the notifications it is interested in, as opposed to receiving every single message posted to the topic.

### Durability

Pub/Sub messaging services often provide very high durability, and at least once delivery, by storing copies of the same message on multiple servers.

### Security

Message topics authenticate applications that try to publish content, this allows us to use encrypted endpoints and encrypt messages in transit over the network.