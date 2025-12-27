# OOP Concepts - Learning Guide for LLD & Machine Coding

**Language-Agnostic Learning Order** - Concepts apply to C++, Java, Python, etc.

## Learning Path

1. **Basic Concepts** → 2. **Encapsulation** → 3. **Constructors/Destructors** → 4. **Inheritance** → 5. **Polymorphism** → 6. **Abstraction** → 7. **Advanced Keywords**

---

## 1. Basic OOP Concepts

### Class and Object
**Concept**: Class is a blueprint, Object is an instance

```cpp
// Class definition (blueprint)
class Person {
private:
    string name;
    int age;
    
public:
    void setName(string n) { name = n; }
    string getName() { return name; }
};

// Object creation (instance)
Person p1;  // p1 is an object of Person class
```

**Key Points:**
- Class = Template/Blueprint
- Object = Real instance created from class
- Multiple objects can be created from one class

---

## 2. Encapsulation

**Concept**: Bundling data and methods together, hiding internal details

### Access Modifiers
**Concept**: Control visibility of class members

```cpp
class Example {
private:
    // Only accessible within this class
    int secret;
    
protected:
    // Accessible in this class and derived classes
    int shared;
    
public:
    // Accessible everywhere
    int publicData;
};
```

**Access Levels:**
- **Private**: Only within class (data hiding)
- **Protected**: Within class + derived classes (inheritance)
- **Public**: Everywhere (interface)

**Why Encapsulation?**
- Data protection (prevent invalid states)
- Controlled access (getters/setters)
- Implementation hiding (change internals without affecting users)

### Getters and Setters
**Concept**: Controlled access to private data

```cpp
class BankAccount {
private:
    double balance;  // Hidden data
    
public:
    // Getter - read access
    double getBalance() {
        return balance;
    }
    
    // Setter - write access with validation
    void setBalance(double amount) {
        if (amount >= 0) {
            balance = amount;
        }
    }
};
```

---

## 3. Constructors and Destructors

### Constructor
**Concept**: Special method called when object is created - initializes object

```cpp
class Person {
private:
    string name;
    int age;
    
public:
    // Default constructor (no parameters)
    Person() {
        name = "";
        age = 0;
    }
    
    // Parameterized constructor
    Person(string n, int a) {
        name = n;
        age = a;
    }
    
    // Copy constructor
    Person(const Person& other) {
        name = other.name;
        age = other.age;
    }
};

// Usage
Person p1;                    // Default constructor
Person p2("John", 25);        // Parameterized constructor
Person p3(p2);               // Copy constructor
```

**Constructor Types:**
1. **Default**: No parameters, sets default values
2. **Parameterized**: Takes parameters, custom initialization
3. **Copy**: Creates copy of existing object

**Initialization List** (C++ specific but concept applies):
```cpp
// Efficient initialization
Person(string n, int a) : name(n), age(a) {
    // Members initialized before body executes
}
```

### Destructor
**Concept**: Special method called when object is destroyed - cleanup resources

```cpp
class Resource {
private:
    int* data;
    
public:
    Resource(int size) {
        data = new int[size];  // Allocate memory
    }
    
    // Destructor - cleanup
    ~Resource() {
        delete[] data;  // Free memory
    }
};
```

**When Destructor Called:**
- Object goes out of scope
- `delete` is called (for pointers)
- Program ends

**Important for LLD:**
- Always clean up resources (memory, files, connections)
- Virtual destructor needed for polymorphism (see section 7 - virtual keyword)

---

## 4. Inheritance

**Concept**: Create new class based on existing class - "is-a" relationship

### Basic Inheritance
```cpp
// Base class (Parent)
class Animal {
protected:
    string name;
    
public:
    void setName(string n) { name = n; }
    void eat() {
        cout << name << " is eating" << endl;
    }
};

// Derived class (Child)
class Dog : public Animal {
public:
    void bark() {
        cout << name << " is barking" << endl;
    }
};

// Usage
Dog d;
d.setName("Buddy");  // Inherited from Animal
d.eat();             // Inherited from Animal
d.bark();            // Own method
```

