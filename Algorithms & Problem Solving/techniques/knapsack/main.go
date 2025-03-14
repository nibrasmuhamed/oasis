package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	A := []int{
		// 3, 3, 9, 1, 10, 5, 8, 1, 10,
		// 1, 3, 2, 4, 5,
		4, 8, 6, 5, 1, 10, 2,
	}
	fmt.Println(y(A, 5))
}
func solve(A []int, B int) int {
	N := len(A)
	dp := make([]int, N)
	dp[0] = A[0]
	dp[1] = A[1]
	if B == 0 {
		dp[1] += dp[0]
	}
	for i := 2; i < N; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + A[i]
	}
	if B > 1 {
		for i := B + 1; i < N; i++ {
			if i == B+1 {
				dp[i] = dp[i-1] + A[i]
			} else {
				dp[i] = min(dp[i-1], dp[i-1]) + A[i]
			}
		}
	}
	return min(dp[N-1], dp[N-2])
}
func y(A []int, B int) int {
	N := len(A)
	dp := make([]int, N)
	dp[0] = A[0]
	dp[1] = A[1]
	if B == 0 {
		dp[1] += dp[0]
	}
	for i := 2; i <= B; i++ {
		dp[i] = A[i] + min(dp[i-1], dp[i-2])
	}
	for i := B + 1; i < len(A); i++ {
		if i == B+1 {
			dp[i] = A[i] + dp[i-1]
		} else {
			dp[i] = A[i] + min(dp[i-1], dp[i-2])
		}
	}
	return min(dp[N-1], dp[N-2])
}

func minCost(costs []int) int {
	min := costs[0]
	for _, cost := range costs[1:] {
		if cost < min {
			min = cost
		}
	}
	return min
}
func graph(A [][]int) [][]int {
	N := len(A)
	M := len(A[0])

	// Initialize the result matrix with maximum possible distance
	B := make([][]int, N)
	for i := 0; i < N; i++ {
		B[i] = make([]int, M)
		for j := 0; j < M; j++ {
			B[i][j] = N + M
		}
	}

	// Use BFS to find the distance of the nearest 1 for each cell
	queue := list.New()

	// Add all the 1's to the queue and set their distance to 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if A[i][j] == 1 {
				B[i][j] = 0
				queue.PushBack([2]int{i, j})
			}
		}
	}

	// Define possible directions to move (up, down, left, right)
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// Perform BFS
	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		current := front.Value.([2]int)
		x, y := current[0], current[1]

		// Explore all possible directions
		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]

			// Check if the next position is within the matrix bounds
			if nx >= 0 && nx < N && ny >= 0 && ny < M {
				// Update the distance if it's shorter than the current distance
				if B[nx][ny] > B[x][y]+1 {
					B[nx][ny] = B[x][y] + 1
					queue.PushBack([2]int{nx, ny})
				}
			}
		}
	}

	return B
}
func x(A []int, B int) int {
	N := len(A)

	// Initialize an array to store the minimum cost to reach each stair
	dp := make([]int, N)

	// Base case: the cost to reach the first two stairs
	dp[0] = A[0]
	dp[1] = A[1]

	// Iterate from the third stair onwards
	for i := 2; i < N; i++ {
		// Update the minimum cost to reach the current stair based on the previous two stairs
		dp[i] = min(dp[i-1], dp[i-2]) + A[i]

		// If the current stair is the B-th stair, ensure that it is not skipped
		if i == B {
			dp[i] = min(dp[i], dp[i-1])
		}
	}

	// Return the minimum cost to reach the top of the staircase
	return min(dp[N-1], dp[N-2])
}

