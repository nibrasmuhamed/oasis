package main

import "fmt"

func main() {
	a := []int{1, 1, 2}
	fmt.Println(CheckMonotonic(a))
}

func CheckMonotonic(a []int) bool {
	if len(a) == 0 {
		return true
	}
	f, l := a[0], a[len(a)-1]
	if f == l {
		for i := 0; i <= len(a)-2; i++ {
			if a[i+1] != a[i] {
				return false
			}
		}
	} else if f > l {
		for i := 0; i <= len(a)-2; i++ {
			if a[i+1] > a[i] {
				return false
			}
		}
	} else {
		for i := 0; i <= len(a)-2; i++ {
			if a[i+1] < a[i] {
				return false
			}
		}
	}
	return true
}
