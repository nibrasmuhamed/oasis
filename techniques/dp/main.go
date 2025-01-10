package main

import (
	"fmt"
	"math"
	"time"
)

// notes: https://notability.com/n/0vnJV81q49Ocp_P8mzzUDX

// 1. suboptimal structure
// if a problem can be divided in to subproblems
// and solving those will help to solve main problem

// 2. problem has overlapping subproblem
// while solving the problem, if same calculation
// done twice, it has  redundant work/ overlapping subproblem

// top-down dp/ recursive dp (memoization)
// start with current problem. solve, store subproblem ans
// use it to solve current problem. similar to cache.

// bottom-up dp/ iterative dp
// start with smallest subproblem.
// use it to calculate bigger problem till we get
// final result.
// if no recursion space, there is a possibility
// to reduce space complexity.

func main() {
	// x := 12
	// x := make([]int, 8)\
	// x := [][]int{
	// 	{1, -1, 0},  // {-2, -3, 3},  //  {3, 6, 3}
	// 	{-1, 1, -1}, // {-5, -10, 1}, //  {8, 16, 2}
	// 	{1, 0, -1},  //{10, 30, -5}, //  {1, 1, 6}
	// }
	ti := time.Now()
	fmt.Println(numTrees(18))
	fmt.Println(time.Since(ti))
	// fmt.Println(maxSumWithoutAdjacentSlice(0, []int{1, 20, 3, 4, 5, 20, 27, 89}, x))
}
func numTrees(A int) int {
	dpArray := make([]int, A+1)
	dpArray[0] = 1
	dpArray[1] = 1
	for i := 2; i <= A; i++ {
		sum := 0
		for j := 1; j <= i; j++ {
			left := dpArray[i-1]
			right := dpArray[i-j]
			sum += left * right
		}
		dpArray[i] = sum
	}
	return dpArray[A]
}
func MaxSumVal(A []int, B, C, D int) int {
	n := len(A)
	dp := make([][3]int, n+1)
	dp[0][0] = -1000000000
	dp[0][1] = -1000000000
	dp[0][2] = -1000000000

	for i := 1; i <= n; i++ {
		dp[i][0] = max(dp[i-1][0], A[i-1]*B)
		dp[i][1] = max(dp[i-1][1], dp[i][0]+A[i-1]*C)
		dp[i][2] = max(dp[i-1][2], dp[i][1]+A[i-1]*D)
	}
	return dp[n][2]

}
func maxSumVal(A []int, B, C, D, i, j, k int, memo map[[3]int]int) int {
	if i > j || i > k || j > k || k >= len(A) {
		return math.MinInt64
	}
	if val, ok := memo[[3]int{i, j, k}]; ok {
		return val
	}

	max1 := maxSumVal(A, B, C, D, i, j, k+1, memo)
	max2 := maxSumVal(A, B, C, D, i, j+1, k, memo)
	max3 := maxSumVal(A, B, C, D, i+1, j, k, memo)
	res := A[i]*B + A[j]*C + A[k]*D

	memo[[3]int{i, j, k}] = max(max1, max2, max3, res)
	return memo[[3]int{i, j, k}]
}
func calculateMinimumHP(A [][]int) int {
	rows := len(A)
	cols := len(A[0])

	// Create a 2D slice (matrix) with the same dimensions as A
	mat := make([][]int, rows)
	for i := range mat {
		mat[i] = make([]int, cols)
	}

	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if i == rows-1 && j == cols-1 {
				mat[i][j] = max(1-A[i][j], 1)
			} else if i == rows-1 {
				mat[i][j] = max(mat[i][j+1]-A[i][j], 1)
			} else if j == cols-1 {
				mat[i][j] = max(mat[i+1][j]-A[i][j], 1)
			} else {
				mat[i][j] = max(min(mat[i+1][j], mat[i][j+1])-A[i][j], 1)
			}
		}
	}
	return mat[0][0]
}
func numOfWaysRecursive(i, j int, dp [][]int) int {
	if i == 0 || j == 0 {
		return 1
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	dp[i][j] = numOfWaysRecursive(i-1, j, dp) + numOfWaysRecursive(i, j-1, dp)
	return dp[i][j]
}
func numberOfWays(A [][]int) int {
	rows := len(A)
	cols := len(A[0])

	// Create a 2D slice (matrix) with the same dimensions as A
	mat := make([][]int, rows)
	for i := range mat {
		mat[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; i < cols; j++ {
			if A[i][j] == 1 {
				mat[i][j] = 0
			} else if i == 0 && j == 0 {
				mat[i][j] = 1
			} else if i == 0 {
				mat[i][j] = mat[i][j-1]
			} else if j == 0 {
				mat[i][j] = mat[i-1][j]
			} else {
				mat[i][j] = mat[i-1][j] + mat[i][j-1]
			}
		}
	}
	return mat[rows-1][cols-1]
}
func numDecodings(A string) int {
	const mod = 1000000007
	return 0
}
func maxSumWithoutAdjacentSlice(i int, A []int, memo []int) int {
	if i >= len(A) {
		return 0
	}

	// Check if the result for the current index is already memoized
	if memo[i] != 0 {
		fmt.Println("cache found", memo[i])
		return memo[i]
	}

	// Calculate the maximum sum for the current index
	s1 := A[i] + maxSumWithoutAdjacentSlice(i+2, A, memo)
	s2 := maxSumWithoutAdjacentSlice(i+1, A, memo)

	// Store the result in the memo slice
	memo[i] = max(s1, max(s2))

	return memo[i]
}
func maxSumWithoutAdjacentMemoize(i int, A []int, memo map[int]int) int {
	if i >= len(A) {
		return 0
	}

	// Check if the result for the current index is already memoized
	if val, found := memo[i]; found {
		fmt.Println("cache found :", val)
		return val
	}

	// Calculate the maximum sum for the current index
	s1 := A[i] + maxSumWithoutAdjacentMemoize(i+2, A, memo)
	s2 := maxSumWithoutAdjacentMemoize(i+1, A, memo)

	// Store the result in the memo map
	memo[i] = max(s1, s2)

	return memo[i]
}
func maxSumWithoutAdjacent(i int, A []int) int {
	if i >= len(A) {
		return 0
	}
	s1 := A[i] + maxSumWithoutAdjacent(i+2, A)
	s2 := maxSumWithoutAdjacent(i+1, A)
	return max(s1, s2)
}

func maxProduct(A []int) int {
	suff, pref := 1, 1
	maxI := A[0]
	n := len(A)
	for i := 0; i < len(A); i++ {
		if pref == 0 {
			pref = 1
		}
		if suff == 0 {
			suff = 1
		}
		pref = pref * A[i]
		suff = suff * A[n-i-1]
		maxI = max(maxI, max(pref, suff))
	}
	return maxI
}
func adjacent(A [][]int) int {
	totalCols := len(A[0])
	arr := make([]int, totalCols)
	arr[0] = max(A[0][0], A[0][1])
	if len(A) == 1 {
		return arr[0]
	}
	arr[1] = max(arr[0], max(A[1][0], A[1][1]))
	for i := 2; i < totalCols; i++ {
		arr[i] = max(arr[i-2]+max(A[0][i], A[1][i]), arr[i-1])
	}
	return arr[totalCols-1]
}

func largestInColumns(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0 {
		// Empty matrix, return an empty array
		return []int{}
	}

	cols := len(matrix[0])
	result := make([]int, cols)

	for j := 0; j < cols; j++ {
		maximum := matrix[0][j]

		for i := 1; i < rows; i++ {
			if matrix[i][j] > maximum {
				maximum = matrix[i][j]
			}
		}

		result[j] = maximum
	}

	return result
}

var memo = make(map[int]int)

func countMinSquares(A int) int {
	if A == 0 {
		return 0
	}

	if val, ok := memo[A]; ok {
		return val
	}

	m := A

	for i := 1; i*i <= A; i++ {
		m = min(m, 1+countMinSquares(A-i*i))
	}

	memo[A] = m
	return m
}

// [[16,5 ,54,55,36,82,61,77,66,61],
// [[31,30,36,70,9 ,37,1 ,11,68,14]]
func climbStairs(A int) int {
	if A <= 2 {
		return A
	}

	mod := 1000000007
	a, b := 1, 2

	for i := 3; i <= A; i++ {
		c := (a + b) % mod
		a = b
		b = c
	}

	return b
}

func fibo(n int) int {
	a := 0
	b := 1
	for i := 2; i <= n; i++ {
		c := a + b
		a = b
		b = c
	}
	return a + b
}
