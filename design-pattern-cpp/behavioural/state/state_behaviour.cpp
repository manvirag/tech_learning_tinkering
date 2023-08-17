#include <iostream>

class Fan;  // Forward declaration for Fan class

class FanState {
public:
    virtual void pullChain(Fan* fan) = 0;
};

class FanOffState : public FanState {
public:
    void pullChain(Fan* fan) override;
};

class FanLowState : public FanState {
public:
    void pullChain(Fan* fan) override;
};

class FanHighState : public FanState {
public:
    void pullChain(Fan* fan) override;
};

class Fan {
private:
    FanState* currentState;

public:
    Fan() {
        currentState = new FanOffState();
    }

    void setState(FanState* state) {
        delete currentState;  // Clean up the previous state
        currentState = state;
    }

    void pullChain() {
        currentState->pullChain(this);
    }

    ~Fan() {
        delete currentState;
    }
};

void FanOffState::pullChain(Fan* fan) {
    std::cout << "Turning fan on to low." << std::endl;
    fan->setState(new FanLowState());
}

void FanLowState::pullChain(Fan* fan) {
    std::cout << "Turning fan on to high." << std::endl;
    fan->setState(new FanHighState());
}

void FanHighState::pullChain(Fan* fan) {
    std::cout << "Turning fan off." << std::endl;
    fan->setState(new FanOffState());
}

int main() {
    Fan* fan = new Fan();
    fan->pullChain(); // Turning fan on to low.
    fan->pullChain(); // Turning fan on to high.
    fan->pullChain(); // Turning fan off.

    delete fan;

    return 0;
}
