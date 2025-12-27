# C++ OOP - Quick Reference

## Headers
- **Base classes**: Standard class syntax
- **Virtual functions**: `virtual` keyword
- **Abstract classes**: Pure virtual functions `= 0`

## Abstract Classes and Interfaces

### Abstract Class (Pure Virtual Functions)
```cpp
class Shape {
public:
    // Pure virtual function - must be overridden
    virtual double area() = 0;            // Abstract method
    
    // Virtual function - can be overridden
    virtual void draw() {                 // Has default implementation
        std::cout << "Drawing shape" << std::endl;
    }
    
    // Regular function - cannot be overridden
    void print() {
        std::cout << "Shape" << std::endl;
    }
    
    // Virtual destructor (important!)
    virtual ~Shape() { }
};

// Cannot create object of abstract class
// Shape s;  // ❌ Error: cannot instantiate abstract class
```

### Interface (All Pure Virtual)
```cpp
class Drawable {
public:
    virtual void draw() = 0;              // Pure virtual
    virtual void resize(int factor) = 0;   // Pure virtual
    virtual ~Drawable() { }                // Virtual destructor
};

// All methods must be implemented in derived class
class Circle : public Drawable {
public:
    void draw() override {                // Must implement
        std::cout << "Drawing circle" << std::endl;
    }
    
    void resize(int factor) override {     // Must implement
        // Resize logic
    }
};
```

### Implementation
```cpp
class Circle : public Shape {
private:
    double radius;
    
public:
    Circle(double r) : radius(r) { }
    
    // Must implement pure virtual function
    double area() override {
        return 3.14159 * radius * radius;
    }
    
    // Can override virtual function
    void draw() override {
        std::cout << "Drawing circle" << std::endl;
    }
};

// Usage
Shape* shape = new Circle(5.0);
shape->area();                            // Calls Circle::area()
shape->draw();                            // Calls Circle::draw()
```

## Constructors and Destructors

### Constructor Types
```cpp
class MyClass {
private:
    int x;
    std::string name;
    
public:
    // Default constructor
    MyClass() : x(0), name("") { }
    
    // Parameterized constructor
    MyClass(int val, std::string n) : x(val), name(n) { }
    
    // Copy constructor
    MyClass(const MyClass& other) : x(other.x), name(other.name) { }
    
    // Move constructor (C++11)
    MyClass(MyClass&& other) noexcept : x(other.x), name(std::move(other.name)) {
        other.x = 0;
    }
    
    // Copy assignment
    MyClass& operator=(const MyClass& other) {
        if (this != &other) {
            x = other.x;
            name = other.name;
        }
        return *this;
    }
};
```

### Destructor
```cpp
class Resource {
private:
    int* data;
    
public:
    Resource(int size) {
        data = new int[size];
    }
    
    // Destructor - called automatically when object goes out of scope
    ~Resource() {
        delete[] data;                    // Clean up resources
        std::cout << "Resource destroyed" << std::endl;
    }
};

// Virtual destructor (important for polymorphism)
class Base {
public:
    virtual ~Base() { }                   // Virtual destructor
};

class Derived : public Base {
public:
    ~Derived() {                          // Called when deleting Base*
        // Cleanup
    }
};

Base* ptr = new Derived();
delete ptr;                               // ✅ Calls Derived destructor (if Base has virtual destructor)
```

### Constructor/Destructor Order
```cpp
class Base {
public:
    Base() { std::cout << "Base constructor" << std::endl; }
    ~Base() { std::cout << "Base destructor" << std::endl; }
};

class Derived : public Base {
public:
    Derived() { std::cout << "Derived constructor" << std::endl; }
    ~Derived() { std::cout << "Derived destructor" << std::endl; }
};

// Order:
// Base constructor → Derived constructor → Derived destructor → Base destructor
```

## Important Keywords

### virtual
**Purpose**: Enables runtime polymorphism (late binding)

```cpp
class Base {
public:
    virtual void func() {                 // Virtual function
        std::cout << "Base" << std::endl;
    }
};

class Derived : public Base {
public:
    void func() override {                // Override (C++11)
        std::cout << "Derived" << std::endl;
    }
};

Base* ptr = new Derived();
ptr->func();                              // Calls Derived::func() (runtime binding)
```

**Rules:**
- Virtual functions use runtime binding (vtable)
- Non-virtual functions use compile-time binding
- Virtual destructor needed when deleting through base pointer
- `override` keyword (C++11) ensures you're overriding, not creating new function

### static
**Purpose**: Shared across all instances, belongs to class not object

```cpp
class Counter {
private:
    static int count;                     // Static member variable
    
public:
    Counter() {
        count++;                          // Increment shared counter
    }
    
    static int getCount() {               // Static member function
        return count;                     // Can only access static members
    }
};

int Counter::count = 0;                  // Must define outside class

// Usage
Counter c1, c2, c3;
std::cout << Counter::getCount();        // 3 (shared across all objects)
```

