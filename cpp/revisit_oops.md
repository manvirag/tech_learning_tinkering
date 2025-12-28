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
public:
    void testAccess() {
        publicVar = 10;      // ✅ OK: publicVar is public in Derived1
        protectedVar = 20;   // ✅ OK: protectedVar is protected in Derived1
        // privateVar = 30;  // ❌ Error: privateVar not accessible
    }
};

// Usage
Derived1 d1;
d1.publicVar = 10;        // ✅ OK: publicVar remains public
// d1.protectedVar = 20;  // ❌ Error: protectedVar is protected (not accessible outside)
// d1.privateVar = 30;    // ❌ Error: privateVar not accessible

// Protected inheritance
class Derived2 : protected Base {
public:
    void testAccess() {
        publicVar = 10;      // ✅ OK: publicVar becomes protected in Derived2
        protectedVar = 20;   // ✅ OK: protectedVar remains protected
        // privateVar = 30;  // ❌ Error: privateVar not accessible
    }
};

// Usage
Derived2 d2;
// d2.publicVar = 10;      // ❌ Error: publicVar is now protected (not accessible outside)
// d2.protectedVar = 20;   // ❌ Error: protectedVar is protected
// d2.privateVar = 30;     // ❌ Error: privateVar not accessible

// Private inheritance
class Derived3 : private Base {
public:
    void testAccess() {
        publicVar = 10;      // ✅ OK: publicVar becomes private in Derived3
        protectedVar = 20;   // ✅ OK: protectedVar becomes private in Derived3
        // privateVar = 30;  // ❌ Error: privateVar not accessible
    }
};

// Usage
Derived3 d3;
// d3.publicVar = 10;      // ❌ Error: publicVar is now private (not accessible outside)
// d3.protectedVar = 20;   // ❌ Error: protectedVar is now private
// d3.privateVar = 30;     // ❌ Error: privateVar not accessible
```

**What This Means:**
- **publicVar: public** → Can be accessed from outside Derived1 class
- **protectedVar: protected** → Can be accessed in Derived1 and its derived classes, but not from outside
- **privateVar: not accessible** → Cannot be accessed in Derived1 at all (only in Base)
- Access mode in inheritance changes how base class members are accessible in derived class

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

**Alternative Syntax for Calling Parent Constructor:**

```cpp
// Method 1: Initialization list (most common)
Student(string n, int a, int id) : Person(n, a), studentId(id) {
    // Parent constructor called via initialization list
}

// Method 2: If parent has default constructor, it's called automatically
class Person {
public:
    Person() { }  // Default constructor
    Person(string n, int a) : name(n), age(a) { }
};

class Student : public Person {
public:
    Student(int id) : studentId(id) {
        // Person() default constructor called automatically
        // Can then set parent members if they're protected/public
        name = "Unknown";
        age = 0;
    }
};

// Method 3: Delegating constructor (C++11)
class Student : public Person {
public:
    Student(string n, int a) : Person(n, a), studentId(0) { }
    
