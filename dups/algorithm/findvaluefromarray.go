package main

import "fmt"

func main() {
	a := []int{4, 3, 1, 2, 5, 6, 7, 8, 8}
	fmt.Println(findrep(a))
}

func findrep(a []int) int {
	hare := 0
	tortoise := 0
	x := 0
	for {
		hare = a[a[hare]]
		tortoise = a[tortoise]
		if hare == tortoise {
			for x != tortoise {
				tortoise = a[tortoise]
				x = a[x]

			}
			break
		}
	}
	return x
}
