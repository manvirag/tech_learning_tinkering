/*

**Design Question:** Implement a simplified **Randomized Set** data structure

- Functional Requirements:
    - Insert, delete, search — all in average constant time
    - Should support random element access
    - Should handle duplicates if required
- Proposed a hybrid structure using:
    - HashMap (for O(1) lookups)
    - Array/List (for O(1) random access)
- Implemented insert and getRandom
- Covered:
    - Tradeoffs between space and time
    - Handling edge cases (like deleting the last element)
    - Testing strategy for random behavior

    
*/