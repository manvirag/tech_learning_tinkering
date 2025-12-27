#include <iostream>
#include <vector>
#include <algorithm>
#include <thread>
#include <mutex>
#include <chrono>
#include <functional>
#include <string>
#include <numeric>
#include <memory>
using namespace std;

/*
================================================================================
LAMBDA FUNCTIONS IN C++
================================================================================

Lambdas were introduced in C++11 (2011)

Lambda Syntax:
[ capture-list ] ( parameters ) -> return-type { body }

Simplified Syntax:
[ capture-list ] ( parameters ) { body }  // Return type deduced automatically

Key Components:
1. Capture List: Specifies which variables from the enclosing scope are captured
   - []        : Capture nothing
   - [=]       : Capture all by value
   - [&]       : Capture all by reference
   - [x]       : Capture x by value
   - [&x]      : Capture x by reference
   - [=, &x]   : Capture all by value, but x by reference
   - [&, x]    : Capture all by reference, but x by value

2. Parameters: Function parameters (optional)
3. Return Type: Optional, can be deduced
4. Body: Function body

================================================================================
*/

// ============================================================================
// Example 1: Basic Lambda Syntax
// ============================================================================
void example1_BasicLambda() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 1: Basic Lambda Syntax" << endl;
    cout << string(70, '=') << endl;
    
    // Simple lambda with no captures
    auto greet = []() {
        cout << "Hello from lambda!" << endl;
    };
    
    greet();
    
    // Lambda with parameters
    auto add = [](int a, int b) {
        return a + b;
    };
    
    cout << "5 + 3 = " << add(5, 3) << endl;
    
    // Lambda with explicit return type
    auto multiply = [](int a, int b) -> int {
        return a * b;
    };
    
    cout << "4 * 7 = " << multiply(4, 7) << endl;
    
    cout << "\n✓ Basic lambda examples completed\n" << endl;
}

// ============================================================================
// Example 2: Capture by Value [=]
// ============================================================================
void example2_CaptureByValue() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 2: Capture by Value [=]" << endl;
    cout << string(70, '=') << endl;
    
    int x = 10;
    int y = 20;
    
    // Capture all variables by value
    auto lambda1 = [=]() {
        cout << "Inside lambda: x = " << x << ", y = " << y << endl;
        // x = 100;  // ERROR: Cannot modify captured by-value variables
    };
    
    lambda1();
    cout << "Outside lambda: x = " << x << ", y = " << y << endl;
    
    // Capture specific variables by value
    auto lambda2 = [x, y]() {
        cout << "Captured x = " << x << ", y = " << y << endl;
    };
    
    lambda2();
    
    cout << "\n✓ Capture by value examples completed\n" << endl;
}

// ============================================================================
// Example 3: Capture by Reference [&]
// ============================================================================
void example3_CaptureByReference() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 3: Capture by Reference [&]" << endl;
    cout << string(70, '=') << endl;
    
    int x = 10;
    int y = 20;
    
    // Capture all variables by reference
    auto lambda1 = [&]() {
        x = 100;  // Can modify because captured by reference
        y = 200;
        cout << "Inside lambda: x = " << x << ", y = " << y << endl;
    };
    
    cout << "Before lambda: x = " << x << ", y = " << y << endl;
    lambda1();
    cout << "After lambda: x = " << x << ", y = " << y << endl;
    
    // Capture specific variables by reference
    int z = 5;
    auto lambda2 = [&x, &z]() {
        x = 999;
        z = 888;
        cout << "Modified x = " << x << ", z = " << z << endl;
    };
    
    lambda2();
    cout << "After lambda2: x = " << x << ", z = " << z << endl;
    
    cout << "\n✓ Capture by reference examples completed\n" << endl;
}

