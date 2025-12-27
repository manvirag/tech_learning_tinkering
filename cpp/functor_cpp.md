# C++ Functors - Quick Reference

## Basic Usage
- **Definition**: Class that overloads `operator()` to make objects callable like functions
- **Also called**: Function objects

```cpp
class Adder {
public:
    int operator()(int a, int b) {
        return a + b;
    }
};

Adder add;
int result = add(3, 4);  // Calls operator(), returns 7
```

## When to Use

### ✅ Use functors when:
- **Stateful operations**: Need to store data between calls
- **STL algorithms**: `std::sort`, `std::transform`, `std::for_each`
- **Reusable logic**: Same logic with different parameters

### ❌ Prefer lambdas when:
- **Simple one-time use**: `std::sort(vec.begin(), vec.end(), [](int a, int b) { return a > b; });`
- **Local scope only**: No need to reuse elsewhere
- **C++11+**: Lambdas are more concise for simple cases

## Important Rules

### Stateful Functors
Functors can maintain state between calls, unlike regular functions:

```cpp
class Counter {
public:
    int count = 0;
    void operator()(int x) {
        count++;
        std::cout << "Processed: " << x << std::endl;
    }
};

Counter counter;
std::for_each(vec.begin(), vec.end(), counter);
// counter.count now holds total processed items
```

### STL Algorithm Usage
```cpp
class Greater {
public:
    bool operator()(int a, int b) const {
        return a > b;
    }
};

std::vector<int> vec = {3, 1, 4, 1, 5};
std::sort(vec.begin(), vec.end(), Greater());  // Descending order
```

## Competitive Programming Tips

1. **Custom comparators**: `std::sort`, `std::priority_queue`, `std::set`
   ```cpp
   class Compare {
   public:
       bool operator()(pair<int,int> a, pair<int,int> b) {
           return a.second > b.second;  // Min-heap by second element
       }
   };
   priority_queue<pair<int,int>, vector<pair<int,int>>, Compare> pq;
   ```

2. **Transform operations**: `std::transform` with state
   ```cpp
   class Multiply {
   public:
       int factor;
       Multiply(int f) : factor(f) {}
       int operator()(int x) { return x * factor; }
   };
   std::transform(vec.begin(), vec.end(), vec.begin(), Multiply(2));
   ```

3. **Predicates**: `std::find_if`, `std::count_if`
   ```cpp
   class IsEven {
   public:
       bool operator()(int x) const { return x % 2 == 0; }
   };
   auto it = std::find_if(vec.begin(), vec.end(), IsEven());
   ```

## Common Pitfalls

```cpp
// Wrong: Forgetting const in comparator
class Compare {
public:
    bool operator()(int a, int b) {  // Missing const
        return a < b;
    }
};
// Container comparators (std::set, std::map, std::priority_queue) require const
// Compiler will error: "passing 'const Compare' as 'this' argument discards qualifiers"

// Correct:
class Compare {
public:
    bool operator()(int a, int b) const {  // Add const
        return a < b;
    }
};
// Rule: Use const for comparators/predicates that don't modify state
// How to know: Compiler error or best practice - always use const for pure functions

// Wrong: Modifying state in const context
class Accumulator {
public:
    int sum = 0;
    int operator()(int x) const {
        // sum += x;  // ❌ Can't modify in const function
        return sum;
    }
};

// Correct: Remove const if state needs modification
class Accumulator {
public:
    mutable int sum = 0;  // mutable allows modification in const
    int operator()(int x) const {
        sum += x;
        return sum;
    }
};
```

## Functor vs Lambda

```cpp
// Functor - reusable, can be stored
class Square {
public:
    int operator()(int x) const { return x * x; }
};
Square square;
std::transform(vec.begin(), vec.end(), vec.begin(), square);

// Lambda - concise, one-time use
std::transform(vec.begin(), vec.end(), vec.begin(), 
               [](int x) { return x * x; });

// Lambda with capture - similar to stateful functor
int factor = 2;
std::transform(vec.begin(), vec.end(), vec.begin(),
               [factor](int x) { return x * factor; });
```

