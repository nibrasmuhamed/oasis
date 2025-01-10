package main

import "fmt"

func main() {
	a := []int{3, 2, 4}
	fmt.Println(twoSum(a, 6))
}
func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, val := range nums {
		_, ok := hash[val]
		if !ok {
			hash[val] = i
		}
		b := target - val
		_, ok = hash[b]
		if ok && hash[b] != i {
			return []int{hash[b], i}
		}
	}

	// fmt.Println(hash)
	return []int{0, 0}
}
