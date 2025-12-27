# C++ STL - Quick Reference

## Headers
- **Containers**: `#include <vector>`, `#include <set>`, `#include <map>`, etc.
- **Algorithms**: `#include <algorithm>`
- **Utilities**: `#include <utility>` (for pair), `#include <tuple>`

## Essential Containers

### 1. vector
**Dynamic array, random access, fast insertion at end**

```cpp
#include <vector>

std::vector<int> vec;
std::vector<int> vec2{1, 2, 3};           // Initialize
std::vector<int> vec3(5, 0);              // 5 elements, all 0

// Operations
vec.push_back(10);                        // Add at end
vec.pop_back();                           // Remove from end
vec.size();                               // Size
vec.empty();                              // Check if empty
vec[i];                                   // Access element
vec.at(i);                                // Access with bounds check
vec.front();                              // First element
vec.back();                               // Last element
vec.clear();                              // Remove all
vec.insert(vec.begin() + 2, 5);           // Insert at position
vec.erase(vec.begin() + 1);               // Erase at position
vec.resize(10);                           // Resize

// Iterator Operations (Random Access Iterator)
auto it = vec.begin();                    // Get iterator
auto it2 = vec.end();                      // End iterator
auto it3 = vec.begin() + 3;               // ✅ Can add/subtract integers
it++;                                     // ✅ Increment
++it;                                     // ✅ Pre-increment
it--;                                     // ✅ Decrement
--it;                                     // ✅ Pre-decrement
it += 2;                                  // ✅ Add integer
it -= 2;                                  // ✅ Subtract integer
int diff = it2 - it;                      // ✅ Distance between iterators
int val = it[2];                          // ✅ Random access
```

### 2. string
**String container, similar to vector<char>**

```cpp
#include <string>

std::string str = "hello";
std::string str2("world");

// Operations
str += " world";                          // Concatenate
str.size();                               // Length
str.empty();                              // Check empty
str[i];                                   // Access character
str.substr(0, 3);                         // Substring (start, length)
str.find("lo");                           // Find substring, returns position or npos
str.replace(0, 2, "hi");                  // Replace (pos, len, new_str)
str.insert(2, "x");                       // Insert at position
str.erase(1, 2);                          // Erase (pos, len)

// Iterator Operations (Random Access Iterator - same as vector)
auto it = str.begin();                    // Get iterator
auto it2 = str.end();                      // End iterator
it++;                                     // ✅ Increment
it--;                                     // ✅ Decrement
it += 2;                                  // ✅ Add integer
it -= 2;                                  // ✅ Subtract integer
int diff = it2 - it;                      // ✅ Distance
char c = *it;                             // ✅ Dereference
```

### 3. set
**Sorted unique elements, O(log n) operations**

```cpp
#include <set>

std::set<int> s;
std::set<int> s2{3, 1, 4, 1, 5};          // {1, 3, 4, 5} - sorted, unique

// Operations
s.insert(10);                             // Insert
s.erase(10);                               // Erase
s.find(10);                                // Find, returns iterator or s.end()
s.count(10);                               // Count (0 or 1 for set)
s.size();                                  // Size
s.empty();                                 // Check empty
s.lower_bound(5);                          // First element >= 5
s.upper_bound(5);                         // First element > 5
s.begin();                                 // Smallest element
s.rbegin();                                // Largest element (reverse)

// Iterator Operations (Bidirectional Iterator)
auto it = s.begin();                      // Get iterator
auto it2 = s.find(5);                     // Get iterator to element
it++;                                     // ✅ Increment (forward)
++it;                                     // ✅ Pre-increment
it--;                                     // ✅ Decrement (backward)
--it;                                     // ✅ Pre-decrement
// it += 2;                               // ❌ Cannot add integers (tree structure, must traverse step-by-step)
// it -= 2;                               // ❌ Cannot subtract integers (tree structure, must traverse step-by-step)
int val = *it;                            // ✅ Dereference
if (it != s.end()) { }                    // ✅ Compare with end()
```

### 4. map
**Key-value pairs, sorted by key, O(log n) operations**

