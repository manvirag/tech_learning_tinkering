#include"atm-machine.h"


class AtmStatesInterface
{
    public:
        virtual void processRequest(AtmMachine atmMachine) = 0;

};
