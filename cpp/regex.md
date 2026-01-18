# C++ Regex - Quick Reference

## Setup

```cpp
#include <regex>
#include <string>
using namespace std;
```

---

## 1. Match

```cpp
string text = "Hello World";
regex pattern("Hello");
bool matches = regex_search(text, pattern);
```

---

## 2. Extract Match

```cpp
string text = "Price: $100";
regex pattern(R"(\$(\d+))");
smatch matches;
if (regex_search(text, matches, pattern)) {
    string full = matches[0];  // "$100"
    string num = matches[1];   // "100"
}
```

---

## 3. Find All Matches

```cpp
string text = "cat bat hat";
regex pattern(R"(\w+at)");
sregex_iterator iter(text.begin(), text.end(), pattern);
for (; iter != sregex_iterator(); ++iter) {
    cout << iter->str() << endl;
}
```

---

## 4. Replace

```cpp
string text = "Hello World";
regex pattern("World");
string result = regex_replace(text, pattern, "C++");
```

---

## 5. Match Entire String

```cpp
string text = "12345";
regex pattern(R"(\d+)");
bool full_match = regex_match(text, pattern);
```

---

## 6. Common Patterns

```cpp
regex digits(R"(\d+)");
regex word(R"(\w+)");
regex whitespace(R"(\s+)");
```

---

## 7. Metacharacters

```cpp
\d      // digit
\w      // word character
\s      // whitespace
.       // any character
*       // zero or more
+       // one or more
?       // zero or one
{n}     // exactly n times
^       // start
$       // end
[abc]   // any of a,b,c
```

---

## 8. Raw String (R"()")

```cpp
regex pattern1("\\d+");      // need escape
regex pattern2(R"(\d+)");    // no escape needed
```
