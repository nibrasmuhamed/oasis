package main

import "fmt"

func hourglassSum(arr [][]int32) int32 {
	// Write your code here
	var max_sum, sum int32 = 0, 0
	for i := 0; i < len(arr)-2; i++ {
		for j := 0; j < len(arr)-2; j++ {
			sum = (arr[i][j] + arr[i][j+1] + arr[i][j+2]) +
				(arr[i+1][j+1]) +
				(arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2])

			if sum > max_sum {
				max_sum = sum
			}
		}

	}
	return max_sum
}

func main() {
	a := [][]int32{{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0}}
	fmt.Println(hourglassSum(a))
}
