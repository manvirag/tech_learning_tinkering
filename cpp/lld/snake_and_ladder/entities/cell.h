#include "Coordinates.h"

class Cell
{
    public:
      Coordinates position = Coordinates();
      Cell(int x, int y){
        this->position.update(x,y);
      }
};
