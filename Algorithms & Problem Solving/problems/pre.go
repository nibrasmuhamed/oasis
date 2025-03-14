package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func main() {
	x := &Node{1, nil}
	fmt.Println(preorder(x))
}
func preorder(root *Node) []int {
	a := &[]int{}
	preOrderHelper(root, a)
	return *a

}

func preOrderHelper(node *Node, res *[]int) {
	if node != nil {

		*res = append(*res, node.Val)

		for _, n := range node.Children {
			preOrderHelper(n, res)
		}
	}

}
