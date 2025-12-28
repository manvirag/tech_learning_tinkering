/*

This problem asks you to design an in-memory file system data structure that supports basic file and directory operations.

The file system should support four main operations:

ls(path): List the contents at a given path

If the path points to a file, return a list containing only that file's name
If the path points to a directory, return all files and subdirectories within it in lexicographic (alphabetical) order

=> extract path
-> go word by word 
-> if its file -> return faile name  -> check via flag 
-> elss go to each child and take its name reutrn ; 


Assumptions: 
- path existing always and format is correct. 



mkdir(path): Create a new directory
The path given will not already exist
If intermediate directories in the path don't exist, create them automatically (similar to mkdir -p in Unix)

-> extract all word 
-> if exist go , else create and go . 

Assumption: 
- path format is correct. 

addContentToFile(filePath, content): Add content to a file
If the file doesn't exist, create it with the given content
If the file already exists, append the new content to the existing content

=> go to n-1 word 
-> if nth word not exit -> create child as file node. 
-> else append. 

Assumption: 
- path alwasy existing before that file 
    - like /a/b/c.txt -> then a -> b existing , its just that file doesn't exist. 
- path format is correct. 



readContentFromFile(filePath): Read and return all content from a file
=> extract the path - >traveser through thrie -> when you reach last word -> return content of file. 

/ path is legit and correct and exist.


................................................

path -> file  
or directory -> contain files and directory

content -> can assume its list of string , become to a file. 

path -> /a/b/c/d 

extract -> a, b, c, d 

Assumption: 
root path -> / 


*/



#include<iostream> 
#include<string> 
#include<map> 
using namespace std;


class FSTrie {
    public: 
        bool isFile; 
        string name;
        vector<string> content; 
        map<string, FSTrie*> child; 
    FSTrie(){}
    FSTrie(string name, bool isFile) {
        this -> name = name; 
        this -> isFile = isFile; 
    }
    
    void insert(vector<string> words, bool isFile) {
        FSTrie *temp = this; 
        for(int i=0;i<words.size()-1;i++) {

            
            if(temp->child.find(words[i]) == temp -> child.end()) {
                temp -> child[words[i]] = new FSTrie();
                temp -> child[words[i]]->name = words[i];
                temp -> child[words[i]]->isFile = false;
            } 
            temp = temp -> child[words[i]];

        }
        if(words.size()> 0) {
            int lix = words.size()-1;
            if(temp->child.find(words[lix]) == temp -> child.end()) {
                temp -> child[words[lix]] = new FSTrie();
                temp -> child[words[lix]]->name = words[lix];
                temp -> child[words[lix]]->isFile = isFile;
            } 
        }
    }
    vector<string> search(vector<string> words) {
            FSTrie* temp = this; 
            vector<string> names;
            for(int i=0;i<words.size();i++) {
                if(temp->child.find(words[i]) == temp -> child.end()) {
                    return names; 
                } 
                temp = temp -> child[words[i]];
            }
            if(temp -> isFile) {
                names.push_back(temp->name);
                return names; 
            }

            for(auto dir: temp -> child) {
                names.push_back(dir.first);
            }
            return names; 

    }
    FSTrie* getNode(vector<string> words) {
            FSTrie* temp = this; 
            vector<string> names;
            for(int i=0;i<words.size();i++) {
                temp = temp -> child[words[i]];
            }
            return temp;

    }
    vector<string> getContents(vector<string> words){
            FSTrie* temp = this -> getNode(words);
            vector<string> c= temp -> content; 
            // for(auto x: c) cout<<x<<" ";
            // cout<<endl;
            return c;

    }
    void addContent(vector<string> words, string c) {
        this -> insert(words, true);
        FSTrie* temp =this -> getNode(words);
        temp -> content.push_back(c);
    }
};
class FileSystem {
    private: 
      FSTrie * root; 
      vector<string> extractPath(string path ) {
        vector<string> words;
        string temp = "";
        for(int i =1;i<path.size();i++) {

            if(path[i] == '/') {
                words.push_back(temp);
                // cout<<temp<<endl;
                temp = "";
            } else {
                temp+=path[i];
            }
        }
        if(temp.size()>0)  words.push_back(temp);
        // for(auto x: words){
        //     cout<<x<<" ";
        // }
        // cout<<endl;
        return words; 

        
      }

