// same for all
// +,-,/,*,%,^
#include<bits/stdc++.h>
using namespace std;
class par
{
     public:
      int x,y;
      
      par(int a,int b)
      {
         x=a; y=b;
      }
      par operator *(const par &temp)
      {
           return par(x*temp.x, y*temp.y);
      }

      void display()
      {
           cout<<x<<" "<<y<<endl;
      }
};
int main()
{
      par t(1,2),t1(1,2);
      par t2=t*t1;
      
      t2.display();
         return 0;
}
