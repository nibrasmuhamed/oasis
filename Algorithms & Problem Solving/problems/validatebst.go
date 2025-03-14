package main

import (
	"fmt"
	"math"
)

type bst struct {
	root *Node
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

type QueueNode struct {
	data *Node
	next *QueueNode
}

type Queue struct {
	front *QueueNode
	rear  *QueueNode
}

func findMin(current *Node) int {
	for current.left != nil {
		current = current.left
	}
	return current.data
}

func (q Queue) Print() {
	if q.rear == nil {
		fmt.Println("queue is empty")
		return
	}
	temp := q.front
	for temp != nil {
		fmt.Printf("%d - \n", temp.data.data)
		temp = temp.next
	}
}
func (q *Queue) IsEmpty() bool {
	if q.rear == nil && q.front == nil {
		return false
	}
	return true
}

func (q *Queue) Enqueue(data *Node) {
	newNode := &QueueNode{data, nil}
	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
		return
	}
	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue) Dequeue() *QueueNode {
	if q.front == nil {
		fmt.Println("queue empty")
		return nil
	} else if q.front == q.rear {
		x := q.front
		q.front = nil
		q.rear = nil
		return x
	}
	x := q.front
	q.front = q.front.next
	return x
}

func (b *bst) bfs() []int {
	if b.root == nil {
		return []int{}
	}

	// temp := b.root
	x := []int{}
	q := Queue{}
	q.Enqueue(b.root)
	for q.IsEmpty() {
		a := q.Dequeue()
		x = append(x, a.data.data)
		if a.data.left != nil {
			q.Enqueue(a.data.left)
		}
		if a.data.right != nil {
			q.Enqueue(a.data.right)
		}
	}
	return x
}

func (b *bst) removeNode(data int) {
	b.removeHelper(data, nil, b.root)
}

func (b *bst) removeHelper(data int, parent, current *Node) {
	if current == nil {
		return
	}
	for current != nil {
		if current.data < data {
			parent = current
			current = current.right
		} else if current.data > data {
			parent = current
			current = current.left
		} else {
			// found the node which has to be deleted.
			if current.left != nil && current.right != nil {
				current.data = findMin(current.right)
				b.removeHelper(current.data, current, current.right)
			} else if parent == nil {
				if current.left != nil {
					b.root = current.left
				} else if current.right != nil {
					b.root = current.right
				} else {
					b.root = nil
				}
			} else if current == parent.left {
				if current.left == nil {
					parent.left = current.right
				} else {
					parent.left = current.left
				}
			} else if current == parent.right {
				if current.left == nil {
					parent.right = current.right
				} else {
					parent.right = current.left
				}
			}
			break
		}
	}
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
	b := new(bst)
	b.insert(6)
	b.insert(9)
	b.insert(3)
	b.insert(8)
	b.insert(10)
	b.insert(2)
	b.insert(4)
	// b.root.right.right.right = &Node{1, nil, nil}
	fmt.Println(b.bfs())
	fmt.Println(validatebst(b.root))
}

func validatebst(n *Node) bool {
	return validateHelper(n, math.MinInt, math.MaxInt)
}

func validateHelper(n *Node, min, max int) bool {
	if n == nil {
		return true
	}
	// math.Max()
	if n.data < min || n.data > max {
		return false
	}
	isLeftBst := validateHelper(n.left, min, n.data)
	isRightBst := validateHelper(n.right, n.data, max)
	return isLeftBst && isRightBst
}

func (b *bst) ValidateBst() bool {
	if b.root == nil {
		return false
	}
	var valid bool
	q := new(Queue)
	q.Enqueue(b.root)
	for q.IsEmpty() {
		x := q.Dequeue()
		if x.data.left != nil {
			if x.data.data > x.data.left.data {
				valid = true
			} else {
				return false
			}
			q.Enqueue(x.data.left)
		}
		if x.data.right != nil {
			if x.data.data < x.data.right.data {
				valid = true
			} else {
				return false
			}
			q.Enqueue(x.data.right)
		}
	}
	return valid
}
