# C++ Constructor Initialization List - Quick Reference

## Basic Usage
- **Definition**: Syntax `: member(value)` used in constructors to initialize member variables
- **Also called**: Member initializer list
- **Purpose**: Initialize members before constructor body executes

## Vector Initialization (Normal Usage)

### Before (C++98/03) - Old Way
```cpp
// ❌ OLD WAY: Multiple push_back calls
std::vector<int> vec;
vec.push_back(1);
vec.push_back(2);
vec.push_back(3);
vec.push_back(4);
vec.push_back(5);

// OR using array
int arr[] = {1, 2, 3, 4, 5};
std::vector<int> vec2(arr, arr + 5);  // Need array + size

// OR using constructor with size
std::vector<int> vec3(5, 0);  // 5 elements, all initialized to 0
```

### After (C++11) - New Way with Brace Initialization
```cpp
// ✅ NEW WAY: Single line initialization using initializer_list
std::vector<int> vec{1, 2, 3, 4, 5};
std::vector<int> vec2 = {10, 20, 30};  // Also works

// Other STL containers
std::set<int> s{5, 2, 8, 1};
std::map<int, string> m{{1, "one"}, {2, "two"}};
std::list<int> l{1, 2, 3};
```

### How It Works
STL containers (like `std::vector`) have constructors that accept `std::initializer_list`, allowing brace initialization syntax. The compiler automatically converts `{1, 2, 3}` to `std::initializer_list<int>`.

```cpp
// Behind the scenes, this:
std::vector<int> vec{1, 2, 3};

// Is equivalent to:
std::initializer_list<int> list = {1, 2, 3};
std::vector<int> vec(list);
```

### Different Initialization Methods
```cpp
// Method 1: Brace initialization (uses initializer_list)
std::vector<int> v1{1, 2, 3};  // Creates vector with elements 1, 2, 3

// Method 2: Constructor with size and value
std::vector<int> v2(5, 10);  // Creates vector with 5 elements, each = 10

// Method 3: Constructor with size only
std::vector<int> v3(5);  // Creates vector with 5 elements, default initialized (0 for int)

// Method 4: Empty then push_back
std::vector<int> v4;
v4.push_back(1);
v4.push_back(2);
```

### Important: Brace vs Parentheses
```cpp
// Brace initialization - uses initializer_list
std::vector<int> v1{5, 10};  // Vector with 2 elements: [5, 10]

// Parentheses - constructor call
std::vector<int> v2(5, 10);  // Vector with 5 elements, each = 10: [10, 10, 10, 10, 10]

// Single value
std::vector<int> v3{5};  // Vector with 1 element: [5]
std::vector<int> v4(5);  // Vector with 5 elements: [0, 0, 0, 0, 0]
```

## 1. Using in Constructor

### Basic Syntax
```cpp
class MyClass {
private:
    int x;
    int y;
    std::string name;
    
public:
    // Constructor with initialization list
    MyClass(int a, int b, std::string n) : x(a), y(b), name(n) {
        // Constructor body (members already initialized)
    }
};
```

### Before (Without Initialization List) - Old Way
```cpp
class MyClass {
private:
    int x;
    int y;
    const int value;  // const member
    
public:
    // ❌ OLD WAY: Assignment in constructor body
    MyClass(int a, int b) {
        x = a;        // Assignment, not initialization
        y = b;        // Assignment, not initialization
        // value = 10;  // ❌ ERROR: Can't assign to const member
    }
};
```

### After (With Initialization List) - New Way
```cpp
class MyClass {
private:
    int x;
    int y;
    const int value;  // const member
    
public:
    // ✅ NEW WAY: Initialization list
    MyClass(int a, int b) : x(a), y(b), value(10) {
        // All members initialized before body executes
        // value is initialized, not assigned
    }
};
```

## 2. When Initialization List is Required

### Must Use Initialization List For:
1. **const members**: Can't be assigned, must be initialized
2. **Reference members**: Must be initialized, can't be assigned
3. **Member objects without default constructor**: Must be initialized
4. **Base class constructors**: Must call base constructor

