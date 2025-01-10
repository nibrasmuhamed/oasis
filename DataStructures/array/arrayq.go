package main

import (
	"fmt"
)

func main() {
	a := []int{-4, 0, 2, 3}
	fmt.Println(SquareAndSort(a))
}

func SquareAndSort(a []int) *[]int {
	first, end := 0, a[len(a)-1]
	res := make([]int, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		ls := a[first] * a[first]
		rs := a[end] * a[end]
		if ls > rs {
			res[i] = ls
			first++
		} else {
			res[i] = rs
			end--
		}
	}
	// sort.Slice((res), func(i, j int) bool { return (res)[i] < (res)[j] })
	return &res
}
