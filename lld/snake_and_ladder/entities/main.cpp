#include<iostream>
#include "gameFacade.h"
#include<string>
#include<vector>
using namespace std;

/*
https://anubhavgupta606417.invisionapp.com/freehand/Snake-Ladder-Design-7bQtu2quX?dsid_h=8ffb811f2a05dff740c48edb7cfb88f505be5d440e358008898f743918d4a9a9&uid_h=373fa89dda27bd6c7475f430a6bbd494a1d3935e03b6160caf1b1674109ee030
*/

int main(){
    vector<string> arr;
    arr.push_back("anubhav");
    arr.push_back("manvirag");
   GameFacade obj = GameFacade(arr);
   obj.GameRun();
}