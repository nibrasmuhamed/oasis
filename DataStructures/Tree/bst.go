package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type bst struct {
	root *Node
}

func main() {
	myBst := bst{}
	myBst.insert(10)
	myBst.insert(7)
	myBst.insert(9)
	myBst.insert(5)
	myBst.insert(3)
	myBst.insert(15)
	myBst.insert(13)
	myBst.insertIterative(20)

	myBst.removeNode(10)
	myBst.print()
	// fmt.Println(myBst.findClosest(7))
	// fmt.Println(myBst.find(15))
	// fmt.Println(myBst.findIterative(15))

	// fmt.Println(myBst.getMin(myBst.root.right))
}

// func (b *bst)bfs() []int {
// 	if b.root == nil{
// 		return []int{}
// 	}

// }

func (b *bst) findClosest(data int) int {
	var closest int
	root := b.root
	for root != nil {
		if root.data < data {
			if closest-data < closest {
				closest = root.data
			}
			root = root.right
		} else if root.data > data {
			if closest-data < closest {
				closest = root.data
			}
			closest = root.data
			root = root.left
		} else {
			return root.data
		}
	}
	return closest
}

func (b *bst) removeNode(data int) {
	b.removeHelper(data, b.root, nil)
}

func (b *bst) removeHelper(data int, current *Node, parent *Node) {
	for current != nil {
		if current.data > data {
			parent = current
			current = current.left
		} else if current.data < data {
			parent = current
			current = current.right
		} else {
			// found the element.
			if current.left != nil && current.right != nil {
				// if the child has right and left child.
				current.data = b.getMin(current.right)
				b.removeHelper(current.data, current.right, parent)
			} else if parent == nil {
				// condition where root node only have single child. in this
				// case every condition will be false as root doesn't have
				// any parent node.
				if current.left != nil {
					// if root node has only left child
					//                 current.data = current.left.data
					//                 current.left = current.left.left
					b.root = current.left
				} else if current.right != nil {
					// if root has only right child
					b.root = current.right
				} else {
					// single node tree
					current = nil
				}
			} else if current == parent.left {
				// if node has only single child.
				// and the child is in left of parent.
				// we still don't know which side of the current node is null
				if current.left != nil {
					// if the current node has left child.
					// we assign current.left to parent.left
					parent.left = current.left
				} else {
					// if the current node has right side.
					// we assign parent right as current right
					parent.left = current.right
				}
			} else if current == parent.right {
				// if node has only single child.
				// and the child is in right side of parent.
				if current.left != nil {
					// if the current node has left child.
					// we assign current.left to parent.left
					parent.right = current.left
				} else {
					// if the current node has right side.
					// we assign parent right as current right
					parent.right = current.right
				}
			}
			break
		}
	}
}

func (b *bst) getMin(root *Node) int {
	for root.left != nil {
		root = root.left
	}
	return root.data
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

func (b *bst) find(data int) bool {
	return findHelper(b.root, data)
}

func findHelper(root *Node, data int) bool {
	if root == nil {
		return false
	}
	if root.data == data {
		return true
	} else if root.data > data {
		return findHelper(root.left, data)
	}
	return findHelper(root.right, data)

}

func (b *bst) findIterative(data int) bool {
	if b.root == nil {
		return false
	}
	temp := b.root
	for temp != nil {
		if data < temp.data {
			temp = temp.left
		} else if data > temp.data {
			temp = temp.right
		} else {
			return true
		}
	}
	return false
}

func (b bst) print() {
	printHelper(b.root)
}

func printHelper(node *Node) {
	if node == nil {
		return
	}
	printHelper(node.left)
	printHelper(node.right)
	fmt.Println(node.data)

}

func (b *bst) insert(data int) {
	newNode := &Node{data, nil, nil}
	if b.root == nil {
		b.root = newNode
		return
	}
	if newNode.data > b.root.data {
		b.insertHelper(newNode, b.root)
	} else {
		b.insertHelper(newNode, b.root)
	}
}

func (b *bst) insertHelper(node *Node, root *Node) {

	if node.data > root.data {
		if root.right == nil {
			root.right = node
			return
		}
		b.insertHelper(node, root.right)
	} else {
		if root.left == nil {
			root.left = node
			return
		}
		b.insertHelper(node, root.left)
	}
}
