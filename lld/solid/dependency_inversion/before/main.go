package main

import "fmt"

// SimpleGreetService is a concrete implementation of a greeting service.
type SimpleGreetService struct{}

func (s *SimpleGreetService) Greet() {
	fmt.Println("Hello, without DI using concrete type!")
}

// App is an application that uses the SimpleGreetService.
type App struct {
	greeter SimpleGreetService
}

func (a *App) Run() {
	a.greeter.Greet()
}

func main() {
	app := App{}

	app.Run()
}
