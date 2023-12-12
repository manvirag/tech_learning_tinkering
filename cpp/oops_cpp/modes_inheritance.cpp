#include<bits/stdc++.h>
using namespace std;
class A
{
    public:
	  int l;
	protected:
	  int b;
	private:
	  int h;	
};
class B : public A
{
	 
};
class C: protected A{
	
};
class D: private A{
	
};
int main(){
	
//      B p; 
//      p.b=2; // error protected
//      p.l=2; // fine
//      p.h=2; // error private
      
//      C q;
//      q.l=2;  // protected error
//      q.b=2;  // protected error
//      q.h=2; // private error
	 	
	 	D r;
//	 	r.l=2; // inaccessable
//        r.h=2;    // error
//        r.b=2;       // error
}
