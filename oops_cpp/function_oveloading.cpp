

#include <bits/stdc++.h> 

using namespace std; 
class G
{ 
	public: 
	
	// function with 1 int parameter 
	void func(int x) 
	{ 
		cout << "value of x is " << x << endl; 
	} 
	
	// function with same name but 1 double parameter 
	void func(double x) 
	{ 
		cout << "value of x is " << x << endl; 
	} 
	
	// function with same name and 2 int parameters 
	void func(int x, int y) 
	{ 
		cout << "value of x and y is " << x << ", " << y << endl; 
	} 
}; 

int main() { 
	
	G obj1; 
	
	// Which function is called will depend on the parameters passed 
	// The first 'func' is called 
	obj1.func(7); 
	
	// The second 'func' is called 
	obj1.func(9.132); 
	
	// The third 'func' is called 
	obj1.func(85,64); 
	return 0; 
} 



#include<iostream> 
using namespace std; 

class Test 
{ 
protected: 
	int x; 
public: 
	Test (int i):x(i) { } 
	void fun() const
	{ 
		cout << "fun() const called " << endl; 
	} 
	void fun() 
	{ 
		cout << "fun() called " << endl; 
	} 
}; 

int main() 
{ 
	Test t1 (10); 
	const Test t2 (20); 
	t1.fun(); 
	t2.fun(); 
	return 0; 
} 





// PROGRAM 2 (Compiles and runs fine) 
#include<iostream> 
using namespace std; 

void fun(char *a) 
{ 
cout << "non-const fun() " << a; 
} 

void fun(const char *a) 
{ 
cout << "const fun() " << a; 
} 

int main() 
{ 
const char *ptr = "GeeksforGeeks"; 
fun(ptr); 
return 0; 
} 


// C++ allows functions to be overloaded on the basis of const-ness of parameters only if the const parameter is a reference or a pointer.
// That is why the program 1 failed in compilation, but the program 2 worked fine. This rule actually makes sense. In program 1, 
// the parameter ‘i’ is passed by value, so ‘i’ in fun() is a copy of ‘i’ in main(). Hence fun() cannot modify ‘i’ of main(). Therefore, 
// it doesn’t matter whether ‘i’ is received as a const parameter or normal parameter. When we pass by reference or pointer, we can modify
// the value referred or pointed, so we can have two versions of a function, one which can modify the referred or pointed value, other
// which can not.