// ============================================================================
// Example 4: Mixed Capture [=, &x] or [&, x]
// ============================================================================
void example4_MixedCapture() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 4: Mixed Capture" << endl;
    cout << string(70, '=') << endl;
    
    int a = 1, b = 2, c = 3, d = 4;
    
    // Capture all by value, but 'a' by reference
    auto lambda1 = [=, &a]() {
        a = 100;  // Can modify 'a' because captured by reference
        cout << "a = " << a << ", b = " << b << ", c = " << c << ", d = " << d << endl;
        // b = 200;  // ERROR: Cannot modify 'b' (captured by value)
    };
    
    lambda1();
    cout << "After lambda1: a = " << a << " (modified), b = " << b << " (unchanged)" << endl;
    
    // Capture all by reference, but 'c' by value
    auto lambda2 = [&, c]() {
        a = 200;
        b = 300;
        d = 400;
        cout << "a = " << a << ", b = " << b << ", c = " << c << ", d = " << d << endl;
        // c = 500;  // ERROR: Cannot modify 'c' (captured by value)
    };
    
    lambda2();
    cout << "After lambda2: a = " << a << ", b = " << b << ", c = " << c << ", d = " << d << endl;
    
    cout << "\n✓ Mixed capture examples completed\n" << endl;
}

// ============================================================================
// Example 5: Mutable Lambda
// ============================================================================
void example5_MutableLambda() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 5: Mutable Lambda" << endl;
    cout << string(70, '=') << endl;
    cout << "mutable keyword allows modifying captured-by-value variables\n";
    cout << "Note: Changes are local to the lambda, don't affect original\n" << endl;
    
    int x = 10;
    
    // Without mutable - cannot modify captured by-value variables
    auto lambda1 = [x]() {
        // x = 20;  // ERROR: Cannot modify
        cout << "x = " << x << endl;
    };
    
    // With mutable - can modify, but changes are local
    auto lambda2 = [x]() mutable {
        x = 20;  // OK: Can modify, but original 'x' is unchanged
        cout << "Inside lambda: x = " << x << endl;
    };
    
    cout << "Before lambda: x = " << x << endl;
    lambda2();
    cout << "After lambda: x = " << x << " (unchanged)" << endl;
    
    // Counter example with mutable
    auto counter = [count = 0]() mutable {
        return ++count;
    };
    
    cout << "\nCounter lambda (mutable):" << endl;
    for (int i = 0; i < 5; i++) {
        cout << "Call " << (i + 1) << ": " << counter() << endl;
    }
    
    cout << "\n✓ Mutable lambda examples completed\n" << endl;
}

// ============================================================================
// Example 6: Lambda with STL Algorithms
// ============================================================================
void example6_LambdaWithSTL() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 6: Lambda with STL Algorithms" << endl;
    cout << string(70, '=') << endl;
    
    vector<int> numbers = {5, 2, 8, 1, 9, 3, 7, 4, 6};
    
    cout << "Original vector: ";
    for (int n : numbers) cout << n << " ";
    cout << endl;
    
    // Sort using lambda
    sort(numbers.begin(), numbers.end(), [](int a, int b) {
        return a > b;  // Sort in descending order
    });
    
    cout << "Sorted (descending): ";
    for (int n : numbers) cout << n << " ";
    cout << endl;
    
    // Find elements greater than 5
    auto it = find_if(numbers.begin(), numbers.end(), [](int n) {
        return n > 5;
    });
    
    if (it != numbers.end()) {
        cout << "First number > 5: " << *it << endl;
    }
    
    // Count even numbers
    int evenCount = count_if(numbers.begin(), numbers.end(), [](int n) {
        return n % 2 == 0;
    });
    cout << "Even numbers count: " << evenCount << endl;
    
    // Transform: square each number
    vector<int> squared;
    transform(numbers.begin(), numbers.end(), back_inserter(squared), [](int n) {
        return n * n;
    });
    
    cout << "Squared: ";
    for (int n : squared) cout << n << " ";
    cout << endl;
    
    // For each: print with prefix
    cout << "For each: ";
    for_each(numbers.begin(), numbers.end(), [](int n) {
        cout << "[" << n << "] ";
    });
    cout << endl;
    
    cout << "\n✓ Lambda with STL examples completed\n" << endl;
}

