package main

import (
	"fmt"
)

func main() {
	a, b := "bba", "aas"
	fmt.Println(checkIsomorphic(a, b))
}

func checkIsomorphic(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	aHash := make(map[string]string)
	bHash := make(map[string]string)
	for i := 0; i < len(a); i++ {
		charA := a[i]
		charB := b[i]
		_, isPresent := aHash[string(charA)]
		if !isPresent {
			aHash[string(charA)] = string(charB)
		}
		_, isPresent = bHash[string(charB)]
		if !isPresent {
			bHash[string(charB)] = string(charA)
		}
		fmt.Println(aHash, bHash)
		if aHash[string(charA)] != string(charB) || bHash[string(charB)] != string(charA) {
			return false
		}

	}
	return true
}
