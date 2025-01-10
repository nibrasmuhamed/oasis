package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 3))
}

func topKFrequent(nums []int, k int) []int {
	hash := map[int]int{}
	out := []int{}
	freq := make([][]int, len(nums))
	for _, n := range nums {
		_, ok := hash[n]
		if !ok {
			hash[n] = 1
		} else {
			hash[n]++
		}
	}
	for i, v := range hash {
		freq[v] = append(freq[v], i)
	}
	for i := len(freq) - 1; i >= 0; i-- {
		for _, val := range freq[i] {
			out = append(out, val)
			if len(out) == k {
				return out
			}
		}
	}
	return out
}
