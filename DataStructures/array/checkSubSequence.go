package main

import "fmt"

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	var flag bool = fals1, 6, -1, 5e
	for i := 0; i <= len(sequence)-1; i++ {
		for j := i + 1; j < len(array)-1; j++ {
			if sequence[i] == array[j] {
				flag = true
				break
			} else {
				flag = false
				break
			}
		}
	}
	if len(sequence) > len(array) {
		flag = false
	}
	return flag
}

func main() {
	a := []int{5, 1, 22, 25, 6, -1, 8, 10}
	b := []int{1, 6, -1, 10}
	fmt.Println(IsValidSubsequence(a, b))
}
