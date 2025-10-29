Perfect — below are **complete, deep-dive, non-summarized notes** dedicated entirely to **`std::map`**, **`std::unordered_map`**, **`std::multimap`**, and **`std::unordered_multimap`**, with every technical and conceptual detail explained clearly, including rules for using **classes as keys**, how comparisons and hashing work, key constraints, pitfalls, and real examples.
This is written to be your **one-stop reference** for map-type containers in C++.

---

# 🧭 Deep Dive Notes — `std::map`, `std::unordered_map`, `std::multimap`, `std::unordered_multimap`

---

## 🔹 1. Overview: What “map” Containers Are

A **map** is an associative container that stores key–value pairs.
Each key is unique (except in multimap variants), and the container automatically manages:

* **Storage** (insertion, deletion)
* **Lookup** (by key)
* **Ordering or hashing** (depending on type)

There are two broad families:

| Type                      | Internally uses               | Key uniqueness     | Order maintained    | Complexity   | Requires                  |
| ------------------------- | ----------------------------- | ------------------ | ------------------- | ------------ | ------------------------- |
| `std::map`                | Red–Black tree (balanced BST) | Unique             | Sorted order by `<` | O(log n)     | `operator<` or comparator |
| `std::multimap`           | Red–Black tree                | Duplicates allowed | Sorted order        | O(log n)     | `operator<` or comparator |
| `std::unordered_map`      | Hash table                    | Unique             | No order            | Average O(1) | `operator==` and hash     |
| `std::unordered_multimap` | Hash table                    | Duplicates allowed | No order            | Average O(1) | `operator==` and hash     |

---

## 🔸 2. `std::map` — Ordered Key–Value Store

### ⚙️ Internal Mechanism

* Uses a **self-balancing binary search tree** (Red-Black Tree).
* Every node stores `(const Key, Value)` pair.
* The key part is **immutable (`const`)** once inserted to preserve tree order.
* All keys are automatically kept **sorted** by `<` or a provided comparator.

### ✅ Requirements on Key

1. Must be **comparable** — either via:

   * `operator<` defined in the class, **or**
   * A custom comparator functor passed as the 3rd template argument.
2. `<` must define a **strict weak ordering**, meaning:

   * Irreflexive: `!(a < a)`
   * Transitive: if `a < b` and `b < c`, then `a < c`
   * Consistent: `a == b` if and only if both `!(a < b)` and `!(b < a)`

### 📘 Example — Class as key using `operator<`

```cpp
class Box {
public:
    int volume;
    Box(int v = 0): volume(v) {}
    bool operator<(const Box &other) const {
        return volume < other.volume;
    }
};

std::map<Box, std::string> m;
m[Box(10)] = "Small";
m[Box(20)] = "Medium";
m[Box(30)] = "Large";
```

Internally, the map will sort by `Box::volume`.

---

### 📘 Example — Using custom comparator

```cpp
struct CompareBox {
    bool operator()(const Box &a, const Box &b) const {
        return a.volume < b.volume;
    }
};

std::map<Box, std::string, CompareBox> m;
```

Custom comparators are often used when:

* You cannot modify the class.
* You want a different sorting logic (e.g., descending order).

---

### ⚠️ Key Immutability

In `map`, the key is **const**:

```cpp
typedef std::pair<const Key, T> value_type;
```

That means:

```cpp
auto it = m.find(Box(10));
it->first.volume = 50;  // ❌ not allowed (const)
```

Even if you use `const_cast` to force-change it:

```cpp
const_cast<Box&>(it->first).volume = 50;  // ❌ undefined behavior
```

the internal BST structure breaks — the node’s position no longer matches its key ordering.
This can cause:

* Elements becoming “lost”
* Wrong lookup results
* Crashes in future insert/find operations

🔒 **Rule:** Once inserted, a key is immutable.

---

### ⚙️ Lookup and Range Search

`std::map` offers ordered search methods:

| Method             | Description                     | Uses `<` |
| ------------------ | ------------------------------- | -------- |
| `find(key)`        | Finds element with given key    | ✅        |
| `lower_bound(key)` | First element not less than key | ✅        |
| `upper_bound(key)` | First element greater than key  | ✅        |
| `equal_range(key)` | Both bounds as a pair           | ✅        |

All rely only on `<` — not `==`.
Equality in map logic means `!(a < b)` and `!(b < a)`.

