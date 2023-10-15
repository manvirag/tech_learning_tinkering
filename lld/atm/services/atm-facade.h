#include"../models/atm-machine/atm-machine.h"

class AtmFacade{

    AtmMachine* atmMachine;
    public:
        AtmFacade() {
            this->atmMachine = new AtmMachine();  
        }
        void PowerOnTheAtm();
        void process();
        
        
};
