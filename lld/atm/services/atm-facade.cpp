#include"atm-facade.h"
#include<iostream>
void AtmFacade::PowerOnTheAtm(){
        std::cout<<"Wait......powering on the atm...." << std::endl;
        this->atmMachine->SetAtmState(new AtmStale());
        std::cout<<"Done......powered on the atm...." << std::endl;

        while(true){
            this->process();
        }
}
void AtmFacade :: process(){
    this->atmMachine->GetAtmState()->processRequest(this->atmMachine);
}