package main

import "fmt"

type list struct {
	val  int
	next *list
}

func roundtwosecond() {
	x1 := &list{val: 1, next: nil}
	x2 := &list{val: 2, next: nil}
	x3 := &list{val: 3, next: nil}
	x4 := &list{val: 4, next: nil}

	x1.next = x2
	x2.next = x3
	x3.next = x4
	fmt.Println(findCycleinList(x1))
}

func findCycleinList(head *list) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

func roundtwofirst() {
	votes := [][]int{
		{0, 1, 1, 1, 1},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
	}
	fmt.Println(findSupremeLeader(votes))
}

func findSupremeLeader(votes [][]int) int {
	n := len(votes)
	candidate := 0

	for i := 1; i < n; i++ {
		if votes[candidate][i] == 1 {
			candidate = i
		}

	}
	for i := 0; i < n; i++ {
		if i != candidate {
			if votes[candidate][i] == 1 || votes[i][candidate] == 0 {
				return -1
			}
		}
	}
	return candidate
}

// Suppose you are given number of people voting in the election
// of supreme leader n and a n*n matrix. where if a[i][j]=0 it means
// i doesn’t want j to be supreme leader, if it is 1 it denotes i want
// j to be supreme leader. Find all the supreme leaders. Supreme leader
// is a person who is voted by everyone and they don’t vote anyone.
//
// -------------------------------------------------------------------
//
//  | 0 |1 |2 |3 |4
//  | -------------
// 0| 0 |1 |1 |1 |1
// 1| 1 |0 |0 |1 |0
// 2| 1 |0 |0 |1 |0
// 3| 0 |0 |0 |0 |0
// 4| 0 |1 |0 |1 |0
