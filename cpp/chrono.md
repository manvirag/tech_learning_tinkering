# C++ Chrono Library - Practical Reference

## Quick Setup: Using Shortcuts (Recommended)

```cpp
#include <chrono>
#include <iostream>

// Shortcuts for cleaner code
using namespace std::chrono;
using Clock = steady_clock;
using TimePoint = Clock::time_point;
using Seconds = duration<long long>;
using Milliseconds = duration<long long, std::milli>;
using Microseconds = duration<long long, std::micro>;

// Or use predefined types
using namespace std::chrono;
// Now you can use: seconds, milliseconds, hours, minutes directly
```

---

## 1. Getting Current Time

```cpp
// Most common: steady_clock (monotonic, good for measurements)
auto now = std::chrono::steady_clock::now();
auto start = std::chrono::steady_clock::now();

// For wall-clock time (can be adjusted by system)
auto systemNow = std::chrono::system_clock::now();

// With shortcuts:
using namespace std::chrono;
auto now = steady_clock::now();
auto systemNow = system_clock::now();
```

---

## 2. Creating Specific Durations

```cpp
// Common duration types
std::chrono::seconds sec(5);           // 5 seconds
std::chrono::milliseconds ms(500);      // 500 milliseconds
std::chrono::microseconds us(1000);     // 1000 microseconds
std::chrono::minutes min(30);          // 30 minutes
std::chrono::hours hr(24);             // 24 hours

// With shortcuts:
using namespace std::chrono;
seconds sec(5);
milliseconds ms(500);
hours hr(1);
```

---

## 3. Time Point Operations (Add/Subtract)

```cpp
// Add duration to time point
auto future = std::chrono::steady_clock::now() + std::chrono::seconds(5);
auto deadline = std::chrono::steady_clock::now() + std::chrono::minutes(10);

// Subtract duration from time point
auto past = std::chrono::steady_clock::now() - std::chrono::hours(1);

// With shortcuts:
using namespace std::chrono;
auto future = steady_clock::now() + seconds(5);
auto past = steady_clock::now() - hours(1);
```

---

## 4. Time Difference (Duration Between Two Times)

```cpp
auto start = std::chrono::steady_clock::now();
// ... do work ...
auto end = std::chrono::steady_clock::now();

// Get difference (returns duration)
auto elapsed = end - start;

// Convert to specific duration type
auto elapsed_ms = std::chrono::duration_cast<std::chrono::milliseconds>(end - start);
auto elapsed_sec = std::chrono::duration_cast<std::chrono::seconds>(end - start);

// Get count value
long long ms_count = elapsed_ms.count();  // Get numeric value

// With shortcuts:
using namespace std::chrono;
auto elapsed = end - start;
auto elapsed_ms = duration_cast<milliseconds>(end - start);
long long ms = elapsed_ms.count();
```

---

## 5. Duration Types and Conversions

```cpp
// Predefined duration types
std::chrono::nanoseconds   ns(1000000);
std::chrono::microseconds  us(1000);
std::chrono::milliseconds  ms(500);
std::chrono::seconds       sec(5);
std::chrono::minutes       min(30);
std::chrono::hours         hr(24);

// Convert between duration types
std::chrono::seconds sec(5);
auto ms = std::chrono::duration_cast<std::chrono::milliseconds>(sec);  // 5000 ms
auto us = std::chrono::duration_cast<std::chrono::microseconds>(sec); // 5000000 us

// Get count (numeric value)
long long seconds_value = sec.count();        // 5
long long milliseconds_value = ms.count();    // 5000

// With shortcuts:
using namespace std::chrono;
seconds sec(5);
auto ms = duration_cast<milliseconds>(sec);
long long value = ms.count();
```

---

## 6. Timestamp Operations (System Clock)

```cpp
// Get current timestamp
auto now = std::chrono::system_clock::now();

// Convert to time_t (for logging/printing)
std::time_t time_t = std::chrono::system_clock::to_time_t(now);
std::cout << "Time: " << std::ctime(&time_t);

// Create timestamp from time_t
std::time_t t = 1234567890;
auto tp = std::chrono::system_clock::from_time_t(t);

// Convert to string
std::string timeStr = std::ctime(&time_t);

// With shortcuts:
using namespace std::chrono;
auto now = system_clock::now();
std::time_t t = system_clock::to_time_t(now);
std::cout << std::ctime(&t);
```

---

## 7. Comparing Time Points

```cpp
auto start = std::chrono::steady_clock::now();
auto deadline = start + std::chrono::seconds(10);
auto now = std::chrono::steady_clock::now();

// Comparisons
if (now < deadline) {
    // deadline hasn't passed
}

if (now >= deadline) {
    // deadline has passed
}

if (start < now) {
    // start is before now
}

// With shortcuts:
using namespace std::chrono;
auto deadline = steady_clock::now() + seconds(10);
auto now = steady_clock::now();
if (now < deadline) { /* ... */ }
```

---

## 8. Practical Examples

### Measure Execution Time
```cpp
using namespace std::chrono;
auto start = steady_clock::now();
// ... your code ...
auto end = steady_clock::now();
auto elapsed = duration_cast<milliseconds>(end - start);
std::cout << "Time taken: " << elapsed.count() << " ms" << std::endl;
```

### Check if Timeout Occurred
```cpp
using namespace std::chrono;
auto start = steady_clock::now();
auto timeout = start + seconds(5);

// Later...
auto now = steady_clock::now();
if (now >= timeout) {
    std::cout << "Timeout!" << std::endl;
}
```

### Calculate Time Until Deadline
```cpp
using namespace std::chrono;
auto deadline = steady_clock::now() + minutes(30);
auto now = steady_clock::now();
auto remaining = duration_cast<seconds>(deadline - now);
std::cout << "Time remaining: " << remaining.count() << " seconds" << std::endl;
```

### Log with Timestamp
```cpp
using namespace std::chrono;
auto now = system_clock::now();
std::time_t t = system_clock::to_time_t(now);
std::cout << "[" << std::ctime(&t) << "] Log message" << std::endl;
```

---

## Summary: Most Common Patterns

```cpp
#include <chrono>
using namespace std::chrono;

// 1. Get current time
auto now = steady_clock::now();

// 2. Create duration
seconds sec(5);
milliseconds ms(500);

// 3. Add/subtract time
auto future = now + seconds(10);
auto past = now - minutes(5);

// 4. Calculate difference
auto elapsed = end - start;
auto elapsed_ms = duration_cast<milliseconds>(elapsed);
long long ms = elapsed_ms.count();

// 5. Compare times
if (now < deadline) { /* ... */ }

// 6. Timestamp for logging
auto sysNow = system_clock::now();
std::time_t t = system_clock::to_time_t(sysNow);
std::cout << std::ctime(&t);
```
