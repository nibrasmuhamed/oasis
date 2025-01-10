package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(reverse(a, 102))
}

func reverse(a []int, k int) []int {
	k = k % len(a)
	Reverse(a, 0, (len(a) - 1))
	Reverse(a, 0, k-1)
	Reverse(a, k, len(a)-1)
	return a
}

func Reverse(b []int, start int, end int) {
	for start < end {
		temp := b[start]
		b[start] = b[end]
		b[end] = temp
		start++
		end--
	}
}
