package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rotateArray(a, 3))
}

func rotateArray(a []int, num int) []int {
	rotation := num % len(a)
	for i := 1; i <= rotation; i++ {
		temp := a[len(a)-1]
		for j := len(a) - 1; j > 0; j-- {
			a[j] = a[j-1]
		}
		a[0] = temp
	}
	return a
}
