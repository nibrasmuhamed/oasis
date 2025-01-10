package main

import "fmt"

func main() {
	fmt.Println(checkPalindrome("malaylalam"))
}

func checkPalindrome(s string) bool {
	j, pali := len(s)-1, false
	for i := 0; i < len(s)/2; i++ {
		if s[i] == s[j] {
			pali = true
			j--
		} else {
			pali = false
			break
		}
	}
	return pali
}