func coinSelection(A []int, sum int, idx int) int {
	if sum == 0 {
		return 1
	}
	if sum < 0 || idx >= len(A) {
		return 0
	}

	include := coinSelection(A, sum-A[idx], idx)
	exclude := coinSelection(A, sum, idx+1)

	return include + exclude
}
func offOnSignalBottomUp(A int) int {
	const mod = 1000000007

	dp := make([][2]int, A)
	dp[0][1] = 1
	dp[0][0] = 1
	for i := 1; i < A; i++ {
		dp[i][0] = ((dp[i-1][0] % mod) + (dp[i-1][1] % mod)) % mod
		dp[i][1] = dp[i-1][1]
	}
	return (dp[A-1][0]%mod + dp[A-1][1]%mod) % mod
}

func onOffSignal(A int, B bool) int {
	if A <= 0 {
		return 1
	}
	var on, off int
	if B {
		off = onOffSignal(A-1, false)
	}
	on = onOffSignal(A-1, true)
	return on + off
}
func minFlipsRecursive(A []int, i int, flip int) int {
	if i >= len(A) {
		return flip
	}

	// Try both possibilities: not flipping and flipping the current element
	exclude := A[i] + minFlipsRecursive(A, i+1, flip)
	include := -A[i] + minFlipsRecursive(A, i+1, flip+1)

	// Return the minimum number of flips between the two possibilities
	return minNonNegative(min(exclude, include), flip+1)
}
func minNonNegative(a, b int) int {
	if a >= 0 && b >= 0 {
		return min(a, b)
	} else if a >= 0 {
		return a
	} else if b >= 0 {
		return b
	}
	return -1 // Bot
}
func unboundedKnapSack(A, B []int, C int) int {
	n := len(A)
	dp := make([]int, C+1)
	for i := 1; i < C+1; i++ {
		for j := 0; j < n; j++ {
			if i >= B[j] {
				dp[i] = max(dp[i], A[j]+dp[i-B[j]])
			}
		}
	}
	return dp[C]
}
func unboundedKnapSackR(A, B []int, C int) int {
	if C <= 0 {
		return 0
	}
	val := -1
	for j := 0; j < len(A); j++ {
		if B[j] <= C {
			val = max(val, A[j]+unboundedKnapSackR(A, B, C-B[j]))
		}
	}
	return val
}
func zeroOneKnapSack(A, B []int, C int) int {
	n := len(A)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, C+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= C; j++ {
			if j < B[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], A[i-1]+dp[i-1][j-B[i-1]])
			}
		}
	}
	return dp[n][C]
}

func zeroOneKnapSackRecursive(A, B []int, i, C int) int {
	if i >= len(A) || C <= 0 {
		return 0
	}
	// include the current item if its weight is less than or equal to remaining capacity
	if B[i] <= C {
		include := A[i] + zeroOneKnapSackRecursive(A, B, i+1, C-B[i])
		notInclude := zeroOneKnapSackRecursive(A, B, i+1, C)
		return max(include, notInclude)
	}
	// if weight is greater than remaining capacity, skip to the next item
	return zeroOneKnapSackRecursive(A, B, i+1, C)
}

type Item struct {
	Value, Weight float64
}

func FractionKnapSack(A []int, B []int, C int) int {
	n := len(A)
	items := make([]Item, n)

	// Create a slice of items with values and weights
	for i := 0; i < n; i++ {
		items[i] = Item{Value: float64(A[i]), Weight: float64(B[i])}
	}

	// Sort the items in descending order of value-to-weight ratio
	sort.Slice(items, func(i, j int) bool {
		return items[i].Value/items[i].Weight > items[j].Value/items[j].Weight
	})

	var totalValue float64 = 0.000
	var currentWeight float64 = 0.000

	// Iterate through the sorted items and add them to the knapsack
	for i := 0; i < n; i++ {
		if currentWeight+items[i].Weight <= float64(C) {
			currentWeight += items[i].Weight
			totalValue += items[i].Value
		} else {
			remainingWeight := float64(C) - currentWeight
			totalValue += (items[i].Value / items[i].Weight) * remainingWeight
			break
		}
	}

	// Return the floor of (totalValue * 100)
	totalValue *= 1000
	return int(totalValue / 10)
}
