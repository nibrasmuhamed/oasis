package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{32, 543, 892, 523, 223, 23}
	fmt.Println(radixSort(a))
}

func radixSort(a []int) []int {
	x := findMax(a)
	totalDigit := int(math.Floor(math.Log10(float64(x))) + 1)
	for i := 0; i < totalDigit; i++ {
		countingSort(a, i)
	}
	return a
}

func countingSort(a []int, q int) {
	newArray := make([]int, len(a))
	temp := make([]int, 10)
	// selecting the digit, (unit, tenth place or hundreadth place)
	digitPlace := int(math.Pow10(q))
	for _, val := range a {
		digit := int(math.Floor(float64(val/digitPlace))) % 10
		temp[digit]++
	}

	for j := 0; j < len(temp); j++ {
		if j == 0 {
			continue
		}
		temp[j] = temp[j] + temp[j-1]
	}
	for j := len(a) - 1; j >= 0; j-- {
		val := int(math.Floor(float64(a[j]/digitPlace))) % 10
		temp[val]--
		insertPos := temp[val]
		newArray[insertPos] = a[j]
	}
	a = newArray
}

func findMax(a []int) int {
	max := 0
	for _, val := range a {
		if val > max {
			max = val
		}
	}
	return max
}
