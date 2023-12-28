package main

import "fmt"

func main() {
	normalPizza := Pizza{}
	fmt.Println(normalPizza.Price())

	capsicumPizza := CapsicumTopping{pizza: normalPizza}
	fmt.Println(capsicumPizza.Price())

	capsicumOnionPizza := OnionTopping{pizza: capsicumPizza}
	fmt.Println(capsicumOnionPizza.Price())

	doubleCapsicumOnion := CapsicumTopping{pizza: capsicumOnionPizza}
	fmt.Println(doubleCapsicumOnion.Price())
}
