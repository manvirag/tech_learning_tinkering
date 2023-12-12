/*
 check again and again

 In simple term: we define some method in abstract class which are common and those which are not common they will have to be made abstract
*/


#include <iostream>

class Vehicle {
public:
    void start() {
        checkFuel();
        igniteEngine();
        startMoving();
    }

protected:
    void checkFuel() {
        std::cout << "Checking fuel..." << std::endl;
    }

    virtual void igniteEngine() = 0;

    void startMoving() {
        std::cout << "Starting to move..." << std::endl;
    }
};

// Car class
class Car : public Vehicle {
protected:
    void igniteEngine() override {
        std::cout << "Igniting engine..." << std::endl;
    }
};

// Bike class
class Bike : public Vehicle {
protected:
    void igniteEngine() override {
        std::cout << "Kicking start..." << std::endl;
    }
};

// Bus class
class Bus : public Vehicle {
protected:
    void igniteEngine() override {
        std::cout << "Turning on ignition..." << std::endl;
    }
};

// Main function
int main() {
    Vehicle* car = new Car();
    Vehicle* bike = new Bike();
    Vehicle* bus = new Bus();
    
    car->start();
    bike->start();
    bus->start();
    
    delete car;
    delete bike;
    delete bus;
    
    return 0;
}

