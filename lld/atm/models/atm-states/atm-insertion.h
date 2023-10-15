#include"atm-machine.h"
#include"atm-states.h"


class AtmInsertion : public AtmStatesInterface
{
    public:
        void processRequest(AtmMachine atmMachine);
};