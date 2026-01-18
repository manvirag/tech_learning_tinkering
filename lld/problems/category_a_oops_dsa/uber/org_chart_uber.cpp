/*
→Design similar to org chart which we have on teams. 
There were multiple requirement. 
Like 
allocate budget. 
Add manager, 
add IC, 
remove manager, 
remove IC 
and few more requirements



→  https://leetcode.com/discuss/post/7313588/uber-sse-technical-phone-round-by-get_si-w5km/

1. You are given a list of edges [[a,b], [b,c], [d,f], [f,g]]. Here a,b means a is 
manager of b. You have to implement 3 functions:
a. Return total number of employees under a manager (both direct and indirect)
b. Change the manager of an employee. If that employee has some reportees then those will also be affected.
c. Add new employees. Eg: [c,d] needs to be added.

Input format was not hardcoded. 
You have to write the main function on how input will be given to these apis.

Basically if you create a tree structure
 you would be able to solve this question. 
 The key was to have optimized solution such that (a) part should have TC - O(1)

I couldn't come up with optimized solution in the given time.




→ https://leetcode.com/discuss/post/6880181/uber-sde-2-interview-experience-verdict-8e53t/

Asked to design an Employee & Team Directory Management System with:

- Add employee
- Assign manager
- Add to team
- Hierarchical view from any employee
- Calculate employee and team CTC

Was able to implement the core logic and run it on given test cases.

Due to time limits, couldn’t add all validations and edge-case handling.
*/




#include <iostream>
#include <vector>
#include <map>
#include <string>
using namespace std;

/* =============================
   Employee Node
   ============================= */
class Employee {
public:
    string id;
    Employee* manager;
    vector<Employee*> children; // multiple reportees
    int totalReportees;         // direct + indirect

    Employee(string _id) {
        id = _id;
        manager = nullptr;
        totalReportees = 0;
    }
};

/* =============================
   Org Chart
   ============================= */
class OrgChart {
private:
    map<string, Employee*> empMap;

    // Update totalReportees up the manager chain by delta
    void updateManagerChain(Employee* node, int delta) {
        Employee* mgr = node->manager;
        while (mgr != nullptr) {
            mgr->totalReportees += delta;
            mgr = mgr->manager;
        }
    }

    // Count total nodes in a subtree
    int subtreeSize(Employee* node) {
        return node->totalReportees + 1;
    }

public:
    /* Add a new employee under a manager (multiple reportees allowed) */
    void addEmployee(string empId, string managerId = "") {
        if (empMap.find(empId) != empMap.end()) return; // already exists

        Employee* emp = new Employee(empId);
        empMap[empId] = emp;

        if (managerId != "") {
            if (empMap.find(managerId) == empMap.end()) {
                addEmployee(managerId); // create manager if missing
            }
            Employee* mgr = empMap[managerId];
            emp->manager = mgr;
            mgr->children.push_back(emp);

            // update reportee counts up the chain
            updateManagerChain(emp, 1);
        }
    }

    /* Return total reportees under employee (direct + indirect) */
    int getTotalReportees(string empId) {
        if (empMap.find(empId) == empMap.end()) return 0;
        return empMap[empId]->totalReportees;
    }

    /* Change manager of employee (subtree move) */
    void changeManager(string empId, string newManagerId) {
        if (empMap.find(empId) == empMap.end()) return;
        if (empMap.find(newManagerId) == empMap.end()) addEmployee(newManagerId);

        Employee* emp = empMap[empId];
        Employee* oldMgr = emp->manager;
        Employee* newMgr = empMap[newManagerId];

        if (oldMgr != nullptr) {
            // remove from old manager's children
            for (auto it = oldMgr->children.begin(); it != oldMgr->children.end(); ++it) {
                if (*it == emp) {
                    oldMgr->children.erase(it);
                    break;
                }
            }
            // subtract entire subtree size from old manager chain
            updateManagerChain(emp, -subtreeSize(emp));
        }

        // set new manager
        emp->manager = newMgr;
        newMgr->children.push_back(emp);
        // add subtree size to new manager chain
        updateManagerChain(emp, subtreeSize(emp));
    }

    /* Print hierarchy from any employee */
    void printHierarchy(string empId, int level = 0) {
        if (empMap.find(empId) == empMap.end()) return;
        Employee* emp = empMap[empId];

        for (int i = 0; i < level; i++) cout << "  ";
        cout << emp->id << " (" << emp->totalReportees << ")" << endl;

        for (size_t i = 0; i < emp->children.size(); i++) {
            printHierarchy(emp->children[i]->id, level + 1);
        }
    }
};

/* =============================
   Demo / Main
   ============================= */
int main() {
    OrgChart org;

    // Build initial org
    org.addEmployee("a");          // root
    org.addEmployee("b", "a");
    org.addEmployee("c", "a");
    org.addEmployee("d", "b");
    org.addEmployee("f", "d");
    org.addEmployee("g", "d");

    cout << "Total reportees under a: " << org.getTotalReportees("a") << endl; // 6
    cout << "Total reportees under b: " << org.getTotalReportees("b") << endl; // 3
    cout << "Total reportees under d: " << org.getTotalReportees("d") << endl; // 2

    cout << "\nHierarchy before changeManager:\n";
    org.printHierarchy("a");

    // Move d under c
    org.changeManager("d", "c");

    cout << "\nHierarchy after changeManager d->c:\n";
    org.printHierarchy("a");

    cout << "\nTotal reportees under a: " << org.getTotalReportees("a") << endl; // 6
    cout << "Total reportees under b: " << org.getTotalReportees("b") << endl; // 0
    cout << "Total reportees under c: " << org.getTotalReportees("c") << endl; // 3
    cout << "Total reportees under d: " << org.getTotalReportees("d") << endl; // 2

    return 0;
}