**Inheritance Types:**
- **Single**: One base class
- **Multiple**: Multiple base classes (C++ supports)
- **Multilevel**: Chain of inheritance
- **Hierarchical**: Multiple derived from one base

### Access Modes in Inheritance
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

### Calling Parent Constructor from Child
**Concept**: Child class must initialize parent class - pass attributes to parent constructor

```cpp
class Person {
protected:
    string name;
    int age;
    
public:
    // Parent constructor
    Person(string n, int a) : name(n), age(a) {
        cout << "Person constructor: " << name << endl;
    }
};

class Student : public Person {
private:
    int studentId;
    
public:
    // Child constructor - calls parent constructor
    Student(string n, int a, int id) : Person(n, a), studentId(id) {
        // Person(n, a) calls parent constructor with name and age
        // studentId(id) initializes child's own member
        cout << "Student constructor: " << name << ", ID: " << studentId << endl;
    }
};

// Usage
Student s("John", 20, 12345);
// Output:
// Person constructor: John
// Student constructor: John, ID: 12345
```

**Key Points:**
- Child constructor must call parent constructor (if parent has no default constructor)
- Use initialization list: `: ParentClass(params)`
- Parent constructor called before child constructor body
- Can pass child's parameters to parent constructor

### Order of Constructor and Destructor in Inheritance
**Concept**: Constructors called top-down (base to derived), Destructors called bottom-up (derived to base)

```cpp
class Grandparent {
public:
    Grandparent() {
        cout << "Grandparent constructor" << endl;
    }
    ~Grandparent() {
        cout << "Grandparent destructor" << endl;
    }
};

class Parent : public Grandparent {
public:
    Parent() {
        cout << "Parent constructor" << endl;
    }
    ~Parent() {
        cout << "Parent destructor" << endl;
    }
};

class Child : public Parent {
public:
    Child() {
        cout << "Child constructor" << endl;
    }
    ~Child() {
        cout << "Child destructor" << endl;
    }
};

// Usage
Child c;
// Constructor order (top-down):
// Grandparent constructor
// Parent constructor
// Child constructor

// When object goes out of scope, destructor order (bottom-up):
// Child destructor
// Parent destructor
// Grandparent destructor
```

**Order Rules:**
- **Constructors**: Base → Derived (top to bottom in hierarchy)
- **Destructors**: Derived → Base (bottom to top in hierarchy)
- Parent must be fully constructed before child
- Child must be fully destroyed before parent

### Multilevel Inheritance - Detailed Operations
**Concept**: Chain of inheritance - Grandparent → Parent → Child

```cpp
class Animal {
protected:
    string name;
    int age;
    
public:
    Animal(string n, int a) : name(n), age(a) {
        cout << "Animal created: " << name << endl;
    }
    
    void eat() {
        cout << name << " is eating" << endl;
    }
    
    void display() {
        cout << "Animal: " << name << ", Age: " << age << endl;
    }
    
    ~Animal() {
        cout << "Animal destroyed: " << name << endl;
    }
};

class Mammal : public Animal {
protected:
    string furColor;
    
public:
    Mammal(string n, int a, string color) : Animal(n, a), furColor(color) {
        cout << "Mammal created: " << name << ", Color: " << furColor << endl;
    }
    
    void breathe() {
        cout << name << " is breathing" << endl;
    }
    
    void display() {  // Override parent method
        Animal::display();  // Call parent method
        cout << "Fur Color: " << furColor << endl;
    }
    
    ~Mammal() {
        cout << "Mammal destroyed: " << name << endl;
    }
};

class Dog : public Mammal {
private:
    string breed;
    
public:
    Dog(string n, int a, string color, string b) 
        : Mammal(n, a, color), breed(b) {
        cout << "Dog created: " << name << ", Breed: " << breed << endl;
    }
    
    void bark() {
        cout << name << " is barking" << endl;
    }
    
    void display() {  // Override parent method
        Mammal::display();  // Call parent method
        cout << "Breed: " << breed << endl;
    }
    
    ~Dog() {
        cout << "Dog destroyed: " << name << endl;
    }
};

// Usage
Dog d("Buddy", 3, "Brown", "Labrador");

// Constructor order:
// Animal created: Buddy
// Mammal created: Buddy, Color: Brown
// Dog created: Buddy, Breed: Labrador

// Method calling
d.eat();        // From Animal (inherited)
d.breathe();    // From Mammal (inherited)
d.bark();       // From Dog (own method)
d.display();    // From Dog (overridden, calls parent's display)

// Destructor order (when object goes out of scope):
// Dog destroyed: Buddy
// Mammal destroyed: Buddy
// Animal destroyed: Buddy
```

