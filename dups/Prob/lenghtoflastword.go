package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World "))
}

func lengthOfLastWord(s string) int {
	lenght := len(s) - 1
	count := 0
	sp := false
	for i := lenght; i >= 0; i-- {
		if unicode.IsSpace(rune(s[i])) {
			sp = true
			if count >= 1 && sp {
				break
			}
			continue

		}
		if unicode.IsLetter(rune(s[i])) {
			count++
		}
	}
	return count
}
