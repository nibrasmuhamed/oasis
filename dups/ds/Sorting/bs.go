package main

import (
	"fmt"
	"strconv"
)

func bs(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j+1]
				a[j+1] = a[j]
				a[j] = temp
			}
		}
		fmt.Print(strconv.Itoa(i) + " ")
		fmt.Println(a)
	}
	return a
}

func main() {
	a := []int{4, 5, 3, 2, 6, 1}
	fmt.Println(bs(a))
}
