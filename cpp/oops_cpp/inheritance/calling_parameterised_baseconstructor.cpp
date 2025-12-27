/*class Base
{ 
    int x;
    public:
    // default constructor
    Base() 
    { 
        cout << "Base default constructor\n"; 
    }
};

class Derived : public Base
{ 
    int y;
    public:
    // default constructor
    Derived() 
    { 
        cout << "Derived default constructor\n"; 
    }
    // parameterized constructor
    Derived(int i) 
    { 
        cout << "Derived parameterized constructor\n"; 
    }
};

int main()
{
    Base b;        
    Derived d1;    
    Derived d2(10);
}

Base default constructor
Base default constructor
Derived default constructor
Base default constructor
Derived parameterized constructor*/

class Base
{ 
    int x;
    public:
    // parameterized constructor
    Base(int i)
    { 
        x = i;
        cout << "Base Parameterized Constructor\n";
    }
};

class Derived : public Base
{ 
    int y;
    public:
    // parameterized constructor
    Derived(int j):Base(j)
    { 
        y = j;
        cout << "Derived Parameterized Constructor\n";
    }
};

int main()
{
    Derived d(10) ;
}

Base Parameterized Constructor
Derived Parameterized Constructor
