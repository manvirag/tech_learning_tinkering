#include<iostream>
#include"atm-keypad.h"


void AtmKeypad::ShowKeyOnScreen(){
    for(auto key: this->keys){
        std::cout<<key<<" ";
    }
    std::cout<<std::endl;
};

std::string AtmKeypad::GetSelectedKey(){
    this->ShowKeyOnScreen();
    std::string selectedKey;
    std::cout<<"Enter the selected key: ";
    std::cin>>selectedKey;
    return selectedKey;
}