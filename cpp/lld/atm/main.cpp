/*
design link:
https://anubhavgupta606417.invisionapp.com/freehand/ATM-BlgsgWEtn

*/

#include"./services/atm-facade.h"

int main(){

    AtmFacade * atmFacade = new AtmFacade();
    atmFacade -> PowerOnTheAtm();
    
    return 0;
}