```cpp
class Example {
private:
    const int id;           // const - must use init list
    int& ref;              // reference - must use init list
    std::vector<int> vec;  // No default constructor - must use init list
    
public:
    Example(int i, int& r) : id(i), ref(r), vec(10) {
        // All must be in initialization list
    }
};
```

### Performance Benefit
```cpp
class WithInitList {
private:
    std::string name;
    
public:
    // ✅ Efficient: Direct initialization
    WithInitList(std::string n) : name(n) { }
    // name is constructed directly with n
};

class WithoutInitList {
private:
    std::string name;
    
public:
    // ❌ Less efficient: Default construct + assign
    WithoutInitList(std::string n) {
        name = n;  // name default constructed, then assigned
    }
};
```

## Initialization vs Assignment

### Initialization (Init List)
- **When**: Constructor initialization list `: member(value)`
- **What happens**: Object is constructed directly with the value
- **Efficiency**: One operation - direct construction
- **Timing**: Happens before constructor body executes

```cpp
class Example {
private:
    std::string name;
    int x;
    
public:
    Example(std::string n, int val) : name(n), x(val) {
        // name is INITIALIZED with n (constructor called directly)
        // x is INITIALIZED with val (direct assignment to uninitialized memory)
        // Both happen BEFORE this body executes
    }
};
```

### Assignment (Constructor Body)
- **When**: Inside constructor body `member = value`
- **What happens**: Object is first default-constructed, then assigned a new value
- **Efficiency**: Two operations - default construct + assign
- **Timing**: Happens during constructor body execution

```cpp
class Example {
private:
    std::string name;
    int x;
    
public:
    Example(std::string n, int val) {
        // name is first default-constructed (empty string)
        // Then ASSIGNED the value n (operator= called)
        name = n;  // Two steps: construct + assign
        
        // x is first default-initialized (garbage/zero)
        // Then ASSIGNED the value val
        x = val;   // Two steps: initialize + assign
    }
};
```

### Key Difference
```cpp
// Initialization (Init List)
MyClass(int a) : x(a) { }
// Step 1: x is created and set to 'a' directly
// Total: 1 operation

// Assignment (Body)
MyClass(int a) {
    x = a;  // Step 1: x is default-initialized (0 or garbage)
            // Step 2: x is assigned value 'a'
}
// Total: 2 operations
```

### Why It Matters
- **Performance**: Initialization is faster (one operation vs two)
- **const/References**: Can only be initialized, not assigned
- **Objects without default constructor**: Must be initialized, can't default-construct then assign

### Objects Without Default Constructor - Examples

#### Example 1: Custom Class Without Default Constructor
```cpp
// Class without default constructor
class NoDefault {
private:
    int value;
    
public:
    NoDefault(int v) : value(v) { }  // Only parameterized constructor
    // No default constructor NoDefault() = delete; (implicitly)
};

class Container {
private:
    NoDefault obj;  // Member without default constructor
    
public:
    // ❌ WRONG: Can't default-construct then assign
    Container(int val) {
        // obj = NoDefault(val);  // ERROR: obj must be initialized first
        // Compiler error: "no matching function for call to 'NoDefault::NoDefault()'"
    }
    
    // ✅ CORRECT: Must use initialization list
    Container(int val) : obj(val) {
        // obj is initialized directly with val
    }
};
```

#### Example 2: std::vector with Size (No Default Constructor for That Size)
```cpp
class Matrix {
private:
    std::vector<std::vector<int>> data;
    
public:
    // ❌ WRONG: Can't default-construct vector then assign
    Matrix(int rows, int cols) {
        // data = std::vector<std::vector<int>>(rows, std::vector<int>(cols));
        // This works but inefficient: default construct + assign
        
        // Better approach with init list:
    }
    
    // ✅ CORRECT: Initialize directly
    Matrix(int rows, int cols) : data(rows, std::vector<int>(cols)) {
        // data is initialized directly with specified size
    }
};
```

#### Example 3: Reference Member (No Default Constructor)
```cpp
class ReferenceHolder {
private:
    int& ref;  // Reference has no default constructor
    
public:
    // ❌ WRONG: Can't default-construct reference
    ReferenceHolder(int value) {
        // ref = value;  // ERROR: Reference must be initialized
        // Compiler error: "uninitialized reference member"
    }
    
    // ✅ CORRECT: Must initialize in init list
    ReferenceHolder(int& r) : ref(r) {
        // ref is initialized to refer to r
    }
};
```

