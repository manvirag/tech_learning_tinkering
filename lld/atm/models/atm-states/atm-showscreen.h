#include"../atm-machine/atm-machine.h"
#include"atm-states.h"
#include"../atm-screen/atm-screen.h"


class AtmShowScreen : public AtmStatesInterface
{
    private:
       AtmScreen * atmScreen ;
    public:
        AtmShowScreen(){
            this->atmScreen  = new AtmScreen();
        }
        void processRequest(AtmMachine* atmMachine);
};