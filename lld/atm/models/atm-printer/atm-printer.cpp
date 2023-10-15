#include "atm-printer.h"
#include<iostream>
#include<string>

void Printer::PrintBalanceInquiry( std::string cardNumber, double balance)  {
    std::cout << "--------------------------------------------"<< std::endl;
    std::cout<< "Balance Inquiry for Card " << cardNumber << std::endl;
    std::cout << "Current Balance: $" << balance << std::endl;
    std::cout << "--------------------------------------------"<<std::endl;
    std::cout << "THANKS FOR VISITING ATM." << std::endl << "आपका दिन शुभ हो! धन्यवाद !" <<  std::endl;
    
}

void Printer::PrintWithdrawal( std::string cardNumber, double amount, double newBalance)  {
    std::cout << "--------------------------------------------"<<std::endl;
    std::cout << "Withdrawal for Card " << cardNumber << std::endl;
    std::cout << "Withdrawn Amount: $" << amount << std::endl;
    std::cout << "New Balance: $" << newBalance << std::endl;
    std::cout << "--------------------------------------------" << std::endl;
    std::cout << "THANKS FOR VISITING ATM." << std::endl << "आपका दिन शुभ हो! धन्यवाद !" <<  std::endl;
}