---

### 🧠 Duplicate Keys — `std::multimap`

Same as `std::map`, but allows **multiple elements with equal keys**.
Still ordered by `<`.
Useful when one key corresponds to multiple values.

```cpp
std::multimap<int, std::string> mm;
mm.insert({1, "A"});
mm.insert({1, "B"}); // allowed
```

Operations like `equal_range` and `count` help retrieve all duplicates.

---

## 🔹 3. `std::unordered_map` — Hash-Based Key–Value Store

### ⚙️ Internal Mechanism

* Uses a **hash table** (array of buckets).
* Each key is hashed into a bucket index using a hash function.
* Keys in the same bucket are compared using `operator==`.

Unlike `std::map`, this one does **not maintain order**, but is faster on average.

---

### ✅ Requirements on Key

1. Must be **hashable** → needs a hash function.
2. Must be **equality-comparable** → needs `operator==`.
3. The hash and equality must be **consistent**:

   * If `a == b`, then `hash(a) == hash(b)` (always true).
   * Different keys can have same hash (collision), but equality resolves them.

---

### 📘 Example — Using class as key

```cpp
class Box {
public:
    int volume;
    Box(int v = 0): volume(v) {}
    bool operator==(const Box &o) const {
        return volume == o.volume;
    }
};

struct BoxHash {
    std::size_t operator()(const Box &b) const {
        return std::hash<int>()(b.volume);
    }
};

std::unordered_map<Box, std::string, BoxHash> m;
m[Box(10)] = "Ten";
```

Here:

* `BoxHash` defines how to compute hash value.
* `Box::operator==` defines how to compare equality when hashes collide.

---

### ⚙️ Hash Function Mechanics

The standard library provides `std::hash` for:

* All built-in types (`int`, `char`, `double`, etc.)
* `std::string`, `std::u16string`, `std::u32string`
* `std::nullptr_t`

For custom classes, you must write your own.

#### Common pattern for multiple fields

```cpp
std::size_t operator()(const std::pair<int, int> &p) const {
    return std::hash<int>()(p.first) ^ (std::hash<int>()(p.second) << 1);
}
```

Explanation:

* `std::hash<int>()(p.first)` → hash of first part.
* `<< 1` shifts bits to avoid overlap.
* `^` XOR combines both into one number.
  This spreads out bits and reduces collisions.

You can combine more fields the same way:

```cpp
h ^= std::hash<int>()(some_field) + 0x9e3779b9 + (h << 6) + (h >> 2);
```

This is known as **boost-style hash mixing** — balances hash values better.

---

### 🧩 Consistency Rule

Your `operator==` and hash must always agree:

```cpp
if (a == b)
    assert(hash(a) == hash(b)); // must hold true
```

If not, the map may lose elements or fail to find existing ones.

---

### ⚠️ Mutable Keys

Never modify keys after insertion — same rule as `map`.
In `unordered_map`, changing a key can change its hash bucket, breaking internal structure.

```cpp
auto it = m.find(Box(10));
it->first.volume = 50;  // ❌ breaks hash invariants
```

---

### ⚙️ Duplicate Keys — `std::unordered_multimap`

Same as `unordered_map`, but allows multiple equal keys.
Internally still hashed into buckets.
When you use `insert`, duplicates go into the same bucket and are linked together.

---

### 🧠 Collisions and Load Factor

* **Collision:** Different keys produce same hash → stored in same bucket.
* **Load factor:** ratio = (number of elements) / (number of buckets).
  When the load factor grows beyond a threshold, the container **rehashes** — increases bucket count and redistributes all elements.

You can control this:

```cpp
m.max_load_factor(0.7);
m.rehash(100);
```

---

## 🔍 4. Why Key Requirements Differ

| Container            | Internally Needs | Why                                            |
| -------------------- | ---------------- | ---------------------------------------------- |
| `map`                | `<` comparison   | Tree must know where to place each key node    |
| `unordered_map`      | `==` + hash      | Hash table uses equality to resolve collisions |
| `multimap`           | `<`              | Same as map, but allows duplicates             |
| `unordered_multimap` | `==` + hash      | Same as unordered_map, allows duplicates       |

This difference stems from **storage design**:

* A tree needs a **total ordering** (who goes left, who goes right).
* A hash table just needs **bucket membership** and equality checks.

---

## ⚖️ 5. Which Key Types Are Suitable

