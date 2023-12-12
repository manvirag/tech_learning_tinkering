#include<bits/stdc++.h>
using namespace std;
#define int long long int
const int mxN =150000;
class base1{
	public:
		int x;
        void prin()	{
        	cout<<" hi i am in base1"<<endl;
		}	
};
class base2{
	public:
		int z;
		
};
class derived : public base1,public base2{       // can be private , protected  order matters first base1 constructor will call then base2
	public:
		int y;
		void prin(){
			cout<<" hi i am in derived"<<endl;
		}
};
int32_t main(){
  ios_base::sync_with_stdio(0);
   cin.tie(0); cout.tie(0);

      derived temp; // constructor of base1--->base2--->derived




 return 0;}

