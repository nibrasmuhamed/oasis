package main

import "fmt"

func main() {
	a := []int{3, 5, 4, 7, 5, 9, 5, 2}
	buildHeap(a)
	fmt.Println(a)
	extractMax(&a)
	fmt.Println(a)
}

func insert(a []int, data int) {
	a = append(a, data)
	bubbleUp(a)
}

func bubbleUp(a []int) {
	idx := len(a) - 1
	val := a[idx]
}

func buildHeap(a []int) []int {
	length := len(a)
	lastParent := a[(length/2)-1]
	for i := lastParent; i >= 0; i-- {
		bubbleDown(a, i)
	}
	return a
}

func extractMax(a *[]int) int {
	max := (*a)[0]
	last := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	if len(*a) > 1 {
		(*a)[0] = last
		bubbleDown(*a, 0)
	}
	return max
}

func bubbleDown(a []int, idx int) {
	length := len(a)
	current := a[idx]
	for {
		leftChildIndex := 2*idx + 1
		rightChildIndex := 2*idx + 2
		var left, right int
		var large *int = nil
		if leftChildIndex < length {
			left = a[leftChildIndex]
			if left > current {
				large = &leftChildIndex
			}
		}
		if rightChildIndex < length {
			right = a[rightChildIndex]
			if (large == nil && right > current) || (large != nil && right > left) {
				large = &rightChildIndex
			}
		}
		if large == nil {
			break
		}
		a[idx] = a[*large]
		a[*large] = current
		idx = *large
	}
}
