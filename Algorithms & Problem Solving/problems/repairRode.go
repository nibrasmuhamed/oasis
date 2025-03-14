package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	a := []int32{5, 5, 3, 1, 4, 6}
	b := []int32{5, 9, 8, 3, 15, 1}
	fmt.Println(getMinCost(a, b))
}

func getMinCost(a []int32, b []int32) int64 {
	var max int64 = 0
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	for i := len(a) - 1; i >= 0; i-- {
		max += int64(math.Abs(float64(b[i] - a[i])))
	}
	return max
}

//  SOLUTION 1

// func getMinCost(crew_id []int32, job_id []int32) int64 {
// 	// Write your code here
// 	dataC, dataJ := 10000000000, 10000000000
// 	var sum int64
// 	var maxOfA, maxofB int64
// 	for range crew_id {
// 		maxOfA = findmax(crew_id, int64(dataC))
// 		maxofB = findmax(job_id, int64(dataJ))
// 		dataC = int(maxOfA)
// 		dataJ = int(maxofB)
// 		sum += int64(math.Abs(float64(maxOfA - maxofB)))
// 	}
// 	return sum
// }

// func findmax(a []int32, data int64) int64 {
// 	var max int64
// 	j := 0
// 	for i, val := range a {
// 		if int64(val) > max && int64(val) <= data {
// 			max = int64(val)
// 			j = i
// 		}

// 	}
// 	a[j] = -1
// 	return max
// }