```cpp
#include <map>

std::map<int, std::string> m;
std::map<int, std::string> m2{{1, "one"}, {2, "two"}};

// Operations
m[5] = "five";                            // Insert/update
m.insert({3, "three"});                   // Insert
m.erase(5);                                // Erase by key
m.find(5);                                 // Find, returns iterator or m.end()
m.count(5);                                // Count (0 or 1)
m.size();                                  // Size
m.empty();                                 // Check empty
m.lower_bound(5);                          // First key >= 5
m.upper_bound(5);                          // First key > 5
for (auto& [key, val] : m) { }            // Iterate (C++17)

// Iterator Operations (Bidirectional Iterator)
auto it = m.begin();                      // Get iterator
auto it2 = m.find(5);                     // Get iterator to key
it++;                                     // ✅ Increment (forward)
++it;                                     // ✅ Pre-increment
it--;                                     // ✅ Decrement (backward)
--it;                                     // ✅ Pre-decrement
// it += 2;                               // ❌ Cannot add integers (tree structure, must traverse step-by-step)
int key = it->first;                      // ✅ Access key
std::string val = it->second;             // ✅ Access value
if (it != m.end()) { }                    // ✅ Compare with end()
```

### 5. unordered_set
**Hash set, O(1) average, O(n) worst case**

```cpp
#include <unordered_set>

std::unordered_set<int> us;
std::unordered_set<int> us2{3, 1, 4, 1, 5};

// Operations
us.insert(10);                            // Insert
us.erase(10);                              // Erase
us.find(10);                               // Find, returns iterator or us.end()
us.count(10);                              // Count (0 or 1)
us.size();                                 // Size
us.empty();                                // Check empty

// Iterator Operations (Forward Iterator)
auto it = us.begin();                     // Get iterator
auto it2 = us.find(10);                   // Get iterator to element
it++;                                     // ✅ Increment (forward only)
++it;                                     // ✅ Pre-increment
// it--;                                  // ❌ Cannot decrement (hash table uses forward-only linked lists, no reverse pointers)
// it += 2;                               // ❌ Cannot add integers (hash buckets not in order, must traverse one by one)
int val = *it;                            // ✅ Dereference
if (it != us.end()) { }                   // ✅ Compare with end()
```

### 6. unordered_map
**Hash map, O(1) average, O(n) worst case**

```cpp
#include <unordered_map>

std::unordered_map<int, std::string> um;
std::unordered_map<int, std::string> um2{{1, "one"}, {2, "two"}};

// Operations
um[5] = "five";                           // Insert/update
um.insert({3, "three"});                  // Insert
um.erase(5);                               // Erase by key
um.find(5);                                // Find, returns iterator or um.end()
um.count(5);                               // Count (0 or 1)
um.size();                                 // Size
um.empty();                                // Check empty
for (auto& [key, val] : um) { }           // Iterate (C++17)

// Iterator Operations (Forward Iterator)
auto it = um.begin();                     // Get iterator
auto it2 = um.find(5);                    // Get iterator to key
it++;                                     // ✅ Increment (forward only)
++it;                                     // ✅ Pre-increment
// it--;                                  // ❌ Cannot decrement (hash table uses forward-only linked lists, no reverse pointers)
int key = it->first;                      // ✅ Access key
std::string val = it->second;             // ✅ Access value
if (it != um.end()) { }                   // ✅ Compare with end()
```

### 7. queue
**FIFO (First In First Out)**

```cpp
#include <queue>

std::queue<int> q;

// Operations
q.push(10);                                // Add to back
q.pop();                                   // Remove from front
q.front();                                 // Front element
q.back();                                  // Back element
q.size();                                  // Size
q.empty();                                 // Check empty

// Iterator Operations
// ❌ Queue does not support iterators
// Must use front() and pop() to access elements
// Reason: Queue is a container adapter that intentionally hides iteration
//         to enforce FIFO discipline (only front/back access allowed)
```

### 8. stack
**LIFO (Last In First Out)**

```cpp
#include <stack>

std::stack<int> st;

// Operations
st.push(10);                               // Add to top
st.pop();                                  // Remove from top
st.top();                                  // Top element
st.size();                                 // Size
st.empty();                                // Check empty

// Iterator Operations
// ❌ Stack does not support iterators
// Must use top() and pop() to access elements
```

