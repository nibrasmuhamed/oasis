package main

import "fmt"

func main() {
	a := []int{4, 4, 4, 4, 4}
	fmt.Println(Find(a, 4))
}

func findStart(a []int, target int) int {
	start, end := 0, len(a)-1
	fIndex := -1
	for start <= end {
		middle := (start + end) / 2
		if a[middle] == target {
			if middle == 0 {
				fIndex = 0
				break
			}
			if a[middle-1] == target {
				end = middle - 1
				continue
			}
			fIndex = middle
			break
		}
		if target < a[middle] {
			end = middle - 1
		} else if target > a[middle] {
			start = middle + 1
		}
	}
	return fIndex
}

func Find(a []int, target int) []int {
	left := findStart(a, target)
	right := findEnd(a, target)
	return []int{left, right}
}

func findEnd(a []int, target int) int {
	start, end := 0, len(a)-1
	fIndex := -1
	for start <= end {
		middle := (start + end) / 2
		if a[middle] == target {
			if middle == end {
				fIndex = middle
				break
			}
			if a[middle+1] == target {
				start = middle + 1
				continue
			}
			fIndex = middle
			break
		}
		if target < a[middle] {
			end = middle - 1
		} else if target > a[middle] {
			start = middle + 1
		}
	}
	return fIndex
}
