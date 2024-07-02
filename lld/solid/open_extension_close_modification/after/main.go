package main

import "fmt"

type Details interface {
	printName()
}

type Cow struct {
	name string
}

func (c *Cow) printName() {

	fmt.Println("i am cow")
}

type Dog struct {
	name string
}

func (c *Dog) printName() {
	fmt.Println("i am dog")
}

func facade(a Details) {
	fmt.Print("%T", a)
	a.printName()
}

func main() {
	cow := &Cow{
		name: "abcd",
	}
	facade(cow)
}
