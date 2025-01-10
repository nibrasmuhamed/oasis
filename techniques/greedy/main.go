package main

import (
	heaps "container/heap"
	"fmt"
	"math"
	"sort"
)

type MinHeapInt []int

func (h MinHeapInt) Len() int           { return len(h) }
func (h MinHeapInt) Less(i, j int) bool { return h[i] < h[j] }
func (h *MinHeapInt) Swap(i, j int) {
	if len(*h) > 0 {
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
	}
}

func (h *MinHeapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeapInt) Pop() interface{} {
	if len(*h) == 0 {
		return -1
	}
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findMinimumNoOfCoins(A int) int {
	ans := 0
	for A > 0 {
		val := findPowerOf5(A)
		ans += A / val
		A = A % val
	}
	return ans
}

func findPowerOf5(value int) int {
	if value <= 0 {
		return 0
	}

	exponent := int(math.Log(float64(value)) / math.Log(5))
	powerOf5 := int(math.Pow(5, float64(exponent)))

	return powerOf5
}
func MaxNumberOfJobs(A []int, B []int) int {
	maxJobs := 0
	sortByA(B, A)
	lastJob := 0
	for i := 0; i < len(B); i++ {
		if A[i] > lastJob {
			maxJobs += 1
			lastJob = B[i]
		}
	}
	return maxJobs
}

func greedyCarSelling(A []int, B []int) int {
	sortByA(A, B)
	mod := 1000000007
	h := &MinHeapInt{}
	res := 0
	t := 0
	for i := 0; i < len(B); i++ {
		if A[i] > t {
			heaps.Push(h, B[i])
			res = ((res % mod) + (B[i] % mod)) % mod
			t++
		} else {
			j := heaps.Pop(h).(int)
			if B[i] > j {
				res -= j
				res = ((res % mod) + (B[i] % mod)) % mod
				heaps.Push(h, B[i])
			} else {
				heaps.Push(h, j)
			}
		}
	}
	return res
}

type Pair struct {
	ValueA int
	ValueB int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].ValueA < p[j].ValueA }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sortByA(A, B []int) {
	pairs := make(PairList, len(A))

	for i := range A {
		pairs[i] = Pair{ValueA: A[i], ValueB: B[i]}
	}

	sort.Sort(pairs)

	for i := range A {
		A[i] = pairs[i].ValueA
		B[i] = pairs[i].ValueB
	}
}
func candy(A []int) int {
	arr := make([]int, len(A))
	n := len(A)
	for i := 0; i < n; i++ {
		arr[i] = 1
	}
	for i := 1; i < n; i++ {
		if A[i] > A[i-1] {
			arr[i] = arr[i-1] + 1
		}
	}
	ans := arr[n-1]
	for i := n - 2; i >= 0; i-- {
		if A[i] > A[i+1] && arr[i] <= arr[i+1] {
			arr[i] = arr[i+1] + 1
		}
		ans += arr[i]
	}
	return ans
}

func main() {
	a := []int{-4, 2, 3}
	b := []int{0, -2, 4}
	fmt.Println(mice(a, b))
}

func mice(A []int, B []int) int {
	// A Position of mice.
	// B is the position of holes.
}

// made a great progress here!
func seats(A string) int {
	mod := 10000003
	out := []int{}
	for i := 0; i < len(A); i++ {
		if A[i] == byte('x') {
			out = append(out, i)
		}
	}
	med := len(out) / 2
	left, right := med-1, med+1
	x := 1
	ans := 0
	for left >= 0 {
		ans = (ans + (out[med]-out[left]-x)%mod + mod) % mod
		left--
		x++
	}
	x = 1
	for right < len(out) {
		ans = (ans + (out[right]-out[med]-x)%mod + mod) % mod
		right++
		x++
	}
	return ans
}
