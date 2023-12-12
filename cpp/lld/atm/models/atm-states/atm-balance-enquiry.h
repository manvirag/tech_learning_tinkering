#include"../atm-machine/atm-machine.h"
#include"atm-states.h"
#include"../atm-printer/atm-printer.h"


class AtmBalanceEnquiry : public AtmStatesInterface
{
    private:
      Printer* atmPrinter;
    public:
        AtmBalanceEnquiry(){
        this->atmPrinter = new Printer();
        }
        void processRequest(AtmMachine* atmMachine);
};