**Multilevel Inheritance Operations:**

1. **Constructor Calling:**
   ```cpp
   Dog(string n, int a, string c, string b) 
       : Mammal(n, a, c), breed(b) { }
   // Mammal constructor calls Animal constructor
   // Order: Animal → Mammal → Dog
   ```

2. **Method Access:**
   ```cpp
   d.eat();           // ✅ Direct access to Animal method
   d.breathe();       // ✅ Direct access to Mammal method
   d.bark();          // ✅ Direct access to Dog method
   d.Animal::eat();   // ✅ Explicit parent method call
   ```

3. **Method Overriding:**
   ```cpp
   // Each level can override parent method
   void display() {
       Mammal::display();  // Call immediate parent
       // Add own functionality
   }
   ```

4. **Attribute Access:**
   ```cpp
   // Protected members accessible in derived classes
   // name, age (from Animal) - accessible in Mammal and Dog
   // furColor (from Mammal) - accessible in Dog
   // breed (from Dog) - only in Dog
   ```

5. **Pointer/Reference Operations:**
   ```cpp
   Dog d("Buddy", 3, "Brown", "Labrador");
   Animal* aPtr = &d;      // ✅ Base pointer to derived object
   Mammal* mPtr = &d;      // ✅ Intermediate base pointer
   Dog* dPtr = &d;         // ✅ Same type pointer
   
   aPtr->eat();            // ✅ Calls Animal::eat()
   aPtr->display();        // ⚠️ Calls Animal::display() (if not virtual)
   // If display() is virtual, calls Dog::display()
   ```

**Key Points:**
- Each level adds its own attributes and methods
- Can override methods at any level
- Constructor chain: Top to bottom
- Destructor chain: Bottom to top
- Protected members accessible down the chain

**Key Points:**
- Inheritance enables code reuse
- Derived class "is-a" base class
- Can add new features or override existing ones
- Child must call parent constructor (if no default)
- Constructors: Base → Derived, Destructors: Derived → Base

---

## 5. Polymorphism

**Concept**: Same interface, different implementations - "one interface, many forms"

### Compile-time Polymorphism (Static Binding)
**Concept**: Function/operator overloading - resolved at compile time

```cpp
// Function Overloading
class Calculator {
public:
    int add(int a, int b) {
        return a + b;
    }
    
    double add(double a, double b) {
        return a + b;
    }
    
    int add(int a, int b, int c) {
        return a + b + c;
    }
};

// Operator Overloading
class Vector {
    int x, y;
public:
    Vector(int x, int y) : x(x), y(y) {}
    
    Vector operator+(const Vector& other) {
        return Vector(x + other.x, y + other.y);
    }
};
```

### Runtime Polymorphism (Dynamic Binding)
**Concept**: Virtual functions - resolved at runtime

