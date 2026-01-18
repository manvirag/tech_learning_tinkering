/*

Implement a thread-safe File System API mimicking Linux commands (mkdir, cd with regex paths, pwd) in Java for Uber's 60-min specialization round. Focus on OOP design with a directory tree structure.
​

Functional Requirements
text
class FileSystem {
    void mkdir(String path);      // "a/b/c" → create directories (idempotent)
    void cd(String path);         // "/a/b.*" → change to matching directory (regex support)
    String pwd();                 // Return current path "/a/b/c"
}
Paths: Absolute (/a/b), relative (../a), regex (dir.*, te?t).

mkdir: Create nested dirs if parent exists; ignore if exists.

cd: Regex matches directory name (not full path); change to first match or fail.

pwd: Always absolute path from root.

Key Constraints
Directory tree (no files needed); case-sensitive names.

Regex on basename only: cd "a/b.*" matches a/b1, a/batch but not aX/b.

Handle edge cases: cd ".." , cd "/", cd "nonexistent.*" → error.

Thread-safe for concurrent operations (use locks on tree nodes).
​

Data Structure
text
class Directory {
    Map<String, Directory> children;
    Directory parent;
    String name;
}
Root = /, current pointer tracks working directory.

Examples
text
fs.mkdir("/a/b/c");
fs.mkdir("/a/b/d");
fs.cd("a/b*");     // → /a/b
fs.pwd();           // "/a/b"
fs.cd("te?t");      // Matches nothing → error
Follow-ups (10 mins)
Add ls(path) with regex filtering.

Support symlinks or file creation.

Persistence (serialize tree).

For cd(String path), implement directory navigation with regex matching 
on directory names (not full path). Change the current working directory to the first
 matching directory found via DFS traversal from root/current position.

Exact Behavior
text
Current: /
fs.mkdir("/a/b1/c");
fs.mkdir("/a/batch/d");  
fs.mkdir("/x/test");

fs.cd("a/b.*");    // Regex matches "b1" or "batch" → cd to /a/b1 (first match)
fs.pwd();          // Returns "/a/b1"

fs.cd("te?t");     // Matches "test" → cd to /x/test  
fs.pwd();          // Returns "/x/test"

fs.cd("../b.*");   // From /x/test, cd to /a/b1 (first match)
Key Rules
Regex scope: Apply regex to basename only (b.* matches b1, batch; not ab1)

Search order: DFS from root (for absolute) or current dir (for relative)

Relative paths: cd(".."), cd("./sub") work normally

Absolute paths: Start with /

First match wins: Return first directory whose name matches regex

No match → throw exception

path
absolution -> /x/x 
relative -> ./../  , ../../s/f -> can start with . -> after that list of string, .. 

string -> can be regex 
cd only changes current path. 






.... 

class directory
    - name
    - map<string, directory*> mp; 

string currPath; 

path 
1. absolution -> list of string {a,b,c*} -> can be regex 

2. relative -> 
    . -> current 
    ./a/b -> list currentPath+{a,b}
    ../b/../c -> list currentPath+{.., b, .., c}

    list can cnotain -> .. , nmae or regex 

    regex -> cd -> will go any matching 




*/


// NOTE too >45 mins;
#include<iostream> 
#include<map> 
#include<string> 
using namespace std ;
/*
'' -> a(a) -> .. 

*/

bool isMatching(string a, string b) {
    int i=0;
    int j = 0;
    int preStarInx = -1;
    while(i<a.size() && j<b.size()) {
        if(a[i] == b[j]) {
            i++;
            j++;
        } else if(b[j] == '.' || b[j] == '?') {
            i++;
            j++;
        } else if(b[j] == '*') {
            preStarInx = j;
            j++;
        } else if(preStarInx == -1){
            return false;
        } else {
            i++;
            j = preStarInx; 
            preStarInx = -1;
        }
        // cout<<i<<" "<<j<<endl;
    }
    if(i!=a.size() && preStarInx ==-1) {
        return false;
    } 
    while(j<b.size()) {
        if(b[j] == '*') {
            j++;
        }else {
            return false;
        }
    }
    return true;
    
}
class Directory {
    public: 
        string name = "";
        map<string, Directory*> directories;
    Directory(): name(""){}
    Directory(string name): name(name){}
    
    void Insert(vector<string> absolutePath) { // no regex in mkdir
        Directory* temp = this;
        for(auto dname: absolutePath){
            if(!temp->directories[dname]){
                temp->directories[dname] = new Directory(dname);
            }
            temp = temp->directories[dname];
        }
    }
    string validateAndMove(vector<string> absolutePath, string currentPath){
        string resultedPath = "";
        Directory* temp = this;
        for(auto dname: absolutePath){
            bool foundMatching = false; 
            for(auto existingD: temp->directories){
                if(isMatching(existingD.first, dname)) {
                    resultedPath+="/" + existingD.first;
                    temp = temp->directories[existingD.first];
                    foundMatching = true;
                    break;
                }
            }
            if(!foundMatching)
            return currentPath;       
        }
        
        return resultedPath;

    }
};
class FileSystem {
    public: 
        string currentPath; 
        Directory* root;
    ~FileSystem(){ delete root;}
    FileSystem(): currentPath("/"),root(new Directory()) {}

