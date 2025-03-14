package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 13))
}

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		middle := (start + end) / 2
		if nums[middle] == target {
			return middle
		}
		if target > nums[middle] {
			start = middle + 1
		} else if target < nums[middle] {
			end = middle - 1
		}

	}
	return -1
}
