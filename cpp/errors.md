# C++ Error Handling - DSA & Machine Coding Reference

Concise reference for error handling in competitive programming and machine coding.

---

## 0. Simple Mental Model (Go vs C++)

**Go's Simple Way:**
```go
result, err := divide(a, b)
if err != nil {
    return err
}
// Use result
```

**C++ Equivalent:**
```cpp
auto [result, err] = divide(a, b);
if (err) {
    return err;
}
// Use result
```

**Key Difference:**
- **Go**: Built-in `(value, error)` return pattern
- **C++**: Use `std::optional` or `std::expected` (C++23) or return error codes

---

## 1. Method 1: Return Error Code (Simplest - Like Go)

**Pattern:** Return `int` error code (0 = success, non-zero = error)

```cpp
// Function returns error code, result via reference/pointer
int divide(int a, int b, int& result) {
    if (b == 0) {
        return 1;  // Error: division by zero
    }
    result = a / b;
    return 0;  // Success
}

// Usage
int result;
int err = divide(10, 2, result);
if (err != 0) {
    cout << "Error: Division by zero" << endl;
    return;
}
cout << "Result: " << result << endl;
```

**With Enum (Better):**
```cpp
enum class ErrorCode {
    SUCCESS = 0,
    DIVISION_BY_ZERO = 1,
    INVALID_INPUT = 2,
    NOT_FOUND = 3
};

ErrorCode divide(int a, int b, int& result) {
    if (b == 0) {
        return ErrorCode::DIVISION_BY_ZERO;
    }
    result = a / b;
    return ErrorCode::SUCCESS;
}

// Usage
int result;
if (auto err = divide(10, 2, result); err != ErrorCode::SUCCESS) {
    cout << "Error occurred" << endl;
    return;
}
cout << "Result: " << result << endl;
```

**Pros:**
- Simple, explicit, no exceptions
- Fast (no overhead)
- Easy to understand

**Cons:**
- Must remember to check error code
- Result passed via reference (can forget to check)

---

## 2. Method 2: std::optional (C++17) - For Simple Cases

**Pattern:** Return `std::optional<T>`, `nullopt` = error

```cpp
#include <optional>

std::optional<int> divide(int a, int b) {
    if (b == 0) {
        return std::nullopt;  // Error
    }
    return a / b;  // Success
}

// Usage
auto result = divide(10, 2);
if (!result.has_value()) {  // or: if (!result)
    cout << "Error: Division by zero" << endl;
    return;
}
cout << "Result: " << *result << endl;  // or: result.value()
```

**With Default Value:**
```cpp
auto result = divide(10, 2);
int value = result.value_or(-1);  // -1 if error
```

**Pros:**
- Clean syntax
- Type-safe
- No exceptions needed

**Cons:**
- Only works when "no value" = error
- Can't return error details (just success/failure)

---

## 3. Method 3: std::expected (C++23) - Best (Like Go)

**Pattern:** Return `std::expected<T, E>`, exactly like Go's `(value, error)`

```cpp
#include <expected>

enum class Error {
    DIVISION_BY_ZERO,
    INVALID_INPUT
};

std::expected<int, Error> divide(int a, int b) {
    if (b == 0) {
        return std::unexpected(Error::DIVISION_BY_ZERO);
    }
    return a / b;  // Success
}

// Usage
auto result = divide(10, 2);
if (!result.has_value()) {  // Error occurred
    if (result.error() == Error::DIVISION_BY_ZERO) {
        cout << "Error: Division by zero" << endl;
    }
    return;
}
cout << "Result: " << *result << endl;  // or: result.value()
```

**With Structured Binding (C++17):**
```cpp
auto result = divide(10, 2);
if (result.has_value()) {
    cout << "Result: " << result.value() << endl;
} else {
    cout << "Error: " << static_cast<int>(result.error()) << endl;
}
```

**Pros:**
- Exactly like Go's `(value, error)` pattern
- Type-safe
- Can return error details
- No exceptions

**Cons:**
- Requires C++23 (not available everywhere yet)

---

## 4. Method 4: Pair/Struct Return (Pre-C++23)

**Pattern:** Return `pair<T, bool>` or custom struct

```cpp
#include <utility>

// Option 1: pair<value, success>
pair<int, bool> divide(int a, int b) {
    if (b == 0) {
        return {0, false};  // Error
    }
    return {a / b, true};  // Success
}

// Usage
auto [result, success] = divide(10, 2);
if (!success) {
    cout << "Error: Division by zero" << endl;
    return;
}
cout << "Result: " << result << endl;
```

**Option 2: Custom Result Struct (More Flexible)**
```cpp
template<typename T>
struct Result {
    T value;
    bool success;
    string error_msg;
    
    static Result<T> ok(T val) {
        return {val, true, ""};
    }
    
    static Result<T> error(string msg) {
        return {T{}, false, msg};
    }
};

Result<int> divide(int a, int b) {
    if (b == 0) {
        return Result<int>::error("Division by zero");
    }
    return Result<int>::ok(a / b);
}

// Usage
auto result = divide(10, 2);
if (!result.success) {
    cout << "Error: " << result.error_msg << endl;
    return;
}
cout << "Result: " << result.value << endl;
```

