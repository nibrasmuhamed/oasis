package main

import "math"

func ThreeFactorCount(arr []int, queries [][]int) []int {
	result := make([]int, len(queries))

	for i, query := range queries {
		start := query[0]
		end := query[1]
		count := 0

		for j := start; j <= end; j++ {
			if isThreeFactor(arr[j]) {
				count++
			}
		}

		result[i] = count
	}

	return result
}

func isThreeFactor(num int) bool {
	sqrt := math.Sqrt(float64(num))
	return sqrt == float64(int(sqrt))
}
