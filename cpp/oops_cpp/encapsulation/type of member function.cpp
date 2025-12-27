// different types of class member function
#include<bits/stdc++.h>
using namespace std;
class base{
	public:
		int data;
	 void functionbase();
	 
};

class test{
	 private:
	 	int z;
	 public:
	 	int y;
	 	
	 static int x;
	 static void staticfunction(){      // it will give error since static member can only acces static member and that should be define
//	 	y=2;
       cout<<x<<endl;
	 }
	 void constantfunction() const
	 {
	 	x++;
	 	//y=3;                 // give error since can not update data member except static data member
	 }
	 friend void friendfunction(int y);
     friend void base::functionbase();
}; 
void base::functionbase()
{
	 test tt;
	 tt.z=7;
	 
}
void friendfunction(int y)
{
	 cout<<y<<endl;
}
int test::x;   // if we donot define here it will give compilation error
void simplefuncition(){
   cout<<"i am simple function";	 		
}
int main(){
	
//	test va;
//	va.x=4;
//	va.staticfunction();
	test vs;
	vs.x=4;
	vs.y=2;
	cout<<vs.x;
	vs.constantfunction();
	cout<<vs.x;
	int z=3;
	friendfunction(z);
	// all the function inside the class are inline
	
	/*
	  1- simple
	  2- static
	  3- const
	  4- friend
	  5- inline
	
	*/
	
	
	
	
	
}
