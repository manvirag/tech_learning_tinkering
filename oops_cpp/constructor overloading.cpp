#include<bits/stdc++.h>
using namespace std;
#define int long long int
const int mxN =150000;
class test{
	public:
		int x;
		int y;
	test(int x);
	test(int x,int y);
	void prin();
};
test::test(int y)
{
	 y=3;
}
test::test(int x,int y)
{
	 x=2;y=5;
}
void test::prin()
{
	 cout<<x<<" "<<y<<endl;
}
int32_t main(){
  ios_base::sync_with_stdio(0);
   cin.tie(0); cout.tie(0);

     test val(2);
     test val1(2,3);
     val.prin();
     val1.prin();
     
     




 return 0;}

