package main

import (
	"fmt"
	"math"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type bst struct {
	root *Node
}

func (b *bst) insert(data int) {
	node := &Node{data, nil, nil}
	if b.root == nil {
		b.root = node
		return
	}
	current := b.root
	for current != nil {
		if data < current.data {
			if current.left == nil {
				current.left = node
				return
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = node
				return
			}
			current = current.right
		}
	}
}

func main() {
	b := &bst{}

	b.insert(40)
	b.insert(23)
	b.insert(57)
	b.insert(34)
	b.insert(22)
	b.insert(67)
	b.insert(33)
	b.insert(78)
	b.insert(43)
	fmt.Println(b.DiameeterBinaryTree())
}

var max int = 0

func (b *bst) DiameeterBinaryTree() int {
	if b.root == nil {
		return 0
	}
	dfs(b.root)
	return max
}

func dfs(n *Node) int {
	if n == nil {
		return -1
	}
	lh := dfs(n.left)
	rh := dfs(n.right)
	diaMeeter := lh + rh + 2
	max = int(math.Max(float64(diaMeeter), float64(max)))
	return int(math.Max(float64(lh), float64(rh)) + 1)
}
