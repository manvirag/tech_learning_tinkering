# aws lambda

> Documentation: https://docs.aws.amazon.com/lambda/latest/dg/welcome.html
> 
1. Assume aws lambda as black box (like a function) as of now, that is first invoked , then it run the code with given input.
2. It has to be invoked , like its api from our code or adding the trigger for e.g kinesis, sqs etc. (and this methoding of running is also called serverless computing)
3. We will discuss in this flow:
A. How it calculates its price.
B. Lambda configurations.
C. Monitoring and scalability.
D. Again price calcuations
4. How its calculate the price: ([https://aws.amazon.com/lambda/pricing/](https://aws.amazon.com/lambda/pricing/)

> Sure. AWS Lambda pricing is based on two factors:
> 
> 
> **Invocations**: The number of times your Lambda function is called.
> 
> **Duration**: The amount of time your Lambda function takes to execute.( Unit GB-second the number of seconds of compute you’ve done, multiplied by the number of GB of memory that compute allocated.)
> 
> Lambda offers a free tier that includes 1 million invocations and 3.2 million seconds of compute time per month. After you exceed the free tier, you will be charged $0.20 per 1 million invocations and $0.00001667 per GB-second of compute time.
> 
> To calculate the cost of your Lambda function, you can use the following formula: (Per month after free tier)
> 
> **Cost = (Invocations * $0.20) + (Duration * $0.00001667)**
> 
> Finding invocation is easy right. But how to find duration
> 
> **GB-second = (Memory in GB) * (Execution time in seconds)**
> 
> For example, if a Lambda function has 1 GB of memory and executes for 10 seconds, it will consume 10 GB-second of compute time.
> 
> What is Memory in GB = this the memory that we allocation for lambda , we can change it.
> 
> Configuration -> General Setting - > Memory
> 
> ![aws%20lambda%20df9bad6c6d174ef78ea8c0eee8f503e5/image6.png](aws%20lambda%20df9bad6c6d174ef78ea8c0eee8f503e5/image6.png)
> 
1. Okk So Now we know how it calculate its price.
Let talk about lambda specific configuration. (When request comes what we can tweak to meet requirement like scalability etc)
And also the configuration of a trigger in lambda**Lambda configuration:Throttle (at the top right) : ?**[https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-memory-console](https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-memory-console)

> Read from doc, only thing which are not in doc or doubt will be writtern
> 
> 
> **General Settings:**
> 
> Memory:
> 
> – suppose if we allocate 2000MB memory and in single invocation it uses only 100MB then cost will be calculated with 100 or 2000 ? (Will be 2000 , [https://aws.amazon.com/lambda/pricing/#AWS_Lambda_Pricing](https://aws.amazon.com/lambda/pricing/#AWS_Lambda_Pricing))
> 
> Timeout:
> 
> – Doubt does it equivalent to getting error while executing lambda ? , if we have retry attempts in trigger will it retry again, or if dql then will again process it if its enabled ?
> 

Ephemeral storage: Price: [https://aws.amazon.com/lambda/pricing/#Lambda_Ephemeral_Storage_Pricing](https://aws.amazon.com/lambda/pricing/#Lambda_Ephemeral_Storage_Pricing)

Doc: [https://docs.aws.amazon.com/lambda/latest/dg/configuration-concurrency.html](https://docs.aws.amazon.com/lambda/latest/dg/configuration-concurrency.html)

![aws%20lambda%20df9bad6c6d174ef78ea8c0eee8f503e5/image3.png](aws%20lambda%20df9bad6c6d174ef78ea8c0eee8f503e5/image3.png)

\

![aws%20lambda%20df9bad6c6d174ef78ea8c0eee8f503e5/image5.png](aws%20lambda%20df9bad6c6d174ef78ea8c0eee8f503e5/image5.png)

> By default, your account has a concurrency limit of 1,000 across all functions in a region
> 
> 
> **Concurrency**
> 
> Reserved Concurrency:: Fixed value can scale up to limit. After that it will throttle. Use reserved concurrency to reserve a portion of your account's concurrency for a function. This is useful if you don't want other functions taking up all the available unreserved concurrency
> 
> ![aws%20lambda%20df9bad6c6d174ef78ea8c0eee8f503e5/image4.png](aws%20lambda%20df9bad6c6d174ef78ea8c0eee8f503e5/image4.png)
> 
> (if allocated 6 , using 4 ,how price ? → Configuring reserved concurrency counts towards your overall account concurrency limit. There is no charge for configuring reserved concurrency for a function.
> 
> ) (then how its calculate ? my guess each concurrent count as different invocations for suppose we reserved 6 ,using 4 , then 4 will be used to calculated price, i mean it can be different a/c to concurrency using in that time)
> 
> —> What if we use 1000 concurrency in a region ?
> 
> Provisional Concurrency:
> 
> ![aws%20lambda%20df9bad6c6d174ef78ea8c0eee8f503e5/image2.png](aws%20lambda%20df9bad6c6d174ef78ea8c0eee8f503e5/image2.png)
> 
> ![aws%20lambda%20df9bad6c6d174ef78ea8c0eee8f503e5/image1.png](aws%20lambda%20df9bad6c6d174ef78ea8c0eee8f503e5/image1.png)
> 

**You can't allocate more provisioned concurrency than reserved concurrency for a function.**

Both reserved concurrency and provisioned concurrency count towards your account concurrency limit and Regional quotas. In other words, allocating reserved and provisioned concurrency can impact the concurrency pool that's available to other functions. **Configuring provisioned concurrency incurs charges to your AWS account.**Price: [https://aws.amazon.com/lambda/pricing/](https://aws.amazon.com/lambda/pricing/)

Doc: [https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html)

**Asynchornous Invocations: (**Several AWS services, such as Amazon Simple Storage Service (Amazon S3) and Amazon Simple Notification Service (Amazon SNS), invoke functions asynchronously to process events. When you invoke a function asynchronously, you don't wait for a response from the function code. You hand off the event to Lambda and Lambda handles the rest**)**Maximum age of eventRetry AttemptsDead Letter Queue Service
Some other configuration no learning them. Since not require as of now.

**Trigger Configuration:**

> 1. State: This field indicates whether the configuration is enabled or disabled.
> 
> 
> 2. **Activate trigge**r: This field indicates whether to activate the trigger for this configuration.
> 
> 3. **Batch size**: This field specifies the number of records to be processed in a single batch. (when lambda invoke we can run that function with more than one record in batch benefit: only one time it will init that’s it.)
> 
> 4. **Batch window**: This field specifies the duration of the batch window, within which records will be processed. (suppose batch size is 10 , but we have only 1 record then function wait upto batch window seconds (value in second) then starts executing whatever records they are)
> 
> 5. **Concurrent batches per shard**: (there is concept of shard in kinesis , mean concurrency in kinesis and each can have no. of batches ) This field specifies the maximum number of batches that can be processed concurrently per shard.(means how many batch we want our kinesis have in each shard, what if we increase this benefit: suppose we have extra reserved concurrency then increase batch can help since other instance will handle this , but if we have maximum oncurrency limit then it will not help)
> 
> 6**. Last processing result:** tells the status of event source mapping , just afyer when we update or create lambda. ok = active and lambda processing data other option dropped = active but not processing lambda testing purpose and disable = unactive stop flow from event source.
> 
> 7. **Maximum age of record**: The Maximum age of record configuration is a setting in AWS Lambda's event source mapping that determines the maximum time that a record can stay in an event source (such as a Kinesis data stream) before it is considered expired and is no longer processed by the function. (-1 infinite)
> 
> 8. **On-failure destination**: The "On-failure destination" configuration allows you to specify where events that fail to be processed by the Lambda function should be sent for further analysis or handling. It can be set to an SQS queue, an SNS topic, or another Lambda function. (iguess instead of this we directly send fromcode)
> 
> 9. **Report batch item failures:? why,d we enbaled thia** When "Report batch item failures" is set to "Yes," Lambda will provide detailed information about failed records, allowing for granular error handling and troubleshooting. This is beneficial when you need to handle failures and successes separately. On the other hand, setting it to "No" means Lambda won't report individual failures, which can be useful when you prioritize overall batch processing and don't require specific details about failed records.
> 
> [https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#example-standard-queue-message-event](https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#example-standard-queue-message-event)
> 
> 10. **Retry attempts:** This field specifies the number of times to retry processing a failed record.
> 
> 11. **Split batch on error**: When "Split batch on error" is enabled, Lambda will automatically split a batch into smaller sub-batches when an error occurs during processing. The successful records will be processed independently, while the failed records can be retried separately or handled differently. This feature helps isolate and manage failures more effectively, allowing for partial processing and easier error recovery. It provides flexibility in handling different records within a batch based on their individual success or failure. split the failed and its subsequent in second batch.
> 
> Sure! Let's consider a batch of 10 records [R1, R2, R3, R4, R5, R6, R7, R8, R9, R10], and the "Split batch on error" option is enabled.
> 
> Lambda starts processing the batch of 10 records.
> 
> During processing, an error occurs while processing the 5th record, R5.
> 
> Instead of failing the entire batch, Lambda splits the batch into two sub-batches:
> 
> Sub-batch 1: [R1, R2, R3, R4]
> 
> Sub-batch 2: [R6, R7, R8, R9, R10]
> 
> Sub-batch 1 is processed independently, excluding the failed record R5. It can be retried, logged, or handled differently based on your application's logic.
> 
> Sub-batch 2 is also processed independently, without any impact from the failed record in Sub-batch 1.
> 
> The results of processing Sub-batch 1 and Sub-batch 2 are handled separately, allowing for separate error handling, retries, or destination configurations.
> 
> By splitting the batch on failure, the impact of a failed record is contained, and processing can continue for the remaining records in the batch. It provides a more granular approach to handle failures within a batch, improving fault tolerance and allowing for more efficient error handling.
> 
> Case 1: Split batch on error (enabled)
> 
> Suppose during processing, Record5 encounters an error. With the "Split batch on error" option enabled, Lambda will split the batch into two smaller batches:
> 
> Batch 1: [Record1, Record2, Record3, Record4]
> 
> Batch 2: [Record5, Record6, Record7, Record8, Record9, Record10]
> 
> Lambda will then retry processing Batch 2, which includes the failed record (Record5). The successful records from Batch 1 will be processed separately and independently.
> 
> Case 2: Split batch on error (disabled)
> 
> In this case, if an error occurs with Record5, Lambda will not split the batch. The entire batch of 10 records will be retried together as a single unit.
> 
> [https://medium.com/srcecde/handle-sqs-message-failure-in-batch-with-partial-batch-response-b858ad212573](https://medium.com/srcecde/handle-sqs-message-failure-in-batch-with-partial-batch-response-b858ad212573)
> 
> 11. **Starting position:** This field specifies the position in the stream from which to start processing records. our Latest.
> 
> 12. **Tumbling window duration:** ?
> 
> The tumbling window duration and the batch window are both concepts related to processing events or records in AWS Lambda triggers, but they have different purposes:
> 
> Tumbling Window Duration: It refers to a fixed time duration used for aggregating events or records before processing them as a batch. The tumbling window duration determines how long events are collected before triggering the Lambda function. Once the duration expires, the window is closed, and the accumulated events within that window are processed together as a batch. This allows you to control the time-based grouping of events or records.
> 
> Batch Window: The batch window, on the other hand, is a configuration setting specific to AWS Lambda Kinesis and DynamoDB stream triggers. It determines the maximum time window for collecting events or records before triggering the Lambda function. If the number of events specified by the batch size is not reached within the batch window, the Lambda function will still be triggered once the batch window expires. It ensures that events are processed within a certain time frame, even if the desired batch size is not met.
> 
> In summary, the tumbling window duration is a time-based grouping mechanism for aggregating events or records within a fixed duration, whereas the batch window is a time threshold within which events are accumulated before triggering the Lambda function, regardless of the batch size.
> 
> ask what happen when it caused failure
> 
1. **Now let’s talk about pricing with considering batch, reserved concurrency , provisional concurrency**If request coming in batch invocation will be one , but execution time will increase
Reserved concurrency have no price for configuration. It will calculate a/c to concurrent at time
Provisional have price : [https://aws.amazon.com/lambda/pricing/](https://aws.amazon.com/lambda/pricing/)
2. **Now We know what are the different ways lambda take input or different input with which lambda invocation works like batch etc.**Now talk about the failure case and how to handle it.
If retry have count → send it to dql → make it retry attemp 1 → will remain in this then.
3. **Now Lets’ talk about the monitoring, or in simple Some cloudwatch metric for aws lambda**Invocations by batch, throttle by batch, error, duration, age of record
4. **Later: Deep Dive the implementation of aws lambda.**