// ============================================================================
// Example 7: Lambda in Threads
// ============================================================================
void example7_LambdaInThreads() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 7: Lambda in Threads" << endl;
    cout << string(70, '=') << endl;
    
    int sharedValue = 0;
    mutex mtx;
    
    // Lambda in thread with capture
    thread t1([&sharedValue, &mtx]() {
        for (int i = 0; i < 5; i++) {
            lock_guard<mutex> lock(mtx);
            sharedValue += 10;
            cout << "Thread 1: sharedValue = " << sharedValue << endl;
            this_thread::sleep_for(chrono::milliseconds(100));
        }
    });
    
    thread t2([&sharedValue, &mtx]() {
        for (int i = 0; i < 5; i++) {
            lock_guard<mutex> lock(mtx);
            sharedValue += 20;
            cout << "Thread 2: sharedValue = " << sharedValue << endl;
            this_thread::sleep_for(chrono::milliseconds(100));
        }
    });
    
    t1.join();
    t2.join();
    
    cout << "\nFinal sharedValue: " << sharedValue << endl;
    cout << "✓ Lambda in threads examples completed\n" << endl;
}

// ============================================================================
// Example 8: Lambda as Function Parameter
// ============================================================================
void example8_LambdaAsParameter() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 8: Lambda as Function Parameter" << endl;
    cout << string(70, '=') << endl;
    
    // Function that accepts a lambda
    auto applyOperation = [](int a, int b, function<int(int, int)> op) {
        return op(a, b);
    };
    
    // Using lambda as parameter
    int result1 = applyOperation(10, 5, [](int x, int y) { return x + y; });
    cout << "10 + 5 = " << result1 << endl;
    
    int result2 = applyOperation(10, 5, [](int x, int y) { return x * y; });
    cout << "10 * 5 = " << result2 << endl;
    
    // Template function with lambda
    auto process = [](vector<int>& vec, function<void(int&)> processor) {
        for (auto& v : vec) {
            processor(v);
        }
    };
    
    vector<int> numbers = {1, 2, 3, 4, 5};
    cout << "Before: ";
    for (int n : numbers) cout << n << " ";
    cout << endl;
    
    process(numbers, [](int& n) { n *= 2; });
    
    cout << "After (doubled): ";
    for (int n : numbers) cout << n << " ";
    cout << endl;
    
    cout << "\n✓ Lambda as parameter examples completed\n" << endl;
}

// ============================================================================
// Example 9: Lambda with Initialization Capture (C++14)
// ============================================================================
void example9_InitCapture() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 9: Initialization Capture [x = expr] (C++14)" << endl;
    cout << string(70, '=') << endl;
    
    int x = 10;
    
    // Initialize capture: create new variable in lambda
    auto lambda1 = [y = x * 2]() {
        cout << "y (initialized as x*2) = " << y << endl;
    };
    
    lambda1();
    
    // Move capture
    string str = "Hello";
    auto lambda2 = [moved_str = move(str)]() {
        cout << "Moved string: " << moved_str << endl;
    };
    
    lambda2();
    cout << "Original string after move: \"" << str << "\" (empty)" << endl;
    
    // Unique pointer capture
    auto lambda3 = [ptr = make_unique<int>(42)]() {
        cout << "Captured unique_ptr value: " << *ptr << endl;
    };
    
    lambda3();
    
    cout << "\n✓ Initialization capture examples completed\n" << endl;
}

