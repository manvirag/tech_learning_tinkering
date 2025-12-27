/* 
  
  In C++, the concept of friend functions and friend classes allows access to private and protected members of a class from outside that class. This enables flexibility in designing classes and can be useful for specific scenarios where controlled access to internal members is required. Let's explore friend functions and friend classes with code examples:

Friend Function in C++:
A friend function is a function that is not a member of the class but is allowed to access the private and protected members of the class. Here's how you declare and use friend functions:

Example 1: Working of a friend function

*/


#include <iostream>
using namespace std;

class Distance {
private:
    int meter;
    
    // Declare friend function
    friend int addFive(Distance);

public:
    Distance() : meter(0) {}
};

// Define friend function
int addFive(Distance d) {
    d.meter += 5;
    return d.meter;
}

int main() {
    Distance D;
    cout << "Distance: " << addFive(D);
    return 0;
}

/*
  
In this example, the friend function addFive can access the private member meter of the Distance class.

Friend Class in C++:
A friend class is a class that is granted access to the private and protected members of another class. Here's how you declare and use friend classes:

Example 2: Using a friend class

*/

#include <iostream>
using namespace std;

// Forward declaration
class ClassB;

class ClassA {
private:
    int numA;

    // Declare ClassB as a friend class
    friend class ClassB;

public:
    ClassA() : numA(12) {}
};

class ClassB {
private:
    int numB;

public:
    ClassB() : numB(1) {}

    int add(ClassA objectA) {
        return objectA.numA + numB;
    }
};

int main() {
    ClassA objectA;
    ClassB objectB;
    cout << "Sum: " << objectB.add(objectA);
    return 0;
}
In this example, ClassB is a friend class of ClassA. This allows ClassB to access the private member numA of ClassA.

Friend functions and classes should be used with caution as they can compromise encapsulation, which is an important aspect of object-oriented programming. They should only be used when necessary and justified.
