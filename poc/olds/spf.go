package main

import "sort"

func sieve() [1000005]int {
	// spf[x] = smallest prime factor of x
	var spf [1000005]int
	for i := 1; i <= 1000000; i++ {
		spf[i] = i
	}
	for i := 2; i*i <= 1000000; i++ {
		if spf[i] == i {
			for j := i * i; j <= 1000000; j += i {
				if spf[j] == j {
					spf[j] = i
				}
			}
		}
	}
	return spf
}
func solvesieve(A []int) []int {
	spf := sieve()
	// Using prime factorization to get the number of divisors for every integer
	var sol []int
	for i := 0; i < len(A); i++ {
		temp := A[i]
		ans := 1
		for temp != 1 {
			cnt := 1
			d := spf[temp]
			for temp != 1 && temp%d == 0 {
				cnt++
				temp /= d
			}
			ans *= cnt
		}
		sol = append(sol, ans)
	}
	return sol
}

type Interval struct {
	start, end int
}

func merge(A []Interval) []Interval {
	if len(A) <= 1 {
		return A
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i].start < A[j].start
	})
	result := []Interval{A[0]}
	for i := 1; i < len(A); i++ {
		if A[i].start <= result[len(result)-1].end {
			result[len(result)-1].end = maxnum(A[i].end, result[len(result)-1].end)
		} else {
			result = append(result, A[i])
		}
	}
	return result
}

func maxnum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
