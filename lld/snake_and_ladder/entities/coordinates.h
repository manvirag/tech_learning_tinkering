#ifndef COORDINATE
#define COORDINATE
class Coordinates
{
public:
    int x;
    int y;
    Coordinates()
    {
        this->x = 0;
        this->y = 0;
    }
    Coordinates(int x, int y)
    {
        this->x = x;
        this->y = y;
    }
    void update(int x, int y)
    {
        this->x = x;
        this->y = y;
    }
     bool operator<(const Coordinates& other) const {
        if (x == other.x) {
            return y < other.y;
        }
        return x < other.x;
    }
    
    bool operator==(const Coordinates& other) const {
        return x == other.x && y == other.y;
    }
};

#endif