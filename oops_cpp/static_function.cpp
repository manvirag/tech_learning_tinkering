/*
https://www.scaler.com/topics/static-member-function-in-cpp/
The term "static" is a C++ keyword that gives special characteristics to variables, functions, and elements within the C++ programming language. It is used to allocate storage for elements that remain valid throughout the program's scope. In C++, static elements can persist in memory for the entire program duration. This explanation will focus on static data members and static member functions in C++.

Static Data Member:
Static data members in C++ are used to store values shared among all instances of a class. Unlike regular (non-static) data members, static data members maintain a single copy shared by all class objects. They remain in memory throughout the program but are accessible only within the class or the function in which they are defined. To declare a static data member, you use the keyword "static" within the class definition, and its definition is placed outside the class definition.

Example:

cpp
Copy code
class Scaler {
    static int number; // Declaration of static data member
};
int Scaler::number; // Definition of static data member
Static Member Function:
Static member functions in C++ are functions that can access only static data members. These functions share a single copy of themselves across different class objects. You can declare a function as static by using the "static" keyword before the function name in the class definition. Static member functions are used to perform operations that are not dependent on object-specific data.

Example:

cpp
Copy code

*/

  
class Scaler {
    static int number; // Static data member
    static void getNoOfTopics() { // Static member function declaration
         cout << number << "\n";
    }
};
int Scaler::number; // Definition of static data member

int main() {
    Scaler S;
    S.getNoOfTopics(); // Accessing static member function using object
    return 0;
}


/*
Properties of Static Member Functions:

Static members can only be accessed by static member functions.
To invoke a static member function, you use the class name and scope resolution operator (::).
Static member functions cannot be declared as virtual functions.
The "this" pointer is not accessible in static member functions.
Static member functions cannot be declared as inline functions.
Example using Static Data Member and Member Function:

cpp
Copy code

*/


class Book {
    int bookId;
    float price;
    static int count; // Static data member
public:
    void getData(int id, float p) {
        bookId = id;
        price = p;
        ++count; // Counting number of objects
    }
    void showData() {
        cout << "Book ID: " << bookId << "\t";
        cout << "Price: " << price << "\n";
    }
    static void showCount() {
        cout << "Count of Book objects: " << count << "\n";
    }
};
int Book::count = 0; // Initializing static data member

int main() {
    Book B1, B2;
    B1.getData(198, 2550.3);
    B2.getData(174, 3756.89);
    Book::showCount();
    B1.showData();
    B2.showData();
    Book B3;
    B3.getData(345, 5432.4);
    Book::showCount();
    B1.showData();
    B2.showData();
    B3.showData();
    return 0;
}

/*
In this example, static data member count keeps track of the number of Book objects created. The static member function showCount displays the count, and regular member functions getData and showData operate on object-specific data. Static member functions can only access static data members.

In conclusion, static data members and static member functions are valuable in C++ for sharing data across class instances and performing operations that don't require object-specific data.
  */
