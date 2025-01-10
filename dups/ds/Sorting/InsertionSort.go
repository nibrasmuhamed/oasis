package main

import "fmt"

func main() {
	var a = []int{1, 6, 4, 8, 3, 6, 7, 9}
	insertionSort(a)
	fmt.Println(a)
}

func insertionSort(a []int) {
	for i := 1; i < len(a)-1; i++ {
		current := a[i]
		var j int
		for j = i - 1; j >= 0 && current < a[j]; j-- {
			a[j+1] = a[j]
			// a[j] = current
		}
		a[j+1] = current
		fmt.Println(a)
	}
}
