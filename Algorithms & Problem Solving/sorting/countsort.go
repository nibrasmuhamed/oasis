package main

// count sort
import (
	"fmt"
)

func main() {
	a := []int{3, 5, 2, 6, 8, 9, 3, 0}
	fmt.Println(countSort(a))
}

func countSort(a []int) []int {
	// max := findMax(a)
	// maxDigit := math.Floor(math.Log10(float64(max)))
	output := make([]int, 10)
	index := make([]int, 10)
	for _, val := range a {
		index[val] += 1
	}
	cumilative := make([]int, len(index))
	for i := 0; i <= len(cumilative)-1; i++ {
		if i == 0 {
			cumilative[i] = index[i]
			continue
		}
		cumilative[i] = index[i] + cumilative[i-1]
	}
	for i := len(a) - 1; i >= 0; i-- {
		val := a[i]
		val2 := cumilative[val]
		cumilative[val] = cumilative[val] - 1
		output[val2-1] = val
	}
	return output
}

func findMax(a []int) int {
	max := 0
	for _, val := range a {
		if val > max {
			val = max
		}
	}
	return max
}