### 9. priority_queue
**Max-heap by default, O(log n) insert/extract**

```cpp
#include <queue>

std::priority_queue<int> pq;              // Max-heap
std::priority_queue<int, std::vector<int>, std::greater<int>> minPq;  // Min-heap

// Operations
pq.push(10);                               // Insert
pq.pop();                                  // Remove top
pq.top();                                  // Top element (largest for max-heap)
pq.size();                                 // Size
pq.empty();                                // Check empty

// Iterator Operations
// ❌ Priority queue does not support iterators
// Must use top() and pop() to access elements

// Custom comparator
class Compare {
public:
    bool operator()(pair<int, int> a, pair<int, int> b) {
        return a.second > b.second;        // Min-heap by second element
    }
};
priority_queue<pair<int, int>, vector<pair<int, int>>, Compare> customPq;
```

### 10. deque
**Double-ended queue, fast insertion at both ends**

```cpp
#include <deque>

std::deque<int> dq;

// Operations
dq.push_back(10);                         // Add at end
dq.push_front(5);                         // Add at front
dq.pop_back();                            // Remove from end
dq.pop_front();                           // Remove from front
dq[i];                                     // Access element
dq.size();                                // Size
dq.empty();                               // Check empty

// Iterator Operations (Random Access Iterator - same as vector)
auto it = dq.begin();                     // Get iterator
auto it2 = dq.end();                      // End iterator
it++;                                     // ✅ Increment
it--;                                     // ✅ Decrement
it += 2;                                  // ✅ Add integer
it -= 2;                                  // ✅ Subtract integer
int diff = it2 - it;                      // ✅ Distance
int val = it[2];                          // ✅ Random access
```

**Why deque supports iterators but queue doesn't:**
- **Deque**: Stores elements in multiple small arrays (chunks). The deque knows which chunk contains each element, so it can calculate positions and support iterators. Like having multiple shelves where you can find any item by knowing which shelf and position.
- **Queue**: Wrapper container that intentionally hides all elements except front/back. Even though it may use deque internally, it doesn't expose iterators to enforce FIFO behavior - you can only access the front element, not iterate through all elements.

### 11. pair
**Two values together**

```cpp
#include <utility>

std::pair<int, std::string> p{1, "one"};
std::pair<int, std::string> p2 = std::make_pair(2, "two");

// Operations
p.first;                                   // First element
p.second;                                 // Second element
p = {3, "three"};                         // Assign (C++11)

// Structured bindings (C++17)
auto [a, b] = p;                           // Unpack pair
```

## Essential Algorithms

### Sorting and Searching

```cpp
#include <algorithm>

std::vector<int> vec = {3, 1, 4, 1, 5};

// Sort
std::sort(vec.begin(), vec.end());        // Ascending
std::sort(vec.begin(), vec.end(), std::greater<int>());  // Descending
std::sort(vec.begin(), vec.end(), [](int a, int b) { return a > b; });  // Lambda

// Binary search (requires sorted container)
bool found = std::binary_search(vec.begin(), vec.end(), 4);

// Lower bound (first element >= value)
auto it = std::lower_bound(vec.begin(), vec.end(), 4);

// Upper bound (first element > value)
auto it2 = std::upper_bound(vec.begin(), vec.end(), 4);
```

### Finding Elements

```cpp
// Find
auto it = std::find(vec.begin(), vec.end(), 4);
if (it != vec.end()) { /* found */ }

// Find if (with condition)
auto it = std::find_if(vec.begin(), vec.end(), [](int x) { return x > 3; });

// Count
int cnt = std::count(vec.begin(), vec.end(), 4);

// Count if
int cnt = std::count_if(vec.begin(), vec.end(), [](int x) { return x > 3; });
```

### Min/Max

```cpp
// Min/Max element
auto it = std::min_element(vec.begin(), vec.end());
auto it2 = std::max_element(vec.begin(), vec.end());
int minVal = *it;
int maxVal = *it2;

// Min/Max of two values
int m = std::min(a, b);
int M = std::max(a, b);
```

### Modifying Algorithms

