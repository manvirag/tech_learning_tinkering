package main

import "fmt"

// SimpleGreetService is a concrete implementation of a greeting service.
type SimpleGreetService struct{}

// Greet prints a simple greeting message.
func (s *SimpleGreetService) Greet() {
	fmt.Println("Hello, without DI using concrete type!")
}

// App is an application that uses the SimpleGreetService.
type App struct {
	greeter SimpleGreetService
}

// Run runs the application.
func (a *App) Run() {
	a.greeter.Greet()
}

func main() {
	// Create an instance of App.
	app := App{}

	// Run the application.
	app.Run()
}
