Like Redis(Not distributed):

Requirements:

We have to do low-level design for a Cache system. The cache that we will design will have to support the following operations:

1. Put: This will allow the user to put a value against a key in the cache.
2. Get: This will allow the user to get the previously saved value using a key.
3. Eviction: The cache should also support the removal of some keys in case the cache is full, and we try to add a new key value.
4. PrintCache: Prints the size, least recently used key, and most recently used key.
5. FlushAll: Cleans up the cache.

Expectations:

1. Code should be functionally correct.
2. Code should be modular and readable. Clean and professional level code.
3. Code should be extensible and scalable. Means it should be able to accommodate new requirements with minimal changes.
4. Code should have good OOPs design.