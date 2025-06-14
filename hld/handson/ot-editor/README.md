# 📝 Operational Transform (OT) Collaborative Editor

A real-time collaborative document editor built with **Operational Transform (OT)** algorithms for conflict-free concurrent editing. This implementation demonstrates how multiple users can edit the same document simultaneously while maintaining consistency and preserving user intentions.

## 🎯 What is Operational Transform?

**Operational Transform (OT)** is a technology for supporting real-time collaborative editing of shared documents. It allows multiple users to edit a document concurrently while automatically resolving conflicts and ensuring all users converge to the same final state.

### Key Principles:
- **Concurrent Operations**: Multiple users can edit simultaneously
- **Conflict Resolution**: Automatic handling of conflicting edits
- **Intention Preservation**: User's intended changes are maintained
- **Convergence**: All users eventually see the same document

## 🚀 Features

- **Real-time Collaboration**: Multiple users editing simultaneously
- **WebSocket Communication**: Low-latency real-time updates
- **Conflict Resolution**: Automatic OT-based conflict handling
- **Persistent Storage**: Auto-save with backup system
- **Operation History**: Complete edit history tracking
- **Client Management**: User tracking and cursor positions
- **Modern UI**: Clean, responsive web interface

## 🏗️ Architecture Overview

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Web Client    │    │   Go Server      │    │   File Storage  │
│                 │    │                  │    │                 │
│ - Frontend JS   │◄──►│ - WebSocket      │◄──►│ - JSON Files    │
│ - Diff Engine   │    │ - OT Engine      │    │ - Auto-backup   │
│ - UI Updates    │    │ - Doc Manager    │    │ - Versioning    │
└─────────────────┘    └──────────────────┘    └─────────────────┘
```

## 🔧 Technology Stack

- **Backend**: Go 1.21
- **WebSocket**: Gorilla WebSocket
- **HTTP Router**: Gorilla Mux
- **Storage**: JSON File System
- **Frontend**: Vanilla JavaScript
- **UI**: Modern CSS with responsive design

## 📋 Quick Start

### Prerequisites
- Go 1.21 or higher
- Modern web browser

### Installation & Running

```bash
# Clone the repository
cd ot-editor

# Install dependencies
go mod tidy

# Run the server
go run main.go

# Open your browser
open http://localhost:8080
```

### Multiple Users Testing
1. Open multiple browser tabs/windows to `http://localhost:8080`
2. Enter different usernames for each tab
3. Start editing simultaneously to see OT in action!

## 🧮 How OT Resolves Conflict Operations

### The Conflict Resolution Process

When multiple users edit simultaneously, conflicts arise when operations are based on outdated document versions. OT resolves these through **transformation**.

#### Example Conflict Scenario:
```
Initial Document: "Hello" (Version 1)

User A: Insert " World" at position 5 (based on v1)
User B: Insert " Beautiful" at position 5 (based on v1)
```

#### Resolution Steps:

1. **Operation Arrival**: Server receives operations in some order
2. **Concurrent Detection**: `getConcurrentOperations()` finds conflicting ops
3. **Transformation**: Later operations transform against earlier ones
4. **Application**: Transformed operations applied to document
5. **Broadcasting**: All clients receive updates

#### Transformation Algorithm:
```
Operation A: Insert " World" at pos 5
Operation B: Insert " Beautiful" at pos 5

If A processes first:
- Document becomes: "Hello World" (v2)
- B's operation transforms: pos 5 → pos 11 (after " World")
- Final result: "Hello World Beautiful"

All users converge to the same state! ✅
```

### OT Transformation Rules

#### Insert vs Insert:
```
Op1: Insert "X" at pos 5
Op2: Insert "Y" at pos 5

Transform Op2: pos 5 → pos 6 (after "X")
Result: Both insertions preserved in consistent order
```

#### Insert vs Delete:
```
Op1: Insert "X" at pos 5  
Op2: Delete 3 chars at pos 7

Transform Op2: pos 7 → pos 8 (account for inserted "X")
Result: Insertion and deletion both preserved
```

#### Delete vs Delete:
```
Op1: Delete 2 chars at pos 5
Op2: Delete 1 char at pos 6

Transform Op2: Handle overlapping deletes
Result: Smart conflict resolution avoiding duplicate deletions
```

## 🔄 Current Local Code Flow with Concurrent Clients

### System Architecture Flow

