// -x !x

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
      par operator-()
      {
           x=-x;
           y=-y;

           return par(x,y);
      }

      void display()
      {
           cout<<x<<" "<<y<<endl;
      }
};
int main()
{
      par t(1,2);
      
      t.display();
      -t;
      t.display();
         return 0;
}

// can also do with !
