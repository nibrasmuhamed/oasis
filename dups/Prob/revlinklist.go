package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// h := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	// h := &ListNode{}
	// h := &ListNode{1, &ListNode{2, nil}}
	h := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	x := middleNode(h)

	printlist(x)
}

func middleNode(head *ListNode) *ListNode {
	var length int
	temp := head
	for temp != nil {
		temp = temp.Next
		length++
	}

	middle := (length / 2) + 1
	// if length%2 == 0 {
	// 	middle = middle + 1
	// }
	x := 1
	for x != middle {
		head = head.Next
		x++
	}
	return head
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	// var next *ListNode = head.Next
	next := head
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
		// next = next.Next
	}
	// head.Next = prev
	return prev
}

func printlist(x *ListNode) {
	for x != nil {
		fmt.Printf("%d ->", x.Val)
		x = x.Next
	}

}
