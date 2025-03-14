package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("012"))
}

var m map[rune][]string = map[rune][]string{
	'1': []string{"1"},
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
	'0': []string{"0"},
}

func letterCombinations(A string) []string {
	j := make([]string, 0)
	generateLetterCombinations(A, 0, "", &j)
	return j
}

func generateLetterCombinations(a string, i int, s string, q *[]string) {
	if len(a) == i {
		*q = append(*q, s)
		return
	}

	d := rune(a[i])
	letters := m[d]

	for _, j := range letters {
		generateLetterCombinations(a, i+1, s+j, q)
	}
}

func generateParenthesis(A int) []string {
	return generate(A, "", 0, 0, 0)
	// return []string{}
}

func generate(a int, s string, current, opening, closing int) []string {
	if len(s) == a*2 {
		// k := s
		return []string{s}
		// fmt.Println(s)
	}
	res := make([]string, 0)

	if opening < a {
		res = append(res, generate(a, s+"(", current+1, opening+1, closing)...)
	}
	if closing < opening {
		res = append(res, generate(a, s+")", current+1, opening, closing+1)...)
	}

	return res
}

func permute(A []int) [][]int {
	return permuteBackTrack(A, 0, make([]int, len(A)), make([]bool, len(A)))
}

func permuteBackTrack(A []int, i int, curr []int, visited []bool) [][]int {
	if i == len(A) {
		// Create a copy of the current permutation before returning
		perm := make([]int, len(curr))
		copy(perm, curr)
		return [][]int{perm}
	}

	var result [][]int
	for j := 0; j < len(A); j++ {
		// Only visit unvisited members
		if !visited[j] {
			curr[i] = A[j]
			visited[j] = true
			result = append(result, permuteBackTrack(A, i+1, curr, visited)...)
			// Backtracking
			visited[j] = false
		}
	}
	return result

}
func printSubSets(a []int, i int, current []int) [][]int {
	if i == len(a) {
		// fmt.Println(current)
		return [][]int{current}
	}

	// 2 options available
	// 	1. add current element
	// 	2. do not add current element

	// adding element
	// 	current = append(current, a[i])
	// 	printSubSets(a, i+1, current)
	x := printSubSets(a, i+1, append(current, a[i]))
	// removing added element to go back to original set
	// 	current = current[:len(current)-1]
	// not adding the element
	x = append(x, printSubSets(a, i+1, current)...)

	return x
}

func deduplicate(subsets [][]int) [][]int {
	seen := make(map[string]bool)
	result := make([][]int, 0, len(subsets))

	for _, subset := range subsets {
		key := fmt.Sprint(subset)
		if !seen[key] {
			seen[key] = true
			result = append(result, subset)
		}
	}

	return result
}
