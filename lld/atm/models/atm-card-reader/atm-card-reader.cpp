#include "atm-card-reader.h"
#include <iostream>

bool AtmCardReader::ValidateFromBank(AtmCard* atm) {
    if (atm->GetCardNumber() == "1234567890123456") {
        std::cout << "Validating card with the bank...\n";
        std::cout << "Card " << atm->GetCardNumber() << " is valid.\n";
        return true;
    } 
    std::cout << "Validating card with the bank...\n";
    std::cout << "Card " << atm->GetCardNumber() << " is not valid.\n";
    return false;

}
