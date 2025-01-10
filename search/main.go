package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = []int{1}
	fmt.Println(searchTwo(a, 1))
	// fmt.Println(books(a, 26))
}
func searchTwo(A []int, B int) []int {
	out := make([]int, 2)
	s, e := 0, len(A)-1
	for s <= e {
		mid := (s + e) / 2
		if A[mid] == B && (mid <= 0 || A[mid-1] != B) {
			out[0] = mid
			break
		}
		if A[mid] >= B {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	s, e = 0, len(A)-1
	for s <= e {
		mid := (s + e) / 2
		if A[mid] == B && (mid >= len(A)-1 || A[mid+1] != B) {
			out[1] = mid
			break
		}
		if A[mid] > B {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	return out
}

func majorityElement(nums []int) int {
	majority := nums[0]
	count := 0
	for i := 0; i < len(nums); i++ {
		if majority != nums[i] {
			count--
		}
		if count <= 0 {
			majority = nums[i]
		}
		if majority == nums[i] {
			count += 1
		}
	}
	return majority
}
func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	freqMap := make(map[int]int)

	// Create frequency map for nums1
	for _, num := range nums1 {
		freqMap[num]++
	}

	// Iterate through nums2 and check for intersection
	for _, num := range nums2 {
		if freqMap[num] > 0 {
			result = append(result, num)
			freqMap[num]--
		}
	}

	return result
}
func books(A []int, B int) int {
	low, high := 9223372036854775807, 0
	result := -1

	for i := 0; i < len(A); i++ {
		if A[i] < low {
			low = A[i]
		}
		high += A[i]
	}
	for low <= high {
		mid := (low + high) / 2
		cnt := findNumberOfStudents(A, mid)
		if cnt <= B {
			result = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return result
}

func findNumberOfStudents(V []int, curMin int) int {
	students, curSum := 1, 0
	for i, _ := range V {
		if curSum+V[i] > curMin {
			students++
			curSum = V[i]
		} else {
			curSum += V[i]
		}
	}
	return students
}
func subArraySumProblem(A []int, B int) int {
	s, e := 1, len(A)
	res := 0
	for s <= e {
		mid := (s + e) / 2
		if findNumberOfSubArrays(A, B, mid) >= 1 {
			e = mid - 1
		} else {
			res = mid
			s = mid + 1
		}
	}
	return res
}

func findNumberOfSubArrays(a []int, maxSum, length int) int {
	// if len(a) < length {
	// 	return 1
	// }

	counter := 0
	sum := 0
	for i := 0; i < length; i++ {
		sum += a[i]
	}
	if sum > maxSum {
		counter++
	}
	for i := length; i < len(a); i++ {
		sum -= a[i-length]
		sum += a[i]
		if sum > maxSum {
			counter++
		}
	}
	return counter
}

func aggressiveCows(A []int, B int) int {
	sort.Ints(A)
	s, e := 0, A[len(A)-1]-A[0]
	res := 0
	for s <= e {
		mid := (s + e) / 2
		if numberOfCows(A, mid) >= B {
			res = mid
			s = mid + 1
		} else {
			e = mid - 1
		}
	}
	return res
}

func numberOfCows(A []int, B int) int {
	count, last := 1, A[0]
	for i := 1; i < len(A); i++ {
		if A[i]-last >= B {
			count++
			last = A[i]
		}
	}
	return count
}

func paintersProblemBinarySearch(A, B int, C []int) int {
	const mod = 10000003
	s, e := 0, 0

	// Find the maximum and total time required based on the painters' speed
	for _, v := range C {
		if v*B > s {
			s = v * B
		}
		e += v * B
	}

	res := 0

	// Binary search to find the minimum time required
	for s <= e {
		mid := (s + e) / 2
		painters := countPainters(mid, B, C)

		if painters <= A {
			res = mid
			e = mid - 1
		} else {
			s = mid + 1
		}
	}

	return res % mod
}
func countPainters(totalTime int, unitOfTime int, brands []int) int {
	painters, x := 1, 0
	for i := 0; i < len(brands); i++ {
		// if brands[i]*unitOfTime > totalTime {
		// 	return -1
		// }
		x += (brands[i] * unitOfTime)
		if x > totalTime {
			painters++
			x = brands[i] * unitOfTime
		}
	}
	return painters
}
func findMin(nums []int) int {
	s, e := 0, len(nums)-1

	for s <= e {
		m := (s + e) / 2

		// Check if the current element is smaller than both its neighbors or if it's the only element
		if (m == 0 || nums[m] < nums[m-1]) && (m == len(nums)-1 || nums[m] < nums[m+1]) {
			return nums[m]
		}

		// If the middle element is greater than the last element, search in the right half
		if nums[m] > nums[e] {
			s = m + 1
		} else {
			// Otherwise, search in the left half
			e = m - 1
		}
	}

	// The array is already sorted, return the first element
	return nums[0]
}

func nthnumberoptimise(a, b, c int) int {
	s, e := 1, a*min(b, c)
	result := 0
	for s <= e {
		mid := (s + e) / 2
		p := int(mid/b) + int(mid/c) - (mid / lcm(b, c))

		if p == a {
			result = mid
			e = mid - 1
		} else if p > a {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	return result
}

func lcm(a, b int) int {
	d := (a * b) / gcd(a, b)
	return d
}

func firstBadVersion(n int) int {
	s, e := 1, n
	for s <= e {
		m := (s + e) / 2
		if isBadVersion(m) && !isBadVersion(m-1) {
			return m
		}
		if isBadVersion(m) {
			e = m - 1
		} else {
			s = m + 1
		}
	}
	return 0
}

func isBadVersion(i int) bool {
	return i == 4
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func nthnumberrecursion(a, b, c, d, e int) int {
	if d == a {
		return e - 1
	}
	if e%b == 0 || e%c == 0 {
		return nthnumberrecursion(a, b, c, d+1, e+1)
	}
	return nthnumberrecursion(a, b, c, d, e+1)
}

func nthnumber(A, B, C int) int {
	s, i := 0, 0
	for s < A {
		i++
		if i%B == 0 || i%C == 0 {
			s++
		}
	}
	return i
}

func oddone(a []int) int {
	x := 0
	for _, v := range a {
		x ^= v
	}
	return x
}
func sqrt(A int) int {
	s, e := 1, (A)
	for s <= e {
		mid := (s + e) / 2
		if mid*mid == A || ((mid+1)*(mid+1) > A && (mid-1)*(mid-1) < A) {
			if mid*mid <= A {
				return mid
			} else {
				return mid - 1
			}
		}
		if mid*mid > A {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	return 0
}

func findIndex(a []int, b int) int {
	s, e := 0, len(a)-1
	for s <= e {
		mid := (s + e) / 2
		if (mid == 0 || a[mid-1] > a[mid]) && (mid == len(a)-1 || a[mid+1] > a[mid]) {
			s = mid
			break
		}
		if a[mid] >= a[0] {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}
	if b >= a[0] {
		e = s - 1
		s = 0
	} else {
		e = len(a) - 1
	}
	for s <= e {
		mid := (s + e) / 2
		if a[mid] == b {
			return mid
		}
		if a[mid] > b {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	return -1
}
func searchMatrix(A [][]int, B int) int {
	k := 0
	s, e := 0, len(A)-1
	for s <= e {
		mid := s + (e-s)/2
		if s == 0 && e == 0 {
			k = 0
			break
		}
		if B >= A[mid][0] && B <= A[mid][len(A[mid])-1] {
			k = mid
			break
		}
		if B > A[mid][len(A[mid])-1] {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}
	s, e = 0, len(A[k])-1

	for s <= e {
		mid := (s + e) / 2
		if A[k][mid] == B {
			return 1
		}
		if A[k][mid] > B {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	return 0
}
