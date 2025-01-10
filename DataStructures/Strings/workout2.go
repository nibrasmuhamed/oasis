package main

import "fmt"

func main() {
	fmt.Println(findFirst("abca"))
}

func findFirst(s string) int {
	set := make(map[string]int)
	for i := 0; i < len(s); i++ {
		_, isPresent := set[string(s[i])]
		if !isPresent {
			set[string(s[i])] = 1
		} else {
			set[string(s[i])] = set[string(s[i])] + 1
		}
	}
	for i := 0; i < len(s); i++ {
		if set[string(s[i])] == 1 {
			return i
		}
	}
	return -1
}
