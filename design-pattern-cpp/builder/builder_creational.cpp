/*

I understand your confusion. The code you've provided is using the this pointer to access the member variables of the class within the member function of the class. Let's break down the usage of this and the -> arrow operator:

this Pointer:
In C++, this is a reserved keyword that refers to the current instance of the class. It's a pointer that holds the memory address of the object on which the member function is being called. It is used to disambiguate between member variables of the class and function parameters that might have the same name.

Arrow Operator (->):
The arrow operator is used to access members of a class or structure through a pointer. When you have a pointer to an object, you use the arrow operator to access members of the object pointed to by the pointer. It's a shorthand way to dereference the pointer and access the members.


return *this; returns a reference to the current instance of the Builder class. The *this syntax is used to dereference the this pointer and return the actual object that the pointer is pointing to.

if we are return original object why using Builder& instead of Builder

ChatGPT
You raise a good point. The use of Builder& versus Builder in the return type is a design decision that affects how the builder methods can be chained together. Let's discuss the differences between these two approaches:

Returning Builder&:
Returning a reference to the current instance of the Builder class (Builder&) allows you to chain multiple setter calls together in a fluent interface. Each setter method modifies the same object and returns the same object by reference, enabling you to call the next method immediately. This is often referred to as method chaining.
Example:

Builder builder;
builder.setName("Product 1").setDescription("Description").setPrice(100);
Returning Builder:
Returning just the Builder type (without the reference) still allows you to chain methods, but it would create a new temporary object for each chained method call. This means that each method call would operate on a separate copy of the builder, and you wouldn't be able to set multiple attributes of the same object in a single chain.
Example (with non-reference return):


Builder builder;
builder = builder.setName("Product 1"); // This creates a new object, doesn't modify original builder
builder = builder.setDescription("Description"); // This operates on a different copy
builder = builder.setPrice(100); // This operates on yet another copy
Returning a reference (Builder&) is generally more efficient because it doesn't involve creating temporary copies of the object for each method call. Additionally, it closely matches the intention of method chaining, where you're progressively modifying a single object.

So, in summary, returning Builder& allows for a more intuitive and efficient method chaining experience where you can modify multiple attributes of the same object in a single chain.
*/

#include <iostream>
#include <string>
using namespace std;

class Product {
private:
    string name;
    string description;
    int price;

public:
    //The syntax you're seeing is a C++ feature known as a member initializer list. 
    Product(string name, string description, int price)
        : name(name), description(description), price(price) {}

    string getName() {
        return name;
    }

    string getDescription() {
        return description;
    }

    int getPrice() {
        return price;
    }


    // nested class check out oops repo
    class Builder {
    private:
        string name;
        string description;
        int price;

    public:
        Builder& setName(string name) {
            this->name = name;
            return *this;
        }

        Builder& setDescription(string description) {
            this->description = description;
            return *this;
        }

        Builder& setPrice(int price) {
            this->price = price;
            return *this;
        }

        Product build() {
            return Product(name, description, price);
        }
    };
};

int main() {
    Product::Builder builder;
    Product product = builder.setName("Sample Product")
                           .setDescription("This is a sample product")
                           .setPrice(100)
                           .build();

    cout << "Name: " << product.getName() << endl;
    cout << "Description: " << product.getDescription() << endl;
    cout << "Price: " << product.getPrice() << endl;

    return 0;
}