    Student(string n, int a, int id) : Student(n, a) {
        // Delegates to other constructor, then modifies
        studentId = id;
    }
};
```

**Key Points:**
- Child constructor must call parent constructor (if parent has no default constructor)
- Use initialization list: `: ParentClass(params)` - most common and efficient
- If parent has default constructor, it's called automatically (no need to specify)
- Parent constructor called before child constructor body
- Can pass child's parameters to parent constructor
- Delegating constructor (C++11): Can call another constructor of same class

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
   class Animal {
   protected:
       string name;    // Protected - accessible in derived classes
       int age;        // Protected - accessible in derived classes
   };
   
   class Mammal : public Animal {
   protected:
       string furColor;  // Protected - accessible in derived classes
   public:
       void setAnimalInfo(string n, int a) {
           name = n;      // ✅ OK: name from Animal is accessible
           age = a;       // ✅ OK: age from Animal is accessible
       }
       void setFurColor(string color) {
           furColor = color;  // ✅ OK: own member
       }
   };
   
   class Dog : public Mammal {
   private:
       string breed;  // Private - only in Dog
   public:
       void setAllInfo(string n, int a, string color, string b) {
           name = n;         // ✅ OK: name from Animal (grandparent) accessible
           age = a;          // ✅ OK: age from Animal (grandparent) accessible
           furColor = color; // ✅ OK: furColor from Mammal (parent) accessible
           breed = b;        // ✅ OK: own member
       }
       
       void displayInfo() {
           cout << "Name: " << name << endl;        // ✅ From Animal
           cout << "Age: " << age << endl;          // ✅ From Animal
           cout << "Color: " << furColor << endl;   // ✅ From Mammal
           cout << "Breed: " << breed << endl;      // ✅ Own member
       }
   };
   
   // Usage
   Dog d("Buddy", 3, "Brown", "Labrador");
   // d.name = "Max";        // ❌ Error: name is protected, not accessible outside
   // d.furColor = "Black"; // ❌ Error: furColor is protected, not accessible outside
   // d.breed = "Poodle";   // ❌ Error: breed is private, not accessible outside
   d.setAllInfo("Max", 4, "Black", "Poodle");  // ✅ OK: using public method
   ```
   
   **Attribute Access Summary:**
   - **name, age (from Animal)**: Accessible in Mammal and Dog (protected members)
   - **furColor (from Mammal)**: Accessible in Dog (protected member)
   - **breed (from Dog)**: Only accessible in Dog (private member)
   - Protected members flow down the inheritance chain
   - Private members stay in their own class

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

### Can Base Class Object/Pointer Call Child Class Methods?

**Concept**: Base class pointer/reference can only call methods that exist in the base class. If a method only exists in the child class, you cannot call it through a base pointer/reference.

#### Case 1: Normal Inheritance

```cpp
class Animal {
public:
    void eat() {
        cout << "Animal is eating" << endl;
    }
    
    virtual void makeSound() {
        cout << "Animal makes sound" << endl;
    }
};

class Dog : public Animal {
public:
    void makeSound() override {
        cout << "Dog barks" << endl;
    }
    
    // Method only in Dog (not in Animal)
    void fetch() {
        cout << "Dog fetches ball" << endl;
    }
};

// Usage
Dog d;
Animal* aPtr = &d;        // Base pointer to derived object
Animal& aRef = d;         // Base reference to derived object

// ✅ Can call methods that exist in base class
aPtr->eat();              // ✅ OK: eat() exists in Animal
aPtr->makeSound();        // ✅ OK: makeSound() exists in Animal (calls Dog::makeSound() if virtual)

// ❌ Cannot call methods that only exist in child class
// aPtr->fetch();         // ❌ Error: 'class Animal' has no member named 'fetch'
// aRef.fetch();          // ❌ Error: 'class Animal' has no member named 'fetch'

// ✅ Solution: Cast to derived type
Dog* dPtr = dynamic_cast<Dog*>(aPtr);
if (dPtr != nullptr) {
    dPtr->fetch();        // ✅ OK: Now can call fetch()
}

// Or use static_cast if you're sure about the type
static_cast<Dog*>(aPtr)->fetch();  // ✅ OK: But unsafe if not actually Dog
```

**Key Points:**
- Base pointer/reference can only see methods declared in base class
- Cannot call methods that only exist in derived class
- Must cast to derived type to call derived-only methods
- Virtual methods can be called through base pointer (polymorphism works)

#### Case 2: Abstract Class

