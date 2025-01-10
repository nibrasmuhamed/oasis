package main

// import (
// 	"fmt"
// )

// // Bird interface
// type Bird interface {
// 	Fly() string
// }

// // Sparrow struct
// type Sparrow struct{}

// // Fly method for Sparrow
// func (s Sparrow) Fly() string {
// 	return "Sparrow is flying"
// }

// // Ostrich struct
// type Ostrich struct{}

// // Fly method for Ostrich
// func (o Ostrich) Fly() string {
// 	return "Ostrich can't fly"
// }

// func main() {
// 	var bird Bird

// 	bird = Sparrow{}
// 	fmt.Println(bird.Fly())

// 	bird = Ostrich{}
// 	fmt.Println(bird.Fly())
// }

// In this example:

// 	•	We define a Bird interface with a Fly method.
// 	•	Both Sparrow and Ostrich structs implement the Bird interface.
// 	•	The main function demonstrates that both Sparrow and Ostrich can be used interchangeably as Bird.

// However, this example violates LSP because an ostrich cannot fly, which changes the expected behavior of the Fly method. To adhere to LSP, you would typically design your interfaces and types in such a way that they make logical sense and don’t lead to unexpected behavior.

// Let’s refactor this example to better adhere to LSP:

import (
	"fmt"
)

// Bird interface
type Bird interface {
	Move() string
}

// FlyingBird interface extends Bird interface
type FlyingBird interface {
	Bird
	Fly() string
}

// Sparrow struct
type Sparrow struct{}

// Move method for Sparrow
func (s Sparrow) Move() string {
	return "Sparrow is moving"
}

// Fly method for Sparrow
func (s Sparrow) Fly() string {
	return "Sparrow is flying"
}

// Ostrich struct
type Ostrich struct{}

// Move method for Ostrich
func (o Ostrich) Move() string {
	return "Ostrich is moving"
}

func main() {
	var bird Bird
	var flyingBird FlyingBird

	bird = Ostrich{}
	fmt.Println(bird.Move())

	flyingBird = Sparrow{}
	fmt.Println(flyingBird.Move())
	fmt.Println(flyingBird.Fly())
}

// In this refactored example:

// 	•	We separate the Fly method into a new interface FlyingBird, which extends the Bird interface.
// 	•	Now, Ostrich only implements the Bird interface, and Sparrow implements both Bird and FlyingBird.
// 	•	The main function uses Bird and FlyingBird interfaces appropriately, adhering to LSP.
