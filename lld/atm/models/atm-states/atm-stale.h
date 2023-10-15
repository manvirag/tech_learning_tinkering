#pragma once   // so that only once import or one declaration
#include"atm-states.h"

class AtmMachine;  // if include the header it will go in cyclic dependency

class AtmStale : public AtmStatesInterface
{
    public:
        void processRequest(AtmMachine* atmMachine);
};