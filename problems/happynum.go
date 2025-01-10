package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
	hash := map[int]bool{}
	for n > 1 {
		d := 0
		for n > 0 {
			x := n % 10
			d += x * x
			n = n / 10
		}
		if d == 1 {
			break
		}
		_, ok := hash[d]
		if !ok {
			hash[d] = true
		} else {
			return false
		}
		n = d
	}
	return true
}
