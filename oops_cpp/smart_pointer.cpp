#include<iostream>
using namespace std;
// A generic smart pointer class
template <class T>
class Smartpointer {
   T *p; // Actual pointer
   public:
      // Constructor
      Smartpointer(T *ptr = NULL) {
         p = ptr;
      }
   // Destructor
   ~Smartpointer() {
      delete(p);
   }
   // Overloading dereferencing operator
   T & operator * () {
      return *p;
   }
   // Overloding arrow operator so that members of T can be accessed
   // like a pointer
   T * operator -> () {
      return p;
   }
};
int main() {
   Smartpointer<int> p(new int());
   *p = 26;
   cout << "Value is: "<<*p;
   return 0;
}
