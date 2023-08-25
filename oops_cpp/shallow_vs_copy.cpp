#include <iostream>
using namespace std;

class shallow{
	public:
		int* p;
	shallow(int x)
	{
	  p=new int(x);	
	}	
};
class Copy{
	public:
		int* p;
	Copy(int x)
	{
	  p=new int(x);	
	}	
	Copy(const Copy &t)
	{
	     this->p=new int(*(t.p));
	}
	
};
int main() {

      shallow temp(5);
      shallow temp1(temp);
      cout<<*(temp.p);
      cout<<*(temp1.p);
      cout<<endl;
      *(temp.p)=6;
      cout<<*(temp.p);
      cout<<*(temp1.p);
      cout<<endl;
      Copy temp2(5);
      Copy temp3(temp2);
      cout<<*(temp2.p);
      cout<<*(temp3.p);
      cout<<endl;
      *(temp2.p)=6;
      cout<<*(temp2.p);
      cout<<*(temp3.p);
      
      
}


/*output
55
66
55
65*/
