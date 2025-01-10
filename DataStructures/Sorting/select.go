package main

import "fmt"

func main() {
	var a = []int{9, 3, 5, 4, 8, 1, 6, 7}
	selectSort(a)
	fmt.Println(a)
}

func selectSort(a []int) {
	temp, j := 4, 0
	for i := 0; i < len(a); i++ {
		min := 9999999
		for j = i; j < len(a); j++ {
			if a[j] < min {
				min = a[j]
				temp = j
			}
		}
		val := a[i]
		a[i] = min
		a[temp] = val
		fmt.Println(a)
	}
}
