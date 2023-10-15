#include <iostream>

class Printer {
public:
    void PrintBalanceInquiry(std::string cardNumber, double balance);
    void PrintWithdrawal(std::string cardNumber, double amount, double newBalance);
};
