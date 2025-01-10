package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(isAnagram("anagram", "naagram"))
}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make([]int, 26)
	for _, ch := range s {
		fmt.Printf("tyep is %T", ch)
		ch -= 'a'
		fmt.Println(reflect.TypeOf('a'))

		count[ch]++
	}
	for _, ch := range t {
		ch -= 'a'
		count[ch]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}
