/*

The storage team is developing a version management system. Every night we built one or a few new releases. 
The versions could be compatible, which means the system could be upgraded without the data migration; otherwise, 
we have to upgrade the data before deploying the new release binaries (typically with a MapReduce job to rebuild the data). 
The compatibility is transitive, which means if version a is compatible with version b, version b is compatible with c, 
then a is compatible with c. If any of them are incompatible, it makes a to c incompatible. We already have a compatible 
check test (Jenkins job). After each release, we only run the compatible test against the previous version.

public void addNewVersion(int ver, boolean isCompatibleWithPrev)
{
}

```
public boolean isCompatible(int srcVer, int targetVer)
{
}

```

1 -> 2 -> 3 -> 4 -> 5 -> 6
T    T    F    T    T

```
    VersionCompatibilityManagement versions = new VersionCompatibilityManagement();
    versions.addNewVersion(1, false);
    versions.addNewVersion(2, true);
    versions.addNewVersion(3, true);
    versions.addNewVersion(4, false);
    versions.addNewVersion(5, true);
    versions.addNewVersion(6, true);
    assert(versions.isCompatible(1, 3) == true);
    assert(versions.isCompatible(3, 5) == false);
    assert(versions.isCompatible(4, 2) == false); // downgrade
    assert(versions.isCompatible(3, 3) == true); // upgrade to itself, always compatible

```

so its kine of linear right ? 
1 version -> new version -> ...   

1 2 3 4 

let say one is false in betwee 
then before and not is not compatible any right ? 
x,y -> compatible -> no false from x --> y or y-> x small -> large 

1 -> alwasy false
we can club all near true and when xy become other other club they are not compatible -> 

logn -> to see compatibility

nearly constaint -> require to maintain all version state.

*/

#include<iostream> 
#include<algorithm>
using namespace std; 
class VersionCompatability {
    public: 
        vector<int> version; // assume version can fil in int. 
        int latestVersion; 

    VersionCompatability(): latestVersion(0) {}
    void addNewVersion(int id, bool isPrevCompatible) {
        if(id <= latestVersion) {
            cout<<"wrong input"<<endl;
            return ; 
        }
        if(id != latestVersion + 1) {
            cout<<"wrong version relation"<<endl;
        }
        if(!isPrevCompatible){
            version.push_back(id);
        }
        latestVersion++;
        cout<<"done version addition"<<endl;
    }   

    bool isCompatible(int x ,int y) { 
        if(y<x) {
            swap(x,y);
        }
        if(x==y){
            return true; 
        }
        if(x<=0 || y<=0 || x>latestVersion || y>latestVersion) {
            cout<<"invalid input"<<endl;
            return false;
        }
        auto ptrX = upper_bound(version.begin(), version.end(),x);
        auto ptrY = upper_bound(version.begin(), version.end(),y);

        return ptrX == ptrY;
    }

};

int main() {
    
    VersionCompatability *versions = new VersionCompatability(); 
    versions->addNewVersion(1, false);
    versions->addNewVersion(2, true);
    versions->addNewVersion(3, true);
    versions->addNewVersion(4, false);
    versions->addNewVersion(5, true);
    versions->addNewVersion(6, true);
    cout<<versions->isCompatible(1, 3)<<endl;
    cout<<versions->isCompatible(3, 5)<<endl;
    cout<<versions->isCompatible(4, 2)<<endl;
    cout<<versions->isCompatible(3, 3)<<endl;
    return 0; 
}

/*
by gpt: 
#include <iostream>
#include <unordered_map>
#include <shared_mutex>
#include <stdexcept>

using namespace std;

class VersionCompatibilityManagement {
private:
    int latestVersion;
    int currentSegment;
    unordered_map<int, int> versionToSegment;
    mutable shared_mutex mu;

public:
    VersionCompatibilityManagement()
        : latestVersion(0), currentSegment(0) {}

    // Adds a new version; must be added in increasing order
    void addNewVersion(int ver, bool isCompatibleWithPrev) {
        unique_lock lock(mu);

        if (ver != latestVersion + 1) {
            throw invalid_argument("Versions must be added sequentially");
        }

        // First version OR incompatible with previous → new segment
        if (ver == 1 || !isCompatibleWithPrev) {
            currentSegment++;
        }

        versionToSegment[ver] = currentSegment;
        latestVersion = ver;
    }

    // Checks compatibility only for upgrades
    bool isCompatible(int srcVer, int targetVer) {
        if (srcVer == targetVer) return true;
        if (srcVer > targetVer) return false;  // downgrade not allowed

        shared_lock lock(mu);

        auto it1 = versionToSegment.find(srcVer);
        auto it2 = versionToSegment.find(targetVer);

        if (it1 == versionToSegment.end() || it2 == versionToSegment.end()) {
            return false;
        }

        return it1->second == it2->second;
    }
};

int main() {
    VersionCompatibilityManagement versions;

    versions.addNewVersion(1, false);
    versions.addNewVersion(2, true);
    versions.addNewVersion(3, true);
    versions.addNewVersion(4, false);
    versions.addNewVersion(5, true);
    versions.addNewVersion(6, true);

    cout << versions.isCompatible(1, 3) << endl; // true
    cout << versions.isCompatible(3, 5) << endl; // false
    cout << versions.isCompatible(4, 2) << endl; // false (downgrade)
    cout << versions.isCompatible(3, 3) << endl; // true

    return 0;
}



*/