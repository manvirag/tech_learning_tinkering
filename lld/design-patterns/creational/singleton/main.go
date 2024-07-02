package main

import "fmt"

func main() {

	firstInstance := GetInstance()
	count := firstInstance.AddOne()
	fmt.Println(count)
	secondInstance := GetInstance()
	count = secondInstance.AddOne()
	fmt.Println(count)

	// 1
	// 2
}
