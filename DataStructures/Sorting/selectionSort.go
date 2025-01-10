package main

import "fmt"

func main() {
	var a = []int{1, 6, 4, 8, 3, 2, 7, 9}
	selectionSort(a)
	fmt.Println(a)
}

func selectionSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a)-1; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
		fmt.Println(a)
	}
}
