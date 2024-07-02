package main

import "fmt"

// Eater interface
type Eater interface {
	Eat()
}

// Sleeper interface
type Sleeper interface {
	Sleep()
}

// Mover interface
type Mover interface {
	Move()
}

// Dog type implementing Eater, Sleeper, and Mover interfaces
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

// Bird type implementing Eater, Sleeper, and Mover interfaces
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

	// Using the specific interfaces
	dog.Eat()
	dog.Sleep()
	dog.Move()

	bird.Eat()
	bird.Sleep()
	bird.Move()
}
