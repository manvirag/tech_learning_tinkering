//* Ever incoming request is passed to the chain and each of the handler:
//1. Processes the request or skips the processing.
//2. Decides whether to pass the request to the next handler in the chain or not.

package main

import "fmt"

// Payment represents a payment with a specific amount.
type Payment struct {
	Amount float64
}

// PaymentHandler defines the interface for handling payments.
type PaymentHandler interface {
	HandlePayment(payment Payment)
	SetSuccessor(successor PaymentHandler)
}

// PaypalHandler is a concrete implementation of PaymentHandler for Paypal payments.
type PaypalHandler struct {
	successor PaymentHandler
}

func (p *PaypalHandler) SetSuccessor(successor PaymentHandler) {
	p.successor = successor
}

func (p *PaypalHandler) HandlePayment(payment Payment) {
	if payment.Amount < 10 {
		fmt.Printf("Processing payment of $%.2f with Paypal.\n", payment.Amount)
	} else if p.successor != nil {
		p.successor.HandlePayment(payment)
	}
}

// CreditCardHandler is a concrete implementation of PaymentHandler for Credit Card payments.
type CreditCardHandler struct {
	successor PaymentHandler
}

func (c *CreditCardHandler) SetSuccessor(successor PaymentHandler) {
	c.successor = successor
}

func (c *CreditCardHandler) HandlePayment(payment Payment) {
	if payment.Amount < 100 {
		fmt.Printf("Processing payment of $%.2f with Credit Card.\n", payment.Amount)
	} else if c.successor != nil {
		c.successor.HandlePayment(payment)
	}
}

// BankTransferHandler is a concrete implementation of PaymentHandler for Bank Transfer payments.
type BankTransferHandler struct {
	successor PaymentHandler
}

func (b *BankTransferHandler) SetSuccessor(successor PaymentHandler) {
	b.successor = successor
}

func (b *BankTransferHandler) HandlePayment(payment Payment) {
	if payment.Amount >= 100 {
		fmt.Printf("Processing payment of $%.2f with Bank Transfer.\n", payment.Amount)
	} else {
		fmt.Printf("Invalid payment amount: $%.2f\n", payment.Amount)
	}
}

func main() {
	paypalHandler := &PaypalHandler{}
	creditCardHandler := &CreditCardHandler{}
	bankTransferHandler := &BankTransferHandler{}

	paypalHandler.SetSuccessor(creditCardHandler)
	creditCardHandler.SetSuccessor(bankTransferHandler)

	payment1 := Payment{Amount: 5}
	paypalHandler.HandlePayment(payment1)

	payment2 := Payment{Amount: 50}
	paypalHandler.HandlePayment(payment2)

	payment3 := Payment{Amount: 500}
	paypalHandler.HandlePayment(payment3)
}
