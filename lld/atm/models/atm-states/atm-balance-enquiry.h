#include"atm-machine.h"
#include"atm-states.h"


class AtmBalanceEnquiry : public AtmStatesInterface
{
    public:
        void processRequest(AtmMachine atmMachine);
};