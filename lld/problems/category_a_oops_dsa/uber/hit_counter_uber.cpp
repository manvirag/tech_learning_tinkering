/*



https://leetcode.com/discuss/post/5838801/uber-sde-2-phone-screen-by-anonymous_use-99jw/

Question: Implement a Counter class that has the following methods:

1. `put(number)`: put the number to the data structure
2. `count(number)`: count the number of times `number` was put during the last `window=5 minutes`
3. `countAll()`: count the number of times any number was put during the last `window=5 minutes`.

Example:

At t = 10PM, `put(2)`

At t = 10:02PM, `put(2)`

At t = 10:03PM, `put(3)`

At t = 10:04PM, `count(2)` should return 2

At t = 10:04PM, `countAll()` should return 3

At t = 10:06PM, `count(2)` should return 1 (the one that was put at 10:02PM)

At t = 10:06PM, `countAll()` should return 2

Follow-ups:

1. If you were to write unit tests, what would they be
2. If your code was to be in production, what issues may it cause?

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




*/



// problem 1

/*
 
  map<int,int> freq 
  dequeue<pair<ts, key> 
  window -> 300 

*/


#include<iostream> 
#include<unordered_map> 
#include<deque>

using namespace std; 

class HitCounter {
    public: 
        long long window; 
        unordered_map<long long, long long> keyFreq;
        deque<pair<long long, long long>> orderedHits;
    HitCounter(long long w): window(w) {}
    void hit(long long key, long long ts) {
        
        long long prev = ts- window;
        while(!orderedHits.empty() && orderedHits.front().first <=  prev){
            keyFreq[orderedHits.front().second]--;
            if(keyFreq[orderedHits.front().second] == 0) keyFreq.erase(orderedHits.front().second);
            orderedHits.pop_front();
            
        }
        keyFreq[key]++;
        orderedHits.push_back({ts, key});
    }
    long long getCounts(long long key, long long ts) { // ts = now
        long long prev = ts- window;
        while(!orderedHits.empty() && orderedHits.front().first <=  prev){
            keyFreq[orderedHits.front().second]--;
            if(keyFreq[orderedHits.front().second] == 0) keyFreq.erase(orderedHits.front().second);
            orderedHits.pop_front();
        }
        if(keyFreq.find(key) != keyFreq.end()) {
            return keyFreq[key];
        }
        return 0; 
    }
    long long getAllCounts(long long ts) { // ts = now
        long long prev = ts- window;
        while(!orderedHits.empty() && orderedHits.front().first <=  prev){
            keyFreq[orderedHits.front().second]--;
            if(keyFreq[orderedHits.front().second] == 0) keyFreq.erase(orderedHits.front().second);
            orderedHits.pop_front();
        }
        return orderedHits.size(); 
    }
        
};
int main() {
    HitCounter *hitCounter = new HitCounter(300); // seconds
    hitCounter->hit(3,1000); // seconds. // sorted 
    hitCounter->hit(3,1002);
    hitCounter->hit(2,1004);
    hitCounter->hit(2,1300);
    hitCounter->hit(3,1302);
    cout<<hitCounter->getCounts(3,1303)<<endl;
    cout<<hitCounter->getAllCounts(1303)<<endl;
    return 0; 
}


// problem 2 

/*


*/