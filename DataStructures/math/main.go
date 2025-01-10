package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// fmt.Println(solve([]int{12, 15, 18}))

	// fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	// fmt.Println(convertToNumber("AB"))

	fmt.Println(isPowerOfTwo(12))
}
func isPowerOfTwo(n int) bool {
	return po(n, 0)
}
func po(n, a int) bool {
	if power(2, a) == n {
		return true
	}
	if power(2, a) > n {
		return false
	}
	return po(n, a+1)
}
func reverseString(s string) string {
	return rev(s)
}

func rev(s string) string {
	if len(s) <= 1 {
		return s
	}
	return s[len(s)-1:] + rev(s[:len(s)-1])
}

func t(i int) {

}

func palindrome(s string, b, e int) bool {
	if b >= e {
		return true
	}
	if s[b] != s[e] {
		return false
	}
	return palindrome(s, b+1, e-1)
}
func validParenthesis(A int) []string {
	return valpar(A, 0, 0, "")
}

func valpar(n, o, c int, s string) []string {
	result := make([]string, 0)

	if len(s) == 2*n {
		result = append(result, s)
	}

	if o < n {
		result = append(result, valpar(n, o+1, c, s+"(")...)
	}
	if c < o {
		result = append(result, valpar(n, o, c+1, s+")")...)
	}

	return result
}
func toh(n int, source, destination, helper string) {
	if n == 0 {
		return
	}
	toh(n-1, source, helper, destination)
	fmt.Println(n, " is "+source, destination)
	toh(n-1, helper, destination, source)
}
func convertToNumber(n string) int {
	result := 0
	j := 0
	for i := len(n) - 1; i >= 0; i-- {
		result += int(((n[i] - 'A') + 1)) * int(power(26.0, j))
		j++
	}
	return result
}

func power(a, b int) int {
	res := 1
	for b != 0 {
		if b&1 == 1 {
			res *= a
		}
		a *= a
		b >>= 1
	}
	return res
}

func convertToTitle(n int) string {
	result := ""
	for n > 0 {
		remainder := (n - 1) % 26
		result = string('A'+remainder) + result
		n = (n - 1) / 26
	}
	return result
}

func findRank(A string) int {
	result := []string{}
	permute([]rune(A), 0, &result)
	sort.Strings(result)
	for i := 0; i < len(result); i++ {
		if A == result[i] {
			return i + 1
		}
	}
	return -1
}

func permute(s []rune, start int, result *[]string) {
	if start == len(s)-1 {
		// Convert the rune slice to a string and add to the result
		*result = append(*result, string(s))
		return
	}

	for i := start; i < len(s); i++ {
		// Swap characters
		s[start], s[i] = s[i], s[start]
		// Recursively generate permutations
		permute(s, start+1, result)
		// Swap back to backtrack
		s[start], s[i] = s[i], s[start]
	}
}

func binaryExponentiation(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res *= a
		}
		a *= a
		b >>= 1
	}
	return res
}
func strStr(haystack string, needle string) int {

	p := 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[p] {
			p++
			if p == len(needle) {
				return i - p + 1
			}
			continue
		}
		if p > 0 {
			i = i - p
		}
		p = 0
	}
	return -1
}

func LuckyNumbers(A int) int {
	s := make([]int, A+1)
	for i := 2; i <= A; i++ {
		if isPrime(i) {
			for j := i; j <= A; j = j + i {
				s[j]++
			}
		}
	}
	count := 0
	for i := 0; i <= A; i++ {
		if s[i] == 2 {
			count++
		}
	}
	return count
	// s := generateSieve(50001)
	// count := 0
	// for i := 5; i <= A; i++ {
	// 	c := 0
	// 	for j := i; j >= 2; j-- {
	// 		if s[j] && i%j == 0 {
	// 			c++
	// 		}
	// 	}
	// 	if c == 2 {
	// 		count++
	// 	}
	// }
	// return count
}
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	// Check for divisibility from 2 to the square root of n
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
func primesum(A int) []int {
	s := generateSieve(A)
	i, l := 2, A-2
	for {
		if s[l] && s[i] {
			if i+l == A {
				return []int{i, l}
			}
			if i+l > A {
				l--
			} else {
				i++
			}
		}
		if !s[i] {
			i++
		}
		if !s[l] {
			l--
		}
	}
	return []int{}
}
func generateSieve(length int) []bool {
	sieve := make([]bool, length)

	// Initialize the sieve array
	for i := 2; i < length; i++ {
		sieve[i] = true
	}

	// Iterate through numbers starting from 2
	for i := 2; i*i < length; i++ {
		// If i is prime
		if sieve[i] {
			// Mark multiples of i as non-prime
			for j := i * i; j < length; j += i {
				sieve[j] = false
			}
		}
	}
	return sieve
}
func generateSPFArray(length int) []int {
	spfa := make([]int, length)

	// Initialize the SPF array with the identity function
	for i := 0; i < length; i++ {
		spfa[i] = i
	}

	// Iterate through numbers starting from 2
	for i := 2; i*i < length; i++ {
		// If i is prime
		if spfa[i] == i {
			// Update SPF for multiples of i
			for j := i * i; j < length; j += i {
				if spfa[j] > i {
					spfa[j] = i
				}
			}
		}
	}

	return spfa
}

func longestCommonPrefix(strs []string) string {
	s := strs[0]
	for i := 0; i < len(strs); i++ {
		if len(s) > len(strs[i]) {
			s = strs[i]
		}
	}
	for i := 0; i < len(strs); i++ {
		a := strs[i]
		for j := 0; j < len(s); j++ {
			if a[j] == s[j] {
				continue
			}
			s = s[:j]
			break
		}
	}
	return s
}

func divisor_game(A int, B int, C int) int {
	ans := 0
	for i := A; i > 1; i-- {
		if i%B == 0 && i%C == 0 {
			ans++
		}
	}
	return ans
}

// gcd calculates the greatest
// common divisor of two integers
// using Euclid's algorithm.
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func gcd_array(a []int) int {
	ans := 0
	for _, v := range a {
		ans = gcd(ans, v)
	}
	return ans
}

func calculatePrefixGCD(arr []int) []int {
	n := len(arr)
	prefixGCD := make([]int, n)

	prefixGCD[0] = arr[0]
	for i := 1; i < n; i++ {
		prefixGCD[i] = gcd(prefixGCD[i-1], arr[i])
	}

	return prefixGCD
}

func calculateSuffixGCD(arr []int) []int {
	n := len(arr)
	suffixGCD := make([]int, n)

	suffixGCD[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixGCD[i] = gcd(suffixGCD[i+1], arr[i])
	}

	return suffixGCD
}

func solve(A []int) int {
	prefix := calculatePrefixGCD(A)
	suffix := calculateSuffixGCD(A)
	m := 0
	for i := 0; i < len(A); i++ {
		gd := 0
		if i == 0 {
			gd = suffix[i+1]
		} else if i+1 == len(A) {
			gd = prefix[i]
		} else {
			gd = gcd(prefix[i-1], suffix[i+1])
		}

		m = max(gd, m)
	}
	return m
}
