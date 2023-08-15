/*
 In heritance when child have same functionality of override method, then we use this pattern. Because it will be inefficient if we have lots of child with same override functionality.
*/

#include <iostream>

class DriveStrategy {
public:
    virtual void drive() = 0;
};

class NormalDriverStrategy : public DriveStrategy {
public:
    void drive() override {
        std::cout << "Normal drive strategy" << std::endl;
    }
};

class SportDriveStrategy : public DriveStrategy {
public:
    void drive() override {
        std::cout << "Sports drive strategy" << std::endl;
    }
};

class Vehicle {
private:
    DriveStrategy* driveStrategy;

public:
    Vehicle(DriveStrategy* strategy) : driveStrategy(strategy) {}

    void drive() {
        driveStrategy->drive();
    }
};

class SportsVehicle : public Vehicle {
public:
    SportsVehicle() : Vehicle(new SportDriveStrategy()) {}
};

class OffRoadVehicle : public Vehicle {
public:
    OffRoadVehicle() : Vehicle(new NormalDriverStrategy()) {}
};

class GoodsVehicle : public Vehicle {
public:
    GoodsVehicle() : Vehicle(new NormalDriverStrategy()) {}
};

int main() {
    SportsVehicle sportsCar;
    OffRoadVehicle offRoadVehicle;
    GoodsVehicle goodsTruck;

    sportsCar.drive();      // Output: Sports drive strategy
    offRoadVehicle.drive(); // Output: Normal drive strategy
    goodsTruck.drive();     // Output: Normal drive strategy

    return 0;
}
