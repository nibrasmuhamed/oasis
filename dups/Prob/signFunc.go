package main

import "fmt"

func main() {
	fmt.Println(arraySign([]int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}))
}

func arraySign(nums []int) int {
	ans := 1
	for _, v := range nums {
		if v == 0 {
			return 0
		}
		if v < 0 {
			ans *= -1
		}
	}
	return ans
}