// ============================================================================
// Example 10: Lambda Returning Lambda (Higher-Order Function)
// ============================================================================
void example10_LambdaReturningLambda() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 10: Lambda Returning Lambda" << endl;
    cout << string(70, '=') << endl;
    
    // Function that returns a lambda
    auto makeMultiplier = [](int factor) {
        return [factor](int n) {
            return n * factor;
        };
    };
    
    auto multiplyBy2 = makeMultiplier(2);
    auto multiplyBy5 = makeMultiplier(5);
    
    cout << "3 * 2 = " << multiplyBy2(3) << endl;
    cout << "3 * 5 = " << multiplyBy5(3) << endl;
    
    // Lambda that returns lambda with different operations
    // Note: Must use std::function because each lambda has a unique type
    auto makeOperation = [](string op) -> function<int(int, int)> {
        if (op == "add") {
            return [](int a, int b) { return a + b; };
        } else if (op == "multiply") {
            return [](int a, int b) { return a * b; };
        } else {
            return [](int a, int b) { return a - b; };
        }
    };
    
    auto addOp = makeOperation("add");
    auto mulOp = makeOperation("multiply");
    
    cout << "10 + 5 = " << addOp(10, 5) << endl;
    cout << "10 * 5 = " << mulOp(10, 5) << endl;
    
    cout << "\n✓ Lambda returning lambda examples completed\n" << endl;
}

// ============================================================================
// Example 11: Lambda with Generic Parameters (C++14)
// ============================================================================
void example11_GenericLambda() {
    cout << "\n" << string(70, '=') << endl;
    cout << "EXAMPLE 11: Generic Lambda (C++14)" << endl;
    cout << string(70, '=') << endl;
    
    // Generic lambda with auto parameters
    auto add = [](auto a, auto b) {
        return a + b;
    };
    
    cout << "5 + 3 = " << add(5, 3) << endl;
    cout << "5.5 + 3.2 = " << add(5.5, 3.2) << endl;
    cout << "string + string = " << add(string("Hello "), string("World")) << endl;
    
    // Generic lambda in algorithm
    vector<int> intVec = {1, 2, 3};
    vector<double> doubleVec = {1.1, 2.2, 3.3};
    
    auto print = [](auto& vec) {
        for (auto& v : vec) {
            cout << v << " ";
        }
        cout << endl;
    };
    
    cout << "Int vector: ";
    print(intVec);
    
    cout << "Double vector: ";
    print(doubleVec);
    
    cout << "\n✓ Generic lambda examples completed\n" << endl;
}

int main() {
    cout << "\n" << string(70, '=') << endl;
    cout << "LAMBDA FUNCTIONS IN C++ - COMPREHENSIVE EXAMPLES" << endl;
    cout << string(70, '=') << endl;
    cout << "\nLambdas were introduced in C++11 (2011)" << endl;
    cout << "Enhanced in C++14 with generic lambdas and initialization capture" << endl;
    cout << "Further enhanced in C++17 and C++20\n" << endl;
    
    example1_BasicLambda();
    example2_CaptureByValue();
    example3_CaptureByReference();
    example4_MixedCapture();
    example5_MutableLambda();
    example6_LambdaWithSTL();
    example7_LambdaInThreads();
    example8_LambdaAsParameter();
    example9_InitCapture();
    example10_LambdaReturningLambda();
    example11_GenericLambda();
    
    cout << "\n" << string(70, '=') << endl;
    cout << "SUMMARY:" << endl;
    cout << string(70, '=') << endl;
    cout << "Lambda Syntax: [capture](params) -> return-type { body }" << endl;
    cout << "\nCapture Options:" << endl;
    cout << "  []        - Capture nothing" << endl;
    cout << "  [=]       - Capture all by value" << endl;
    cout << "  [&]       - Capture all by reference" << endl;
    cout << "  [x]       - Capture x by value" << endl;
    cout << "  [&x]      - Capture x by reference" << endl;
    cout << "  [=, &x]   - Capture all by value, x by reference" << endl;
    cout << "  [&, x]    - Capture all by reference, x by value" << endl;
    cout << "  [x = expr] - Initialize capture (C++14)" << endl;
    cout << "\nKey Features:" << endl;
    cout << "  - Introduced in C++11" << endl;
    cout << "  - Generic lambdas (auto params) in C++14" << endl;
    cout << "  - Perfect for STL algorithms" << endl;
    cout << "  - Great for callbacks and event handlers" << endl;
    cout << "  - Useful in multithreading" << endl;
    cout << "  - Can be stored in std::function" << endl;
    cout << string(70, '=') << endl;
    
    return 0;
}

