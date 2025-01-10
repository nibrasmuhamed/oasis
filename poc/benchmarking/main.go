package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(randomGenerator(100))
}

func randomGenerator(n int) int {
	a := int(time.Now().UnixMilli())
	return a % n
}
