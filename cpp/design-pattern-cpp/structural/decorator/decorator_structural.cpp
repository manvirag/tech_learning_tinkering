/*
 have 3 functionally , 
 we have combination of this function
 will be hard to create class
 instead 
 create wrapper

 which take same type of class decorate it and return same type of class

*/


/*
check below example
have two type of pizzas
and have two types of toppings
now can create pizza of two types and with combination of toppings

*/

#include <iostream>
#include <string>
using namespace std;

// Base class for Pizza
class Pizza {
public:
    virtual string getDescription() const = 0;
    virtual double cost() const = 0;
};

// Concrete pizza classes
class MargheritaPizza : public Pizza {
public:
    string getDescription() const override {
        return "Margherita Pizza";
    }

    double cost() const override {
        return 8.99;
    }
};

class PepperoniPizza : public Pizza {
public:
    string getDescription() const override {
        return "Pepperoni Pizza";
    }

    double cost() const override {
        return 9.99;
    }
};

// Decorator base class
class ToppingDecorator : public Pizza {
protected:
    Pizza* pizza;

public:
    ToppingDecorator(Pizza* pizza) : pizza(pizza) {}

    string getDescription() const override {
        return pizza->getDescription();
    }

    double cost() const override {
        return pizza->cost();
    }
};

// Concrete decorator classes
class CheeseTopping : public ToppingDecorator {
public:
    CheeseTopping(Pizza* pizza) : ToppingDecorator(pizza) {}

    string getDescription() const override {
        return pizza->getDescription() + ", Cheese";
    }

    double cost() const override {
        return pizza->cost() + 2.0;
    }
};

class MushroomTopping : public ToppingDecorator {
public:
    MushroomTopping(Pizza* pizza) : ToppingDecorator(pizza) {}

    string getDescription() const override {
        return pizza->getDescription() + ", Mushrooms";
    }

    double cost() const override {
        return pizza->cost() + 1.5;
    }
};

int main() {
    Pizza* margherita = new MargheritaPizza();
    margherita = new CheeseTopping(margherita);
    margherita = new MushroomTopping(margherita);

    cout << "Description: " << margherita->getDescription() << endl;
    cout << "Cost: $" << margherita->cost() << endl;

    delete margherita;

    return 0;
}
