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
	b.insertArray([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(b.bfs())
}

func (b *bst) insertArray(values []int) {
	j := (len(values) / 2)
	for i := 0; j < (len(values)); i++ {
		b.insert(values[j])
		j++
		if i >= len(values)/2 {
			continue
		}
		b.insert(values[i])
	}
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

type QueueNode struct {
	data *Node
	next *QueueNode
}

type Queue struct {
	front *QueueNode
	rear  *QueueNode
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
