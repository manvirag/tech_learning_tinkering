#include"atm-cash-dispencer.h"
#include<string>
#include<vector>
#include<iostream>

void AtmCashDispencer::CollectCash(std::vector<std::string> notes){
    for(auto note: notes){
        std::cout<<"Please collect cash of price: "<<note<<std::endl;
    }
}