```cpp
class Shape {
public:
    // Virtual function - can be overridden
    virtual void draw() {
        cout << "Drawing shape" << endl;
    }
    
    // Virtual destructor (important!)
    virtual ~Shape() {}
};

class Circle : public Shape {
public:
    void draw() override {  // Override virtual function
        cout << "Drawing circle" << endl;
    }
};

class Rectangle : public Shape {
public:
    void draw() override {
        cout << "Drawing rectangle" << endl;
    }
};

// Polymorphism in action
void drawShape(Shape* shape) {
    shape->draw();  // Calls appropriate draw() based on actual object type
}

// Usage
Circle c;
Rectangle r;
drawShape(&c);  // Calls Circle::draw()
drawShape(&r);  // Calls Rectangle::draw()
```

**Virtual Functions:**
- `virtual` keyword enables runtime binding
- Without `virtual`: calls base class function (compile-time binding)
- With `virtual`: calls derived class function (runtime binding)

**Virtual Destructor:**
```cpp
class Base {
public:
    virtual ~Base() {}  // Virtual destructor
};

class Derived : public Base {
    int* data;
public:
    Derived() { data = new int[10]; }
    ~Derived() { delete[] data; }
};

Base* ptr = new Derived();
delete ptr;  // ✅ Calls Derived destructor (if Base has virtual destructor)
```

**Key Points:**
- Polymorphism enables flexible, extensible code
- Runtime polymorphism uses virtual functions
- Always use virtual destructor in base class for polymorphism

---

## 6. Abstraction

**Concept**: Hide implementation details, show only essential features

### Abstract Classes
**Concept**: Class with at least one pure virtual function - cannot be instantiated

```cpp
// Abstract class
class Shape {
public:
    // Pure virtual function - must be implemented by derived class
    virtual double area() = 0;
    
    // Virtual function with default implementation
    virtual void draw() {
        cout << "Drawing shape" << endl;
    }
    
    virtual ~Shape() {}
};

// Cannot create object
// Shape s;  // ❌ Error: abstract class

// Derived class must implement pure virtual function
class Circle : public Shape {
    double radius;
public:
    Circle(double r) : radius(r) {}
    
    double area() override {  // Must implement
        return 3.14159 * radius * radius;
    }
};

// Now can create object
Circle c(5.0);  // ✅ OK
```

### Interface (Pure Abstract Class)
**Concept**: All methods are pure virtual - defines contract

```cpp
// Interface - all pure virtual functions
class Drawable {
public:
    virtual void draw() = 0;
    virtual void resize(int factor) = 0;
    virtual ~Drawable() {}
};

// Class implementing interface
class Circle : public Drawable {
public:
    void draw() override {
        // Implementation
    }
    
    void resize(int factor) override {
        // Implementation
    }
};
```

**Key Points:**
- Abstract class: Cannot instantiate, defines interface
- Pure virtual function: Must be implemented by derived class
- Interface: Contract that classes must follow
- Enables design by contract

---

## 7. Important Keywords & Concepts

### virtual
**Concept**: Enables runtime polymorphism - function binding happens at runtime

```cpp
class Base {
public:
    // Non-virtual function - compile-time binding
    void nonVirtual() {
        cout << "Base::nonVirtual" << endl;
    }
    
    // Virtual function - runtime binding
    virtual void virtualFunc() {
        cout << "Base::virtualFunc" << endl;
    }
    
    // Pure virtual function - must be overridden
    virtual void pureVirtual() = 0;
    
    // Virtual destructor - important!
    virtual ~Base() {}
};

class Derived : public Base {
public:
    void nonVirtual() {  // Hides base function (not override)
        cout << "Derived::nonVirtual" << endl;
    }
    
    void virtualFunc() override {  // Overrides base function
        cout << "Derived::virtualFunc" << endl;
    }
    
    void pureVirtual() override {  // Must implement
        cout << "Derived::pureVirtual" << endl;
    }
};

// Usage
Base* ptr = new Derived();
ptr->nonVirtual();    // Calls Base::nonVirtual() (compile-time)
ptr->virtualFunc();    // Calls Derived::virtualFunc() (runtime)
ptr->pureVirtual();   // Calls Derived::pureVirtual() (runtime)
delete ptr;            // Calls Derived destructor (if Base has virtual destructor)
```

