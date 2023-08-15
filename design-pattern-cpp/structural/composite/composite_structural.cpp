
/*
i mean don't know why this is a different design pattern , its very simple
intuitive also,

Let's say we have a tree structure that represents a company's organizational structure. The root node of the tree represents the company, and the child nodes of the root node represent the company's departments. The child nodes of the departments can be either other departments or employees.

The Composite Design Pattern allows us to treat all of the nodes in the tree in a uniform way, regardless of whether they are departments or employees. This means that we can write code that can traverse the tree and perform operations on all of the nodes, without having to worry about whether the node is a department or an employee.

For example, we could write code that traverses the tree and prints the name of each department and employee. This code would be the same, regardless of whether the node is a department or an employee.

*/

#include <iostream>
#include <vector>

using namespace std;

// The base component class
class Component {
public:
  virtual void traverse() = 0;
};

// The leaf node class
class Leaf : public Component {
public:
  int value;

  Leaf(int val) { value = val; }

  void traverse() {
    cout << value << " ";
  }
};

// The composite node class
class Composite : public Component {
public:
  vector<Component*> children;

  void add(Component* ele) { children.push_back(ele); }

  void traverse() {
    for (int i = 0; i < children.size(); i++) {
      children[i]->traverse();
    }
  }
};

int main() {
  Composite containers[4];
  for (int i = 0; i < 4; i++) {
    for (int j = 0; j < 3; j++) {
      containers[i].add(new Leaf(i * 3 + j));
    }
  }
  for (int i = 1; i < 4; i++) {
    containers[0].add(&(containers[i]));
  }
  for (int i = 0; i < 4; i++) {
    containers[i].traverse();
    cout << endl;
  }
}