package main

import "fmt"

type listNode struct {
	val   int
	right *listNode
	down  *listNode
}

type node struct {
	val        int
	next, prev *node
}

// func flatten(root *listNode) *listNode {
// 	if root == nil {
// 		return nil
// 	}

// }
func newnode(val int) *node {
	return &node{val: val, next: nil, prev: nil}
}
func main() {
	// n := newnode(7)
	var n *node
	addNode(1, &n)
	addNode(2, &n)
	addNode(3, &n)
	addNodeAtPosition(0, 0, n)

	print(n)
}

func addNode(val int, n **node) {
	x := newnode(val)
	if *n == nil {
		*n = x
		return
	}

	y := *n
	for y.next != nil {
		y = y.next
	}

	y.next = x
	x.prev = y
}

type mc map[string]int

func (m mc) get(s string) int {
	return m[s]
}
func addNodeAtPosition(val, pos int, n *node) {
	newNode := newnode(val)
	if n == nil {
		return
	}
	if pos == 0 {
		newNode.next = n
		n.prev = newNode
		n = newNode
		return
	}

	curr := n
	for i := 1; i <= pos; i++ {
		curr = curr.next
	}
	curr.prev.next = newNode
	newNode.next = curr
	newNode.prev = curr.prev
	curr.prev = newNode
}

func print(n *node) {
	var curr = n
	for curr != nil {
		fmt.Printf("%d ->", curr.val)
		curr = curr.next
	}
	fmt.Print("nil")
}
