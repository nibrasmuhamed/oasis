package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	var maxProfit int
	previousLowestPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		if price < previousLowestPrice {
			previousLowestPrice = price
		} else {
			profit := price - previousLowestPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}
