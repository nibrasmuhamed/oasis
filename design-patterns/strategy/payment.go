package main

import "fmt"

// PaymentStrategy defines the interface for payment strategies
type PaymentStrategy interface {
	Pay(amount float64) string
}

// CreditCardPayment is a concrete strategy for credit card payments
type CreditCardPayment struct {
	Name   string
	CardNo string
}

func (c CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card (%s)", amount, c.CardNo)
}

// PayPalPayment is a concrete strategy for PayPal payments
type PayPalPayment struct {
	Email string
}

func (p PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal (%s)", amount, p.Email)
}

// BitcoinPayment is a concrete strategy for Bitcoin payments
type BitcoinPayment struct {
	WalletAddress string
}

func (b BitcoinPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Bitcoin (%s)", amount, b.WalletAddress)
}

// PaymentContext is the context class that uses the strategy
type PaymentContext struct {
	Strategy PaymentStrategy
}

func (pc *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	pc.Strategy = strategy
}

func (pc *PaymentContext) ExecutePayment(amount float64) string {
	return pc.Strategy.Pay(amount)
}

func mainPayment() {
	// Create a PaymentContext
	paymentContext := &PaymentContext{}

	// Use Credit Card Payment Strategy
	creditCard := CreditCardPayment{Name: "Alice", CardNo: "1234-5678-9876-5432"}
	paymentContext.SetStrategy(creditCard)
	fmt.Println(paymentContext.ExecutePayment(250.00))

	// Use PayPal Payment Strategy
	paypal := PayPalPayment{Email: "alice@example.com"}
	paymentContext.SetStrategy(paypal)
	fmt.Println(paymentContext.ExecutePayment(350.00))

	// Use Bitcoin Payment Strategy
	bitcoin := BitcoinPayment{WalletAddress: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"}
	paymentContext.SetStrategy(bitcoin)
	fmt.Println(paymentContext.ExecutePayment(500.00))
}
