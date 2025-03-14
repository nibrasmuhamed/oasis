package main

import (
	"fmt"
	"math"
)

func main() {
	x, y, points := 3, 4, [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}
	fmt.Println(nearestValidPoint(x, y, points))
}
func nearestValidPoint(x int, y int, points [][]int) int {
	nearest := math.MaxInt
	index := -1
	for i, v := range points {
		if x != v[0] && y != v[1] {
			continue
		}
		j := math.Abs(float64(x-v[0])) + math.Abs(float64(y-v[1]))
		if nearest > int(j) {
			nearest = int(j)
			index = i
		}
	}
	return index
}
