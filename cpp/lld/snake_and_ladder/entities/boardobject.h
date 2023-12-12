#include "coordinates.h"
class BoardObject{
    public:
    Coordinates start;
    Coordinates end;
    BoardObject(Coordinates start, Coordinates end){
        this->start = start;
        this->end = end;
    }
     virtual Coordinates updateCoordinate(Coordinates pos) = 0;
       
};

class Snake : public BoardObject {
    public:
     Snake(Coordinates start, Coordinates end) : BoardObject(start, end){};
     Coordinates updateCoordinate(Coordinates pos)  override{
       if(pos.x == this->start.x && pos.y == this->start.y){
        cout<<"Ha Ha Snake Got You Baby "<<"\U0001F600"<<endl;
        return this->end;
       }
       return pos;
     }

};

class Ladder : public BoardObject {
    public:
     Ladder(Coordinates start, Coordinates end) : BoardObject(start, end){};
     Coordinates updateCoordinate(Coordinates pos)  override{
      if(pos.x == this->start.x && pos.y == this->start.y){
       cout<<"Nice baby you got the ladder "<<"\U0001F600"<<endl;
        return this->end;
       }
       return pos;
    }

};