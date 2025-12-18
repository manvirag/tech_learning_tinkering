# C++ Chrono Library Reference

## 1. Most Used Duration Types

- `std::chrono::seconds`
- `std::chrono::milliseconds`
- `std::chrono::microseconds`
- `std::chrono::nanoseconds`
- `std::chrono::minutes`
- `std::chrono::hours`

### Examples:

```cpp
std::chrono::milliseconds timeout(500);
std::chrono::seconds retryDelay(2);
std::chrono::microseconds preciseDelay(1000);
std::chrono::hours dayDuration(24);
std::chrono::minutes meetingDuration(30);
std::chrono::nanoseconds veryPreciseDelay(1000000);
```

---

## 2. Most Used Time Point Types (VERY IMPORTANT)

- `std::chrono::steady_clock::time_point`
- `std::chrono::system_clock::time_point`
- `std::chrono::high_resolution_clock::time_point`

### Examples:

```cpp
std::chrono::steady_clock::time_point start =
    std::chrono::steady_clock::now();

std::chrono::system_clock::time_point now =
    std::chrono::system_clock::now();

std::chrono::high_resolution_clock::time_point preciseStart =
    std::chrono::high_resolution_clock::now();

std::chrono::steady_clock::time_point end =
    std::chrono::steady_clock::now();
```

---

## 3. Getting Current Time

- `std::chrono::steady_clock::now()`
- `std::chrono::system_clock::now()`
- `std::chrono::high_resolution_clock::now()`

### Usage:

```cpp
std::chrono::steady_clock::time_point t =
    std::chrono::steady_clock::now();

// More examples:
auto currentSystemTime = std::chrono::system_clock::now();
auto preciseTime = std::chrono::high_resolution_clock::now();
std::chrono::steady_clock::time_point benchmarkStart = 
    std::chrono::steady_clock::now();
```

---

## 4. Duration Conversion (Must-Know)

- `std::chrono::duration_cast<std::chrono::milliseconds>(duration)`
- `std::chrono::duration_cast<std::chrono::seconds>(duration)`

### Example:

```cpp
std::chrono::steady_clock::time_point start =
    std::chrono::steady_clock::now();

std::chrono::steady_clock::time_point end =
    std::chrono::steady_clock::now();

std::chrono::milliseconds elapsed =
    std::chrono::duration_cast<std::chrono::milliseconds>(end - start);

// More examples:
auto elapsedSeconds = 
    std::chrono::duration_cast<std::chrono::seconds>(end - start);
auto elapsedMicroseconds = 
    std::chrono::duration_cast<std::chrono::microseconds>(end - start);
auto elapsedMinutes = 
    std::chrono::duration_cast<std::chrono::minutes>(end - start);
```

---

## 5. Time Point Arithmetic (Very Common)

- `time_point + duration`
- `time_point - duration`
- `time_point - time_point -> duration`

### Example:

```cpp
std::chrono::steady_clock::time_point deadline =
    std::chrono::steady_clock::now() + std::chrono::seconds(5);

// More examples:
auto futureTime = std::chrono::system_clock::now() + 
    std::chrono::hours(1);
auto pastTime = std::chrono::steady_clock::now() - 
    std::chrono::minutes(10);
auto timeDifference = end - start;  // Returns a duration
auto timeout = std::chrono::steady_clock::now() + 
    std::chrono::milliseconds(500);
```

---

## 6. Sleeping / Waiting

- `std::this_thread::sleep_for(std::chrono::seconds(2))`
- `std::this_thread::sleep_for(std::chrono::milliseconds(100))`
- `std::this_thread::sleep_until(time_point)`

### Example:

```cpp
std::this_thread::sleep_until(deadline);

// More examples:
std::this_thread::sleep_for(std::chrono::seconds(2));
std::this_thread::sleep_for(std::chrono::milliseconds(100));
std::this_thread::sleep_for(std::chrono::microseconds(500));
auto wakeTime = std::chrono::steady_clock::now() + 
    std::chrono::seconds(5);
std::this_thread::sleep_until(wakeTime);
```

---

## 7. Comparing Time Points

```cpp
if (std::chrono::steady_clock::now() < deadline) {
    // still time left
}

// More examples:
auto now = std::chrono::steady_clock::now();
if (now < deadline) {
    // deadline hasn't passed
}

if (now >= deadline) {
    // deadline has passed
}

auto time1 = std::chrono::system_clock::now();
std::this_thread::sleep_for(std::chrono::milliseconds(100));
auto time2 = std::chrono::system_clock::now();
if (time2 > time1) {
    // time2 is later than time1
}
```

---

## 8. Converting to Calendar Time (Logs)

- `std::chrono::system_clock::to_time_t(time_point)`

### Example:

```cpp
std::chrono::system_clock::time_point now =
    std::chrono::system_clock::now();

std::time_t tt =
    std::chrono::system_clock::to_time_t(now);

// More examples:
auto currentTime = std::chrono::system_clock::now();
std::time_t timeT = std::chrono::system_clock::to_time_t(currentTime);
std::cout << "Current time: " << std::ctime(&timeT);

// Convert to string
std::string timeStr = std::ctime(&timeT);
```
