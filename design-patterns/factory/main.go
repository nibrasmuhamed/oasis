package main

import "fmt"

// Product is the interface that defines the methods that all concrete products must implement
type Product interface {
	Use() string
}

// ConcreteProductA is one implementation of the Product interface
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Using ConcreteProductA"
}

// ConcreteProductB is another implementation of the Product interface
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Using ConcreteProductB"
}
func CreateProduct(productType string) Product {
	switch productType {
	case "A":
		return &ConcreteProductA{}
	case "B":
		return &ConcreteProductB{}
	default:
		return nil
	}
}
func main() {
	// Create products using the factory function
	productA := CreateProduct("A")
	productB := CreateProduct("B")

	if productA != nil {
		fmt.Println(productA.Use()) // Output: Using ConcreteProductA
	}
	if productB != nil {
		fmt.Println(productB.Use()) // Output: Using ConcreteProductB
	}
}

// The Factory Pattern is a creational design pattern that provides
// an interface for creating objects in a superclass but allows
// subclasses to alter the type of objects that will be created.
// This pattern is useful when the exact type of the object being
// created is not known until runtime.
