package main

import (
	"fmt"
)

func main() {
	var a = []int{8, 3, 9, 4, 8, 1, 6, 7}
	insertSort(a)
	fmt.Println(a)
}

func insertSort(a []int) {
	for i := 1; i < len(a); i++ {
		current := a[i]
		for j := i; j >= 0; j-- {
			if current < a[j] {
				temp := a[j+1]
				a[j+1] = a[j]
				a[j] = temp
			}
		}
		fmt.Println(a)
	}
}
