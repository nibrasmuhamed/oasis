package main

import "fmt"

func main() {
	a := []int32{1, 2, 3, 4, 5, 6}
	fmt.Println(reverseArray(a))
}

func reverseArray(a []int32) []int32 {
	// Write your code here
	j := len(a) - 1
	for i := 0; i < (len(a) / 2); i = i + 1 {

		temp := a[j]
		a[j] = a[i]
		a[i] = temp
		j = j - 1

	}
	return a
}
