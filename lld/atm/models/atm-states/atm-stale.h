#include"atm-machine.h"
#include"atm-states.h"


class AtmStale : public AtmStatesInterface
{
    public:
        void processRequest(AtmMachine atmMachine);
};