    public:

    FileSystem() {
        this -> root = new FSTrie();
    }

    void mkdir(string path) {
        
        vector<string> words = this -> extractPath(path);
        this -> root -> insert(words, false); // directory -> false

    }

    vector<string> ls(string path) {
        vector<string> words = this -> extractPath(path);
        vector<string> names = this -> root -> search(words); 
        for(auto x: names ) cout<<x<<" ";
        cout<<endl;

        return names;

    }

    void addContentToFile(string path, string content) {
        this -> root -> addContent(this->extractPath(path), content);
    }

    vector<string> readContentFromFile(string path ){
        vector<string> words = this -> extractPath(path);
        vector<string> contents = this -> root -> getContents(words);
        for(auto x: contents)cout<<x<<" ";
        cout<<endl;
        return contents;
    }




};


int main() {

    FileSystem * fs = new FileSystem();
    fs -> mkdir("/a/b");
    fs -> mkdir("/a/b/d/p");
    vector<string> names = fs -> ls("/a/b");
    // fs -> addContentToFile("/a/b/c", "cccfilehai");
    // fs -> addContentToFile("/a/b/c", "fdsa");
    // fs -> addContentToFile("/a/b/c", "fdasfads");
    // vector<string> contents = fs -> readContentFromFile("/a/b/c");
    return 0;
}




/*

usiang hashmap:


public class FileSystem {
    class Dir {
        HashMap < String, Dir > dirs = new HashMap < > ();
        HashMap < String, String > files = new HashMap < > ();
    }
    Dir root;
    public FileSystem() {
        root = new Dir();
    }
    public List < String > ls(String path) {
        Dir t = root;
        List < String > files = new ArrayList < > ();
        if (!path.equals("/")) {
            String[] d = path.split("/");
            for (int i = 1; i < d.length - 1; i++) {
                t = t.dirs.get(d[i]);
            }
            if (t.files.containsKey(d[d.length - 1])) {
                files.add(d[d.length - 1]);
                return files;
            } else {
                t = t.dirs.get(d[d.length - 1]);
            }
        }
        files.addAll(new ArrayList < > (t.dirs.keySet()));
        files.addAll(new ArrayList < > (t.files.keySet()));
        Collections.sort(files);
        return files;
    }
public void mkdir(String path) {
        Dir t = root;
        String[] d = path.split("/");
        for (int i = 1; i < d.length; i++) {
            if (!t.dirs.containsKey(d[i]))
                t.dirs.put(d[i], new Dir());
            t = t.dirs.get(d[i]);
        }
    }
public void addContentToFile(String filePath, String content) {
        Dir t = root;
        String[] d = filePath.split("/");
        for (int i = 1; i < d.length - 1; i++) {
            t = t.dirs.get(d[i]);
        }
        t.files.put(d[d.length - 1], t.files.getOrDefault(d[d.length - 1], "") + content);
    }
public String readContentFromFile(String filePath) {
        Dir t = root;
        String[] d = filePath.split("/");
        for (int i = 1; i < d.length - 1; i++) {
            t = t.dirs.get(d[i]);
        }
        return t.files.get(d[d.length - 1]);
    }
}

 * Your FileSystem object will be instantiated and called as such:
 * FileSystem obj = new FileSystem();
 * List<String> param_1 = obj.ls(path);
 * obj.mkdir(path);
 * obj.addContentToFile(filePath,content);
 * String param_4 = obj.readContentFromFile(filePath);
 */

