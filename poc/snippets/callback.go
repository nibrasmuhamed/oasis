package main

import "fmt"

func main() {
	// q := func(x int) int {
	// 	return x + 1
	// }
	fmt.Println(call(func(x int) int {
		return x + 1
	}, 1))
}

func call(f func(int) int, x int) int {

	return f(x) + 1
}

func add(x int) int {
	return x + 1
}
