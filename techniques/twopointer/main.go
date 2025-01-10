package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	a := []int{1, 2, 2, 3, 3, 3, 4, 4, 5} // 1, 2, 3, 4, 5
	fmt.Println(pairWithGivenSum(a, 6))
}
func pairWithGivenSum(A []int, B int) int {
	mod := 1000000007
	i, j := 0, len(A)-1
	count := 0
	for i < j {
		if A[i]+A[j] == B {
			if A[i] == A[j] {
				c := 0
				for j >= i {
					j--
					c++
				}
				count = ((count % mod) + (((c * (c - 1)) / 2) % mod)) % mod
				break
			}
			iCount, jCount := 1, 1
			for i < len(A) && A[i] == A[i+1] {
				iCount++
				i++
			}
			for j >= 0 && A[j] == A[j-1] {
				jCount++
				j--
			}
			count = ((count % mod) + ((iCount * jCount) % mod)) % mod
		}
		if A[i]+A[j] > B {
			j--
		} else {
			i++
		}
	}
	return count
}
func rectArea(A []int, B int) int {
	i, j := 0, len(A)-1
	count := 0
	for i < j {
		if A[i]*A[j] < B {
			count += (j - i) * 2
			// fmt.Println(A[i:j+1], " ", A[j])
		}
		if A[i]*A[j] >= B {
			j--
		} else {
			i++
		}
	}
	for _, v := range A {
		if v*v < B {
			count++
		}
	}
	return count
}

func threeSumClosest(A []int, B int) int {
	sort.Ints(A)
	n := len(A)
	sum := 0
	diff := 9999999999999
	for i := 0; i < len(A)-2; i++ {
		v1 := A[i]
		j := i + 1
		k := n - 1
		for j < k {
			closeSum := v1 + A[j] + A[k]
			if closeSum == B {
				return B
			}
			CurrentDiff := absDiff(B, closeSum)
			if diff > CurrentDiff {
				sum = closeSum
				diff = CurrentDiff
			}
			if closeSum > B {
				k--
			} else {
				j++
			}
		}
	}
	return sum
}

func absDiff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func countRectangles(A []int, B int) int {
	i, j := 0, len(A)-1
	count := 0
	for i <= len(A)-1 && j >= 0 {
		var q int64 = int64(A[i] * A[j])
		if q < int64(B) {
			count += j + 1
			i++
		} else {
			j--
		}
	}
	return count
}

func countPairs(A []int, B int) int {
	p1, p2 := 0, 1
	count := 0
	for p2 < len(A)-1 && p1 < len(A)-1 {
		if A[p1]+A[p2] == B {
			count++
			j, k := 0, 0
			for p1 < len(A)-1 && A[p1] == A[p1+1] {
				j++
				p1++
			}
			for p2 < len(A)-1 && A[p2] == A[p2+1] {
				k++
				p2++
			}
			count += j * k
			if j == 0 {
				p1++
			}
			if k == 0 {
				p2++
			}
		}
		if A[p1]+A[p2] > B {
			p1++
		}
		if A[p1]+A[p2] < B {
			p2++
		}
	}
	return count % 1000000007
}

func subArraySum(A []int, B int) []int {
	i, j := 0, 1
	sum := A[i]
	for j < len(A) {
		if sum == B {
			return A[i:j]
		}
		if sum > B {
			sum -= A[i]
			i++
		} else {
			sum += A[j]
			j++
		}
	}
	if sum == B {
		return A[i:j]
	}
	return []int{-1}
}

func absoluteDifference(A []int, B int) int {
	sort.Ints(A)
	i := 0
	j := 1
	ans := 0
	for j < len(A) {
		if j == i {
			j++
			continue
		}
		x := A[i]
		y := A[j]
		if y-x == B {
			// count the pair A[i], A[j] only once
			ans++
			for i < len(A) && A[i] == x {
				i++
			}
			for j < len(A) && A[j] == y {
				j++
			}
		} else if y-x > B {
			i++
		} else {
			j++
		}
	}
	return ans
}

func subArrayWithGivenSum(A []int, B int) []int {
	target := 0
	p1, p2 := 0, 1
	target = A[p1] + A[p2]
	for p2 < len(A) {
		if target == B {
			return A[p1 : p2+1]
		}
		if target > B {
			// remove the left element and try again
			target -= A[p1]
			p1 += 1
		} else {
			// add the right element and try again
			p2 += 1
			if p2 >= len(A) {
				continue
			}
			target += A[p2]
		}
	}
	return []int{-1}
}