    vector<string> getAbsolutePath(string  path) { // asume correct;
        vector<string> directorNames; 
        if(path.size() <= 0) {  // .
            return directorNames;
        }
        if(path[0] == '/' and path.size() == 1){ // cd /
            return directorNames;
        } 
        // atleast 2 size 
        int ix = -1;
        if(path[0] == '.' && path[1] == '/') {
            directorNames = getAbsolutePath(currentPath);
            ix = 2;
        }  else if(path[0] =='/'){ // /aa
            ix = 1;
        } else { //../ 
            ix = 0;
        }
        string temp = "";
        while(ix<path.size()) {
            if(path[ix] == '/'){
                if(temp == ".."){ // assumption path is correct.
                    directorNames.pop_back();
                } else {
                    directorNames.push_back(temp);
                }
                ix++;
                temp = "";
            } else {
                temp+=path[ix];
                ix++;
            }
        }   
        if(temp.size()){
            directorNames.push_back(temp);
        }

        for(auto x: directorNames) {
            cout<<x<<" ";
        }
        cout<<endl;
        return directorNames;
    } 
    bool validate(string path) {
        return true;
    }
    void mkdir(string path) {  // no regex in mkdir
        if(!validate(path))return;
        vector<string> ap = getAbsolutePath(path); // without .. , have regex
        root -> Insert(ap);
    }
    string pwd() {
        return currentPath;
    }
    void cd(string path){
        if(!validate(path))return;
        vector<string> ap = getAbsolutePath(path); // without .. , have regex
        currentPath = root -> validateAndMove(ap, currentPath);
    }
};
int main() {
    FileSystem fs = FileSystem();
    // cout<<fs.pwd()<<endl;
    // fs.mkdir("./a/../c/da/fdasd/ffa/fdsaf/../fd");
    // fs.cd("/c/da/f???????.*");
    // cout<<fs.pwd()<<endl;

    fs.mkdir("/a/b1/c");
    fs.mkdir("/a/batch/d");  
    fs.mkdir("/x/test");

    fs.cd("a/b.*");    // Regex matches "b1" or "batch" → cd to /a/b1 (first match)
    cout<<fs.pwd()<<endl;          // Returns "/a/b1"

    // cout<<isMatching("batch","b.*h")<<endl;
        

    return 0; 
}




/*
review by gpt: 
1. not thread etc event after telling
2. memory leak in directory. 
3. do dfs search consider -> it is in problem 2 description



#include <bits/stdc++.h>
#include <regex>
#include <mutex>

using namespace std;


class Directory {
    public:
        string name;
        Directory* parent;
        unordered_map<string, Directory*> children;
        mutable mutex mtx;
    
        Directory(string name, Directory* parent = nullptr)
            : name(name), parent(parent) {}
    };
    
    
    class FileSystem {
    private:
        Directory* root;
        Directory* current;
    
    
        vector<string> tokenize(const string& path) {
            vector<string> tokens;
            string token;
            stringstream ss(path);
            while (getline(ss, token, '/')) {
                if (!token.empty())
                    tokens.push_back(token);
            }
            return tokens;
        }
    
        Directory* getStartDir(const string& path) {
            return (!path.empty() && path[0] == '/') ? root : current;
        }
    
        Directory* dfsRegexMatch(Directory* node, const regex& pattern) {
            for (auto& [name, child] : node->children) {
                if (regex_match(child->name, pattern))
                    return child;
            }
            for (auto& [name, child] : node->children) {
                Directory* res = dfsRegexMatch(child, pattern);
                if (res) return res;
            }
            return nullptr;
        }
    
    public:
        FileSystem() {
            root = new Directory("/");
            current = root;
        }
    
        
    
        void mkdir(const string& path) {
            Directory* node = getStartDir(path);
            auto tokens = tokenize(path);
    
            for (const string& part : tokens) {
                if (part == "." || part == "..") continue;
    
                lock_guard<mutex> lock(node->mtx);
                if (!node->children.count(part)) {
                    node->children[part] = new Directory(part, node);
                }
                node = node->children[part];
            }
        }
    
        
    
        void cd(const string& path) {
            Directory* node = getStartDir(path);
            auto tokens = tokenize(path);
    
            for (const string& part : tokens) {
                if (part == ".") continue;
    
                if (part == "..") {
                    if (!node->parent)
                        throw runtime_error("Already at root");
                    node = node->parent;
                } else {
                    regex pattern(part);
                    Directory* match = dfsRegexMatch(node, pattern);
                    if (!match)
                        throw runtime_error("Directory not found: " + part);
                    node = match;
                }
            }
            current = node;
        }
    
        
    
        string pwd() {
            vector<string> path;
            Directory* node = current;
    
            while (node && node != root) {
                path.push_back(node->name);
                node = node->parent;
            }
    
            reverse(path.begin(), path.end());
    
            string result = "/";
            for (int i = 0; i < path.size(); i++) {
                result += path[i];
                if (i + 1 < path.size()) result += "/";
            }
            return result;
        }
    };
    

    
    int main() {
        FileSystem fs;
    
        fs.mkdir("/a/b1/c");
        fs.mkdir("/a/batch/d");
        fs.mkdir("/x/test");
    
        fs.cd("a/b.*");
        cout << fs.pwd() << endl;   // /a/b1
    
        fs.cd("te?t");
        cout << fs.pwd() << endl;   // /x/test
    
        fs.cd("../b.*");
        cout << fs.pwd() << endl;   // /a/b1
    
        return 0;
    }
    
*/




