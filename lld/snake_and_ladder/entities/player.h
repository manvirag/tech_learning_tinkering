#include<string>
#include<iostream>
#include "coordinates.h"
using namespace std;
class Player
{
    private:
      string name;
      Coordinates location = Coordinates();
    
    public:
      Player(string name, int x , int y){
        this->name = name;
        this->location.update(x,y);
      }

      Coordinates getLocation(){
        return location;
      }

      void updateLocation(Coordinates loc){
        this->location = loc;
      }

      void printLocation(){
        cout<<this->location.x<<" "<<this->location.y<<endl;
      }

      string getName() {
        return this->name;
      }
  
};

