package main

import "fmt"

var check int = 55

func main() {
	var name = "anubhav"
	// fmt.Println(name)
	_ = name

	a, b := "anubhav", 22
	fmt.Println(a, b)

	//atleast one new varible should be here can be at any possition
	c, a := "gupta", "new variable"
	fmt.Println(a, c)

	// multiple delaration
	var (
		p int
		q int
		r int
	)

	var i, j int
	fmt.Println(p, q, r)
	fmt.Println(i, j)

	// multiple assigning
	i, j = 5, 6

	// swap
	i, j = j, i

}
