#include"atm-machine.h"
#include"atm-states.h"


class AtmStale : public AtmStatesInterface
{
    public:
        void processRequest(ATM_MACHINE::AtmMachine atmMachine);
};