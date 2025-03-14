package main

import (
	"fmt"
	"math"
)

func main() {
	x := 0
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	nums := []int{}
	if math.Signbit(float64(x)) {
		return false
	}
	if x == 0 {
		return true
	}
	for x > 0 {
		z := x % 10
		nums = append(nums, z)
		x = x / 10
	}
	j, palindrome := len(nums)-1, false
	if len(nums) == 1 {
		return true
	}
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] == nums[j] {
			palindrome = true
		} else {
			palindrome = false
			break
		}
		j--
	}
	return palindrome
}
