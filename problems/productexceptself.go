package main

import (
	"fmt"
)

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	out := make([]int, len(nums))
	out[0] = 1
	for i := 1; i < len(nums); i++ {
		out[i] = out[i-1] * nums[i-1]
	}
	post := 1
	for i := len(nums) - 1; i >= 0; i-- {
		out[i] = post * out[i]
		post = post * nums[i]
	}
	return out
}

// func productExceptSelf(nums []int) []int {
// 	out := []int{}
// 	for i := 0; i < len(nums); i++ {
// 		sum := 1
// 		for j := 0; j < len(nums); j++ {
// 			if i == j {
// 				continue
// 			}
// 			sum = sum * nums[j]
// 		}
// 		out = append(out, sum)
// 	}
// 	return out
// }
