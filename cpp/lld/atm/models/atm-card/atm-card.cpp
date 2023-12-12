#include<string>
#include "atm-card.h"


// Constructor definition
AtmCard::AtmCard(std::string userName, std::string cardNumber, std::string bank)
    : userName(userName), cardNumber(cardNumber), bank(bank) {
}

std::string AtmCard::GetCardNumber() {
    return cardNumber;
}

std::string AtmCard::GetCardBank() {
    return bank;
}

std::string AtmCard::GetCardHolderName() {
    return userName;
}
