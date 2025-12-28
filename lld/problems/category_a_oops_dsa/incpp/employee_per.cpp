#include <iostream>
#include <map>
#include <vector>
using namespace std;

/*

resource
user 
permission -> enum 
grantType -> 
user - permission -> granttype1
user - permission -> granttype2
*/

enum PermissionType {
    READ = 1,
    WRITE = 2,
    EXECUTE = 3
};
class Resource {
    public:
     string resourceId;
     Resource() = default;
     Resource(string resourceId) {
        this->resourceId = resourceId;
     }
};
class GrantType {
    public:
     string grantTypeId;
     GrantType() = default;
     GrantType(string grantTypeId) {
        this->grantTypeId = grantTypeId;
     }
};
class UserPermission {
    public:
     PermissionType permissionType;
     vector<string> grantTypes;
     UserPermission() = default;
     UserPermission(PermissionType perType, string grantTypeId) {
        this->permissionType = perType;
        this->grantTypes.push_back(grantTypeId);
     }
};
class Employee {
    public:
     string employeeId;
     vector<Resource> accessedResources;
     map<string, map<int,UserPermission>> permissions;
     Employee() = default;
     Employee(string employeeId) {
        this->employeeId = employeeId;
     }
};
class EmployeeAccessManagement {
    public:
     map<string, Employee> employees;
     map<string, Resource> resources;
     map<string, GrantType> grantTypes;
     EmployeeAccessManagement() {
        employees.clear();
        resources.clear();
        grantTypes.clear();
     }
     map<int, bool> get(string employeeId, string resourceId) {
        map<int, bool> result;
        result[READ] = false;
        result[WRITE] = false;
        result[EXECUTE] = false;
        if(employees.find(employeeId) == employees.end()) {
            return result;
        }
        if(resources.find(resourceId) == resources.end()) {
            return result;
        }
        for(auto per: employees[employeeId].permissions[resourceId]) {
            result[per.first] = true;
        }
        return result;
     }

     void grant(string employeeId, string resourceId, PermissionType permissionType) {
        if(employees.find(employeeId) == employees.end()) {
            return;
        }
        if(resources.find(resourceId) == resources.end()) {
            return;
        }
        if (employees[employeeId].permissions.find(resourceId) == employees[employeeId].permissions.end()) {    
            employees[employeeId].permissions[resourceId] = {};
            employees[employeeId].permissions[resourceId][permissionType] = UserPermission(permissionType, "1");
            employees[employeeId].accessedResources.push_back(resources[resourceId]);
        } else {
            employees[employeeId].permissions[resourceId][permissionType] = UserPermission(permissionType, "1");
        } 
    }
};
int main() {
    EmployeeAccessManagement employeeAccessManagement;
    Employee employee1("1");
    Employee employee2("2");
    Resource resource1("1");
    Resource resource2("2");
    GrantType grantType1("1");
    employeeAccessManagement.employees["1"] = employee1;
    employeeAccessManagement.resources["1"] = resource1;
    employeeAccessManagement.grantTypes["1"] = grantType1;
    employeeAccessManagement.employees["2"] = employee2;
    employeeAccessManagement.resources["2"] = resource2;
    employeeAccessManagement.grant("1", "2", READ);
    map<int, bool> permissions = employeeAccessManagement.get("1", "2");
    for(auto per : permissions) {
        cout<<per.first<<" "<<per.second<<endl;
    }
    return 0;
}