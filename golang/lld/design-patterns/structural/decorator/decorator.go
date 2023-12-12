package main

type PizzaInterface interface {
	Price() int
}

type Pizza struct{}

func (p Pizza) Price() int {
	return 20
}

type CapsicumTopping struct {
	pizza PizzaInterface
}

func (cp CapsicumTopping) Price() int {
	return cp.pizza.Price() + 5
}

type OnionTopping struct {
	pizza PizzaInterface
}

func (op OnionTopping) Price() int {
	return op.pizza.Price() + 10
}
