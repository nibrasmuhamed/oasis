package main

import "fmt"

// time complexity is O(n)
// iterating through the slice

// space complexity is O(4)
// not using any additional spaces other than constant value.
// O(1)
func main() {
	a := []int{20, 78, 32, 54, 23}
	fmt.Println(findLargest(a))
}

func findLargest(a []int) int {
	fst := findBig(a)
	snd := findBig(a)
	thd := findBig(a)
	frt := findBig(a)
	return fst + snd + thd + frt

}

func findBig(a []int) int {
	// find the largest number in the array
	max, index := 0, 0
	for i, val := range a {
		if val > max {
			max = val
			index = i
		}

	}
	a[index] = 0
	return max
}
