package main

// func main() {

// 	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	fmt.Println(binIterative(a, 6))
// }

func binIterative(a []int, val int) int {
	start := 0
	end := len(a) - 1
	middle := 0
	for start <= end {
		middle = (start + end) / 2
		if a[middle] == val {
			return middle
		}
		if a[middle] > val {
			end = middle - 1
		} else if a[middle] < val {
			start = middle + 1
		}

	}
	return -1
}
