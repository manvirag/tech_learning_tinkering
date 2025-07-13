# Introduction to Apache Flink ( Next step -> read its doc complete )

``The limits of event-driven applications are defined by how well a stream processor can handle time and state``

- The limits of event-driven applications are defined by how well a stream processor can handle time and state. Many of Flink’s outstanding features are centered around these concepts. Flink provides a rich set of state primitives that can manage very large data volumes (up to several terabytes) with exactly-once consistency guarantees. Moreover, Flink’s support for event-time, highly customizable window logic, and fine-grained control of time as provided by the ProcessFunction enable the implementation of advanced business logic. 
- At its core, Flink treats all data processing as a series of events, making it particularly well-suited for stream processing tasks. 

## Key Features of Apache Flink

- **True stream processing:** Processes data on an event-by-event basis for low latency and high throughput.
- **Exactly-once semantics:** Guarantees each event is processed exactly once, even in failures.
- **Stateful computations:** Supports complex event processing and machine learning tasks.
- **Event time processing:** Processes events based on their generation time, handling out-of-order and late data.
- **Windowing support:** Flexible window operations for aggregations over time or count intervals.
- **High availability & fault tolerance:** Checkpointing and savepoints ensure recovery without data loss.
- **Scalability:** Can scale to thousands of nodes, processing millions of events per second.

---

# Flink Architecture

Flink’s architecture supports scalable, fault-tolerant stream processing.

### Job Manager

- Coordinates scheduling, checkpoints, failure recovery, and job control flow.
- Components:
  - **ResourceManager:** Manages resource allocation.
  - **Dispatcher:** REST interface for job submissions, starts JobMasters.
  - **JobMaster:** Manages execution of individual jobs.

### Task Manager

- Worker nodes executing tasks in separate threads.
- Reports status to Job Manager.
- Manages memory and CPU.
- Organized into slots — units of resource scheduling.
- kind of same philosophy like map reduce, since they can shuffle etc. between task manager

### Client

- Prepares and submits dataflows to Job Manager.
- Runs within Java/Scala program or CLI.

### Distributed Execution

- **Parallelism:** Distributes operators across Task Managers.
- **Data Exchange Patterns:** Forward, Shuffle, Rebalance, Broadcast.
- **State Backend:**
  - MemoryStateBackend (in-memory, checkpoints to JobManager memory)
  - FsStateBackend (in-memory, checkpoints to file system)
  - RocksDBStateBackend (disk-based for large states)

### Fault Tolerance and High Availability

- Checkpointing: Asynchronous snapshots.
- Savepoints: Manually triggered for updates.
- State Restoration: From checkpoints/savepoints.
- JobManager HA: ZooKeeper for leader election.
- Optimized network stack with credit-based flow control and zero-copy networking.

---


## APIs

- **DataStream API:** For unbounded stream processing.
- **DataSet API:** For bounded batch processing (being phased out).

## Program Structure

- **Source:** Data origin (Kafka, files).
- **Transformations:** Data operations.
- **Sink:** Output destination.

Programs form a **dataflow graph** with streams and operators.

## Windowing

- Time Windows: Tumbling, sliding, session windows.
- Count Windows: Based on element count.
- Custom user-defined windows.

## Time and Watermarks

- Event Time: When event occurred.
- Processing Time: When event is processed.
- Ingestion Time: When event enters system.
- Watermarks track event time progress, enabling late data handling.

## State Management

- **Keyed State:** Partitioned by key.
- **Operator State:** Bound to operators.
- State primitives: ValueState, ListState, MapState, ReducingState, AggregatingState.

## Checkpointing & Exactly-Once Semantics

- Periodic, barrier-coordinated snapshots.
- Supports large state backends (e.g., RocksDB).

## Event Processing Functions

- MapFunction, FlatMapFunction, FilterFunction.
- KeyedProcessFunction: low-level with full state/time access.
- WindowFunction: for windowed data processing.

## Table API & SQL

- Declarative APIs integrating with DataStream/DataSet.
- Support for complex mixed pipelines.


# Use Cases and Real-World Examples

