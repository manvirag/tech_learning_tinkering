#include<bits/stdc++.h>
using namespace std;
#define int long long int
const int mxN =150000;
class test{
	public:
     int x;
     
	 private:
	 int y;
	 protected:
	 int z;
	 
	 public:
	  void set(int x){
	  	y=x;
	  }		
};
int32_t main(){
  ios_base::sync_with_stdio(0);
   cin.tie(0); cout.tie(0);

       test name;
       // for public
       name.x=5;
       //for private
       name.set(6);
       // for protected same as private but in subclass can be called as (.)
 return 0;}

