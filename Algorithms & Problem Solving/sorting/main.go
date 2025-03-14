package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	// x := []int{3, 1, 2, 4, 0}
	// a := [][]int{
	// 	{9, -10},
	// 	{1, -4},
	// 	{10, -4},
	// 	{5, -8},
	// 	{-9, -1},
	// 	{-7, 7},
	// }
	// quickSort(x, 0, len(x)-1)
	// fmt.Println(sortMat(a))
	b := map[int]int{1: 10}
	change(b)
	fmt.Println(b)
}
func change(a map[int]int) {
	a[1] = 100
}
func sortMat(points [][]int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		dA := dist(points[i][0], points[i][1])
		dB := dist(points[j][0], points[j][1])
		return dA < dB || (dA == dB && (points[i][0] < points[j][0] || (points[i][0] == points[j][0] && points[i][1] < points[j][1])))
	})
	return points
}

func dist(x, y int) float64 {
	return math.Sqrt(float64(x*x + y*y))
}
func demo() bool {
	return "345"+"3" > "3"+"345"
}

var compareByAge = func(i, j int) bool {
	return strconv.Itoa(i)+strconv.Itoa(j) < strconv.Itoa(j)+strconv.Itoa(i)
}

// Use sort.Slice with the comparator function
// sort.Slice(people, compareByAge)

func chunk(A []int) int {
	q := -9223372036854775801
	count := 0
	for i := 0; i < len(A); i++ {
		q = max(q, A[i])
		if q == i {
			count++
		}
	}
	return count
}

func quickSort(A []int, s, e int) {
	if s >= e {
		return
	}
	index := partition(A, s, e)
	quickSort(A, s, index-1)
	quickSort(A, index+1, e)
}

func partition(A []int, start, end int) int {
	pivot := A[start]
	s, e := start+1, end

	for s <= e {
		if A[s] < pivot {
			s++
		} else if A[e] > pivot {
			e--
		} else {
			A[s], A[e] = A[e], A[s]
			s++
			e--
		}
	}

	// Swap pivot with the element at index e (the final position of the pivot)
	A[start], A[e] = A[e], A[start]

	// Return the index of the pivot in its final sorted position
	return e
}

func solvex(A []int) int {
	x := math.MaxInt
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if i == j {
				continue
			}
			x = min(x, abs(A[i]-A[j], A[j]-A[i]))
		}
	}
	return x
}

func abs(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func mergeSortf(arr []int, low, high int) int {
	var j, v, s int
	if low < high {
		mid := (low + high) / 2

		s = mergeSort(arr, low, mid)
		j = mergeSort(arr, mid+1, high)

		v = mergx(arr, low, mid, high)
	}
	return v + j + s
}

func mergx(arr []int, low, mid, high int) int {
	leftSize := mid - low + 1
	rightSize := high - mid

	// Create temporary slices for the left and right halves
	left := make([]int, leftSize)
	right := make([]int, rightSize)
	count := 0
	// Copy data to temporary slices left[] and right[]
	for i := 0; i < leftSize; i++ {
		left[i] = arr[low+i]
	}
	for j := 0; j < rightSize; j++ {
		right[j] = arr[mid+1+j]
	}

	// Merge the temporary slices back into arr[low..high]
	i, j, k := 0, 0, low
	for i < leftSize && j < rightSize {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
			count += (mid - low + 1) - i

		}
		k++
	}

	// Copy any remaining elements of left[], if there are any
	for i < leftSize {
		arr[k] = left[i]
		i++
		k++
	}

	// Copy any remaining elements of right[], if there are any
	for j < rightSize {
		arr[k] = right[j]
		j++
		k++
	}
	return count
}

func mergesort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	s, e := 0, len(a)-1
	m := (e + s) / 2
	f := mergesort(a[:m+1])
	l := mergesort(a[m+1:])
	return merge(f, l)
}

func merge(A, B []int) []int {
	n, m := len(A), len(B)
	i, j := 0, 0
	out := []int{}
	for i < n && j < m {
		if A[i] <= B[j] {
			out = append(out, A[i])
			i++
		} else {
			out = append(out, B[j])
			j++
		}
	}
	for i < n {
		out = append(out, A[i])
		i++
	}
	for j < m {
		out = append(out, B[j])
		j++
	}
	return out
}

func solve(A []int) int {
	x := mergeSort(A, 0, len(A)-1)
	return x
}
func mergeSort(arr []int, low, high int) int {
	var j, v, s int
	if low < high {
		mid := (low + high) / 2

		s = mergeSort(arr, low, mid)
		j = mergeSort(arr, mid+1, high)

		v = merg(arr, low, mid, high)
	}
	return v + j + s
}

func merg(arr []int, low, mid, high int) int {
	leftSize := mid - low + 1
	rightSize := high - mid

	// Create temporary slices for the left and right halves
	left := make([]int, leftSize)
	right := make([]int, rightSize)
	count := 0
	q := mid + 1
	for i := low; i <= mid; i++ {
		for q <= high && arr[i] > (2*arr[q]) {
			q++
		}
		count += q - (mid + 1)
	}
	// Copy data to temporary slices left[] and right[]
	for i := 0; i < leftSize; i++ {
		left[i] = arr[low+i]
	}
	for j := 0; j < rightSize; j++ {
		right[j] = arr[mid+1+j]
	}

	// Merge the temporary slices back into arr[low..high]
	i, j, k := 0, 0, low
	for i < leftSize && j < rightSize {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	// Copy any remaining elements of left[], if there are any
	for i < leftSize {
		arr[k] = left[i]
		i++
		k++
	}

	// // Copy any remaining elements of right[], if there are any
	for j < rightSize {
		arr[k] = right[j]
		j++
		k++
	}
	return count
}
