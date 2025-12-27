/*
 * Pure Virtual Functions and Abstract Classes in C++
 * 
 * Sometimes implementation of all functions cannot be provided in a base class
 * because we don't know the implementation. Such a class is called abstract class.
 * 
 * For example, let Shape be a base class. We cannot provide implementation of
 * function draw() in Shape, but we know every derived class must have
 * implementation of draw(). Similarly an Animal class doesn't have implementation
 * of move() (assuming that all animals move), but all animals must know how to move.
 * 
 * We cannot create objects of abstract classes.
 * 
 * A pure virtual function (or abstract function) in C++ is a virtual function
 * for which we don't have implementation, we only declare it. A pure virtual
 * function is declared by assigning 0 in declaration.
 */

#include <iostream>
using namespace std;

// ============================================================================
// Example 1: Basic Abstract Class
// ============================================================================

// An abstract class with pure virtual function
class Test {
    // Data members of class
public:
    // Pure Virtual Function - must be implemented by derived classes
    virtual void show() = 0;
    
    // Other members can be present
    int getValue() { return 10; }
};

// ============================================================================
// Example 2: Complete Example - Base and Derived
// ============================================================================

class Base {
    int x;
public:
    // Pure virtual function - no implementation
    virtual void fun() = 0;
    
    // Regular function - can have implementation
    int getX() { return x; }
};

// This class inherits from Base and implements fun()
class Derived : public Base {
    int y;
public:
    void fun() override {
        cout << "fun() called" << endl;
    }
};

void example2() {
    Derived d;
    d.fun();  // Output: fun() called
}

// ============================================================================
// Example 3: Cannot Create Object of Abstract Class
// ============================================================================

class AbstractTest {
    int x;
public:
    virtual void show() = 0;  // Pure virtual function
    int getX() { return x; }
};

void example3() {
    // AbstractTest t;  // ❌ Compiler Error: cannot declare variable 't' to be
                        //    of abstract type 'AbstractTest' because the following
                        //    virtual functions are pure within 'AbstractTest':
                        //    virtual void AbstractTest::show()
}

// ============================================================================
// Example 4: Pointers and References of Abstract Class Type
// ============================================================================

class Base2 {
public:
    virtual void show() = 0;
};

class Derived2 : public Base2 {
public:
    void show() override {
        cout << "In Derived" << endl;
    }
};

void example4() {
    // ✅ We can have pointers and references of abstract class type
    Base2* bp = new Derived2();
    bp->show();  // Output: In Derived
    delete bp;
}

// ============================================================================
// Example 5: Derived Class Also Becomes Abstract if Not Overridden
// ============================================================================

class Base3 {
public:
    virtual void show() = 0;
};

// Derived class does NOT override pure virtual function
class Derived3 : public Base3 {
    // No implementation of show()
};

void example5() {
    // Derived3 d;  // ❌ Compiler Error: cannot declare variable 'd' to be
                     //    of abstract type 'Derived3' because the following
                     //    virtual functions are pure within 'Derived3':
                     //    virtual void Base3::show()
    
    // If we do not override the pure virtual function in derived class,
    // then derived class also becomes abstract class.
}

// ============================================================================
// Example 6: Abstract Class Can Have Constructors
// ============================================================================

// An abstract class with constructor
class Base4 {
protected:
    int x;
public:
    virtual void fun() = 0;  // Pure virtual function
    
    // ✅ Abstract class can have constructors
    Base4(int i) {
        x = i;
    }
};

class Derived4 : public Base4 {
    int y;
public:
    // Call base class constructor in initialization list
    Derived4(int i, int j) : Base4(i) {
        y = j;
    }
    
    void fun() override {
        cout << "x = " << x << ", y = " << y << endl;
    }
};

void example6() {
    Derived4 d(4, 5);
    d.fun();  // Output: x = 4, y = 5
}

// ============================================================================
// Main Function - Demonstrating All Examples
// ============================================================================

int main() {
    cout << "=== Example 2: Complete Example ===" << endl;
    example2();
    
    cout << "\n=== Example 4: Pointers and References ===" << endl;
    example4();
    
    cout << "\n=== Example 6: Abstract Class with Constructor ===" << endl;
    example6();
    
    return 0;
}

/*
 * Key Points:
 * 
 * 1. A class is abstract if it has at least one pure virtual function.
 * 
 * 2. We cannot create objects of abstract classes directly.
 * 
 * 3. We CAN have pointers and references of abstract class type.
 * 
 * 4. If derived class does not override pure virtual function,
 *    derived class also becomes abstract.
 * 
 * 5. Abstract classes can have constructors.
 * 
 * 6. Abstract classes can have regular (non-virtual) member functions.
 */
