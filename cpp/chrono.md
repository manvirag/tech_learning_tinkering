# C++ Chrono Library - Quick Reference

## Setup

```cpp
#include <chrono>
using namespace std;

using TimePoint = chrono::steady_clock::time_point;
using Duration = chrono::seconds;
using Clock = chrono::steady_clock;
```

---

## 1. Get Current Time

```cpp
TimePoint now = Clock::now();
auto systemNow = chrono::system_clock::now();
```

---

## 2. Create Durations

```cpp
chrono::seconds sec(5);
chrono::milliseconds ms(500);
chrono::minutes min(30);
chrono::hours hr(24);
Duration duration(10);  // if Duration = chrono::seconds
```

---

## 3. Time Operations (Add/Subtract)

```cpp
TimePoint future = Clock::now() + chrono::seconds(5);
TimePoint past = Clock::now() - chrono::hours(1);
```

---

## 4. Time Difference

```cpp
TimePoint start = Clock::now();
TimePoint end = Clock::now();
auto elapsed = end - start;
auto elapsed_ms = chrono::duration_cast<chrono::milliseconds>(end - start);
long long ms = elapsed_ms.count();
```

---

## 5. Duration Conversion

```cpp
chrono::seconds sec(5);
auto ms = chrono::duration_cast<chrono::milliseconds>(sec);
long long value = ms.count();
```

---

## 6. Timestamp (for Logging)

```cpp
#include <ctime>
auto now = chrono::system_clock::now();
time_t t = chrono::system_clock::to_time_t(now);
cout << ctime(&t);

// Epoch timestamp as long long
long long epoch = chrono::duration_cast<chrono::seconds>(now.time_since_epoch()).count();
```

---

## 7. Store Specific Date/Time (Scenario 1)

```cpp
#include <ctime>
// Store specific date/time as timestamp
// Method 1: From time_t (Unix timestamp)
time_t specific_time = 1234567890;  // specific epoch timestamp
auto tp = chrono::system_clock::from_time_t(specific_time);
long long timestamp = chrono::duration_cast<chrono::seconds>(tp.time_since_epoch()).count();

// Method 2: Create from current time and adjust
auto activity_time = chrono::system_clock::now();
long long activity_timestamp = chrono::duration_cast<chrono::seconds>(activity_time.time_since_epoch()).count();
// Store activity_timestamp in variable
```

---

## 8. Store Time of Day (Scenario 2)

```cpp
// Store time like 10:11 AM
// Method 1: As duration since midnight
auto time_of_day = chrono::hours(10) + chrono::minutes(11);
long long total_minutes = chrono::duration_cast<chrono::minutes>(time_of_day).count();  // 611 minutes

// Method 2: Store as simple variables
int hour = 10;
int minute = 11;
```

---

## 9. Store Date

```cpp
#include <ctime>
// Method 1: Using tm structure
tm date = {};
date.tm_year = 2024 - 1900;  // year - 1900
date.tm_mon = 11;             // month (0-11, so 11 = December)
date.tm_mday = 25;            // day of month

// Method 2: Store as simple variables
int year = 2024;
int month = 12;   // 1-12
int day = 25;

// Method 3: Convert from timestamp to date
auto now = chrono::system_clock::now();
time_t t = chrono::system_clock::to_time_t(now);
tm* date_tm = localtime(&t);
int year = date_tm->tm_year + 1900;
int month = date_tm->tm_mon + 1;
int day = date_tm->tm_mday;
```

---

## 10. Compare Times

```cpp
TimePoint deadline = Clock::now() + chrono::seconds(10);
TimePoint now = Clock::now();
if (now < deadline) { /* ... */ }
if (now >= deadline) { /* ... */ }
```
