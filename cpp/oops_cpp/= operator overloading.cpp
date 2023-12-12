
#include <iostream>
using namespace std;
 
class Distance {
   
     public:
   
      int feet;             // 0 to infinite
      int inches;           // 0 to 12
      
      // required constructors
      Distance() {
         feet = 0;
         inches = 0;
      }
      Distance(int f, int i) {
         feet = f;
         inches = i;
      }
      void display()
      {
          cout<<feet<<" "<<inches<<endl;
      }
      void operator=(const Distance &temp)
      { 
        feet=temp.feet;
        inches=temp.inches;
     }
};

int main() {
   Distance D1(11, 10);
   Distance D2;
   D2=D1;
   D2.display();
 
 
   return 0;
}
