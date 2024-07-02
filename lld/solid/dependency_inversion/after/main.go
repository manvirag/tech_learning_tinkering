package main

import "fmt"

// GreetService defines the interface for a greeting service.
type GreetService interface {
	Greet()
}

// SimpleGreetService is an implementation of GreetService.
type SimpleGreetService struct{}

// Greet prints a simple greeting message.
func (s *SimpleGreetService) Greet() {
	fmt.Println("Hello, with DI using interface!")
}

// App is an application that uses the GreetService.
type App struct {
	greeter GreetService
}

// NewApp creates a new instance of App with the provided greeter.
func NewApp(greeter GreetService) *App {
	return &App{greeter: greeter}
}

// Run runs the application.
func (a *App) Run() {
	a.greeter.Greet()
}

func main() {
	// Create an instance of SimpleGreetService.
	greeter := &SimpleGreetService{}

	// Create an instance of App with the GreetService injected.
	app := NewApp(greeter)

	// Run the application.
	app.Run()
}