```cpp
// Reverse
std::reverse(vec.begin(), vec.end());

// Fill
std::fill(vec.begin(), vec.end(), 0);

// Transform
std::transform(vec.begin(), vec.end(), vec.begin(), [](int x) { return x * 2; });

// Unique (removes consecutive duplicates, requires sorted)
std::sort(vec.begin(), vec.end());
auto it = std::unique(vec.begin(), vec.end());
vec.erase(it, vec.end());                 // Remove duplicates
```

### Accumulate

```cpp
#include <numeric>

// Sum
int sum = std::accumulate(vec.begin(), vec.end(), 0);

// Product
int product = std::accumulate(vec.begin(), vec.end(), 1, std::multiplies<int>());

// Custom operation
int result = std::accumulate(vec.begin(), vec.end(), 0, [](int a, int b) { return a + b * b; });
```

### Permutations

```cpp
// Next permutation
std::vector<int> vec = {1, 2, 3};
do {
    // Process permutation
} while (std::next_permutation(vec.begin(), vec.end()));

// Previous permutation
std::prev_permutation(vec.begin(), vec.end());
```

## Common Patterns

### Iterating Containers

#### Vector
```cpp
std::vector<int> vec = {1, 2, 3, 4, 5};

// Range-based for loop (by value)
for (int x : vec) {
    std::cout << x << " ";
}

// Range-based for loop (by reference - can modify)
for (int& x : vec) {
    x *= 2;  // Modify elements
}

// Range-based for loop (const reference - read only)
for (const int& x : vec) {
    std::cout << x << " ";
}

// With index
for (size_t i = 0; i < vec.size(); i++) {
    std::cout << vec[i] << " ";
}

// Iterators
for (auto it = vec.begin(); it != vec.end(); ++it) {
    std::cout << *it << " ";
}

// Reverse iteration
for (auto it = vec.rbegin(); it != vec.rend(); ++it) {
    std::cout << *it << " ";  // 5, 4, 3, 2, 1
}
```

#### String
```cpp
std::string str = "hello";

// Range-based for loop
for (char c : str) {
    std::cout << c << " ";
}

// With index
for (size_t i = 0; i < str.size(); i++) {
    std::cout << str[i] << " ";
}

// Iterators
for (auto it = str.begin(); it != str.end(); ++it) {
    std::cout << *it << " ";
}
```

#### Set
```cpp
std::set<int> s = {3, 1, 4, 1, 5};

// Range-based for loop (always sorted)
for (int x : s) {
    std::cout << x << " ";  // 1, 3, 4, 5 (sorted, unique)
}

// Iterators
for (auto it = s.begin(); it != s.end(); ++it) {
    std::cout << *it << " ";
}

// Reverse iteration (largest to smallest)
for (auto it = s.rbegin(); it != s.rend(); ++it) {
    std::cout << *it << " ";  // 5, 4, 3, 1
}
```

#### Map
```cpp
std::map<int, std::string> m{{1, "one"}, {2, "two"}, {3, "three"}};

// Range-based for loop (C++17 structured bindings)
for (auto& [key, value] : m) {
    std::cout << key << ": " << value << std::endl;
}

// Range-based for loop (old way)
for (auto& pair : m) {
    std::cout << pair.first << ": " << pair.second << std::endl;
}

// Iterators
for (auto it = m.begin(); it != m.end(); ++it) {
    std::cout << it->first << ": " << it->second << std::endl;
}

// Reverse iteration
for (auto it = m.rbegin(); it != m.rend(); ++it) {
    std::cout << it->first << ": " << it->second << std::endl;
}
```

#### Unordered Set/Map
```cpp
std::unordered_set<int> us = {3, 1, 4, 5};

// Range-based for loop (order is arbitrary)
for (int x : us) {
    std::cout << x << " ";
}

// Iterators
for (auto it = us.begin(); it != us.end(); ++it) {
    std::cout << *it << " ";
}

// Unordered map (C++17)
std::unordered_map<int, std::string> um{{1, "one"}, {2, "two"}};
for (auto& [key, value] : um) {
    std::cout << key << ": " << value << std::endl;
}
```

#### Queue (Cannot iterate directly)
```cpp
std::queue<int> q;
q.push(1);
q.push(2);
q.push(3);

// Queue doesn't support iteration
// Must pop elements to access
while (!q.empty()) {
    std::cout << q.front() << " ";
    q.pop();  // Remove after accessing
}
```

