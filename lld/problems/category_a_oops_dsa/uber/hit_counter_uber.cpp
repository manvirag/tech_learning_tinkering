/*

Uber's Hit Counter (from LeetCode 362 & interview reports) is exactly the standard problem: Design a class to track hits over a sliding 5-minute window using a deque for O(1) operations.
​

Exact API (as asked)
java
class HitCounter {
    void hit(int timestamp);           // Record hit at timestamp (seconds, increasing)
    int getHits(int timestamp);        // # hits in [timestamp-299, timestamp]
}
Precise Requirements
text
counter.hit(1);     // at t=1
counter.hit(2);     
counter.hit(3);     
counter.getHits(4); // 3  (1,2,3 all in [4-299= -295? Wait: [1,4])

counter.hit(300);   
counter.getHits(300); // 4 (1,2,3,300? NO: 1 expired, only 300? Wait example says 1)
counter.getHits(301); // 0 (300 expired)
Key: Window [t-299, t] inclusive. Evict timestamps < t-299.

Constraints (Uber-specific)
Timestamps: 1+ , strictly increasing calls

High QPS: 1000s/sec → use deque of (timestamp, count) pairs

Space: O(300) max entries

Time: hit() O(1), getHits() O(1) amortized

Expected Solution Structure
text
Deque<Pair<Integer, Integer>> queue;  // (ts, count_at_ts)

hit(t):
  if (!empty && front.ts == t) front.count++
  else queue.add(new Pair(t, 1))

getHits(t):
  while (!empty && front.ts <= t-300) queue.pollFirst()
  return sum all counts in queue
Uber Follow-ups
Thread-safe (synchronized deque)

Per-service counters (Map<String, HitCounter>)

Multiple windows (1m/5m/1h)

Rate limiting integration



/*


→ Hit Counter

→ 

Design a data structure with time-to-live (TTL) where you receive (key, timestamp) pairs (in increasing order). You need to provide:

a) Count of all active keys

b) Count of active keys for a particular key

c) CRUD functions

Approach: Map<String, Queue>. At every function call, go through the map and remove expired keys. Also maintain global counters.

This approach worked, but the interviewer wanted something like a Queue where Pair = (timestamp, key). This way, instead of iterating over the map for every key, we can directly poll elements from the queue irrespective of the key, and update counts using Map<String, Integer> and a global totalCounter. The worst-case complexity is the same, but this is simpler and has better average-case complexity.

leanhire

*/

​
*/