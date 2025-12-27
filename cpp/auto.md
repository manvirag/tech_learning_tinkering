# C++ `auto` Keyword - Quick Reference

## Basic Usage
- **Type deduction**: Compiler infers type from initializer
- **C++11 feature**: Requires C++11 or later

```cpp
auto x = 10;           // int
auto y = 3.14;         // double
auto z = "hello";      // const char*
auto vec = std::vector<int>();  // std::vector<int>
```

## When to Use

### ✅ Use `auto` when:
- **Iterators**: `auto it = vec.begin();`
- **Complex types**: `auto result = someComplexFunction();`
- **Lambda functions**: `auto lambda = [](int x) { return x * 2; };`
- **Template-heavy code**: Reduces verbosity
- **Range-based for loops**: `for (auto& item : container)`

### ❌ Avoid `auto` when:
- **Initialization unclear**: `auto x = getValue();` (what type?)
- **API contracts**: Function return types should be explicit
- **Portability**: When type size matters (int vs long)

## Important Rules

### Type Deduction
- **By value**: `auto x = expr;` → removes references, const, volatile
- **By reference**: `auto& x = expr;` → preserves reference and const

```cpp
int i = 10;
const int& ref = i;

auto a = ref;        // int (const and & removed)
auto& b = ref;       // const int& (preserved)
const auto& c = ref; // const int&
```

### Array/Pointer Decay
```cpp
int arr[5];
auto x = arr;        // int* (decays to pointer)
auto& y = arr;       // int(&)[5] (reference to array)
```

## Competitive Programming Tips

1. **STL iterators**: `auto it = s.lower_bound(x);`
2. **Pair/tuple unpacking**: `auto [a, b] = make_pair(1, 2);` (C++17)
3. **Range loops**: `for (auto& [key, val] : map)` (C++17)
4. **Lambda captures**: `auto f = [&](int x) { ... };`

## Common Pitfalls

**size_t vs int**: `size_t` is an unsigned integer type used for sizes/counts (STL containers return `size_t`). Unlike `int` (signed, can be negative), `size_t` is always non-negative and can hold larger values. Mixing them causes signed/unsigned comparison warnings. `size_t` can be 32-bit or 64-bit depending on the platform (32-bit systems use 32-bit, 64-bit systems use 64-bit).

```cpp
// Note: auto deduces to size_t (unsigned), not int
auto size = vec.size();  // size_t (unsigned)
int index = -1;
if (index < size) { }  // ⚠️ Signed vs unsigned comparison warning

// Correct: Use size_t explicitly or cast when comparing with int
std::size_t size = vec.size();
// OR cast: if (index >= 0 && static_cast<std::size_t>(index) < size)

// Wrong: auto removes const
const int x = 5;
auto y = x;  // y is int, not const int

// Correct:
const auto y = x;  // or auto const y = x;
```

## C++14/17 Enhancements

- **Return type deduction**: `auto func() { return 42; }`
- **Lambda parameters**: `[](auto x) { return x * 2; }` (C++14)
- **Structured bindings**: `auto [a, b] = pair;` (C++17)

