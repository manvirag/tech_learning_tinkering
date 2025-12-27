# C++ Math and Random - Quick Reference

## Headers
- **Math**: `#include <cmath>`
- **Random**: `#include <random>` (C++11)
- **Algorithm**: `#include <algorithm>` (for min/max)

## Useful Math Functions

### Basic Operations
```cpp
#include <cmath>

double x = 3.14;
std::abs(x);        // Absolute value: |x|
std::sqrt(x);      // Square root: √x
std::pow(x, 2);    // Power: x^2
std::pow(x, 0.5);  // Square root (alternative)
```

### Rounding Functions
```cpp
std::floor(3.7);   // Round down: 3.0
std::ceil(3.2);    // Round up: 4.0
std::round(3.5);   // Round to nearest: 4.0
std::round(3.4);   // Round to nearest: 3.0
std::trunc(3.7);   // Truncate (remove decimal): 3.0
```

### Min/Max Functions
```cpp
#include <algorithm>

std::min(a, b);           // Minimum of two values
std::max(a, b);           // Maximum of two values
std::min({a, b, c});      // Minimum of multiple values (C++11)
std::max({a, b, c});      // Maximum of multiple values (C++11)
```

### Logarithmic Functions
```cpp
std::log(x);       // Natural logarithm (ln)
std::log10(x);     // Base-10 logarithm
std::exp(x);       // e^x
```

### Trigonometric Functions (Radians)
```cpp
std::sin(x);       // Sine
std::cos(x);       // Cosine
std::tan(x);       // Tangent
std::asin(x);      // Arc sine
std::acos(x);      // Arc cosine
std::atan(x);      // Arc tangent
```

## Useful Random Functions

### Setup (Do Once)
```cpp
#include <random>

std::random_device rd;      // Seed generator - provides random seed from OS
std::mt19937 gen(rd());     // Random number generator - Mersenne Twister engine
```

**Explanation**: `std::random_device` gets a truly random seed from the operating system. `std::mt19937` is a fast, high-quality pseudo-random number generator (Mersenne Twister algorithm). We seed it with `rd()` to ensure different random sequences each run.

### Use Case 1: Random Number
```cpp
// Random integer (any number)
int random = gen();                    // Random integer

// Random double between 0.0 and 1.0
std::uniform_real_distribution<double> dis(0.0, 1.0);
double randomDouble = dis(gen);        // Random 0.0 to 1.0
```

### Use Case 2: Random Number in Range
```cpp
// Random integer in range [min, max]
std::uniform_int_distribution<int> dis(1, 100);
int random = dis(gen);                 // Random from 1 to 100

// Random double in range [min, max)
std::uniform_real_distribution<double> realDis(0.0, 10.0);
double randomReal = realDis(gen);      // Random from 0.0 to 10.0

// Quick helper function (gen must be accessible - global or passed)
int randomInRange(int min, int max) {
    std::uniform_int_distribution<int> dis(min, max);
    return dis(gen);  // Uses gen from setup above
}
int num = randomInRange(1, 100);       // Random from 1 to 100

// Note: gen must be in same scope or global. Alternative - pass gen as parameter:
int randomInRange(std::mt19937& gen, int min, int max) {
    std::uniform_int_distribution<int> dis(min, max);
    return dis(gen);
}
int num2 = randomInRange(gen, 1, 100);  // Pass gen as parameter
```

### Use Case 3: Random Number with Weight
```cpp
// Weighted random selection
// Example: Choose index 0 with 50%, index 1 with 30%, index 2 with 20%
std::vector<double> weights = {0.5, 0.3, 0.2};
std::discrete_distribution<int> weightedDis(weights.begin(), weights.end());
int chosen = weightedDis(gen);         // Returns 0, 1, or 2 based on weights

// Normal distribution (weighted around mean)
std::normal_distribution<double> normalDis(50.0, 10.0);  // mean=50, stddev=10
double value = normalDis(gen);        // Values clustered around 50
```

## Competitive Programming Useful Functions

### GCD and LCM
```cpp
// GCD (Greatest Common Divisor)
int gcd(int a, int b) {
    return b == 0 ? a : gcd(b, a % b);
}
// Or use: std::gcd(a, b) in C++17

// LCM (Least Common Multiple)
int lcm(int a, int b) {
    return (a / gcd(a, b)) * b;
}
// Or use: std::lcm(a, b) in C++17
```

### Modular Exponentiation (Power with Mod)
```cpp
// Fast power: (base^exp) % mod
long long power(long long base, long long exp, long long mod) {
    long long result = 1;
    base %= mod;
    while (exp > 0) {
        if (exp % 2 == 1) result = (result * base) % mod;
        base = (base * base) % mod;
        exp /= 2;
    }
    return result;
}

// Usage: power(2, 10, 1000) = (2^10) % 1000
```

### Floating Point Comparison
```cpp
// ❌ Wrong: Direct comparison
if (a == b) { }  // May fail due to precision

// ✅ Correct: Use epsilon
const double EPS = 1e-9;
if (std::abs(a - b) < EPS) { }  // Consider equal if very close
```

## Common Use Cases



### Check if Number is Integer
```cpp
bool isInteger(double x) {
    return std::abs(x - std::round(x)) < 1e-9;
}
```

## Common Pitfalls

```cpp
// Wrong: Floating point equality
double a = 0.1 + 0.2;
if (a == 0.3) { }  // ❌ May be false

// Correct: Use epsilon
const double EPS = 1e-9;
if (std::abs(a - 0.3) < EPS) { }  // ✅

// Wrong: Integer overflow in power
int result = std::pow(10, 9);  // ❌ May overflow

// Correct: Use long long or custom power
long long result = std::pow(10LL, 9LL);  // ✅
// OR use modular exponentiation for large numbers

```

## C++17/20 Enhancements

- **std::gcd(a, b)** (C++17): Greatest common divisor
- **std::lcm(a, b)** (C++17): Least common multiple
- **std::numbers::pi** (C++20): Mathematical constants
