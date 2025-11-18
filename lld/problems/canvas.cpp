#include <iostream>
#include <map>
#include <vector>
#include <set>
using namespace std;

class Box {
    public:
    int length;
    int width;
    int x;
    int y;
    int value;
    Box() = default;
    Box(int length, int width, int x, int y, int value) {
        this->length = length;
        this->width = width;
        this->x = x;
        this->y = y;
        this->value = value;
    }
};
class Cell {
   public: 
    set<pair<int,int>> values;
    Cell() {
        values.clear();
    }
    int getValue() {
        if(values.size() == 0) {
            return 0;
        }
        return values.begin()->second;
    }
    void addValue(int order, int value) {
        values.insert(make_pair(-order, value));
    }
    void removeValue(int order, int value) {
        values.erase(make_pair(-order, value));
    }
};
class Canvas {
    public: 
      int width;
      int height;
      vector<vector<Cell>> cells;
      map<string, int> boxOrder;
      map<string, Box> previousBoxPosition;
      int boxOrderIndex = 1;
      Canvas(int width, int height) {
        this->width = width;
        this->height = height;
        cells.resize(height, vector<Cell>(width));
      }
      void print() {
        for(int i = 0; i < height; i++) {
            for(int j = 0; j < width; j++) {
                cout<<cells[i][j].getValue()<<" ";
            }
            cout<<endl;
        }
        cout<<endl;
      }
      void draw(string key, int length, int width, int x, int y, int value) {  
        boxOrder[key] = boxOrderIndex++;
        previousBoxPosition[key] = Box(length, width, x, y, value);   
        for(int i = 0; i < length; i++) {
            for(int j = 0; j < width; j++) {
                cells[x+i][y+j].addValue(boxOrderIndex-1, value);
            }
        }
      }
      void move(string key, int x, int y) {
        int order = boxOrder[key];
        Box previousBox = previousBoxPosition[key];
        for(int i = 0; i < previousBox.length; i++) {
            for(int j = 0; j < previousBox.width; j++) {
                cells[previousBox.x+i][previousBox.y+j].removeValue(order, previousBox.value);
            }
        }
        boxOrder[key] = boxOrderIndex++;
        for(int i = 0; i < previousBox.length; i++) {
            for(int j = 0; j < previousBox.width; j++) {
                cells[x+i][y+j].addValue(boxOrderIndex-1, previousBox.value);
            }
        }
        previousBoxPosition[key].x = x;
        previousBoxPosition[key].y = y;
      }
};  
// assume no repeated box draw, if want unique key 
int main() {
    Canvas canvas(10, 10);
    // canvas.print();
    canvas.draw("A", 2, 2, 0, 0, 1);
    canvas.draw("B", 2, 2, 1, 1, 2);
    canvas.draw("C", 2, 2, 0, 1, 3);
    canvas.print();
    canvas.move("B", 1,2);
    canvas.print();
    return 0;
}