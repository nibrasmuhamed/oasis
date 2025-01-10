package main

import (
	"fmt"
	"strconv"
)

func getMax(operations []string) []int32 {
	// Write your code here
	type Stack []int32
	stack := Stack{}
	maxStack := Stack{}
	ret := Stack{}
	var maxx int32
	operation := operations[1:]
	for _, value := range operation {
		if rune(value[0]) == rune('1') {
			newVal := value[2:]
			res, _ := strconv.ParseInt(newVal, 10, 32)
			stack = append(stack, int32(res))
			if len(maxStack) == 0 {
				maxStack = append(maxStack, int32(res))
			}
			if int32(res) > maxStack[len(maxStack)-1] {
				maxStack = append(maxStack, int32(res))
			}
			continue
		}
		if rune(value[0]) == rune('2') {
			length := len(stack)
			if len(stack) <= 0 {
				continue
			}
			stack = stack[:len(stack)-1]
			maxStack = maxStack[0 : length-1]
			continue
		}
		if rune(value[0]) == rune('3') {
			maxx = maxStack[len(maxStack)-1]
			ret = append(ret, maxx)
		}
	}
	// if maxx == stack[len(stack)-1] {
	// 	return stack
	// } else {
	// 	stack = append(stack, maxx)
	// }
	return ret
}

func main() {
	// a := []string{"10", "2", "1 97", "2", "1 20", "2", "1 26", "1 20", "2", "3", "1 91", "3"}
	// a := []string{"2", "1 83", "3", "2", "1 76"}
	fmt.Println(getMax(q))
}

var q = []string{"100000",
	"1 1",
	"3",
	"1 2",
	"3",
	"1 3",
	"3",
	"1 4",
	"3",
	"1 5",
	"3",
	"1 6",
	"3",
	"1 7",
	"3",
	"1 8",
	"3",
	"1 9",
	"3",
	"1 10",
}
