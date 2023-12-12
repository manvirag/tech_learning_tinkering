#pragma once

class AtmMachine;

class AtmStatesInterface
{
    public:
        virtual void processRequest(AtmMachine* atmMachine) = 0;

};
