package main

import (
	"fmt"
	"time"
)

func main() {
	h1, m1, s1 := time.Now().Clock()

	fmt.Println(findFibo(45))
	h2, m2, s2 := time.Now().Clock()
	fmt.Println(h2-h1, m2-m1, s2-s1)
}

func findFibo(a int) int {
	if a <= 1 {
		return a
	}
	return findFibo(a-1) + findFibo(a-2)
}

// 1134903170
// 0 0 14
// input was 45
//
