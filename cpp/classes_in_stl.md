Here’s a detailed, structured, *non-summarized* notes document on **using classes with STL data structures and algorithms** — including everything we discussed: vectors, sets, maps, unordered maps, and priority queues, with deep dives into comparison, hashing, and custom behavior. You can treat this as a full reference for future use.

---

## 🧠 Using Classes in STL Containers and Algorithms

### 1️⃣ General Rule

When using **user-defined classes** with STL containers or algorithms, STL must know **how to compare**, **hash**, or **order** your objects.
You must therefore provide:

* `operator<` — for ordered containers (like `std::set`, `std::map`, `std::priority_queue`, and sorting algorithms).
* `operator==` and hash — for unordered containers (`std::unordered_map`, `std::unordered_set`).
* A **custom comparator struct or lambda** if you don’t want to modify the class.
  These define how STL internally organizes and looks up your class objects.

---

## 🧩 Sequence Containers — `vector`, `list`, `deque`, `queue`, `stack`

These containers just **store** elements. They don’t require ordering or hashing, so you can directly use your class as a type:

```cpp
std::vector<Box> v;
std::list<Box> l;
std::deque<Box> d;
```

All basic operations (`push_back`, `insert`, `pop`, `erase`, etc.) work normally.

### When using algorithms (like `sort`, `find`, `unique`, etc.)

STL algorithms such as `std::sort`, `std::find`, or `std::binary_search` depend on comparisons.
For `std::sort` or `std::binary_search`, your class must define `operator<`:

```cpp
class Box {
public:
    int volume;
    Box(int v = 0): volume(v) {}
    bool operator<(const Box &other) const {
        return volume < other.volume;
    }
};
```

Then you can do:

```cpp
std::sort(v.begin(), v.end());
```

If you don’t want to modify your class, provide a comparator:

```cpp
std::sort(v.begin(), v.end(), [](const Box &a, const Box &b) {
    return a.volume < b.volume;
});
```

Without either `operator<` or a comparator, the compiler cannot sort or compare, and compilation fails.

---

## 🏷️ Associative Containers — `set`, `multiset`

These containers **maintain sorted order** internally using a binary search tree (Red-Black Tree).
To insert and order objects of a class, STL needs a *strict weak ordering*, usually via `operator<`.

```cpp
class Box {
public:
    int volume;
    Box(int v = 0): volume(v) {}
    bool operator<(const Box &other) const {
        return volume < other.volume;
    }
};
```

Now:

```cpp
std::set<Box> s;
s.insert(Box(10));
s.insert(Box(5));
s.insert(Box(20));
```

Elements will automatically stay sorted by `volume`.
Operations like `find`, `lower_bound`, `upper_bound`, `equal_range` all rely on `<`.

### ✅ Important: Only `<` is required

For `set` and `multiset`, STL never uses `==` directly — it just checks both `!(a < b)` and `!(b < a)` to detect equality.
You do **not** need to define `operator==` for ordered sets.

### Custom comparator

You can override the default comparison using:

```cpp
struct CompareBox {
    bool operator()(const Box &a, const Box &b) const {
        return a.volume < b.volume;
    }
};

std::set<Box, CompareBox> s;
```

---

## 🗝️ Ordered Map — `std::map` and `std::multimap`

A `std::map` is a tree-based key–value structure sorted by **key** using `operator<`.
If you use a class as a key, it must define `<` or a custom comparator must be given.

