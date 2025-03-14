package main

import "fmt"

func main() {
	// a := []int{78, 32, 43, 12, 89, 98, 79}
	a := []int{5, 4, 3, 0, 1}

	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func partition(a []int, start, end int) int {
	middle := (start + end) / 2
	swap(a, start, middle)
	pivot := a[start]
	i := start + 1
	j := end
	for i <= j {
		for i <= end && a[i] <= pivot {
			i++
		}
		for j >= start && a[j] > pivot {
			j--
		}
		if i < j {
			swap(a, i, j)
		}
	}
	swap(a, start, j)
	return j
}

func quickSort(a []int, start, end int) {
	for start < end {
		pivot := partition(a, start, end)
		if pivot-start < end-pivot {
			quickSort(a, start, pivot-1)
			start = pivot + 1
		} else {
			quickSort(a, pivot+1, end)
			end = pivot - 1
		}
	}
}
