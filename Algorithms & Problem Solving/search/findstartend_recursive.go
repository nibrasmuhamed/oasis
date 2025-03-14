package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 4, 4, 4, 4, 6}
	rnge := []int{-1, -1}
	findRightRecursive(a, 4, rnge)
	findLeftRecursive(a, 4, rnge)
	fmt.Println(rnge)
}

func findLeftRecursive(a []int, target int, rnge []int) {
	findLeftRecursiveHelper(a, target, 0, len(a)-1, rnge)
}

func findLeftRecursiveHelper(a []int, target int, start, end int, rnge []int) {
	middle := (start + end) / 2
	if start > end {
		return
	}
	if target == a[middle] {
		if target == a[middle-1] {
			findLeftRecursiveHelper(a, target, start, middle-1, rnge)
			return
		}
		rnge[0] = middle
	}
	if target > a[middle] {
		findLeftRecursiveHelper(a, target, middle+1, end, rnge)
	} else if target < a[middle] {
		findLeftRecursiveHelper(a, target, start, middle-1, rnge)
	}
}

func findRightRecursive(a []int, target int, rnge []int) {
	findRightRecursiveHelper(a, target, 0, len(a)-1, rnge)
}
func findRightRecursiveHelper(a []int, target, start, end int, rnge []int) {
	middle := (start + end) / 2
	if start > end {
		return
	}
	if target == a[middle] {
		if target == a[middle+1] {
			findRightRecursiveHelper(a, target, middle+1, end, rnge)
			return
		}
		rnge[1] = middle
	}
	if target > a[middle] {
		findRightRecursiveHelper(a, target, middle+1, end, rnge)
	} else if target < a[middle] {
		findRightRecursiveHelper(a, target, start, middle-1, rnge)
	}
}