**Virtual Function Rules:**
- **Without virtual**: Function call resolved at compile-time (early binding)
- **With virtual**: Function call resolved at runtime (late binding)
- **Pure virtual** (`= 0`): Must be implemented by derived class
- **Virtual destructor**: Always use in base class for polymorphism
- **override** keyword: Ensures you're actually overriding (C++11)

**When to Use Virtual:**
- When you need runtime polymorphism
- When derived classes need different implementations
- Always for destructor in base class
- For functions that may be overridden

**Performance:**
- Virtual functions have slight overhead (vtable lookup)
- Use when polymorphism is needed, avoid when not needed

**Virtual with Constructors/Destructors:**
- **Virtual Destructor**: Always use virtual destructor in base class when using polymorphism
  ```cpp
  class Base {
  public:
      virtual ~Base() {}  // Virtual destructor - ensures correct cleanup
  };
  class Derived : public Base {
      int* data;
  public:
      Derived() { data = new int[10]; }
      ~Derived() { delete[] data; }  // Will be called if Base has virtual destructor
  };
  Base* ptr = new Derived();
  delete ptr;  // ✅ Calls Derived destructor, then Base destructor
  // Without virtual: ❌ Only Base destructor called (memory leak!)
  ```
- **Virtual Constructor**: Not possible in C++ (constructors cannot be virtual)
- **Virtual in Constructor**: Calling virtual function in constructor calls base version (object not fully constructed yet)

### static
**Concept**: Shared across all instances - belongs to class, not object

```cpp
class Counter {
private:
    static int count;  // Shared by all objects
    
public:
    Counter() {
        count++;  // Increment shared counter
    }
    
    static int getCount() {  // Static function
        return count;  // Can only access static members
    }
};

int Counter::count = 0;  // Must define outside class

// Usage
Counter c1, c2, c3;
cout << Counter::getCount();  // 3 (shared across all)
```

**Static Members:**
- Shared by all objects of class
- Only one copy exists (not per object)
- Can be accessed without object: `ClassName::member`

**Static Functions:**
- Can be called without object
- Cannot access non-static members
- No `this` pointer

**Static with Constructors/Destructors:**
- **Static Constructor/Destructor**: Not possible in C++ (constructors/destructors operate on instances)
- **Static Members in Constructor/Destructor**: Static members can be accessed/modified in constructors/destructors
  ```cpp
  class MyClass {
  private:
      static int count;  // Static member
  public:
      MyClass() {
          count++;  // ✅ Can modify static member in constructor
      }
      ~MyClass() {
          count--;  // ✅ Can modify static member in destructor
      }
  };
  int MyClass::count = 0;  // Initialize before any object creation
  ```
- **Static Member Initialization**: Static members initialized before any constructor is called

### const
**Concept**: Prevents modification

```cpp
class MyClass {
private:
    int x;
    mutable int y;  // Can be modified in const function
    
public:
    // Const member function - cannot modify members
    int getValue() const {
        // x = 10;  // ❌ Error
        y = 20;     // ✅ OK (mutable)
        return x;
    }
    
    // Const parameter
    void print(const string& str) {
        // str = "new";  // ❌ Error
        cout << str << endl;
    }
};

// Const object
const MyClass obj(10);
// obj.x = 20;  // ❌ Error
int val = obj.getValue();  // ✅ OK (const function)
```

**Const Rules:**
- Const function: Cannot modify non-mutable members
- Const object: Can only call const functions
- Const parameter: Cannot be modified
- `mutable`: Allows modification in const function