| Key Type                                      | Works with `map`       | Works with `unordered_map` | Notes                    |
| --------------------------------------------- | ---------------------- | -------------------------- | ------------------------ |
| `int`, `std::string`                          | ✅                      | ✅                          | Built-in `<`, hash, ==   |
| Custom class                                  | ✅ (with `<`)           | ✅ (with hash + ==)         | Must define operations   |
| Floating-point (`float`, `double`)            | ⚠️                     | ⚠️                         | Precision and NaN issues |
| Pointer                                       | ✅ (compares addresses) | ✅ (hashes addresses)       | Usually unintended       |
| Mutable (e.g. vector modified after insert)   | ❌                      | ❌                          | Breaks ordering/hash     |
| Containers (`std::vector`, `std::pair`, etc.) | ✅ (lexicographic `<`)  | ⚠️ (must define hash)      | Works but slower         |
| Nested maps                                   | ✅                      | ⚠️ (need custom hash)      | Works but complex        |

---

## 💡 6. STL Containers as Keys

STL containers like `std::vector`, `std::pair`, and `std::tuple` can be keys.

* For `map`: they already define `<` lexicographically.
* For `unordered_map`: you must define hash and equality.

### Example: vector key in unordered_map

```cpp
struct VecHash {
    std::size_t operator()(const std::vector<int> &v) const {
        std::size_t h = 0;
        for (int x : v)
            h ^= std::hash<int>()(x) + 0x9e3779b9 + (h << 6) + (h >> 2);
        return h;
    }
};
struct VecEqual {
    bool operator()(const std::vector<int> &a, const std::vector<int> &b) const {
        return a == b;
    }
};
std::unordered_map<std::vector<int>, std::string, VecHash, VecEqual> m;
```

---

## 🔧 7. When to Prefer Each Map Type

| Use Case                                         | Recommended Container                        | Reason                                 |
| ------------------------------------------------ | -------------------------------------------- | -------------------------------------- |
| Need sorted order of keys                        | `std::map` / `std::multimap`                 | Tree maintains order                   |
| Need fast average O(1) lookup                    | `std::unordered_map` / `unordered_multimap`  | Hash table                             |
| Need range search (`lower_bound`, `upper_bound`) | `std::map` / `std::multimap`                 | Only trees support order-based queries |
| Need duplicate keys                              | `std::multimap` or `std::unordered_multimap` | Allows repetition                      |
| Need stable iteration order                      | `std::map`                                   | Deterministic traversal                |

---

## 🧠 8. Internal Behavior Comparison

| Feature            | `map`                            | `unordered_map`                  |
| ------------------ | -------------------------------- | -------------------------------- |
| Internal structure | Balanced tree                    | Hash table                       |
| Order of keys      | Sorted                           | Unordered                        |
| Search complexity  | O(log n)                         | O(1) average, O(n) worst         |
| Memory usage       | Lower                            | Higher (buckets + hash overhead) |
| Iteration order    | Sorted                           | Bucket-dependent (pseudo-random) |
| Rehashing          | ❌                                | ✅ (on growth)                    |
| Range queries      | ✅ (`lower_bound`, `upper_bound`) | ❌                                |
| Custom key logic   | Comparator                       | Hash + equality                  |
| Duplicate keys     | multimap                         | unordered_multimap               |

---

## 🚫 9. Common Mistakes to Avoid

1. ❌ **Modifying keys** after insertion → breaks map or hash invariants.
2. ❌ **Floating-point keys** — NaN and rounding cause failed lookups.
3. ❌ **Forgetting operator== with hash** in unordered containers.
4. ❌ **Assuming order in unordered_map** — never guaranteed.
5. ⚠️ **Using heavy containers as keys** — hashing or comparing large vectors or maps is slow.

---

## ✅ 10. Key Takeaways

* `map` and `multimap` depend on **ordering** (`<` or comparator).
* `unordered_map` and `unordered_multimap` depend on **hashing** and **equality**.
* Keys are always **const** and **must never be modified**.
* Duplicates allowed only in multimap variants.
* Always ensure hash and equality are consistent.
* Avoid floating-point and mutable keys for stability.
* Use comparator structs or lambdas for custom order.
* Use hash functors for custom hashing logic.

With these, you can confidently use any class, struct, or container as a key in any of the four map types without errors, undefined behavior, or performance pitfalls.
