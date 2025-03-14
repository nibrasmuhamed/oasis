package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) > 1 {
		mergedlists := []*ListNode{}
		for i := 0; i < len(lists); i = i + 2 {
			l1 := lists[i]
			l2 := new(ListNode)
			if i+1 < len(lists) {
				l2 = lists[i+1]
			}
			mergedlists = append(mergedlists, mergeTwoLists(l1, l2))
		}
		lists = mergedlists
	}
	return lists[0]
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummyNode := new(ListNode)
	current := dummyNode

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}
	return dummyNode.Next
}

func roundone() {
	lists := make([]*ListNode, 3)
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	l3 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}
	lists[0] = l1
	lists[1] = l2
	lists[2] = l3
	printNodes(mergeKLists(lists))
}

func printNodes(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, "->")
		l = l.Next
	}
}

// n*k
// O(1)
