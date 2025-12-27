# C++ Lambda Functions - Quick Reference

## Basic Usage
- **Definition**: Anonymous function objects defined inline
- **C++11 feature**: Requires C++11 or later
- **Syntax**: `[capture](parameters) -> return_type { body }`

```cpp
// Basic lambda
auto lambda = [](int x) { return x * 2; };
int result = lambda(5);  // Returns 10

// Lambda with capture
int factor = 3;
auto multiply = [factor](int x) { return x * factor; };
int result2 = multiply(4);  // Returns 12
```

## When to Use

### ✅ Use lambdas when:
- **STL algorithms**: `std::sort`, `std::for_each`, `std::transform`
- **One-time use**: Simple functions used only once
- **Local scope**: Functions needed only in current scope
- **Callbacks**: Event handlers, async operations
- **Short functions**: Quick inline operations


## Important Rules

### Capture Modes
```cpp
int a = 10, b = 20;

// [] - Capture nothing
auto f1 = [](int x) { return x * 2; };

// [=] - Capture all by value
auto f2 = [=](int x) { return x + a + b; };  // Copies a and b

// [&] - Capture all by reference
auto f3 = [&](int x) { a += x; return a; };  // References to a and b

// [a, b] - Capture specific by value
auto f4 = [a, b](int x) { return x + a + b; };

// [&a, b] - Capture mixed (a by ref, b by value)
auto f5 = [&a, b](int x) { a += x; return a + b; };

// [=, &a] - Capture all by value, except a by reference
auto f6 = [=, &a](int x) { a += x; return a + b; };
```

### Lambda Syntax Breakdown
```cpp
[capture](parameters) mutable -> return_type { body }

// Example:
int x = 5;
auto lambda = [x](int y) mutable -> int {
    x += y;  // Can modify captured x (copy) because of mutable
    return x;
};
```

### Return Type Deduction
```cpp
// Auto deduction (C++11)
auto add = [](int a, int b) { return a + b; };  // Returns int

// Explicit return type
auto add2 = [](int a, int b) -> int { return a + b; };

// Complex return type
auto getPair = []() -> std::pair<int, int> {
    return {1, 2};
};
```

## Passing Lambda as Argument

### Method 1: Using std::function (Most Flexible)
```cpp
#include <functional>

// Function accepting lambda
void process(std::function<int(int)> func, int value) {
    int result = func(value);
    std::cout << result << std::endl;
}

// Usage
auto lambda = [](int x) { return x * 2; };
process(lambda, 5);  // Pass lambda

// Or pass inline
process([](int x) { return x * 3; }, 5);
```

### Method 2: Using Templates (Most Efficient)
```cpp
// Template function - works with any callable
template<typename Func>
void process(Func func, int value) {
    int result = func(value);
    std::cout << result << std::endl;
}

// Usage
auto lambda = [](int x) { return x * 2; };
process(lambda, 5);

// Works with lambdas with captures too
int factor = 3;
auto lambda2 = [factor](int x) { return x * factor; };
process(lambda2, 5);
```

### Method 3: Using auto Parameters (C++14)
```cpp
// Function with auto parameter
void process(auto func, int value) {
    int result = func(value);
    std::cout << result << std::endl;
}

// Usage
process([](int x) { return x * 2; }, 5);
```

### Method 4: Function Pointer (Only for No-Capture Lambdas)
```cpp
// Function pointer - only works for lambdas without captures
void process(int (*func)(int), int value) {
    int result = func(value);
    std::cout << result << std::endl;
}

// Usage - only works for [] lambdas
process([](int x) { return x * 2; }, 5);  // ✅ OK

// Won't work with captures
int factor = 3;
// process([factor](int x) { return x * factor; }, 5);  // ❌ Error
```

### Real-World Examples

#### Example 1: Callback Function
```cpp
#include <functional>
#include <vector>

class EventHandler {
public:
    void onEvent(std::function<void(int)> callback) {
        // Store callback and call later
        callback(42);
    }
};

EventHandler handler;
handler.onEvent([](int value) {
    std::cout << "Event: " << value << std::endl;
});
```

#### Example 2: STL Algorithm Wrapper
```cpp
template<typename Container, typename Func>
void applyToAll(Container& container, Func func) {
    for (auto& item : container) {
        item = func(item);
    }
}

std::vector<int> vec = {1, 2, 3, 4, 5};
applyToAll(vec, [](int x) { return x * 2; });
```

