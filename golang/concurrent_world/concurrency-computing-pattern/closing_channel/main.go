package main

import (
	"fmt"
)

func main() {
	// unbufferd channel causes Panic without go routine, or if some one is not listening
	// uch := make(chan int)

	uch := make(chan int, 1)
	uch2 := make(chan int)

	uch <- 5

	close(uch)

	val, ok := <-uch
	val2, ok := <-uch

	fmt.Println(val)  // 5
	fmt.Println(val2) // 0
	fmt.Println(ok)   // false

	close(uch2)

	val3, ok := <-uch2

	fmt.Println(val3) // 5
	fmt.Println(ok)   // false

}
