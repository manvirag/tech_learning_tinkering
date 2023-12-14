/*


In Go, a panic is a run-time error that occurs when the program is in an unrecoverable state.
It could be due to a division by zero, out-of-bounds array access,
or other situations that indicate a serious issue. When a panic occurs,
the normal flow of the program is disrupted, and the program terminates with an error message.


To handle panics and potentially recover from them, Go provides the recover function.
The recover function is only useful when called inside a deferred function.
It allows you to capture and handle panics, preventing them from causing the program to crash.


https://go.dev/blog/defer-panic-and-recover

Panic is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.
Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
*/

package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

//The program will output:
//
//Calling g.
//Printing in g 0
//Printing in g 1
//Printing in g 2
//Printing in g 3
//Panicking!
//Defer in g 3
//Defer in g 2
//Defer in g 1
//Defer in g 0
//Recovered in f 4
//Returned normally from f.
//If we remove the deferred function from f the panic is not recovered and reaches the top of the goroutine’s call stack, terminating the program. This modified program will output:
//
//Calling g.
//Printing in g 0
//Printing in g 1
//Printing in g 2
//Printing in g 3
//Panicking!
//Defer in g 3
//Defer in g 2
//Defer in g 1
//Defer in g 0
//panic: 4
//
//panic PC=0x2a9cd8
//[stack trace omitted]
