// Before

package main

import "fmt"

type Bird interface {
	Fly()
}

type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow is flying")
}

type Ostrich struct{}

func (o Ostrich) Fly() {
	// Ostriches can't fly!
	panic("Ostrich can't fly!") // LSP violation
}

func MakeBirdFly(b Bird) {
	b.Fly()
}

func main() {
	s := Sparrow{}
	o := Ostrich{}

	MakeBirdFly(s) // OK
	MakeBirdFly(o) // Panic! LSP broken
}

// After 

package main

import "fmt"

type Bird interface {
	Walk()
}

type FlyingBird interface {
	Bird
	Fly()
}

type Sparrow struct{}

func (s Sparrow) Walk() {
	fmt.Println("Sparrow is walking")
}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow is flying")
}

type Ostrich struct{}

func (o Ostrich) Walk() {
	fmt.Println("Ostrich is walking")
}

func MakeBirdWalk(b Bird) {
	b.Walk()
}

func MakeBirdFly(fb FlyingBird) {
	fb.Fly()
}

func main() {
	s := Sparrow{}
	o := Ostrich{}

	MakeBirdWalk(s) // OK
	MakeBirdFly(s)  // OK

	MakeBirdWalk(o) // OK
	// MakeBirdFly(o) // Compile error! Ostrich does not implement FlyingBird
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
        printArea(rect);    // Works fine
        printArea(broken);  // Crashes
    } catch (const exception& ex) {
        cout << "Exception caught: " << ex.what() << endl;
    }

    return 0;
}




*/
