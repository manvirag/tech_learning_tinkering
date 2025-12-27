#include<bits/stdc++.h>
using namespace std;
class node{
	private:
		int x;
		node()
		{
			 cout<<"called"<<endl;
		}
		friend class b;
		
};
class b{
	public:
		b()
		{
			 node x;
			 cout<<"by me"<<endl;
		}
};
int main(){
 	
      b x;	
	
}


/*
output 
called
by me
*/
