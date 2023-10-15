#include"atm-showscreen.h"
#include"atm-states.h"
#include"../atm-machine/atm-machine.h"
#include"../atm-card/atm-card.h"
#include"atm-balance-enquiry.h"
#include"atm-withdraw.h"
#include<iostream>

void AtmShowScreen :: processRequest(AtmMachine* atmMachine){
    std::vector<std::string> options(3);
    options[0] = "BALANCE_ENQUIRY";
    options[1] = "WITHDRAWL";
    options[2] = "CANCEL";
    std::string selectedOption = this->atmScreen->GetTheUserSelectedOption(options);

    if(selectedOption == options[2]){
        exit(0);
    }else if(selectedOption == options[0]){
        atmMachine->SetAtmState(new AtmBalanceEnquiry());
    }else if(selectedOption == options[1]){
        atmMachine->SetAtmState(new AtmWithDraw());
    }
    else{
        std::cout<<"Wrong Option entered"<<std::endl;
        exit(0);
    }

}