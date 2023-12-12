#include"../atm-machine/atm-machine.h"
#include"atm-states.h"
#include"../atm-card-reader/atm-card-reader.h"
#include"../atm-keypad/atm-keypad.h"


class AtmInsertion : public AtmStatesInterface
{
    private:
       AtmCardReader * atmCardReader;
       AtmKeypad * atmKeyPad;
    public:
        AtmInsertion(){
            this->atmCardReader = new AtmCardReader();
            this->atmKeyPad = new AtmKeypad();
        }
        void processRequest(AtmMachine* atmMachine);
};