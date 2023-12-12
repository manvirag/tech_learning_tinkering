#include"atm-balance-enquiry.h"
#include"../atm-machine/atm-machine.h"
#include"../atm-card/atm-card.h"
#include"atm-stale.h"
#include"atm-states.h"
#include<iostream>

void AtmBalanceEnquiry :: processRequest(AtmMachine* atmMachine){
    // Assume there is banking api/service with which we are fetch data and we got balance = 5000
    this->atmPrinter->PrintBalanceInquiry(atmMachine->GetAtmCard()->GetCardNumber(), 5045.45);
    exit(0);
}