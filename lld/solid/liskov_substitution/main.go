// Example

package main

import "fmt"

// Liskov ...
type Liskov interface {
	LiskovFunc()
}

// Parent ...
type Parent struct{}

// Child ...   // this is not inheritance this is struct embedding -> https://www.tutorialspoint.com/composition-in-golang
// now child can access parent methods as well as field direct.
type Child struct {
	Parent
}

// LiskovFunc ...
func (p *Parent) LiskovFunc() {
	fmt.Println("It works")
}

// LiskovSubstitution ...
func LiskovSubstitution(lis Liskov) {
	lis.LiskovFunc()
}

func main() {
	ch := &Child{}
	par := &Parent{}
	LiskovSubstitution(ch)
	LiskovSubstitution(par)
}





/*

c++, in short child should have things what parent told, 

voilating one: 




#include <iostream>
#include <stdexcept>
using namespace std;

// Superclass (base class)
class Shape {
public:
    virtual double area() const = 0; // Pure virtual function
    virtual ~Shape() = default;
};

// Subclass 1: Rectangle — Good
class Rectangle : public Shape {
    double width, height;
public:
    Rectangle(double w, double h) : width(w), height(h) {}
    double area() const override {
        return width * height;
    }
};

// Subclass 2: BrokenShape — BAD
class BrokenShape : public Shape {
public:
    double area() const override {
        throw runtime_error("Area calculation not supported!");
    }
};

// Function using Shape
void printArea(const Shape& shape) {
    cout << "Area: " << shape.area() << endl;
}

int main() {
    Rectangle rect(5, 10);
    BrokenShape broken;

    try {
        printArea(rect);    // Works fine ✅
        printArea(broken);  // Crashes ❌
    } catch (const exception& ex) {
        cout << "Exception caught: " << ex.what() << endl;
    }

    return 0;
}




*/
