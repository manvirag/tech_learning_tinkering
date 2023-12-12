package main

import "fmt"

type iBase interface {
	sayHi()
}

type base struct {
	color string
}

type child struct {
	base  // embedding
	style string
}

func (b *base) sayHi() {
	fmt.Println("Hi. What's up?")
}

func check(b iBase) {
	b.sayHi()
}

func main() {
	base := base{color: "yellow"}
	child := &child{
		base:  base,
		style: "oook",
	}
	child.sayHi()
	fmt.Printf("Child color is %s and style is %s\n", child.color, child.style)
	check(child)
	check(&base)
}
