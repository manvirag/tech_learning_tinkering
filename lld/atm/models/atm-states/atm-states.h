#include"atm-machine.h"


class AtmStatesInterface
{
    public:
        virtual void processRequest(ATM_MACHINE::AtmMachine atmMachine) = 0;

};
