#include"../atm-machine/atm-machine.h"
#include"atm-states.h"


class AtmPrintReceipt : public AtmStatesInterface
{
    public:
        void processRequest(AtmMachine* atmMachine);
};