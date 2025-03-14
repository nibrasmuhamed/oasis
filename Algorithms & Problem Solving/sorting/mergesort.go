package main

import "fmt"

func main() {
	a := []int{47, 3, 32, 31, 89, 74, 83, 56}
	// b := []int{}
	fmt.Println(mergesort(a))
	// fmt.Println(b)
}

func mergesort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	middle := len(a) / 2
	left := mergesort(a[:middle])
	right := mergesort(a[middle:])
	return mergesortHelper(left, right)

}

func mergesortHelper(a, b []int) []int {
	i := 0
	j := 0
	res := make([]int, 0)
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	for i < len(a) {
		res = append(res, a[i])
		i++
	}
	for j < len(b) {
		res = append(res, b[j])
		j++
	}
	return res
}
