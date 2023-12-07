package main

import "fmt"

func printName(key string) {
	if key == "Dog" {
		fmt.Println("i am dog")
	} else if key == "Cow" {
		fmt.Println("i am cow")
	}
}

func main() {
	printName("Dog")
	printName("Cow")
}