#### Stack (Cannot iterate directly)
```cpp
std::stack<int> st;
st.push(1);
st.push(2);
st.push(3);

// Stack doesn't support iteration
// Must pop elements to access
while (!st.empty()) {
    std::cout << st.top() << " ";
    st.pop();  // Remove after accessing
}
```

#### Priority Queue (Cannot iterate directly)
```cpp
std::priority_queue<int> pq;
pq.push(3);
pq.push(1);
pq.push(4);

// Priority queue doesn't support iteration
// Must pop elements to access (sorted order)
while (!pq.empty()) {
    std::cout << pq.top() << " ";  // 4, 3, 1 (max to min)
    pq.pop();
}
```

#### Deque
```cpp
std::deque<int> dq = {1, 2, 3, 4, 5};

// Range-based for loop
for (int x : dq) {
    std::cout << x << " ";
}

// With index
for (size_t i = 0; i < dq.size(); i++) {
    std::cout << dq[i] << " ";
}

// Iterators
for (auto it = dq.begin(); it != dq.end(); ++it) {
    std::cout << *it << " ";
}
```

#### Pair
```cpp
std::pair<int, std::string> p{1, "one"};

// Access members
std::cout << p.first << ": " << p.second << std::endl;

// Structured bindings (C++17)
auto [key, value] = p;
std::cout << key << ": " << value << std::endl;
```

**Iteration Summary:**
- **Can iterate**: vector, string, set, map, unordered_set, unordered_map, deque
- **Cannot iterate**: queue, stack, priority_queue (must pop to access)
- **Range-based for**: Simplest, use when you don't need index
- **Index-based**: Use when you need position/index
- **Iterators**: Use when you need to modify container during iteration or need iterator operations

### Iterator Categories Summary

**Random Access Iterator** (vector, string, deque):
- ✅ Increment (++it, it++)
- ✅ Decrement (--it, it--)
- ✅ Add/subtract integers (it + n, it - n)
- ✅ Compound assignment (it += n, it -= n)
- ✅ Distance (it2 - it)
- ✅ Random access (it[n])
- ✅ Comparison (<, >, <=, >=)

**Bidirectional Iterator** (set, map):
- ✅ Increment (++it, it++)
- ✅ Decrement (--it, it--)
- ❌ Cannot add/subtract integers (tree structure, must traverse step-by-step)
- ❌ Cannot use random access
- ✅ Comparison (==, !=)

**Forward Iterator** (unordered_set, unordered_map):
- ✅ Increment (++it, it++)
- ❌ Cannot decrement (hash table uses forward-only linked lists, no reverse pointers)
- ❌ Cannot add/subtract integers (hash buckets not in order, must traverse one by one)
- ❌ Cannot use random access
- ✅ Comparison (==, !=)

**No Iterators** (queue, stack, priority_queue):
- ❌ No iterator support
- Must use container-specific access methods

### Checking Existence
```cpp
// Set/Map
if (s.find(x) != s.end()) { }             // Found
if (s.count(x)) { }                       // Found (0 or 1)

// Vector
if (std::find(vec.begin(), vec.end(), x) != vec.end()) { }  // Found
```

### Custom Comparator
```cpp
// Sort pairs by second element
std::vector<std::pair<int, int>> pairs;
std::sort(pairs.begin(), pairs.end(), 
          [](auto a, auto b) { return a.second < b.second; });
```

## Common Pitfalls

```cpp
// Wrong: Accessing empty container
std::vector<int> vec;
int x = vec[0];                            // ❌ Undefined behavior

// Correct: Check size first
if (!vec.empty()) {
    int x = vec[0];                        // ✅
}

// Wrong: Modifying set/map while iterating
for (auto it = s.begin(); it != s.end(); ++it) {
    s.erase(it);                           // ❌ Invalidates iterator
}

// Correct: Use erase return value
for (auto it = s.begin(); it != s.end(); ) {
    it = s.erase(it);                     // ✅
}

// Wrong: Comparing floating point in set/map
std::set<double> s;
s.insert(0.1 + 0.2);
if (s.find(0.3) != s.end()) { }           // ❌ May not find due to precision

// Correct: Use epsilon or custom comparator
```

