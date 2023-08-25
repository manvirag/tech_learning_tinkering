#include<bits/stdc++.h>
using namespace std;
class parent
{
	 public:
	 	int val;
};
class child : public parent
{
	  public:
	    int key; 
};
int main(){
	
	 child name;
	 name.val=2;
	 cout<<name.val<<endl;
	
}