```cpp
// Abstract class (has pure virtual function)
class Shape {
public:
    // Pure virtual function - must be implemented by derived
    virtual double area() = 0;
    
    // Virtual function with implementation
    virtual void draw() {
        cout << "Drawing shape" << endl;
    }
    
    // Regular function
    void printInfo() {
        cout << "This is a shape" << endl;
    }
    
    virtual ~Shape() {}
};

class Circle : public Shape {
private:
    double radius;
    
public:
    Circle(double r) : radius(r) {}
    
    // Must implement pure virtual function
    double area() override {
        return 3.14159 * radius * radius;
    }
    
    // Override virtual function
    void draw() override {
        cout << "Drawing circle" << endl;
    }
    
    // Method only in Circle (not in Shape)
    void setRadius(double r) {
        radius = r;
    }
    
    double getRadius() {
        return radius;
    }
};

// Usage
Circle c(5.0);
Shape* sPtr = &c;         // Base pointer to derived object
Shape& sRef = c;          // Base reference to derived object

// ✅ Can call methods that exist in base class
sPtr->area();             // ✅ OK: area() exists in Shape (calls Circle::area())
sPtr->draw();             // ✅ OK: draw() exists in Shape (calls Circle::draw())
sPtr->printInfo();        // ✅ OK: printInfo() exists in Shape

// ❌ Cannot call methods that only exist in child class
// sPtr->setRadius(10);   // ❌ Error: 'class Shape' has no member named 'setRadius'
// sPtr->getRadius();     // ❌ Error: 'class Shape' has no member named 'getRadius'
// sRef.setRadius(10);    // ❌ Error: 'class Shape' has no member named 'setRadius'

// ✅ Solution: Cast to derived type
Circle* cPtr = dynamic_cast<Circle*>(sPtr);
if (cPtr != nullptr) {
    cPtr->setRadius(10);  // ✅ OK: Now can call setRadius()
    cout << cPtr->getRadius() << endl;  // ✅ OK
}
```

**Key Points:**
- Same rule applies: base pointer can only call methods in base class
- Even though Shape is abstract, pointer still follows same rules
- Must cast to Circle to call Circle-specific methods

#### Case 3: Interface (Pure Abstract Class)

```cpp
// Interface - all pure virtual functions
class Drawable {
public:
    virtual void draw() = 0;
    virtual void resize(int factor) = 0;
    virtual ~Drawable() {}
};

class Rectangle : public Drawable {
private:
    int width, height;
    
public:
    Rectangle(int w, int h) : width(w), height(h) {}
    
    // Must implement interface methods
    void draw() override {
        cout << "Drawing rectangle: " << width << "x" << height << endl;
    }
    
    void resize(int factor) override {
        width *= factor;
        height *= factor;
    }
    
    // Methods only in Rectangle (not in Drawable interface)
    void setWidth(int w) {
        width = w;
    }
    
    void setHeight(int h) {
        height = h;
    }
    
    int getArea() {
        return width * height;
    }
};

// Usage
Rectangle r(10, 20);
Drawable* dPtr = &r;      // Interface pointer to implementing object
Drawable& dRef = r;       // Interface reference to implementing object

// ✅ Can call methods that exist in interface
dPtr->draw();             // ✅ OK: draw() exists in Drawable (calls Rectangle::draw())
dPtr->resize(2);          // ✅ OK: resize() exists in Drawable (calls Rectangle::resize())

// ❌ Cannot call methods that only exist in implementing class
// dPtr->setWidth(15);    // ❌ Error: 'class Drawable' has no member named 'setWidth'
// dPtr->setHeight(25);   // ❌ Error: 'class Drawable' has no member named 'setHeight'
// dPtr->getArea();       // ❌ Error: 'class Drawable' has no member named 'getArea'
// dRef.setWidth(15);     // ❌ Error: 'class Drawable' has no member named 'setWidth'

// ✅ Solution: Cast to implementing type
Rectangle* rPtr = dynamic_cast<Rectangle*>(dPtr);
if (rPtr != nullptr) {
    rPtr->setWidth(15);   // ✅ OK: Now can call setWidth()
    rPtr->setHeight(25);  // ✅ OK
    cout << rPtr->getArea() << endl;  // ✅ OK
}
```

**Key Points:**
- Same rule applies to interfaces
- Interface pointer can only call methods declared in interface
- Cannot call methods that only exist in implementing class
- Must cast to implementing type to call class-specific methods

#### Summary Table

