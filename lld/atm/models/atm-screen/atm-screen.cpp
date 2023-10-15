#include"atm-screen.h"
#include<string>
#include<vector>
#include<iostream>

void AtmScreen::ShowOptionsToUser(std::vector<std::string> options) {
    for(auto option: options){
        std::cout<<option<<std::endl;
    }
}

std::string AtmScreen::GetTheUserSelectedOption(std::vector<std::string> options){
    this->ShowOptionsToUser(options);
    std::cout<<"Enter the options for further processing: "<<std::endl;
    std::string selectedOption;
    std:: cin>>selectedOption;
    return selectedOption;
}