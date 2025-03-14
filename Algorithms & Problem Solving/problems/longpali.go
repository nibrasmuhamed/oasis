package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("bananas"))
}

func longestPalindrome(s string) int {
	length := 0
	z := true
	hash := make(map[string]int)
	for _, val := range s {
		_, ok := hash[string(val)]
		if !ok {
			hash[string(val)] = 1
		} else {
			hash[string(val)] = hash[string(val)] + 1
		}
	}
	if len(hash) == 1 {
		for _, val := range hash {
			return val
		}
	}
	for x, val := range hash {
		if val%2 == 0 {
			length = length + val
			continue
		}
		if val%2 != 0 {
			length = length + (val - 1)
			hash[x] = val - 1
		}
		if z {
			length = length + 1
			z = false
		}
	}
	return length
}
