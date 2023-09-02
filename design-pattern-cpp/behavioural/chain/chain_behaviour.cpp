/*
 check state design pattern commend

 e.x request  -> ordering system  
 can have middle handler
 
 The request must pass a series of checks before the ordering system itself can handle it.

 sequential -> its like the series of knock games if fail at initial stage no need to check further

 This pattern should be used when there is a clear order of escalation or processing.

 
 below e.g.

 In this example, we have a PaymentProcessor class that processes payments based on the amount. If the amount is less than $10, it uses Paypal, if it is less than $100, it uses a Credit Card, and for amounts over $100, it uses Bank Transfer. The Payment class encapsulates the payment amount.

The problem with this approach is that it is not flexible. If we want to add or remove payment options, we need to modify the PaymentProcessor class, which can lead to tight coupling and code that is difficult to maintain.



*/



#include <iostream>


class Payment {
private:
    const double amount;

public:
    Payment(double amount) : amount(amount) {}

    double getAmount() const {
        return amount;
    }
};


class PaymentHandler {
protected:
    PaymentHandler* successor;

public:
    void setSuccessor(PaymentHandler* successor) {
        this->successor = successor;
    }

    virtual void handlePayment(const Payment& payment) = 0;
};

class PaypalHandler : public PaymentHandler {
public:
    void handlePayment(const Payment& payment) override {
        if (payment.getAmount() < 10) {
            std::cout << "Processing payment of $" << payment.getAmount() << " with Paypal." << std::endl;
        } else {
            if (successor != nullptr) {
                successor->handlePayment(payment);
            }
        }
    }
};

class CreditCardHandler : public PaymentHandler {
public:
    void handlePayment(const Payment& payment) override {
        if (payment.getAmount() < 100) {
            std::cout << "Processing payment of $" << payment.getAmount() << " with Credit Card." << std::endl;
        } else {
            if (successor != nullptr) {
                successor->handlePayment(payment);
            }
        }
    }
};

class BankTransferHandler : public PaymentHandler {
public:
    void handlePayment(const Payment& payment) override {
        if (payment.getAmount() >= 100) {
            std::cout << "Processing payment of $" << payment.getAmount() << " with Bank Transfer." << std::endl;
        } else {
            std::cout << "Invalid payment amount: $" << payment.getAmount() << std::endl;
        }
    }
};


int main() {
    PaypalHandler paypalHandler;
    CreditCardHandler creditCardHandler;
    BankTransferHandler bankTransferHandler;

    paypalHandler.setSuccessor(&creditCardHandler);
    creditCardHandler.setSuccessor(&bankTransferHandler);

    Payment payment1(5);
    paypalHandler.handlePayment(payment1);

    Payment payment2(50);
    paypalHandler.handlePayment(payment2);

    Payment payment3(500);
    paypalHandler.handlePayment(payment3);

    return 0;
}



