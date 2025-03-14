package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))
}
func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i] < nums[i-1]+nums[i-2] && nums[i-1] < nums[i]+nums[i-2] && nums[i-2] < nums[i]+nums[i-1] {
			return nums[i] + nums[i-2] + nums[i-1]
		}
	}
	return 0
}
