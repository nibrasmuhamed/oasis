package main

import "fmt"

func main() {
	var a = []int{1, 6, 4, 8, 3, 6, 7, 9}
	insertionSort(a)
	fmt.Println(a)
}

func insertionSort(a []int) {
	for i := 1; i < len(a)-1; i++ {
		temp, j := a[i], i-1
		for temp < a[j] {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = temp
	}
}
