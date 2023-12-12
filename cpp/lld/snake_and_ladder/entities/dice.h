// require for using rand
#include<cstdlib>

class Dice
{
public:
    static Dice* dicInstance;
    const int MAX = 6;
    Dice(){

    }  
    public:
    static Dice* getDiceInstance() {
        if(dicInstance == nullptr){
            dicInstance = new Dice();
        }
        return dicInstance;
    }  

    int rollDice(){
        int randomNumber = rand();
        return randomNumber%MAX + 1;
    }
};

Dice* Dice::dicInstance = nullptr;