**Const with Constructors/Destructors:**
- **Const Constructor/Destructor**: Not possible in C++ (constructors/destructors need to modify object)
- **Const Members in Constructor**: Const members must be initialized in initialization list, not assigned in body
  ```cpp
  class MyClass {
  private:
      const int id;        // Const member
      const int& ref;      // Const reference
  public:
      // ❌ Wrong: Cannot assign in constructor body
      // MyClass(int i, int& r) {
      //     id = i;      // Error: const cannot be assigned
      //     ref = r;     // Error: reference must be initialized
      // }
      // ✅ Correct: Must use initialization list
      MyClass(int i, int& r) : id(i), ref(r) {
          // id and ref initialized, not assigned
      }
  };
  ```
- **Const Parameters**: Can use const parameters in constructors/destructors

### friend
**Concept**: Grants access to private/protected members

```cpp
class MyClass {
private:
    int x;
    
    friend void friendFunction(MyClass& obj);  // Friend function
    friend class FriendClass;                  // Friend class
    
public:
    MyClass(int val) : x(val) {}
};

// Friend function can access private members
void friendFunction(MyClass& obj) {
    obj.x = 100;  // ✅ Can access private member
}

// Friend class can access private members
class FriendClass {
public:
    void accessPrivate(MyClass& obj) {
        obj.x = 200;  // ✅ Can access private member
    }
};
```

**Friend Rules:**
- Not a member of class
- Friendship is not inherited
- Use sparingly (breaks encapsulation)

**Friend with Constructors/Destructors:**
- **Friend Constructor/Destructor**: Not possible (friend applies to functions/classes, not constructors themselves)
- **Friend Access to Private Constructor/Destructor**: Friend functions/classes can access private constructors/destructors
  ```cpp
  class Singleton {
  private:
      Singleton() { }  // Private constructor
      ~Singleton() { } // Private destructor
      friend class SingletonFactory;  // Friend class
      friend Singleton* createSingleton();  // Friend function
  };
  class SingletonFactory {
  public:
      static Singleton* create() {
          return new Singleton();  // ✅ Can access private constructor
      }
  };
  Singleton* createSingleton() {
      return new Singleton();  // ✅ Friend function can access
  }
  ```
- **Use Case**: Enables Singleton pattern and controlled object creation/destruction

### override (C++11)
**Concept**: Ensures function is actually overriding base class virtual function

```cpp
class Base {
public:
    virtual void func(int x) {}
};

class Derived : public Base {
public:
    void func(int x) override {  // ✅ Correct override
        // Implementation
    }
    
    // void func(double x) override { }  // ❌ Error: not overriding
};
```

### final (C++11)
**Concept**: Prevents overriding or inheritance

```cpp
class Base {
public:
    virtual void func() final {}  // Cannot be overridden
};

class Derived : public Base {
public:
    // void func() { }  // ❌ Error: cannot override final function
};

class FinalClass final {  // Cannot be inherited
    // ...
};

// class Child : public FinalClass { }  // ❌ Error
```

### private
**Concept**: Access modifier that restricts member access to within the class only - core of encapsulation

```cpp
class BankAccount {
private:
    double balance;        // Private data - hidden from outside
    string accountNumber;  // Private data
    
    // Private helper function
    bool validateAmount(double amount) {
        return amount > 0;
    }
    
public:
    // Public interface - controlled access
    void deposit(double amount) {
        if (validateAmount(amount)) {  // Can use private function
            balance += amount;
        }
    }
    
    double getBalance() {
        return balance;  // Can access private member
    }
};

// Usage
BankAccount account;
// account.balance = 1000;        // ❌ Error: private member
// account.validateAmount(100);      // ❌ Error: private function
account.deposit(1000);                // ✅ OK: public function
double bal = account.getBalance();    // ✅ OK: public function
```

**Private Access Rules:**
- Accessible only within the same class
- Not accessible from outside the class
- Not accessible in derived classes (use `protected` for that)
- Can be accessed by friend functions/classes

**Impact and Use Cases:**

