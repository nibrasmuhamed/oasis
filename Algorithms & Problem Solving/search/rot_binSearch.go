package main

import "fmt"

func main() {
	a := []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	fmt.Println(binSearch(a, 5))
}

func binSearch(a []int, target int) int {
	start := 0
	end := len(a) - 1
	for start <= end {
		middle := (start + end) / 2
		if target == a[middle] {
			return middle
		}
		if a[start] < a[middle] {
			// left is sorted.
			if target >= a[start] && target < a[middle] {
				// explore left
				end = middle - 1
			} else {
				start = middle + 1
			}
		} else {
			// right is sorted.
			if target <= a[end] && target > a[middle] {
				start = middle + 1
			} else {
				end = middle - 1
			}
		}
	}
	return -1
}
