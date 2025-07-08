# Async

![Untitled](Async%20e5cddddc209e41fb96bcac8a65a7983a/Untitled.png)

1. Message Broker

    1. It is a concept. 
    2. A message broker is software that enables applications, systems, and services to communicate with each other and exchange information
    3. queue, Kafka, [ sqs, sns, kinesis, msk] all comes in this.
    4. it can be point to point: This is the distribution pattern utilized in message queues with a one-to-one relationship between the message's sender and receiver.  Impl. can be pull/get on consumer
    [ Check sqs ]
    5. pub/sub: In this message distribution pattern, often referred to as *"pub/sub"*, the producer of each message publishes it to a topic, and multiple message consumers subscribe to topics from which they want to receive messages.Impl. can be pull/push on consumer
        
        [ sns  ]