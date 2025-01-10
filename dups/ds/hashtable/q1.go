package main

import "fmt"

func main() {
	var a = []int{9, 3, 5, 4, 8, 1, 6, 7}
	fmt.Println(addUp(a, 8))
}

func addUp(a []int, sum int) (int, int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == sum {
				return i, j
			}
		}
	}
	return 0, 0
}
