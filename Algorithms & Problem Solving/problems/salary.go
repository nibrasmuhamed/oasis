package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(average([]int{1000, 2000, 3000, 4000}))
}

func average(salary []int) float64 {
	var max, min = math.MinInt, math.MaxInt
	sum := 0
	for _, v := range salary {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		sum += v
	}
	var q float64 = ((float64(sum) - float64(min)) - float64(max)) / float64((len(salary) - 2))
	return q
}