**Static Member Function:**
- Can be called without object: `ClassName::function()`
- Cannot access non-static members
- No `this` pointer

### const
**Purpose**: Prevents modification

```cpp
class MyClass {
private:
    int x;
    mutable int y;                        // Can be modified in const function
    
public:
    // Const member function - cannot modify members
    int getValue() const {
        // x = 10;                        // ❌ Error: cannot modify in const function
        y = 20;                           // ✅ OK: y is mutable
        return x;
    }
    
    // Const parameter
    void print(const std::string& str) {
        // str = "new";                   // ❌ Error: str is const
        std::cout << str << std::endl;
    }
    
    // Const object
    const MyClass obj(10);
    // obj.x = 20;                        // ❌ Error: cannot modify const object
    int val = obj.getValue();             // ✅ OK: const function can be called
};
```

**Const Rules:**
- Const member function: cannot modify non-mutable members
- Const object: can only call const member functions
- Const parameter: cannot be modified in function
- `mutable`: allows modification in const function

### friend
**Purpose**: Grants access to private/protected members

```cpp
class MyClass {
private:
    int x;
    
    friend void friendFunction(MyClass& obj);  // Friend function
    friend class FriendClass;                  // Friend class
    
public:
    MyClass(int val) : x(val) { }
};

// Friend function can access private members
void friendFunction(MyClass& obj) {
    obj.x = 100;                           // ✅ Can access private member
}

// Friend class can access private members
class FriendClass {
public:
    void accessPrivate(MyClass& obj) {
        obj.x = 200;                       // ✅ Can access private member
    }
};
```

**Friend Rules:**
- Friend is not a member of class
- Friendship is not inherited
- Friendship is not symmetric (A friend of B doesn't make B friend of A)

### override (C++11)
**Purpose**: Ensures function is actually overriding base class virtual function

```cpp
class Base {
public:
    virtual void func(int x) { }
};

class Derived : public Base {
public:
    void func(int x) override {           // ✅ Correct override
        // Implementation
    }
    
    // void func(double x) override { }   // ❌ Error: not overriding (different signature)
};
```

### final (C++11)
**Purpose**: Prevents overriding or inheritance

```cpp
class Base {
public:
    virtual void func() final { }         // Cannot be overridden
};

class Derived : public Base {
public:
    // void func() { }                    // ❌ Error: cannot override final function
};

class FinalClass final {                  // Cannot be inherited
    // ...
};

// class Child : public FinalClass { }    // ❌ Error: cannot inherit from final class
```

## Inheritance

### Access Specifiers
```cpp
class Base {
public:
    int publicVar;
protected:
    int protectedVar;
private:
    int privateVar;
};

// Public inheritance (most common)
class Derived1 : public Base {
    // publicVar: public
    // protectedVar: protected
    // privateVar: not accessible
};

// Protected inheritance
class Derived2 : protected Base {
    // publicVar: protected
    // protectedVar: protected
    // privateVar: not accessible
};

// Private inheritance
class Derived3 : private Base {
    // publicVar: private
    // protectedVar: private
    // privateVar: not accessible
};
```

### Multiple Inheritance
```cpp
class A {
public:
    void funcA() { }
};

class B {
public:
    void funcB() { }
};

class C : public A, public B {
public:
    void funcC() {
        funcA();                           // From A
        funcB();                           // From B
    }
};
```

## Common Patterns

### Virtual Destructor Pattern
```cpp
class Base {
public:
    virtual ~Base() { }                   // Always make destructor virtual in base class
};

class Derived : public Base {
public:
    ~Derived() {                          // Will be called when deleting Base*
        // Cleanup
    }
};
```

### Interface Pattern
```cpp
class IComparable {
public:
    virtual int compare(const IComparable& other) = 0;
    virtual ~IComparable() { }
};

class Number : public IComparable {
private:
    int value;
public:
    Number(int v) : value(v) { }
    int compare(const IComparable& other) override {
        const Number& num = static_cast<const Number&>(other);
        return value - num.value;
    }
};
```

## Common Pitfalls

```cpp
// Wrong: Missing virtual destructor
class Base {
public:
    ~Base() { }                           // ❌ Not virtual
};

class Derived : public Base {
    int* data;
public:
    Derived() { data = new int[10]; }
    ~Derived() { delete[] data; }
};

Base* ptr = new Derived();
delete ptr;                               // ❌ Only Base destructor called, memory leak!

// Correct: Virtual destructor
class Base {
public:
    virtual ~Base() { }                   // ✅ Virtual
};

// Wrong: Calling virtual function in constructor
class Base {
public:
    Base() {
        func();                           // ❌ Calls Base::func(), not Derived::func()
    }
    virtual void func() { }
};

// Correct: Call after construction
class Base {
public:
    void init() {
        func();                           // ✅ Calls correct virtual function
    }
    virtual void func() { }
};
```