#### Example 3: Comparator Parameter
```cpp
template<typename T, typename Compare>
void customSort(std::vector<T>& vec, Compare comp) {
    std::sort(vec.begin(), vec.end(), comp);
}

std::vector<int> vec = {3, 1, 4, 1, 5};
customSort(vec, [](int a, int b) { return a > b; });  // Descending
```

### Comparison of Methods

| Method | Pros | Cons | Use When |
|--------|------|------|----------|
| `std::function` | Works with any callable, type-safe | Overhead (type erasure) | Need runtime polymorphism |
| Template | Zero overhead, most efficient | Code bloat, compile-time only | Performance critical |
| `auto` (C++14) | Simple syntax | C++14+ only | Modern C++ code |
| Function pointer | Simple | Only no-capture lambdas | C-style callbacks |

## Competitive Programming Tips

1. **STL Algorithms with Lambdas**
   ```cpp
   std::vector<int> vec = {3, 1, 4, 1, 5};
   
   // Sort descending
   std::sort(vec.begin(), vec.end(), [](int a, int b) {
       return a > b;
   });
   
   // Transform
   std::transform(vec.begin(), vec.end(), vec.begin(),
                  [](int x) { return x * 2; });
   
   // Find if
   auto it = std::find_if(vec.begin(), vec.end(),
                          [](int x) { return x > 3; });
   ```

2. **Custom Comparators**
   ```cpp
   // Priority queue (min-heap)
   priority_queue<pair<int, int>, vector<pair<int, int>>,
                  [](pair<int, int> a, pair<int, int> b) {
                      return a.second > b.second;
                  }) pq;
   
   // Set with custom comparator
   set<int, decltype([](int a, int b) { return a > b; })> s;
   ```

3. **Capture Local Variables**
   ```cpp
   int threshold = 10;
   auto count = std::count_if(vec.begin(), vec.end(),
                              [threshold](int x) {
                                  return x > threshold;
                              });
   ```

4. **Range-based for with Lambda**
   ```cpp
   std::vector<int> vec = {1, 2, 3, 4, 5};
   std::for_each(vec.begin(), vec.end(),
                 [](int& x) { x *= 2; });
   ```

## Common Pitfalls

```cpp
// Wrong: Dangling reference capture
int* ptr = new int(10);
auto lambda = [&ptr](int x) { return *ptr + x; };
delete ptr;  // ptr is deleted
// lambda(5);  // ❌ Undefined behavior: dangling reference

// Correct: Capture by value or ensure lifetime
int value = 10;
auto lambda = [value](int x) { return value + x; };  // Copy value
// OR
auto lambda2 = [&value](int x) { return value + x; };  // Ensure value lives

// Wrong: Modifying captured by-value variable
int x = 5;
auto lambda = [x](int y) {
    // x += y;  // ❌ Error: can't modify captured by-value variable
    return x + y;
};

// Correct: Use mutable keyword
auto lambda = [x](int y) mutable {
    x += y;  // ✅ OK: modifies copy of x
    return x;
};

// Wrong: Capturing in loop (all lambdas capture same reference)
std::vector<std::function<int()>> funcs;
for (int i = 0; i < 5; i++) {
    funcs.push_back([&i]() { return i; });  // ❌ All capture same i
}
// All return 5 (final value of i)

// Correct: Capture by value
for (int i = 0; i < 5; i++) {
    funcs.push_back([i]() { return i; });  // ✅ Each captures its own copy
}
// Returns 0, 1, 2, 3, 4

// Wrong: Using auto for complex lambda types
auto complex = [](auto x) { return x * 2; };
// Some templates can't deduce auto lambda types

// Correct: Use std::function or explicit type
std::function<int(int)> func = [](int x) { return x * 2; };
```

## Lambda vs Functor

```cpp
// Lambda - concise, one-time use
std::sort(vec.begin(), vec.end(), [](int a, int b) {
    return a > b;
});

// Functor - reusable, can be stored
class Greater {
public:
    bool operator()(int a, int b) const {
        return a > b;
    }
};
std::sort(vec.begin(), vec.end(), Greater());
```

## C++14/17 Enhancements

- **Generic lambdas** (C++14): `[](auto x) { return x * 2; }`
- **Init capture** (C++14): `[x = getValue()](int y) { return x + y; }`
- **constexpr lambdas** (C++17): `[]() constexpr { return 42; }`

