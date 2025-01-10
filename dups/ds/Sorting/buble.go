package main

import "fmt"

func main() {
	var a = []int{9, 3, 5, 4, 8, 1, 6, 7}
	bubbleSort(a)
	fmt.Println(a)
}

func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
		fmt.Println(a)
	}
}
