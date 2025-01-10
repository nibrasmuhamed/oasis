package main

import "fmt"

func main() {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
		{13, 14, 15},
		{16, 17, 18},
	}
	fmt.Println(checkTwoD(a, 0))
}

func checkTwoD(a [][]int, target int) bool {
	end := len(a) - 1
	start := 0
	for start <= end {
		middle := (start + end) / 2
		if target >= a[middle][0] && target <= a[middle][len(a[middle])-1] {
			x := binIterative(a[middle], target)
			if x == -1 {
				return false
			}
			return true
		}
		if target > a[middle][len(a[middle])-1] {
			start = middle + 1
		} else if target < a[middle][0] {
			end = middle - 1
		}

	}
	return false
}