```
┌─────────────┐                    ┌─────────────┐
│   Client 1  │                    │   Client 2  │
│             │                    │             │
└──────┬──────┘                    └──────┬──────┘
       │                                  │
       │ Insert "Hello"                   │ Insert " World"
       │                                  │
       ▼                                  ▼
┌─────────────────────────────────────────────────────────┐
│              WebSocket Handler                          │
│  ┌─────────────────────────────────────────────────────┐│
│  │           Document Service                          ││
│  │                                                     ││
│  │  1. operation.Author = clientID                     ││
│  │  2. operation.Version = doc.Version                 ││
│  │  3. concurrentOps = getConcurrentOps()              ││
│  │  4. Transform if conflicts exist                    ││
│  │                                                     ││
│  │  ┌─────────────────────────────────────────────────┐││
│  │  │           Document Model                        │││
│  │  │                                                 │││
│  │  │  • d.Content = newContent                       │││
│  │  │  • d.Version++                                  │││
│  │  │  • op.Version = d.Version                       │││
│  │  │  • d.Operations.append(op)                      │││
│  │  │                                                 │││
│  │  └─────────────────────────────────────────────────┘││
│  │                                                     ││
│  │  ┌─────────────────────────────────────────────────┐││
│  │  │           Client Manager                        │││
│  │  │                                                 │││
│  │  │  • BroadcastToDocument()                        │││
│  │  │  • SendToClient()                               │││
│  │  │                                                 │││
│  │  └─────────────────────────────────────────────────┘││
│  │                                                     ││
│  └─────────────────────────────────────────────────────┘│
└─────────────────────────────────────────────────────────┘
       │                                  │
       │ ACK + Broadcast                  │ ACK + Broadcast
       │                                  │
       ▼                                  ▼
┌──────────────┐                  ┌──────────────┐
│   Client 1   │◄────────────────►│   Client 2   │
│   "Hello"    │   Synchronized   │   "Hello"    │
└──────────────┘     Content      └──────────────┘

Result: Both clients converged to same state ✅
```

### Key Components Flow

#### 1. **Client Operation Generation**
```javascript
// Frontend detects changes
const operation = generateOperation(oldText, newText);
// Send via WebSocket
websocket.send(JSON.stringify({
    type: "operation",
    payload: { documentId, operation }
}));
```

#### 2. **Server-Side Processing**
```go
// WebSocket Handler
func (wh *WebSocketHandler) handleOperation(clientID string, message *WebSocketMessage) {
    // Parse operation
    transformedOp, err := wh.documentService.ApplyOperation(clientID, payload.DocumentID, &payload.Operation)
    
    // Broadcast to other clients
    wh.documentService.BroadcastToDocument(payload.DocumentID, clientID, broadcastMessage)
}
```

#### 3. **OT Core Logic**
```go
// Document Service
func (ds *DocumentService) ApplyOperation(clientID, documentID string, operation *Operation) (*Operation, error) {
    // Set operation metadata
    operation.Author = clientID
    operation.Version = doc.Version
    
    // Find concurrent operations
    concurrentOps := ds.getConcurrentOperations(doc, operation.Version)
    
    // Transform if conflicts exist
    if len(concurrentOps) > 0 {
        transformedOp, err = ot.TransformAgainstOperationList(operation, concurrentOps)
    }
    
    // Apply to document
    err = doc.ApplyOperation(transformedOp)
    return transformedOp, nil
}
```

#### 4. **Document Version Management**
```go
// Document Model
func (d *Document) ApplyOperation(op *Operation) error {
    // Update content
    d.Content = newContent
    d.Version++                    // Version increment
    op.Version = d.Version         // Tag operation with new version
    d.Operations = append(d.Operations, *op)  // Add to history
    return nil
}
```

### Concurrent Operation Timeline

```
Time: T1    T2    T3    T4    T5
Client 1: [Edit] ────► [ACK] ────► [Receive B's op]
Client 2:        [Edit] ────► [Transform] ────► [ACK]
Server:   [Op A] [Op B] [Transform B] [Broadcast] [Sync]

Result: Both clients have identical final state
```

## ⚡ Importance of Ordering in OT

### Why Operation Ordering is Critical

**Operational Transform's mathematical correctness depends entirely on processing operations in a consistent order across all clients.**

#### The Fundamental Promise:
> "If all clients process the same set of operations in the same order (after transformation), they will converge to the same final state."

### Problems Without Proper Ordering

#### 1. **Document Divergence**
```
Document: "ABC"
User 1: Delete "B" → sees "AC"  
User 2: Delete "C" → sees "AB"

Without ordering: Users see different documents forever!
With ordering: Both see same result (e.g., "A")
```

#### 2. **Intention Violation**
```
Document: "The cat"
User 1: Insert " big" at pos 4 → wants "The big cat"
User 2: Delete "cat" → wants "The "

Wrong order: "The big" (User 1's intention lost)
Right order: Proper transformation preserves both intentions
```

