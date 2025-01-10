package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	mergeTwoLists(list1, list2)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	outlist := new(ListNode)
	list := &ListNode{}
	list.Next = outlist
	for list1 != nil || list2 != nil {
		if list1 == nil && list2 != nil {
			outlist.Next = list2
			outlist = outlist.Next
			list2 = list2.Next
			continue
		}
		if list2 == nil && list1 != nil {
			outlist.Next = list1
			outlist = outlist.Next
			list1 = list1.Next
			continue
		}
		if list1.Val >= list2.Val {
			fmt.Println(list2.Val)
			outlist.Next = list2
			outlist = outlist.Next
			list2 = list2.Next
		} else {
			fmt.Println(list1.Val)
			outlist.Next = list1
			outlist = outlist.Next
			list1 = list1.Next
		}
	}
	return list.Next.Next
}
