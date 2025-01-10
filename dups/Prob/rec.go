package main

import "fmt"

func main() {
	Hello(1)
}

func Hello(count int) {
	if count == 10 {
		return // This is a base case for this function
	}
	// fmt.Println(count)
	count++
	Hello(count)
	fmt.Println(count)
}
