package main

import "fmt"

func main() {
	fmt.Println(replace("h", "a", "hehlo"))
}
func replace(a, b, s string) string {
	var newString, old string
	old = s
	for i := 0; i < len(s); i++ {
		if string(s[i]) == a {
			old = old[:i] + b + s[i+1:]
			newString = old
		}
		// fmt.Println(newString)
	}
	return newString
}
