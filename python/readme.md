# Python Argument Passing (Pass by Value vs Reference)

## 🧠 Core Concept

Python does **NOT** have traditional pass-by-value or pass-by-reference like C/Go.

**Python uses:**

> **Pass-by-object-reference (pass-by-assignment)**

---

## 🔑 Key Rules

1. Variables are **references (names)** to objects
2. Functions receive **references to the same object**
3. Behavior depends on:
   - **Mutable vs Immutable**
   - **Mutation vs Reassignment**

---

## 📊 Data Types Classification

### 🔒 Immutable Types (Value-like behavior)

- `int`
- `float`
- `str`
- `tuple`
- `bool`

**✔** Cannot be modified in-place  
**✔** Any change creates a new object  

---

### 🔓 Mutable Types (Reference-like behavior)

- `list`
- `dict`
- `set`
- `bytearray`
- custom objects

**✔** Can be modified in-place  

---

## 🧪 Immutable Example (Pass-by-value behavior)

```python
def func(x):
    x = x + 1

a = 10
func(a)
print(a)   # 10
```

**✔** Original value unchanged  
**✔** New object created

## 🧪 Mutable Example (Pass-by-reference behavior)

```python
def func(lst):
    lst.append(4)

a = 
func(a)
print(a)   # 
```

**✔** Same object modified

## ⚠️ Reassignment vs Mutation

### ❌ Reassignment (NO effect outside)

```python
def func(lst):
    lst = 

a = 
func(a)
print(a)   # 
```

### ✅ Mutation (affects original)

```python
def func(lst):
    lst.append(100)

a = 
func(a)
print(a)   # 
```

## 🎯 Forcing Pass-by-Value Behavior

**👉 Create a copy**

**List:**
- `lst.copy()`
- `lst[:]`

**Dict:**
- `d.copy()`

**Set:**
- `s.copy()`

**Deep Copy (nested objects):**
```python
import copy
copy.deepcopy(obj)
```

## 🎯 Forcing Pass-by-Reference Behavior

**👉 Mutate the object directly**

```python
def func(d):
    d["key"] = "value"
```

## 📌 Special Cases

### Tuple with mutable element

```python
t = (1, 2, )
t.append(5)
print(t)   # (1, 2, )
```

**✔** Tuple is immutable  
**✔** But inner objects can be mutable

### String (immutable)

```python
def func(s):
    s += " world"

a = "hello"
func(a)
print(a)   # hello
```

### Custom Object

```python
class A:
    def __init__(self):
        self.x = 10

def func(obj):
    obj.x = 20

a = A()
func(a)
print(a.x)   # 20
```

**✔** Objects are mutable by default

## 📊 Summary Table

| Type    | Mutable | Behavior       |
|---------|---------|----------------|
| `int`   | ❌      | value-like     |
| `float` | ❌      | value-like     |
| `str`   | ❌      | value-like     |
| `tuple` | ❌      | value-like*    |
| `list`  | ✅      | reference-like |
| `dict`  | ✅      | reference-like |
| `set`   | ✅      | reference-like |
| `object`| ✅      | reference-like |
