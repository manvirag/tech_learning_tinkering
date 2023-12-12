#include<iostream> 
#include<stdlib.h> 

using namespace std; 
class student 
{ 
	string name; 
	int age; 
public: 
	student() 
	{ 
		cout<< "Constructor is called\n" ; 
	} 
	student(string name, int age) 
	{ 
		this->name = name; 
		this->age = age; 
	} 
	void display() 
	{ 
		cout<< "Name:" << name << endl; 
		cout<< "Age:" << age << endl; 
	} 
	
	void * operator new(size_t size);
	void  operator delete(void *p);
	void * operator new[](size_t size);
	void operator delete[](void *p);
	 
}; 

void *student::operator new(size_t size)
{
    void *p;
    cout << "In overloaded new.";
    p =  malloc(size);
    
    return p;
}

void student::operator delete(void *p)
{
    cout << "In overloaded delete.\n";
    free(p);
}

void *student::operator new[](size_t size)
{
    void *p;
    cout << "Using overload new[].\n";
    p =  malloc(size);
    
    return p;
}

void student::operator delete[](void *p)
{
    cout << "Free array using overloaded delete[]\n";
    free(p);
}
int main() 
{ 
	student * p = new student[5] ;

	p->display(); 
	delete [] p; 
} 

/*
NOTE: In the above new overloaded function, we have allocated dynamic memory through new operator, but it should be global new operator otherwise it will go in recursion
void *p = new student(); // this will go in recursion asnew will be overloaded again and again
void *p = ::new student(); // this is correct*/
