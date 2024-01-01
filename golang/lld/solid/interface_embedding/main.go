package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type human interface {
	animal
	speak()
}

type student struct {
	name string
}

func (e student) breathe() {
	fmt.Println("student breathes")
}

func (e student) walk() {
	fmt.Println("student walk")
}

func (e student) speak() {
	fmt.Println("student speaks")
}

func facade(hum human) {
	hum.breathe()
	hum.walk()
	hum.speak()

}
func main() {
	var h human
	h = student{name: "Aman Thakur"}
	facade(h)
}
