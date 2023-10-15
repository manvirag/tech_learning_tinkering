#include"atm-stale.h"
#include"../atm-machine/atm-machine.h"
#include"../atm-card/atm-card.h"
#include"atm-insertion.h"
#include"atm-states.h"
#include<iostream>

void AtmStale :: processRequest(AtmMachine* atmMachine){
    std::cout<<"Insert you atm card for processing( write cardnumber, name, bank): "<<std::endl;
    atmMachine->SetAtmState(new AtmInsertion());
}