#### 3. **Transformation Failure**
```
Op A: Insert "X" at pos 5
Op B: Insert "Y" at pos 5  

Different processing orders:
Order A→B: Transform B against A → "...XY..."
Order B→A: Transform A against B → "...YX..."

Without consistent ordering: Clients see "XY" vs "YX"
```

### How Our Implementation Ensures Ordering

#### 1. **Version-Based Conflict Detection**
```go
// Tag operations with current document version
operation.Version = doc.Version

// Find operations that happened after client's version
concurrentOps := getConcurrentOperations(doc, operation.Version)
```

#### 2. **Server-Side Sequencing**
```go
// Document version increments atomically
d.Version++
op.Version = d.Version
d.Operations = append(d.Operations, *op)
```

#### 3. **Sequential Processing**
All operations processed by single server thread ensures consistent ordering within each document.

### Ordering Guarantees

| **Guarantee** | **How Achieved** | **Benefit** |
|---------------|------------------|-------------|
| **Causal Ordering** | Version numbers | Operations see effects of predecessors |
| **Total Ordering** | Server sequencing | All clients see same operation sequence |
| **Conflict-Free** | OT transformation | Mathematical convergence guaranteed |

## 🚀 Production-Level Event-Driven Architecture

When scaling beyond a single server, we need distributed event-driven architecture to maintain OT's ordering guarantees across multiple replicas.

### Scaling Challenges

#### Single Server Limitations:
- **Memory Bottleneck**: All documents in single server's RAM
- **Connection Limits**: WebSocket connections tied to one server
- **Single Point of Failure**: Server crash affects all users
- **Geographic Latency**: Users far from server experience delays

### Production Architecture Pattern

#### **Two-Stage Event Pipeline (Recommended)**

```
┌─────────────────────────────────────────────────────────────────┐
│                     Application Layer                           │
│                                                                 │
│  ┌───────────────┐ ┌───────────────┐ ┌───────────────┐        │
│  │ Server        │ │ Server        │ │ Server        │        │
│  │ Replica 1     │ │ Replica 2     │ │ Replica 3     │        │
│  └───────┬───────┘ └───────┬───────┘ └───────┬───────┘        │
└──────────┼─────────────────┼─────────────────┼────────────────┘
           │                 │                 │
           ▼                 ▼                 ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Unordered Pipeline                           │
│                                                                 │
│           ┌─────────────────────────────────────┐               │
│           │     Kafka Topic:                    │               │
│           │   "unordered-operations"            │               │
│           │   (Race conditions OK)              │               │
│           └─────────────┬───────────────────────┘               │
└─────────────────────────┼───────────────────────────────────────┘
                          │
                          ▼
┌─────────────────────────────────────────────────────────────────┐
│                     Ordering Layer                              │
│                                                                 │
│           ┌─────────────────────────────────────┐               │
│           │      Sequencer Consumer             │               │
│           │   (Single instance per document)    │               │
│           │   (Assigns sequence numbers)        │               │
│           └─────────────┬───────────────────────┘               │
└─────────────────────────┼───────────────────────────────────────┘
                          │
                          ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Ordered Pipeline                             │
│                                                                 │
│           ┌─────────────────────────────────────┐               │
│           │     Kafka Topic:                    │               │
│           │   "ordered-operations"              │               │
│           │   (Guaranteed ordering)             │               │
│           └─────────────┬───────────────────────┘               │
└─────────────────────────┼───────────────────────────────────────┘
                          │
           ┌──────────────┼──────────────┐
           ▼              ▼              ▼
┌─────────────────────────────────────────────────────────────────┐
│                   Processing Layer                              │
│                                                                 │
│  ┌───────────────┐ ┌───────────────┐ ┌───────────────┐        │
│  │ Consumer 1    │ │ Consumer 2    │ │ Consumer 3    │        │
│  │ (Server 1)    │ │ (Server 2)    │ │ (Server 3)    │        │
│  └───────────────┘ └───────────────┘ └───────────────┘        │
└─────────────────────────────────────────────────────────────────┘
```

### Production Components

#### 1. **Operation Sequencer**
**Purpose**: Assign globally consistent sequence numbers

**Technology Options**:
- **Redis INCR**: Fast, simple, single point of failure
- **PostgreSQL Sequences**: ACID guarantees, network latency
- **Kafka Offsets**: Built-in ordering, complex setup
- **etcd/Consensus**: Distributed, fault-tolerant, higher latency

**Recommended**: Redis for <10K ops/sec, PostgreSQL for higher loads

#### 2. **Event Service**
**Purpose**: Distribute operations to all server replicas