**Pros:**
- Works in older C++ versions
- Can include error messages
- Simple to understand

**Cons:**
- More verbose than `std::expected`
- Not standard library

---

## 5. Common Patterns for DSA & Machine Coding

### Pattern 1: Array/Vector Access
```cpp
std::optional<int> getElement(const vector<int>& arr, int index) {
    if (index < 0 || index >= arr.size()) {
        return std::nullopt;
    }
    return arr[index];
}

// Usage
if (auto val = getElement(arr, 5)) {
    cout << *val << endl;
}
```

### Pattern 2: Map/Set Lookup
```cpp
std::optional<string> getValue(const map<int, string>& m, int key) {
    auto it = m.find(key);
    if (it == m.end()) {
        return std::nullopt;
    }
    return it->second;
}

// Usage
if (auto val = getValue(myMap, 42)) {
    cout << *val << endl;
}
```

### Pattern 3: Validation
```cpp
enum class ValidationError {
    EMPTY,
    TOO_LONG,
    INVALID_CHAR
};

std::expected<string, ValidationError> validate(const string& input) {
    if (input.empty()) {
        return std::unexpected(ValidationError::EMPTY);
    }
    if (input.length() > 100) {
        return std::unexpected(ValidationError::TOO_LONG);
    }
    return input;  // Valid
}
```

### Pattern 4: Chaining Operations (Like Go)
```cpp
// Go style:
// result, err := step1()
// if err != nil { return err }
// result2, err := step2(result)
// if err != nil { return err }

// C++ with optional:
auto step1() -> optional<int> { /* ... */ }
auto step2(int x) -> optional<string> { /* ... */ }

auto result1 = step1();
if (!result1) return;

auto result2 = step2(*result1);
if (!result2) return;

// Use *result2
```

---

## 6. Quick Decision Guide

| Scenario | Recommended Method |
|----------|-------------------|
| **Simple success/failure** | `std::optional<T>` |
| **Need error details** | `std::expected<T, E>` (C++23) or `Result<T>` struct |
| **C++17 or earlier** | `std::optional<T>` or `pair<T, bool>` |
| **Performance critical** | Return error code (int/enum) |
| **DSA problems** | `std::optional<T>` (simple, clean) |
| **Machine coding** | `std::expected<T, E>` or custom `Result<T>` |

---

## 7. Common Pitfalls

### ❌ Forgetting to Check Error
```cpp
auto result = divide(10, 0);
cout << *result << endl;  // CRASH! Always check first
```

### ✅ Always Check
```cpp
auto result = divide(10, 0);
if (!result) {
    // Handle error
    return;
}
cout << *result << endl;
```

### ❌ Using Exceptions for Control Flow
```cpp
// Don't do this in DSA/machine coding
try {
    int result = divide(a, b);
} catch (...) {
    // Exceptions are slow, avoid in competitive programming
}
```

### ✅ Use Return Values Instead
```cpp
// Better: explicit error handling
auto result = divide(a, b);
if (!result) {
    // Handle error
}
```

---

## 8. TL;DR - The Essentials

**For DSA & Machine Coding:**

1. **Use `std::optional<T>`** for simple cases (C++17+)
   ```cpp
   optional<int> result = divide(a, b);
   if (!result) return;  // Error
   use(*result);  // Success
   ```

2. **Use `std::expected<T, E>`** if you need error details (C++23)
   ```cpp
   expected<int, Error> result = divide(a, b);
   if (!result) {
       handle(result.error());
       return;
   }
   use(result.value());
   ```

3. **Use error codes** for maximum performance
   ```cpp
   int result;
   if (int err = divide(a, b, result); err != 0) {
       return;  // Error
   }
   use(result);
   ```

**Key Rules:**
- ✅ Always check for errors before using result
- ✅ Prefer return values over exceptions (for DSA/machine coding)
- ✅ Use `std::optional` for simple success/failure
- ✅ Use `std::expected` when you need error details (C++23)

---

## 9. Go vs C++ Comparison

| Feature | Go | C++ |
|---------|----|----|
| **Return Pattern** | `(value, error)` | `std::expected<T, E>` (C++23) |
| **Simple Cases** | `(value, error)` | `std::optional<T>` |
| **Error Type** | `error` interface | `enum class` or custom type |
| **Check Error** | `if err != nil` | `if (!result)` or `if (!result.has_value())` |
| **Get Value** | `value` | `*result` or `result.value()` |
| **Chaining** | Multiple `if err != nil` | Multiple `if (!result) return` |

**Go Example:**
```go
result, err := divide(a, b)
if err != nil {
    return err
}
result2, err := multiply(result, 2)
if err != nil {
    return err
}
```

**C++ Equivalent:**
```cpp
auto result = divide(a, b);
if (!result) return;
auto result2 = multiply(*result, 2);
if (!result2) return;
```

---

**Remember:** Keep it simple, explicit, and fast. No exceptions in competitive programming!

