# 📘 C++ Fundamentals – Final Notes

These notes cover **core C++ concepts** you should understand **before writing real code**.  
They are beginner-friendly, practical, and interview-ready.

---

## 1️⃣ C++ Versioning

- C++ evolves via **standards**, not runtime versions
- Versions are named by year:
  - C++11 – Modern C++ begins
  - C++14 – Refinements
  - C++17 – Widely used
  - C++20 – Concepts, ranges, modules
- Version is selected **at compile time**

```bash
g++ -std=c++17 main.cpp
```

C++ has no runtime environment like Java or Python

## 2️⃣ How C++ Code Is Compiled

C++ is a compiled language.

### Compilation Pipeline

```
.cpp
 ↓
Preprocessor   (#include, #define, #pragma once)
 ↓
Compiler       (syntax + type checking)
 ↓
Object file    (.o)
 ↓
Linker         (combine + libraries)
 ↓
Executable
```

### Important Concepts

**Translation Unit**
- One .cpp file plus all included headers

**Compile-time error**
- Syntax or type errors

**Link-time error**
- Missing function definitions

## 3️⃣ Header Files & #include

### Include Types

```cpp
#include <iostream>   // System headers
#include "MyFile.h"   // Project headers
```

### Golden Rule

Each header must include everything it needs to compile on its own.

### Transitive Includes

If:

```
main.cpp → Controller.h → Service.h → DAO.h → Model.h
```

Then:

`main.cpp` only includes `Controller.h`

## 4️⃣ #pragma once

Prevents multiple inclusion of the same header

- Header is processed only once per compilation

```cpp
#pragma once
```

### Why It Exists

- Avoids redefinition errors
- Replaces include guards
- Supported by all modern compilers

## 5️⃣ C++ Compiler Basics

### Popular Compilers

- **g++** (GCC)
- **clang++**
- **MSVC** (Visual Studio)

### Common Compiler Flags

```bash
-std=c++17   # C++ version
-Wall        # Enable warnings
-Wextra      # Extra warnings
-O2          # Optimization
-g           # Debug symbols
```

## 6️⃣ gcc vs g++

| Feature | gcc | g++ |
|---------|-----|-----|
| Default language | C | C++ |
| Links C++ standard library | ❌ | ✅ |
| Recommended for C++ | ❌ | ✅ |

### Always Use

```bash
g++ main.cpp
```

## 7️⃣ Running C++ Code

C++ does not have a built-in `go run` equivalent.

### Closest One-Liner

```bash
g++ main.cpp && ./a.out
```

### Named Binary

```bash
g++ main.cpp -o app && ./app
```

## 8️⃣ Compiling with Threads

### Linux / macOS / WSL

```bash
g++ -std=c++17 main.cpp -pthread -o app
```

Required for:
- `std::thread`
- `std::mutex`
- `std::condition_variable`

### Windows (MSVC)

```bat
cl /std:c++17 main.cpp
```
