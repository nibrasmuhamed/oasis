package main

func main() {
	// a := [][]int{
	// 	{5, 10},  //0  [5,10],[3,6]
	// 	{10, 10}, //1  [10,10],[9,9]
	// 	{3, 6},   //2
	// 	{9, 9},   //3
	// 	{15, 15}, //4
	// }

	// fmt.Println(similarRectangle(a))
}

func similarRectangle(a [][]int) int64 {
	hash := make(map[float32][][]int)
	var sum int64
	for i := 0; i <= len(a)-1; i++ {
		_, ok := hash[float32(a[i][0])/float32(a[i][1])]
		if !ok {
			hash[float32(a[i][0])/float32(a[i][1])] = append(hash[float32(a[i][0])/float32(a[i][1])], a[i])
		} else {
			hash[float32(a[i][0])/float32(a[i][1])] = append(hash[float32(a[i][0])/float32(a[i][1])], a[i])
		}
	}
	// for i, val := range hash {
	// 	if len(hash[i]) >= 2 {
	// 		fmt.Println(i, "         ", val)
	// 	}
	// }
	for _, val := range hash {
		if len(val) == 1 {
			continue
		} else if len(val) < 3 {
			sum = sum + 1
		} else if len(val) == 3 {
			sum = sum + 3
		} else {
			x := len(val) - 1
			for i := 0; i <= x; i++ {
				sum = sum + int64(i)
			}
		}
	}
	return sum
}

// time complexity is O(n^2)

// func similarRectangle(a [][]int) int {
// 	sum := 0
// 	for i := 0; i <= len(a)-1; i++ {
// 		x := a[i]
// 		for j := i + 1; j <= len(a)-1; j++ {
// 			y := a[j]
// 			if x[0] == y[0] && x[1] == y[1] {
// 				continue
// 			}
// 			if float32(x[0])/float32(y[0]) == float32(x[1])/float32(y[1]) {
// 				sum++
// 				fmt.Println(x[0], x[1], y[0], y[1])
// 			}
// 		}
// 	}
// 	return sum
// }
