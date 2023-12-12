#include<iostream>
using namespace std;

class Shape{

  public:
    virtual void Draw() = 0;
    static Shape* ShapeFactory(string type);

};

class Circle : public Shape{
  private:
    void Draw(){
        cout<<"I am the circle"<<endl;
    }
};

class Rectangle : public Shape{
  private:
    void Draw(){
        cout<<"I am the rectangle"<<endl;
    }
};

Shape* Shape::ShapeFactory(string type){
    if(type == "circle")
    return new Circle();
    else if(type == "rectangle")
    return new Rectangle();

    return NULL;
}

int main(){

    Shape * obj = Shape::ShapeFactory("rectangle");
    obj->Draw();
    cout<<"abcd"<<endl;
    return 0;
}