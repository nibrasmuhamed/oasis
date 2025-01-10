package main

import "fmt"

func main() {
	a := []int{1, 3, 2, 3, 3, 4, 3, 3}
	fmt.Println(removeElement(a, 3))
}

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}
