package main

import "fmt"

func main() {
	var a = [10]int{1, 8, 3, 17, 5, 0, 70, 8, 9}
	// array time and space complexities:
	// init O(1)ST
	// get O(1)ST
	// set O(1)ST
	// traverse O(n)T O(1)S
	// insert O(n)T O(1)S
	// delete O(1)ST
	findRes(a, 10)
	fmt.Println("------------------------------")
	findResT(a, 10)
}

// this function will find two combination which can add to get sum.
// sum should be passed as second argument
// time complexity is O(n^2) and space complexity is O(1)
func findRes(a [10]int, sum int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i; j < len(a)-1; j++ {
			if a[i]+a[j] == sum {
				fmt.Println(a[i], " and ", a[j])
				// return a[i], a[j]
			}
		}
	}
}

func findResT(a [10]int, sum int) {

	var exists = struct{}{}
	mySet := make(map[int]struct{})
	for _, val := range a {
		if _, ok := mySet[val]; ok {
			mySet[val] = exists // This won't be printed
		}
	}
	for i, val := range mySet {
		fmt.Println(i, " and ", val)
	}
}
