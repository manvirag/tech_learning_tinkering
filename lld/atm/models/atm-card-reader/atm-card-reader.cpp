#include "atm-card-reader.h"
#include <iostream>

void AtmCardReader::ValidateFromBank(AtmCard atm) {
    if (atm.GetCardNumber() == "1234 5678 9012 3456") {
        std::cout << "Validating card with the bank...\n";
        std::cout << "Card " << atm.GetCardNumber() << " is valid.\n";
    } else {
        std::cout << "Validating card with the bank...\n";
        std::cout << "Card " << atm.GetCardNumber() << " is not valid.\n";
    }
}
