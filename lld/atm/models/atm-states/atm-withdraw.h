#include"atm-machine.h"
#include"atm-states.h"


class AtmWithDraw : public AtmStatesInterface
{
    public:
        void processRequest(AtmMachine atmMachine);
};