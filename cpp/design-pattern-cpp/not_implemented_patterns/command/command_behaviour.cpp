/*
forgettable, check again and again
*/

#include <iostream>
#include <vector>

// Receiver: The device classes that will perform the actual actions
class Light {
public:
    void turnOn() {
        std::cout << "Light is on" << std::endl;
    }

    void turnOff() {
        std::cout << "Light is off" << std::endl;
    }
};

class Fan {
public:
    void turnOn() {
        std::cout << "Fan is on" << std::endl;
    }

    void turnOff() {
        std::cout << "Fan is off" << std::endl;
    }
};

// Command: Abstract base class for commands
class Command {
public:
    virtual void execute() = 0;
};

// Concrete Commands
class TurnOnCommand : public Command {
private:
    Light& light;

public:
    TurnOnCommand(Light& l) : light(l) {}

    void execute() override {
        light.turnOn();
    }
};

class TurnOffCommand : public Command {
private:
    Light& light;

public:
    TurnOffCommand(Light& l) : light(l) {}

    void execute() override {
        light.turnOff();
    }
};

class FanOnCommand : public Command {
private:
    Fan& fan;

public:
    FanOnCommand(Fan& f) : fan(f) {}

    void execute() override {
        fan.turnOn();
    }
};

class FanOffCommand : public Command {
private:
    Fan& fan;

public:
    FanOffCommand(Fan& f) : fan(f) {}

    void execute() override {
        fan.turnOff();
    }
};

// Invoker: Controls the commands
class RemoteControl {
private:
    Command* onCommand;
    Command* offCommand;

public:
    RemoteControl(Command* on, Command* off) : onCommand(on), offCommand(off) {}

    void pressOnButton() {
        onCommand->execute();
    }

    void pressOffButton() {
        offCommand->execute();
    }
};

int main() {
    Light livingRoomLight;
    Fan ceilingFan;

    TurnOnCommand turnOnLivingRoomLight(livingRoomLight);
    TurnOffCommand turnOffLivingRoomLight(livingRoomLight);

    FanOnCommand fanOn(ceilingFan);
    FanOffCommand fanOff(ceilingFan);

    RemoteControl remote(&turnOnLivingRoomLight, &turnOffLivingRoomLight);

    remote.pressOnButton();  // Turns on the living room light
    remote.pressOffButton(); // Turns off the living room light

    remote = RemoteControl(&fanOn, &fanOff);
    remote.pressOnButton();  // Turns on the fan
    remote.pressOffButton(); // Turns off the fan

    return 0;
}
