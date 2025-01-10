package main

import (
	"fmt"
	"sort"
)

// time complexity is O(n log n) * s times (s refers to total number of inputs)
// we are sorting each string O(n log n), iterating through slice O(n)
// upper bound is O s*(n log n)

// space complexity:
// created map O(s*n), slice O(s*n)
// complexity is O s(s*n)

func main() {
	a := []string{"abc", "car", "cat", "arc", "act", "tca"}
	fmt.Println(anagram(a))
}

func anagram(a []string) [][]string {
	hashMap := make(map[string][]string)

	b := make([]string, len(a))
	// c := make([][]string, 0)
	var c [][]string
	for i, val := range a {
		x := Alphbetval(val)
		sort.Sort(x)
		b[i] = string(x)
		_, ok := hashMap[string(x)]
		if !ok {
			hashMap[b[i]] = append(hashMap[b[i]], a[i])
		} else {
			hashMap[b[i]] = append(hashMap[b[i]], a[i])
		}
	}
	for _, val := range hashMap {
		c = append(c, val)
	}
	return c
}

type Alphbetval []rune

func (a Alphbetval) Len() int           { return len(a) }
func (a Alphbetval) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Alphbetval) Less(i, j int) bool { return a[i] < a[j] }
