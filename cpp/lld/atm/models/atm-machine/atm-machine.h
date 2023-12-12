#pragma once
#include"../atm-card/atm-card.h"
#include"../atm-states/atm-states.h"
#include"../atm-states/atm-stale.h"



class AtmMachine
{

    AtmCard* insertedAtmCard;
    AtmStatesInterface* atmCurrentState;

    public:
        AtmStatesInterface* GetAtmState();
        void SetAtmState(AtmStatesInterface* atmState);
        void SetAtmCard(AtmCard* atmCard);
        AtmCard* GetAtmCard();
};
