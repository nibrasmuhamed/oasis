package main

import (
	"fmt"
	"time"
)

func main() {
	h1, m1, s1 := time.Now().Clock()

	fmt.Println(findFibo(40))
	h2, m2, s2 := time.Now().Clock()
	fmt.Println(h2-h1, m2-m1, s2-s1)
}

func findFibo(a int) int {
	my := map[int]int{0: 0, 1: 1}
	_, isPresent := my[a]
	if isPresent {
		return my[a]
	} else {
		my[a] = findFibo(a-1) + findFibo(a-2)
		return my[a]
	}
	// return findFibo(a-1) + findFibo(a-2)
}