| Scenario | Can Base Pointer Call Child Method? | Example |
|----------|-------------------------------------|---------|
| **Method exists in base (virtual)** | ✅ Yes (polymorphism) | `basePtr->virtualFunc()` calls derived version |
| **Method exists in base (non-virtual)** | ✅ Yes (calls base version) | `basePtr->nonVirtualFunc()` calls base version |
| **Method only in child** | ❌ No (compile error) | `basePtr->childOnlyFunc()` → Error |
| **Solution: Cast to child** | ✅ Yes (after cast) | `static_cast<Child*>(basePtr)->childOnlyFunc()` |

#### Why This Restriction Exists?

```cpp
class Animal {
public:
    void eat() { }
};

class Dog : public Animal {
public:
    void fetch() { }
};

class Cat : public Animal {
public:
    void meow() { }
};

// If base pointer could call child methods, this would be possible:
Animal* ptr = new Dog();
ptr->fetch();  // ✅ Works if Dog

ptr = new Cat();  // Now points to Cat
ptr->fetch();     // ❌ CRASH! Cat doesn't have fetch()

// Compiler prevents this by only allowing base class methods
// You must explicitly cast to show you know the actual type
```

**Key Rules:**
1. **Base pointer/reference can only call methods declared in base class**
2. **Virtual methods**: Calls derived version if overridden (polymorphism)
3. **Non-virtual methods**: Calls base version (no polymorphism)
4. **Child-only methods**: Cannot be called through base pointer/reference
5. **Solution**: Cast to derived type using `dynamic_cast` or `static_cast`
6. **Same rules apply**: Normal inheritance, abstract class, interface - all follow same rules

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
**Concept**: Class with at least one pure virtual function - cannot be instantiated (cannot create objects)

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

// ❌ Cannot create object of abstract class in C++
// Shape s;  // Error: cannot declare variable 's' to be of abstract type 'Shape'
// Shape* ptr = new Shape();  // Error: cannot instantiate abstract class

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
**Concept**: All methods are pure virtual - defines contract. Cannot be instantiated (cannot create objects)

```cpp
// Interface - all pure virtual functions
class Drawable {
public:
    virtual void draw() = 0;
    virtual void resize(int factor) = 0;
    virtual ~Drawable() {}
};

// ❌ Cannot create object of interface in C++
// Drawable d;  // Error: cannot instantiate abstract class
// Drawable* ptr = new Drawable();  // Error: cannot instantiate abstract class

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

// ✅ Can create object of implementing class
Circle c;  // OK
```

**Key Points:**
- **Abstract class**: Cannot instantiate (cannot create objects) in C++, defines interface
- **Interface**: Cannot instantiate (cannot create objects) in C++, all methods pure virtual
- **Pure virtual function**: Must be implemented by derived class
- **Cannot create objects**: Both abstract classes and interfaces cannot be instantiated directly
- **Can use pointers/references**: Can have pointers/references of abstract class/interface type
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

**Example: Cannot Access Non-Static Members**
```cpp
class MyClass {
private:
    static int staticVar;  // Static member
    int instanceVar;       // Non-static member
    
public:
    MyClass(int val) : instanceVar(val) {}
    
    // Static function
    static void staticFunc() {
        staticVar = 10;     // ✅ OK: Can access static member
        // instanceVar = 20; // ❌ Error: Cannot access non-static member
        // No object exists, so no instanceVar to access
    }
    
    // Non-static function
    void instanceFunc() {
        staticVar = 10;     // ✅ OK: Can access static member
        instanceVar = 20;   // ✅ OK: Can access non-static member (has this pointer)
    }
};

int MyClass::staticVar = 0;

// Usage
MyClass::staticFunc();  // ✅ Can call without object
// MyClass::instanceFunc();  // ❌ Error: Need object to call

MyClass obj(5);
obj.instanceFunc();  // ✅ OK: Called with object
```

