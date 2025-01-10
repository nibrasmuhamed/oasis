package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Example usage:
	// inputStr := "babad"
	// result := longestPalindrome(inputStr)
	// a := []int{1, 2, 1, 3, 4, 3}
	fmt.Println(subArrayWithDistinctElements([]int{1, 1, 3}))
}
func subArrayWithDistinctElements(A []int) int {
	mod := 1000000007
	l := 0
	m := make(map[int]int)
	ans := 0
	for r := 0; r < len(A); r++ {
		m[A[r]]++
		if m[A[r]] == 2 {
			for m[A[r]] > 1 {
				m[A[l]]--
				l++
			}
		}
		ans = ((ans % mod) + (r-l+1)%mod) % mod
	}
	return ans
}
func checkMap(m map[rune]int, r rune) bool {
	_, ok := m[r]
	return ok
}
func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]bool)
	start := 0
	res := 0

	for r := 0; r < len(s); r++ {
		for charSet[s[r]] {
			delete(charSet, s[start])
			start++
		}
		charSet[s[r]] = true
		res = max(res, r-start+1)
	}

	return res
}
func sortBasedOn(A []int, B []int) []int {
	m1 := make(map[int]int)
	// m2 := make(map[int]int)
	for _, v := range A {
		m1[v]++
	}
	out := []int{}
	prev := 0
	for _, v := range B {
		for i := prev; i <= v; i++ {
			val := m1[i]
			for j := 0; j < val; j++ {
				out = append(out, i)
			}
		}
		prev = v
	}
	return out
}
func colorful(A int) int {
	x := generateSubstrings(strconv.Itoa(A))
	if x {
		return 1
	}
	return 0
}

func generateSubstrings(s string) bool {
	substrings := map[int]bool{}
	n := len(s)

	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			y, _ := strconv.Atoi((s[i:j]))
			sum := 1
			for y != 0 {
				sum *= y % 10
				y = y / 10
			}
			_, ok := substrings[sum]
			if !ok {
				substrings[sum] = true
			} else {
				return false
			}
		}
	}
	return true
}
func longestSubarrayZero(A []int) int {
	prefix := prefixSum(A)
	ans := 0
	m := make(map[int]int)
	for i, v := range prefix {
		if v == 0 {
			_, ok := m[0]
			if ok {
				ans = max(ans, i-m[0]+1)
			} else {
				ans = max(ans, i+1)
			}
		}
		_, ok := m[v]
		if ok {
			ans = max(ans, i-m[v])
		} else {
			m[v] = i
		}
	}
	return ans
}

func prefixSum(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)

	if n == 0 {
		return prefix
	}

	// Calculate prefix sum
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	return prefix
}
func shaggy(A []int) int {
	m := make(map[int][]int)
	for i, v := range A {
		_, ok := m[v]
		if ok {
			m[v] = append(m[v], i)
		} else {
			m[v] = []int{i}
		}
	}
	minimum := 1<<31 - 1
	for _, v := range m {
		if len(v) >= 2 {
			for i := len(v) - 1; i > 0; i-- {
				k := v[i] - v[i-1]
				minimum = min(k, minimum)
			}
		}
	}
	if minimum == 1<<31-1 {
		return -1
	}
	return minimum
}
func intToRoman(a int) string {
	m := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	val := []int{1, 5, 10, 50, 100, 500, 1000}
	i := 6
	roman := ""
	for a > 0 {
		x := a / val[i]
		for j := x; j > 0; j-- {
			roman += m[val[i]]
		}
		a = a % val[i]
		i--
	}
	return roman
}
func dNums(A []int, B int) []int {
	n := len(A)
	out := []int{}
	if B > n {
		return []int{}
	}
	m := map[int]int{}
	for i := 0; i < B; i++ {
		m[A[i]]++
	}
	out = append(out, len(m))
	for i := B; i < len(A); i++ {
		v, ok := m[A[i-B]]
		if ok && v > 1 {
			m[A[i-B]]--
		} else {
			delete(m, A[i-B])
		}
		m[A[i]]++
		out = append(out, len(m))
	}
	return out
}

func countSubarraySum(A []int, B int) int {
	prefixSum := make(map[int]int)
	count := 0
	sum := 0

	// Initialize prefix sum and count for the sum = B
	prefixSum[0] = 1

	for _, num := range A {
		sum += num

		// Check if there is a subarray with sum - B
		if val, found := prefixSum[sum-B]; found {
			count += val
		}

		// Update the prefix sum count
		prefixSum[sum]++
	}

	return count
}

func diff(A []int, B int) int {
	m := make(map[int]int)
	for _, v := range A {
		m[v]++
	}
	count := 0
	for _, v := range A {
		if x, ok := m[v-B]; ok {
			// fmt.Println(v-B, "  ", v)
			count += x
			if 2*v == B {
				count--
			}
		}
	}
	return count
}

func longestPalindrome(s string) string {
	lastIndex := make(map[byte]int)
	start, maxLength := 0, 0

	for i := 0; i < len(s); i++ {
		if index, found := lastIndex[s[i]]; found && index >= start {
			start = index + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastIndex[s[i]] = i
	}

	return s[start : start+maxLength]
}
