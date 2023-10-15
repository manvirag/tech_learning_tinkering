#include<atm-card.h>
#include<atm-states.h>

class AtmMachine
{

    AtmCard insertedAtmCard;
    AtmStates atmCurrentState;

    public:
    AtmStates GetAtmState();
    void SetAtmState(AtmStates atmState);
    AtmMachine();
    void Init();
};
