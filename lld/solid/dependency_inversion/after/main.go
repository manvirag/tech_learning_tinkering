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

func NewApp(greeter GreetService) *App {
	return &App{greeter: greeter}
}

func (a *App) Run() {
	a.greeter.Greet()
}

func main() {
	greeter := &SimpleGreetService{}

	app := NewApp(greeter)

	app.Run()
}
