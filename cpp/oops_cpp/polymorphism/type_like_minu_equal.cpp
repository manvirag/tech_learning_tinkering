#include<iostream>
using namespace std;

class com{
	
	 public:
	 	int x,y;
	 
	    com(int a,int b)
		{
		    x=a;y=b;	
		}	
		
		void operator +=(int val)  // +=,-=,/=%=,^=,&=,|=,*=,<<=,>>=
		{
		    x+=val;
			y+=val;	
		}
	
	
};
int main(){
	com a(2,3);
	com b(5,7);
	
	com c=a;
	c+=2;
	cout<<c.x<<" "<<c.y<<endl;
	
	
	cout<<c.x<<" "<<c.y	<<endl;
}
