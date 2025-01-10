// The Builder Pattern is a design pattern that allows for
// the step-by-step construction of complex objects.
// It separates the construction of an object from its
// representation, allowing the same construction process
// to create different representations. This is particularly
// useful when an object requires multiple steps to be created or
// when the construction process needs to allow for different
// configurations of the object.

// In Go, the Builder Pattern can be implemented using a
// struct and methods that return the struct itself, allowing
// for method chaining. Here is an example with an explanation:
package main

import "fmt"

type Car struct {
	make   string
	model  string
	year   int
	color  string
	engine string
	wheels int
}

type CarBuilder struct {
	car Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (b *CarBuilder) SetMake(make string) *CarBuilder {
	b.car.make = make
	return b
}

func (b *CarBuilder) SetModel(model string) *CarBuilder {
	b.car.model = model
	return b
}

func (b *CarBuilder) SetYear(year int) *CarBuilder {
	b.car.year = year
	return b
}

func (b *CarBuilder) SetColor(color string) *CarBuilder {
	b.car.color = color
	return b
}

func (b *CarBuilder) SetEngine(engine string) *CarBuilder {
	b.car.engine = engine
	return b
}

func (b *CarBuilder) SetWheels(wheels int) *CarBuilder {
	b.car.wheels = wheels
	return b
}

func (b *CarBuilder) Build() Car {
	return b.car
}

func main() {
	car := NewCarBuilder().
		SetMake("Toyota").
		SetModel("Corolla").
		SetYear(2024).
		SetColor("Red").
		SetEngine("V6").
		SetWheels(4).
		Build()

	fmt.Printf("Car: %+v\n", car)
}

//  1.	Car Struct: This is the struct we want to create instances of. It has several fields like make, model, year, color, engine, and wheels.
// 	2.	CarBuilder Struct: This struct helps in building the Car object step by step. It has a field car of type Car.
// 	3.	CarBuilder Methods: These methods (SetMake, SetModel, SetYear, SetColor, SetEngine, SetWheels) set the respective fields of the Car struct and return the CarBuilder itself to allow method chaining.
// 	4.	Build Method: This method returns the constructed Car object.
// 	5.	Usage: We create a new CarBuilder, set the properties of the car using method chaining, and finally call Build to get the constructed Car object.