**Example: No `this` Pointer**
```cpp
class Example {
private:
    int value;
    static int count;
    
public:
    Example(int v) : value(v) {}
    
    // Non-static function - has this pointer
    void nonStaticFunc() {
        this->value = 10;        // ✅ OK: this pointer exists
        value = 20;              // ✅ OK: Implicit this->value
        cout << this->value;     // ✅ OK: Can use this
    }
    
    // Static function - no this pointer
    static void staticFunc() {
        count = 5;               // ✅ OK: Can access static member
        // this->value = 10;      // ❌ Error: No this pointer in static function
        // value = 20;            // ❌ Error: No this pointer, cannot access instance member
        // cout << this;           // ❌ Error: this doesn't exist in static function
    }
};

int Example::count = 0;

// Usage
Example::staticFunc();  // Called without object - no this pointer
Example obj(5);
obj.nonStaticFunc();     // Called with object - has this pointer
```

**Why No `this` Pointer?**
- `this` pointer points to the current object instance
- Static functions don't operate on objects (called without object)
- No object = no `this` pointer
- Therefore, cannot access instance members (need object to access them)

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

**1. Friendship is Not Inherited - Example:**
```cpp
class MyClass {
private:
    int secret;
    
    friend class Base;  // Base is friend
public:
    MyClass(int val) : secret(val) {}
};

class Base {
public:
    void accessSecret(MyClass& obj) {
        obj.secret = 100;  // ✅ OK: Base is friend
    }
};

class Derived : public Base {
public:
    void tryAccessSecret(MyClass& obj) {
        // obj.secret = 200;  // ❌ Error: Derived is NOT friend
        // Even though Derived inherits from Base, friendship is NOT inherited
    }
};

// Usage
MyClass obj(5);
Base base;
base.accessSecret(obj);  // ✅ OK: Base is friend

Derived derived;
// derived.tryAccessSecret(obj);  // ❌ Error: Derived cannot access
```

**What This Means:**
- If `Base` is friend of `MyClass`, `Derived` (child of `Base`) is **NOT** automatically friend
- Each class must be explicitly declared as friend
- Friendship is one-way: A friend of B doesn't make B friend of A
- Friendship doesn't propagate through inheritance hierarchy

**2. Use Sparingly (Breaks Encapsulation) - Example:**
```cpp
// ❌ BAD: Too many friends breaks encapsulation
class BankAccount {
private:
    double balance;
    string accountNumber;
    
    friend class Logger;           // Friend 1
    friend class TransactionLog;  // Friend 2
    friend class AuditSystem;     // Friend 3
    friend class ReportGenerator; // Friend 4
    friend void debugAccount();   // Friend 5
    // Too many friends = no encapsulation!
};

// ✅ GOOD: Controlled access through public interface
class BankAccount {
private:
    double balance;  // Encapsulated - protected
    string accountNumber;
    
public:
    // Controlled access through public methods
    void deposit(double amount) {
        if (amount > 0) balance += amount;
    }
    
    double getBalance() const {
        return balance;  // Read-only access
    }
    
    // Only one friend for specific use case
    friend class BankAuditor;  // Limited, justified friend
};
```

**Why Friend Breaks Encapsulation:**
- **Encapsulation**: Hide internal details, expose controlled interface
- **Friend**: Bypasses access control, gives direct access to private members
- **Problem**: Too many friends = no privacy, hard to maintain, violates OOP principles

**When to Use Friend (Acceptable Cases):**
1. **Operator Overloading**: `operator<<` for streams
   ```cpp
   class MyClass {
       int value;
       friend ostream& operator<<(ostream& os, const MyClass& obj);
   };
   ```
2. **Singleton Pattern**: Controlled object creation
3. **Performance**: When direct access is critical (rare)
4. **Legacy Code**: Interfacing with C code

**Best Practice:**
- Minimize use of friend
- Prefer public interface (getters/setters)
- Use friend only when absolutely necessary
- Document why friend is needed

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

**Two Uses of `final`:**
1. **Function-level `final`**: Prevents a virtual function from being overridden in derived classes
2. **Class-level `final`**: Prevents a class from being inherited

