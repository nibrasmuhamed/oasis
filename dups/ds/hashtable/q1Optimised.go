package main

import "fmt"

func main() {
	var a = []int{9, 4, 5, 4, 8, 1, 6, 7}
	fmt.Println(addup(a, 8))
}

func addup(a []int, target int) (int, int) {
	mySet := make(map[int]int)
	for i := 0; i < len(a); i++ {
		_, isPresent := mySet[a[i]]
		if !isPresent {
			mySet[a[i]] = i
		}
		_, pres := mySet[target-a[i]]
		if pres && i != mySet[target-a[i]] && a[i] != target-a[i] {
			return i, mySet[target-a[i]]
		}
	}
	return 0, 0
}
