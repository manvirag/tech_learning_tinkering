#include<bits/stdc++.h>
using namespace std;
class Human
{
      // acces specifier 
    public:
      int x;

     // function can be defined two types
     // inside outside
     // all the member functions defined inside the class definition are by default inline (like preprosser macro)
     void inside()
     {
        cout<<"i am inside of class"<<endl;
     } 

     void outside();
     // ...................................................................................
     // contructor -- default, parameterized and copy
      


}
void Human:: outside()  
{
    cout<<"i am outside of class"<<endl;
}
int main()
{

}
