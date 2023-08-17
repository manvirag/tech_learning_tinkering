/*
Mediator is a behavioral design pattern that lets you reduce chaotic dependencies between objects. The pattern restricts direct communications between the objects and forces them to collaborate only via a mediator object.

Pilots of aircraft that approach or depart the airport control area don’t communicate directly with each other.
*/

#include <iostream>
#include <vector>
#include <string>
using namespace std;

class Mediator; 

class Colleague {
protected:
    Mediator* mediator;

public:
    Colleague(Mediator* mediator) : mediator(mediator) {}

    virtual void send(const string& message) = 0;
    virtual void receive(const string& message) = 0;
};

class Mediator {
public:
    void addColleague(Colleague* colleague) {
        colleagues.push_back(colleague);
    }

    void send(const string& message, Colleague* sender) {
        for (Colleague* colleague : colleagues) {
            if (colleague != sender) {
                colleague->receive(message);
            }
        }
    }

private:
    vector<Colleague*> colleagues;
};



class ConcreteColleague : public Colleague {
public:
    ConcreteColleague(Mediator* mediator, const string& name) : Colleague(mediator), name(name) {}

    void send(const string& message) override {
        mediator->send(message, this);
    }

    void receive(const string& message) override {
        cout << name << " received: " << message << endl;
    }

private:
    string name;
};



int main() {
    Mediator* chatMediator = new Mediator();

    Colleague* user1 = new ConcreteColleague(chatMediator, "User 1");
    Colleague* user2 = new ConcreteColleague(chatMediator, "User 2");
    Colleague* user3 = new ConcreteColleague(chatMediator, "User 3");

    chatMediator->addColleague(user1);
    chatMediator->addColleague(user2);
    chatMediator->addColleague(user3);

    user1->send("Hello, everyone!");
    user2->send("Hi there!");
    user3->send("Greetings!");

    delete user1;
    delete user2;
    delete user3;
    delete chatMediator;

    return 0;
}