## Real-Time Analytics

**Example: E-commerce Analytics**

- Track user behavior, product popularity, dynamic pricing.
- Uses windowed counts of product views with sinks to Elasticsearch.

```

DataStream<Event> events = environment
.addSource(new KafkaSource<>(“user-events”, new EventDeserializationSchema()));

DataStream<Tuple2<String, Integer>> productViews = events
.filter(event -> event.getType().equals(“product_view”))
.map(event -> new Tuple2<>(event.getProductId(), 1))
.keyBy(tuple -> tuple.f0)
.window(TumblingEventTimeWindows.of(Time.minutes(5)))
.sum(1);

productViews.addSink(new ElasticsearchSink<>(config, indexer));
```
This example processes a stream of user events, calculates product views in 5-minute windows, and sends the results to Elasticsearch for visualization.

Sliding window with sliding time interval

```

import org.apache.flink.api.common.eventtime.WatermarkStrategy;
import org.apache.flink.api.common.serialization.SimpleStringSchema;
import org.apache.flink.api.java.tuple.Tuple2;
import org.apache.flink.streaming.api.datastream.SingleOutputStreamOperator;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
import org.apache.flink.streaming.api.windowing.time.Time;
import org.apache.flink.streaming.api.windowing.assigners.SlidingEventTimeWindows;
import org.apache.flink.streaming.connectors.kafka.FlinkKafkaConsumer;
import org.apache.flink.streaming.connectors.kafka.FlinkKafkaProducer;
import org.apache.flink.util.OutputTag;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.time.Duration;
import java.util.Properties;

public class SlidingWindowClickCount {

    public static class ClickEvent {
        public String userId;
        public long timestamp;
        public String page;
        public String action;
    }

    public static void main(String[] args) throws Exception {
        final StreamExecutionEnvironment env = StreamExecutionEnvironment.getExecutionEnvironment();

        Properties kafkaProps = new Properties();
        kafkaProps.setProperty("bootstrap.servers", "localhost:9092");
        kafkaProps.setProperty("group.id", "click-counter");

        FlinkKafkaConsumer<String> consumer = new FlinkKafkaConsumer<>("user-clicks", new SimpleStringSchema(), kafkaProps);
        consumer.assignTimestampsAndWatermarks(
            WatermarkStrategy
                .<String>forBoundedOutOfOrderness(Duration.ofSeconds(5))
                .withTimestampAssigner((eventStr, timestamp) -> {
                    try {
                        ClickEvent event = new ObjectMapper().readValue(eventStr, ClickEvent.class);
                        return event.timestamp;
                    } catch (Exception e) {
                        return 0L;
                    }
                })
        );

        SingleOutputStreamOperator<String> result = env
            .addSource(consumer)
            .map(value -> new ObjectMapper().readValue(value, ClickEvent.class))
            .keyBy(event -> event.userId)
            .window(SlidingEventTimeWindows.of(Time.minutes(1), Time.seconds(10)))
            .aggregate(new ClickCountAggregator(), new ClickWindowFunction())
            .map(new ObjectMapper()::writeValueAsString);

        FlinkKafkaProducer<String> producer = new FlinkKafkaProducer<>("user-click-counts", new SimpleStringSchema(), kafkaProps);
        result.addSink(producer);

        env.execute("Sliding Window Click Counter");
    }
}

```

## Fraud Detection

**Example: Credit Card Fraud**

- Detect suspicious transactions using keyed state and timers.
- Sends alerts on potential fraud.

## ETL & Data Pipelines

**Example: Log Data Processing**

- Batch job reading logs, filtering errors, enriching data.
- Writes processed logs back to HDFS.
         |

![](./image-flink.png)

### Reference

- https://medium.com/@22.gautam/apache-flink-unveiled-a-deep-dive-into-next-generation-stream-processing-9267352e1819
- Apache doc https://flink.apache.org/what-is-flink/use-cases/
- https://nightlies.apache.org/flink/flink-docs-release-1.20/docs/learn-flink/overview/  [ ek number must read for deep dive , pending]