```cpp
// 1. Function-level final: Prevents overriding
class Base {
public:
    virtual void func() final {}  // Cannot be overridden
    virtual void func2() {}        // Can be overridden
};

class Derived : public Base {
public:
    // void func() { }  // ❌ Error: cannot override final function
    void func2() { }    // ✅ OK: func2 is not final
};

// 2. Class-level final: Prevents inheritance
class FinalClass final {  // Cannot be inherited
    // ...
};

// class Child : public FinalClass { }  // ❌ Error: cannot inherit from final class
```

**Use Cases and When to Use:**

**1. Function-level `final` - Preventing Override:**
```cpp
// Use Case: Critical function that must not be changed
class PaymentProcessor {
public:
    virtual void processPayment(double amount) final {
        // Critical payment logic - must not be overridden
        validateAmount(amount);
        deductFromAccount(amount);
        logTransaction(amount);
    }
    
    virtual void customizeUI() {
        // Can be overridden by derived classes
    }
};

class CreditCardProcessor : public PaymentProcessor {
public:
    // void processPayment(double amount) { }  // ❌ Error: final function
    void customizeUI() { }  // ✅ OK: can override
};
```

**Why Use Function-level `final`:**
- **Security**: Critical functions (payment, authentication) should not be modified
- **Performance**: Compiler can optimize (devirtualization) - knows function won't be overridden
- **Design Intent**: Clearly marks "this function is complete, don't change it"
- **Prevents Bugs**: Stops accidental overrides that could break functionality

**2. Class-level `final` - Preventing Inheritance:**
```cpp
// Use Case 1: Utility classes that shouldn't be extended
class MathUtils final {
public:
    static double pi() { return 3.14159; }
    static int square(int x) { return x * x; }
};

// class MyMathUtils : public MathUtils { }  // ❌ Error: utility class shouldn't be inherited

// Use Case 2: Leaf classes in inheritance hierarchy
class Animal {
    virtual void makeSound() = 0;
};

class Dog : public Animal {
    void makeSound() override { cout << "Woof!"; }
};

class GoldenRetriever final : public Dog {
    // This is a specific breed - no need to inherit further
    void makeSound() override { cout << "Woof! (Golden Retriever)"; }
};

// class PuppyGoldenRetriever : public GoldenRetriever { }  // ❌ Error: final class

// Use Case 3: Immutable classes
class ImmutableString final {
private:
    string data;
public:
    ImmutableString(const string& s) : data(s) {}
    string get() const { return data; }
    // No setters - immutable
};

// Use Case 4: Performance-critical classes
class Vector3D final {
    // Optimized for performance - no virtual functions
    // Marked final to prevent inheritance overhead
    double x, y, z;
public:
    Vector3D(double x, double y, double z) : x(x), y(y), z(z) {}
};
```

**Why Use Class-level `final`:**
- **Performance**: No virtual function table overhead (if no virtual functions)
- **Design Intent**: "This class is complete, don't extend it"
- **Prevents Mistakes**: Stops inappropriate inheritance (e.g., inheriting from utility classes)
- **Optimization**: Compiler can optimize better knowing class won't be inherited
- **Security**: Prevents malicious or accidental subclassing

