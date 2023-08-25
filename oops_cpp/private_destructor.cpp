// Private Destructor 
#include <iostream> 

// A class with private destuctor 
class Test { 
private: 
	~Test() {} 
	friend void destructTest(Test*); 
}; 

// Only this function can destruct objects of Test 
void destructTest(Test* ptr) 
{ 
	delete ptr; 
} 

int main() 
{ 
	// create an object 
	Test* ptr = new Test; 

	// destruct the object 
	destructTest(ptr); 

	return 0; 
} 
