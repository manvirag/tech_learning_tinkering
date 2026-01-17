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

For cd(String path), implement directory navigation with regex matching on directory names (not full path). Change the current working directory to the first matching directory found via DFS traversal from root/current position.

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




*/




