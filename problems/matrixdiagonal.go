package main

import "fmt"

func main() {
	a := [][]int{
		{7, 3, 1, 9},
		{3, 4, 6, 9},
		{6, 9, 6, 6},
		{9, 5, 8, 5},
	}
	fmt.Println(diagonalSum(a))
}
func diagonalSum(mat [][]int) int {
	left := 0
	right := 0
	for i := 0; i < len(mat); i++ {
		left += mat[i][i]
	}
	j := 0
	for i := len(mat) - 1; i >= 0; i-- {
		right += mat[i][j]
		j++
	}
	if len(mat)%2 != 0 {
		return right + (left - mat[len(mat)/2][len(mat)/2])
	}
	return right + left
}
