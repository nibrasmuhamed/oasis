package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type bst struct {
	root *Node
}

func (b *bst) insertIterative(data int) {
	newNode := &Node{data, nil, nil}
	if b.root == nil {
		b.root = newNode
		return
	}
	temp := b.root
	for true {
		if newNode.data > temp.data {
			if temp.right == nil {
				temp.right = newNode
				return
			}
			temp = temp.right
		} else {
			if temp.left == nil {
				temp.left = newNode
				return
			}
			temp = temp.left
		}
	}
}
func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(gretest(a))
}

func gretest(a []int) int {
	mybst := bst{}
	for _, val := range a {
		mybst.insertIterative(val)
	}
	first := mybst.findHelper(mybst.root, 1000000)
	second := mybst.findHelper(mybst.root, first)
	third := mybst.findHelper(mybst.root, second)
	fourth := mybst.findHelper(mybst.root, third)
	return first + second + third + fourth

}

func (b *bst) findHelper(root *Node, data int) int {
	if root == nil {
		return 0
	}
	var max int
	val := root.data
	if val > max && val < data {
		max = root.data
	}

	b.findHelper(root.left, data)
	b.findHelper(root.right, data)
	fmt.Println(root.data)
	return max
}
