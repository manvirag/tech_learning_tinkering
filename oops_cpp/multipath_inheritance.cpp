#include<bits/stdc++.h>
using namespace std;
#define int long long int
const int mxN =150000;
class A{
	public:
		int x;
};
class B: virtual public A{
	public:
		int y;
};
class C: virtual public A{
	public:
		int z;
};
class D: public C,public B{
	public:
		int p;
};
int32_t main(){
  ios_base::sync_with_stdio(0);
   cin.tie(0); cout.tie(0);

    D temp;
//    temp.x=3;   // error abmiguity occur d-c-a or d-b-a
    
    // therefore
    
//    temp.B::x=2;// no error

//    or virtual make 

  temp.x=2;



 return 0;}

