package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(a, 5))
	fmt.Println(binaryRecursive(a, 5))
}

func binarySearch(a []int, target int) int {
	start, end := 0, len(a)-1
	for start <= end {
		middle := (start + end) / 2
		if a[middle] == target {
			return middle
		} else if a[middle] > target {
			end = middle - 1
		} else if a[middle] < target {
			start = middle + 1
		}
	}
	return -1
}

func binaryRecursive(a []int, target int) int {
	return RecursiveHelper(a, target, 0, len(a)-1)
}

func RecursiveHelper(a []int, target, start, end int) int {
	if start > end {
		return -1
	}
	middle := (start + end) / 2
	if a[middle] == target {
		return middle
	} else if a[middle] > target {
		return RecursiveHelper(a, target, start, middle-1)
	} else {
		return RecursiveHelper(a, target, middle+1, end)
	}
}

// iterative:
// 		Space complexity : O(1)
// 		Time complexity	: O(log n)

// recursive:
//		Space complexity : O(log n)
// 		Time complexity : O(log n)