#### Example 4: Base Class Without Default Constructor
```cpp
class Base {
public:
    Base(int x) { }  // Only parameterized constructor
    // No default constructor
};

class Derived : public Base {
public:
    // ❌ WRONG: Can't default-construct base class
    Derived(int val) {
        // Base(val);  // ERROR: Base must be initialized before body
    }
    
    // ✅ CORRECT: Call base constructor in init list
    Derived(int val) : Base(val) {
        // Base is initialized first, then Derived body executes
    }
};
```

## Using Initialization List with Vector

### Vector as Member Variable

#### Without Initialization List - Old Way
```cpp
class Container {
private:
    std::vector<int> data;
    
public:
    // ❌ OLD WAY: Default construct + assign in body
    Container() {
        data = std::vector<int>{1, 2, 3, 4, 5};  // Default construct + assign
    }
    
    Container(int size) {
        data = std::vector<int>(size, 0);  // Default construct + assign
    }
    
    Container(int size, int value) {
        data = std::vector<int>(size, value);  // Default construct + assign
    }
};
```

#### With Initialization List - New Way
```cpp
class Container {
private:
    std::vector<int> data;
    
public:
    // ✅ NEW WAY: Direct initialization
    Container() : data{1, 2, 3, 4, 5} {
        // data is initialized directly with values
    }
    
    Container(int size) : data(size, 0) {
        // data is initialized directly with size and default value
    }
    
    Container(int size, int value) : data(size, value) {
        // data is initialized directly with size and value
    }
};
```

### Performance Comparison

```cpp
class WithInitList {
private:
    std::vector<int> vec;
    
public:
    // ✅ Efficient: Direct initialization
    WithInitList() : vec(100, 0) {
        // vec is constructed directly with 100 zeros
        // One operation: vector constructor called
    }
};

class WithoutInitList {
private:
    std::vector<int> vec;
    
public:
    // ❌ Less efficient: Default construct + assign
    WithoutInitList() {
        vec = std::vector<int>(100, 0);
        // Step 1: vec is default-constructed (empty vector)
        // Step 2: Temporary vector(100, 0) is created
        // Step 3: vec is assigned from temporary (copy/move)
        // Total: 3 operations
    }
};
```

### Real-World Example: Matrix Class

```cpp
class Matrix {
private:
    std::vector<std::vector<int>> data;
    int rows;
    int cols;
    
public:
    // ❌ WITHOUT Init List: Inefficient
    Matrix(int r, int c) {
        rows = r;
        cols = c;
        // data is default-constructed (empty)
        // Then resized in body
        data.resize(rows);
        for (int i = 0; i < rows; i++) {
            data[i].resize(cols, 0);
        }
    }
    
    // ✅ WITH Init List: Efficient
    Matrix(int r, int c) : rows(r), cols(c), data(r, std::vector<int>(c, 0)) {
        // rows and cols initialized directly
        // data initialized directly with r rows, each row has c zeros
        // All happens in one step
    }
};
```

### Vector with Initial Values

```cpp
class MyClass {
private:
    std::vector<int> numbers;
    std::vector<std::string> names;
    
public:
    // ❌ WITHOUT Init List
    MyClass() {
        numbers = {1, 2, 3, 4, 5};  // Default construct + assign
        names = {"Alice", "Bob"};   // Default construct + assign
    }
    
    // ✅ WITH Init List
    MyClass() : numbers{1, 2, 3, 4, 5}, names{"Alice", "Bob"} {
        // Both vectors initialized directly with values
        // More efficient and cleaner
    }
};
```

### Key Benefits for Vectors
- **Performance**: Direct initialization avoids default construction + assignment
- **Cleaner code**: Initialization list is more readable
- **Efficiency**: Especially important for large vectors or nested vectors
- **Consistency**: Same pattern as other member variables

## Key Points
- **Order matters**: Members initialized in declaration order, not list order
- **Initialization vs Assignment**: Init list initializes, body assigns
- **Performance**: Initialization is more efficient than default construct + assign
- **Required for**: const, references, members without default constructor
