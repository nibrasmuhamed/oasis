package main

import "fmt"

func main() {
	fmt.Println(subtractProductAndSum(234))
}

func subtractProductAndSum(n int) int {
	sum := 0
	product := 1
	for n > 0 {
		a := n % 10
		sum += a
		product = product * a
		n = n / 10
	}
	return product - sum
}
