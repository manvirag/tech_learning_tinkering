package main

import "fmt"

// Animal interface with multiple methods
type Animal interface {
	Eat()
	Sleep()
	Move()
}

// Dog type implementing Animal interface
type Dog struct{}

func (d *Dog) Eat() {
	fmt.Println("Dog is eating")
}

func (d *Dog) Sleep() {
	fmt.Println("Dog is sleeping")
}

func (d *Dog) Move() {
	fmt.Println("Dog is moving")
}

// Bird type implementing Animal interface
type Bird struct{}

func (b *Bird) Eat() {
	fmt.Println("Bird is eating")
}

func (b *Bird) Sleep() {
	fmt.Println("Bird is sleeping")
}

func (b *Bird) Move() {
	fmt.Println("Bird is flying")
}

func main() {
	// Client code
	dog := &Dog{}
	bird := &Bird{}

	// Using the Animal interface methods
	dog.Eat()
	dog.Sleep()
	dog.Move()

	bird.Eat()
	bird.Sleep()
	bird.Move()
}

// In this example, both Dog and Bird implement the Animal interface, which includes methods like Eat, Sleep, and Move. However, this violates the Interface Segregation Principle because not all animals can perform all actions. For example, a Bird cannot move the same way a Dog does.
