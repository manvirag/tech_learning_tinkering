#include<bits/stdc++.h>
using namespace std;
#define int long long int
const int mxN =150000;
class test{
	public:
		test()
		{ 
		    cout<<"i am called right now";
		}
		~test()
		{
			 cout<<"hi i am also called ";
		}
};
void func(void){
    test x;
	return ;   
}
int32_t main(){
  ios_base::sync_with_stdio(0);
   cin.tie(0); cout.tie(0);

     func();


 return 0;}

