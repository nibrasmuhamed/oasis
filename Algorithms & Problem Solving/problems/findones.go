package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var x uint32 = 00000000000000000000000000001011
	fmt.Println(hammingWeight(x))
}
func hammingWeight(num uint32) int {
	return bits.OnesCount32(num)su
}