**Technology Options**:
- **Apache Kafka**: Industry standard, millions of events/sec
- **Redis Streams**: Simpler than Kafka, good performance
- **Google Pub/Sub**: Managed service, auto-scaling
- **AWS SQS/SNS**: Managed, easier ops, higher latency

**Recommended**: Kafka for high scale, Redis Streams for simplicity

#### 3. **Distributed Storage**
**Purpose**: Shared document state across replicas

**Technology Options**:
- **PostgreSQL**: Strong consistency, proven at scale
- **MongoDB**: Document-native, good for JSON operations
- **CockroachDB**: Distributed SQL, global consistency
- **Redis Cluster**: In-memory speed, persistence options

**Recommended**: PostgreSQL + Redis cache for most use cases

### Production Flow Example

```
Client 1                Load Balancer              Server 1
   │                           │                       │
   │ WebSocket: Edit operation │                       │
   ├──────────────────────────►│                       │
   │                           │ Route to Server 1     │
   │                           ├──────────────────────►│
   │                           │                       │
   │                           │                       │ Apply locally
   │                           │                       │ (optimistic)
   │                           │                       │
   │                           │                       ▼
   │                           │              ┌─────────────────┐
   │                           │              │ Unordered Topic │
   │                           │              │ (Kafka)         │
   │                           │              └────────┬────────┘
   │                           │                       │
   │                           │                       ▼
   │                           │              ┌─────────────────┐
   │                           │              │   Sequencer     │
   │                           │              │ Assign seq:1001 │
   │                           │              └────────┬────────┘
   │                           │                       │
   │                           │                       ▼
   │                           │              ┌─────────────────┐
   │                           │              │ Ordered Topic   │
   │                           │              │ Operation #1001 │
   │                           │              └────┬───────┬────┘
   │                           │                   │       │
   │                           │                   ▼       ▼
   │                           │            Server 1   Server 2
   │                           │                   │       │
   │                           │                   │       │ Apply #1001
   │                           │                   │       │
   │                           │                   │       ▼
   │                           │                   │   Client 2
   │ ◄─────────────────────────┼───────────────────┘       │
   │ Synchronized content      │                           │
   │                           │ ◄─────────────────────────┘
   │                           │ Both clients synchronized
   ▼                           ▼
✅ Final State: Both clients have identical content
```

### Production Scaling Characteristics

| **Scale** | **Architecture** | **Technologies** | **Capacity** |
|-----------|------------------|------------------|--------------|
| **Small** (< 1K users) | Single server | Current implementation | 1K concurrent |
| **Medium** (1K-10K) | Redis + PostgreSQL | Redis Streams, PG | 10K concurrent |
| **Large** (10K-100K) | Kafka + DB cluster | Kafka, PostgreSQL cluster | 100K concurrent |
| **Enterprise** (100K+) | Multi-region | Kafka clusters, CockroachDB | Millions concurrent |

### Key Production Considerations

#### **Ordering Guarantees**
- **Within Document**: Strict sequential ordering required
- **Cross Document**: Can be processed independently
- **Network Partitions**: Need consensus or graceful degradation

#### **Performance Targets**
- **Latency**: < 100ms operation propagation
- **Throughput**: 10K+ operations/second per server
- **Availability**: 99.9% uptime with automatic failover

#### **Operational Complexity**
- **Monitoring**: Operation flow, sequence gaps, consumer lag
- **Debugging**: Distributed tracing, operation lineage
- **Deployment**: Rolling updates, schema evolution

## 🏃 Getting Started with Production Migration

### Phase 1: External Storage
Replace file storage with PostgreSQL for persistence and multi-server access.

### Phase 2: Redis Sequencing  
Add Redis for atomic sequence number generation across servers.

### Phase 3: Event Broadcasting
Implement Redis Streams or Kafka for cross-server operation distribution.

### Phase 4: Load Balancing
Add multiple server replicas with shared storage and event bus.

### Phase 5: Regional Scaling
Deploy multiple regions with cross-region event replication.

## 📚 Further Reading

- [Operational Transformation Theory](https://en.wikipedia.org/wiki/Operational_transformation)
- [Google Docs OT Implementation](https://googledrive.googleblog.com/2010/09/whats-different-about-new-google-docs.html)
- [Kafka Event Sourcing Patterns](https://kafka.apache.org/documentation/#patterns)
- [Distributed Systems Ordering](https://lamport.azurewebsites.net/pubs/time-clocks.pdf)

## 🤝 Contributing

Contributions welcome! Please read our contributing guidelines and submit pull requests for improvements.

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

**Built with ❤️ using Operational Transform algorithms for real-time collaborative editing.** 