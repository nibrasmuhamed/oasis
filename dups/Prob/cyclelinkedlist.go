package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// h := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	// h := &ListNode{}
	// h := &ListNode{1, &ListNode{2, nil}}
	var q *ListNode
	h := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, q}}}}}}
	q = &ListNode{7, h}
	fmt.Println(detectCycle(h).Val)
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	h := head
	t := head
	for h != nil && h.Next != nil {
		t = t.Next
		h = h.Next.Next
		if h == t {
			break
		}
	}
	if h != t {
		return &ListNode{-1, nil}
	}
	p := head
	for p != t {
		p = p.Next
		t = t.Next
	}
	return t
}
