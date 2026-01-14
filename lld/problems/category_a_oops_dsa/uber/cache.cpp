/*

https://leetcode.com/discuss/post/2508414/sde-2-uber-august-2022-rejection-by-anon-wmhv/

You are designing a cache system for uber with following requirements

1. Adding a key, value pair to cache.

2. Removing key, value pair from cache.

3. Getting value from cache provided a key

4. Getting random key value pair from cache

Optimizied TC and SC expected for each operation.



- LLD - design a data structure like hashmap which has get, set and delete functions for key, values like a hashmap. additionally it has a getRandom function on it which returns a random key,value paid from the hashmap. I kept the hashmap's keys in an array and generated a random number less than the length of the array and returned the key at that index and the corresponding value. Extension was also to have concurrent operations on this data structure to be used by multiple threads at the same time.

→

*/