### Example

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
m[Box(10)] = "A";
m[Box(20)] = "B";
```

The map automatically sorts by `Box::volume`.

### Custom comparator

```cpp
struct CompareBox {
    bool operator()(const Box &a, const Box &b) const {
        return a.volume < b.volume;
    }
};
std::map<Box, std::string, CompareBox> m;
```

### ❌ Key immutability

Keys in `map` are **const** internally (`pair<const Key, T>`).
You cannot (and must not) modify a key once inserted, or the internal tree order breaks.
This is why:

```cpp
const_cast<Box&>(it->first).volume = 50; // ❌ undefined behavior
```

is illegal — it corrupts the tree.
This also applies to primitive types (e.g., `int` key).

---

## ⚡ Unordered Map — `std::unordered_map`, `std::unordered_multimap`

These are **hash-based** containers.
Instead of `<`, they need two things:

1. `operator==` to check equality
2. A **hash function** to decide which bucket the key belongs to.

### Example class key

```cpp
class Box {
public:
    int volume;
    Box(int v = 0): volume(v) {}
    bool operator==(const Box &other) const {
        return volume == other.volume;
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

### Combining multiple fields in a hash

For composite types:

```cpp
std::size_t operator()(const std::pair<int, int> &p) const {
    return std::hash<int>()(p.first) ^ (std::hash<int>()(p.second) << 1);
}
```

* `std::hash<int>()(x)` computes a hash for `x`.
* `<< 1` shifts bits to spread patterns.
* `^` (XOR) combines the two into one hash.

---

## 🧭 When to use struct vs function

* **Ordered containers (set/map)** → use `operator<` inside class or pass a **comparator struct/lambda**.
* **Unordered containers** → need `operator==` and a **hash functor (struct)**.
  `priority_queue` also uses a **comparator struct or lambda** (like `set`).

---

## 🚫 Keys that are unsafe or problematic

| Type                           | Problem                          | Explanation                                                     |
| ------------------------------ | -------------------------------- | --------------------------------------------------------------- |
| Mutable objects                | ❌ Breaks order or hash           | Changing keys after insertion invalidates container invariants. |
| Floating-point                 | ⚠️ NaN != NaN, rounding errors   | Keys may fail lookup or ordering.                               |
| Pointer types                  | ⚠️ Compare by address, not value | May not match logical equality.                                 |
| Types without `<` or hash      | ❌ Compiler error                 | No valid ordering or hashing defined.                           |
| Large containers (vector, map) | ⚠️ Slow                          | Comparison or hashing O(n).                                     |

---

## ⚙️ STL Containers as Keys

* **In `map`:** Works automatically if the key type defines `<`.
  STL containers like `string`, `pair`, `tuple`, `vector`, `array`, `set` already have lexicographic `<`.
* **In `unordered_map`:** You must define a **hash** and **==** manually.
  STL provides `std::hash` only for basic and string types.

Example:

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

## ⛰️ Priority Queue with Classes

A `std::priority_queue` keeps the *largest* (default) element at the top, using `operator<`.
If you store class objects, you must define how to compare them.

### With `operator<`:

```cpp
class Box {
public:
    int volume;
    bool operator<(const Box &other) const {
        return volume < other.volume;
    }
};
std::priority_queue<Box> pq;
```

### With custom comparator:

```cpp
struct CompareBox {
    bool operator()(const Box &a, const Box &b) const {
        return a.volume < b.volume;
    }
};
std::priority_queue<Box, std::vector<Box>, CompareBox> pq;
```

### With lambda:

```cpp
auto cmp = [](const Box &a, const Box &b){ return a.volume < b.volume; };
std::priority_queue<Box, std::vector<Box>, decltype(cmp)> pq(cmp);
```

* Default: **max-heap** (`largest` on top).
* Reverse comparator to make **min-heap**.

---

## 🔧 Comparison Summary

| Container                        | Needs                      | Comparison Type   | How to Provide                          |
| -------------------------------- | -------------------------- | ----------------- | --------------------------------------- |
| `vector`, `list`, `deque`        | Only for sorting/searching | `<` or comparator | `operator<` or lambda                   |
| `set`, `multiset`                | Required                   | `<`               | `operator<` or custom comparator struct |
| `map`, `multimap`                | Required (on key)          | `<`               | `operator<` or comparator struct        |
| `unordered_set`, `unordered_map` | Required (on key)          | `==` and hash     | `operator==` + hash struct              |
| `priority_queue`                 | Required                   | `<` or custom     | `operator<` or comparator struct        |

---

## 🧩 Algorithms with Classes

* **`std::sort` / `std::stable_sort`** → requires `<` or comparator.
* **`std::find`, `std::count`** → requires `==`.
* **`std::binary_search`, `std::lower_bound`** → requires `<`.
* **`std::max_element`, `std::min_element`** → requires `<`.
* **`std::unique`** → requires `==`.

So, for your class to fully support STL algorithms, it should define both:

```cpp
bool operator<(const Class &other) const;
bool operator==(const Class &other) const;
```

---

✅ **Final Note:**
When you use a class in STL:

* Ordered containers → define `<` or comparator struct/lambda.
* Unordered containers → define `==` and hash functor.
* Algorithms → define `<` and `==`.
* Never mutate keys.
* Be careful with floating points, pointers, and mutable containers as keys.
* `priority_queue` always needs comparison logic.

With these, you can safely use custom classes with any STL container or algorithm without undefined behavior, compile-time errors, or logical bugs.
