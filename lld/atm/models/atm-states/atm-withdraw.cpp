#include"atm-withdraw.h"
#include"../atm-machine/atm-machine.h"
#include"atm-states.h"
#include"atm-stale.h"
#include"../atm-card/atm-card.h"
#include<iostream>

void AtmWithDraw :: processRequest(AtmMachine* atmMachine){
    std::cout<<"Enter Money for withdrawl: "<<std::endl;
    double amount;
    std::cin>>amount;
    // balance = 5045 ,assume fetched from banked api
    if(amount > 5045){
        std::cout<<"Your withdrawl amount exceeding balance: "<<std::endl;
        exit(0);
    }

    this->atmPrinter->PrintWithdrawal(atmMachine->GetAtmCard()->GetCardNumber(), amount, (5045-amount));

    exit(0);

    

}