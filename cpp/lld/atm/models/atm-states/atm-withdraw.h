#include"../atm-states/atm-states.h"
#include"atm-states.h"
#include"../atm-printer/atm-printer.h"


class AtmWithDraw : public AtmStatesInterface
{
    private:
        Printer * atmPrinter;
    public:
        AtmWithDraw(){
            this->atmPrinter = new Printer();
        }
        void processRequest(AtmMachine* atmMachine);
};