#### 1. Data Hiding (Encapsulation)
```cpp
class Temperature {
private:
    double celsius;  // Private - cannot be directly modified
    
public:
    void setCelsius(double c) {
        if (c >= -273.15) {  // Validation
            celsius = c;
        }
    }
    
    double getCelsius() {
        return celsius;
    }
    
    double getFahrenheit() {
        return (celsius * 9/5) + 32;  // Controlled conversion
    }
};

// Prevents invalid states
Temperature t;
// t.celsius = -300;  // ❌ Cannot set invalid temperature directly
t.setCelsius(25);    // ✅ Validated access
```

#### 2. Implementation Hiding
```cpp
class Stack {
private:
    vector<int> data;  // Private implementation - can change without affecting users
    
public:
    void push(int value) {
        data.push_back(value);
    }
    
    int pop() {
        if (!data.empty()) {
            int val = data.back();
            data.pop_back();
            return val;
        }
        return -1;
    }
};

// Users don't need to know about vector - implementation can change
// Could change to array, linked list, etc. without breaking user code
```

#### 3. Private Constructors (Singleton Pattern)
```cpp
class Database {
private:
    static Database* instance;
    
    // Private constructor - prevents direct object creation
    Database() {
        // Initialize database connection
    }
    
public:
    // Static method to get instance
    static Database* getInstance() {
        if (instance == nullptr) {
            instance = new Database();
        }
        return instance;
    }
    
    // Prevent copying
    Database(const Database&) = delete;
    Database& operator=(const Database&) = delete;
};

Database* Database::instance = nullptr;

// Usage
// Database db;  // ❌ Error: private constructor
Database* db = Database::getInstance();  // ✅ OK: controlled creation
```

#### 4. Private Destructor
```cpp
class Resource {
private:
    ~Resource() {  // Private destructor
        // Cleanup
    }
    
    friend class ResourceManager;  // Only friend can delete
};

class ResourceManager {
public:
    void cleanup(Resource* r) {
        delete r;  // ✅ OK: friend can access private destructor
    }
};

// Usage
Resource* r = new Resource();
// delete r;  // ❌ Error: private destructor
ResourceManager manager;
manager.cleanup(r);  // ✅ OK: controlled destruction
```

#### 5. Private Inheritance
```cpp
class Base {
public:
    void publicFunc() {}
protected:
    void protectedFunc() {}
};

// Private inheritance - "implemented in terms of"
class Derived : private Base {
public:
    void useBase() {
        publicFunc();      // ✅ Can access (now private in Derived)
        protectedFunc();   // ✅ Can access (now private in Derived)
    }
};

// Usage
Derived d;
// d.publicFunc();  // ❌ Error: now private in Derived
d.useBase();        // ✅ OK: public interface
```

**Private vs Protected vs Public:**

| Access Modifier | Same Class | Derived Class | Outside Class |
|----------------|------------|---------------|---------------|
| **private**   | ✅ Yes     | ❌ No         | ❌ No         |
| **protected** | ✅ Yes     | ✅ Yes        | ❌ No         |
| **public**    | ✅ Yes     | ✅ Yes        | ✅ Yes        |

**Use Cases for Private:**
1. **Data hiding**: Protect internal state from invalid modifications
2. **Implementation hiding**: Hide internal details, expose only interface
3. **Singleton pattern**: Control object creation
4. **Resource management**: Control object destruction
5. **Helper functions**: Internal utility functions not part of public API
6. **Private inheritance**: "Implemented in terms of" relationship

**Best Practices:**
- Make data members private by default
- Provide public getters/setters for controlled access
- Use private for implementation details
- Use protected when derived classes need access
- Use public only for the interface

---

## 8. Design Principles for LLD

### SOLID Principles

**S - Single Responsibility Principle**
- Class should have one reason to change
- One class = one responsibility

**O - Open/Closed Principle**
- Open for extension, closed for modification
- Use inheritance/polymorphism to extend

**L - Liskov Substitution Principle**
- Derived class should be substitutable for base class
- Don't break base class contract

