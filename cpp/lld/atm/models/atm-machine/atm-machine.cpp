#include<iostream>
#include"atm-machine.h"
#include"../atm-states/atm-stale.h"
using namespace std;

AtmStatesInterface* AtmMachine::GetAtmState(){
    return this->atmCurrentState;
}

void AtmMachine::SetAtmState(AtmStatesInterface* atmState){
    this->atmCurrentState = atmState;
}

void AtmMachine::SetAtmCard(AtmCard* atmCard){
    this->insertedAtmCard = atmCard;
}

AtmCard* AtmMachine::GetAtmCard(){
    return this->insertedAtmCard;
}
 