**Important Notes:**
- `final` is a keyword, not a specifier (unlike `virtual`, `static`)
- Can combine with `override`: `void func() override final { }`
- `final` on function requires `virtual` (or it's meaningless)
- `final` on class doesn't require virtual functions
- Compiler error if you try to override/inherit from `final`

**Comparison: `final` vs `private` for Preventing Inheritance:**
```cpp
// Method 1: Using final (C++11) - Clear and explicit
class A final { };

// Method 2: Using private constructor (old way)
class B {
private:
    B() { }  // Private constructor prevents inheritance
    friend class BFactory;  // Only factory can create
};

// final is preferred: clearer intent, better error messages
```

**Best Practices:**
- Use `final` on functions that are critical and shouldn't be changed
- Use `final` on leaf classes (end of inheritance chain)
- Use `final` on utility/helper classes
- Use `final` for performance optimization when inheritance is not needed
- Don't overuse - only when you have a clear reason

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
```

**Understanding `= delete` Syntax:**

**What is `= delete`?**
- `= delete` is a C++11 feature that explicitly **deletes** (disables) a function
- Makes the function **unavailable** - compiler will error if someone tries to use it
- Better than making it `private` - gives clearer error messages

**Syntax Breakdown:**
```cpp
// 1. Copy Constructor = delete
Database(const Database&) = delete;
//     ^^^^^^^^^^^^^^^^^^   ^^^^^^^
//     Function signature   Delete it
//     Takes reference to   Makes it unavailable
//     another Database

// 2. Copy Assignment Operator = delete
Database& operator=(const Database&) = delete;
// ^^^^^^^^ ^^^^^^^ ^^^^^^^^^^^^^^^^^   ^^^^^^^
// Return  Operator Takes reference     Delete it
// type    name     to another Database
```

**Why Prevent Copying in Singleton?**
```cpp
// Without = delete, this would be possible (BAD):
Database* db1 = Database::getInstance();
Database* db2 = Database::getInstance();

// But someone could accidentally do this:
Database db3 = *db1;  // ❌ Copy constructor called - creates NEW instance!
// Now we have 2 Database objects - breaks Singleton pattern!

// With = delete, compiler prevents this:
Database db3 = *db1;  // ❌ Compiler error: use of deleted function
```

**How `= delete` Works:**
```cpp
class MyClass {
public:
    MyClass() { }
    
    // Delete copy constructor
    MyClass(const MyClass&) = delete;
    
    // Delete copy assignment
    MyClass& operator=(const MyClass&) = delete;
};

// Usage
MyClass obj1;
// MyClass obj2 = obj1;        // ❌ Error: use of deleted function 'MyClass::MyClass(const MyClass&)'
// MyClass obj3;
// obj3 = obj1;                 // ❌ Error: use of deleted function 'MyClass& MyClass::operator=(const MyClass&)'
```

**Alternative: Private (Old Way)**
```cpp
// OLD WAY (C++98/03): Make private
class MyClass {
private:
    MyClass(const MyClass&);           // Declare but don't define
    MyClass& operator=(const MyClass&); // Declare but don't define
public:
    MyClass() { }
};

// Problem: Error message is unclear
// "MyClass::MyClass(const MyClass&) is private"
// vs
// "use of deleted function" (clearer with = delete)
```

**What Can Be Deleted?**
```cpp
class Example {
public:
    // Delete any function
    void func(int x) = delete;           // Delete specific overload
    void func(double x) = delete;         // Delete another overload
    void func() { }                       // This one is OK
    
    // Delete special member functions
    Example() = default;                  // Use default constructor
    Example(const Example&) = delete;     // Delete copy constructor
    Example(Example&&) = delete;          // Delete move constructor
    Example& operator=(const Example&) = delete;  // Delete copy assignment
    Example& operator=(Example&&) = delete;       // Delete move assignment
    
    // Delete conversion operators
    operator int() = delete;              // Prevent conversion to int
};

// Usage
Example e;
// e.func(5);        // ❌ Error: func(int) is deleted
// e.func(3.14);     // ❌ Error: func(double) is deleted
e.func();            // ✅ OK: func() is available
// int x = e;        // ❌ Error: conversion to int is deleted
```

**Common Use Cases:**
1. **Singleton Pattern**: Prevent copying (only one instance)
2. **Resource Management**: Prevent copying of unique resources (file handles, network connections)
3. **Disable Specific Overloads**: Prevent certain function calls
4. **Move-only Types**: Delete copy, allow move (like `std::unique_ptr`)

**Key Points:**
- `= delete` must be in public section (for clearer error messages)
- Can delete any function, not just special member functions
- Compiler error is clearer than private access error
- Better than leaving function undefined (linker error vs compile error)

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