**I - Interface Segregation Principle**
- Many specific interfaces > one general interface
- Clients shouldn't depend on unused methods

**D - Dependency Inversion Principle**
- Depend on abstractions, not concretions
- High-level modules shouldn't depend on low-level modules

### Composition vs Inheritance
**Concept**: "Has-a" vs "Is-a" relationship

```cpp
// Inheritance (Is-a)
class Car : public Vehicle {
    // Car IS-A Vehicle
};

// Composition (Has-a)
class Car {
private:
    Engine engine;      // Car HAS-A Engine
    Wheel wheels[4];    // Car HAS-A Wheels
};
```

**When to Use:**
- **Inheritance**: "Is-a" relationship, code reuse, polymorphism
- **Composition**: "Has-a" relationship, more flexible, less coupling

---

## 9. Common Patterns for Machine Coding

### Factory Pattern (Basic)
**Concept**: Create objects without specifying exact class

```cpp
class Animal {
public:
    virtual void makeSound() = 0;
};

class Dog : public Animal {
public:
    void makeSound() override {
        cout << "Woof" << endl;
    }
};

class Cat : public Animal {
public:
    void makeSound() override {
        cout << "Meow" << endl;
    }
};

// Factory function
Animal* createAnimal(string type) {
    if (type == "dog") return new Dog();
    if (type == "cat") return new Cat();
    return nullptr;
}
```

### Singleton Pattern (Basic)
**Concept**: Only one instance of class exists

```cpp
class Database {
private:
    static Database* instance;
    Database() {}  // Private constructor
    
public:
    static Database* getInstance() {
        if (instance == nullptr) {
            instance = new Database();
        }
        return instance;
    }
};

Database* Database::instance = nullptr;
```

---

## 10. Common Pitfalls

```cpp
// ❌ Missing virtual destructor
class Base {
public:
    ~Base() {}  // Not virtual
};

class Derived : public Base {
    int* data;
public:
    Derived() { data = new int[10]; }
    ~Derived() { delete[] data; }
};

Base* ptr = new Derived();
delete ptr;  // ❌ Only Base destructor called - memory leak!

// ✅ Virtual destructor
class Base {
public:
    virtual ~Base() {}  // Virtual
};
```

```cpp
// ❌ Calling virtual function in constructor
class Base {
public:
    Base() {
        func();  // ❌ Calls Base::func(), not Derived::func()
    }
    virtual void func() {}
};
```

```cpp
// ❌ Slicing problem
class Base { };
class Derived : public Base { };

Derived d;
Base b = d;  // ❌ Slicing - Derived part is lost

// ✅ Use pointers/references
Base* ptr = &d;  // ✅ OK
Base& ref = d;   // ✅ OK
```

---

## Learning Checklist

- [ ] Understand class vs object
- [ ] Master encapsulation (private, protected, public)
- [ ] Learn constructors and destructors
- [ ] Understand inheritance (single, multiple, multilevel)
- [ ] Master polymorphism (compile-time and runtime)
- [ ] Understand abstraction (abstract classes, interfaces)
- [ ] Learn important keywords (static, const, friend, virtual)
- [ ] Understand SOLID principles
- [ ] Know when to use composition vs inheritance
- [ ] Practice common design patterns

---

## Language-Specific Notes

### C++
- `virtual` keyword for runtime polymorphism
- `= 0` for pure virtual functions
- Multiple inheritance supported
- Virtual destructor required for polymorphism

### Java
- `abstract` keyword for abstract classes
- `interface` keyword for interfaces
- Single inheritance, multiple interfaces
- All methods virtual by default

### Python
- No explicit access modifiers (convention: `_` for private)
- Duck typing (polymorphism without inheritance)
- Multiple inheritance supported
- `@abstractmethod` decorator for abstract methods

---

**Remember**: Concepts are language-agnostic, syntax varies. Master the concepts first!
