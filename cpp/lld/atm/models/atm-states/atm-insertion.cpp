#include"atm-insertion.h"
#include"../atm-machine/atm-machine.h"
#include"../atm-card/atm-card.h"
#include"atm-states.h"
#include"atm-showscreen.h"
#include<iostream>
#include<string>

void AtmInsertion :: processRequest(AtmMachine* atmMachine){
    std::string no, name, bank;
    std::cin>>no>>name>>bank;
    AtmCard *atmCard = new AtmCard(name, no ,bank);
    std::cout<<"Thanks for inserting................Validating cards details."<<std::endl;
    
    if(this->atmCardReader->ValidateFromBank(atmCard)){
        std::cout<<"Enter your card pin. with key one by one all 4 letter"<<std::endl;
        std::string pin = ""; // can do more validation here also keeping simple
        for(int i=0;i<4;i++)
        {
            pin+=this->atmKeyPad->GetSelectedKey();
        }
        atmMachine->SetAtmCard(atmCard);
        atmMachine->SetAtmState(new AtmShowScreen());
    }
    else{
        std::cout<<"Incorrect deataits. Card validation failed"<<std::endl;
        exit(0);
    }

    
}