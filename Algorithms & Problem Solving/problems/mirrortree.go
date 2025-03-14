package main

import "fmt"

func main() {
	// q := Queue{}
	b := bst{}
	b.root = &Node{7, nil, nil}
	b.root.left = &Node{11, nil, nil}
	b.root.left.right = &Node{6, nil, nil}
	b.root.right = &Node{1, nil, nil}
	b.root.right.right = &Node{8, nil, nil}
	b.root.right.left = &Node{2, nil, nil}
	b.root.right.left.right = &Node{3, nil, nil}
	b.root.right.left.right.left = &Node{5, nil, nil}
	// b.Mirror()
	b.MirrorRecursive(b.root)
	fmt.Println(b.bfs())
}

func (b *bst) MirrorRecursive(n *Node) {
	if n == nil {
		return
	}
	temp := n.left
	n.left = n.right
	n.right = temp

	b.MirrorRecursive(n.left)
	b.MirrorRecursive(n.right)
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

func (b *bst) Mirror() {
	if b.root == nil {
		return
	}
	q := Queue{}
	q.Enqueue(b.root)
	for q.Lenth() > 0 {
		x := q.Dequeue()
		temp := x.data.right
		x.data.right = x.data.left
		x.data.left = temp
		if x.data.left != nil {
			q.Enqueue(x.data.left)
		}
		if x.data.right != nil {
			q.Enqueue(x.data.right)
		}
	}
}

func (b *bst) rightView() []int {
	output := []int{}
	if b.root == nil {
		return output
	}
	q := Queue{}
	q.Enqueue(b.root)
	var count = 0
	for q.IsEmpty() {
		count = 0
		current := []int{}
		leng := q.Lenth()
		for count < leng {
			x := q.Dequeue()
			current = append(current, x.data.data)
			if x.data.left != nil {
				q.Enqueue(x.data.left)
			}
			if x.data.right != nil {
				q.Enqueue(x.data.right)
			}
			count++
			if count == leng {
				output = append(output, x.data.data)
			}
		}
		// output = append(output, current[len(current)-1])
	}
	return output
}

func (b *bst) levelSearch() [][]int {
	output := [][]int{}
	if b.root == nil {
		return output
	}
	q := Queue{}
	q.Enqueue(b.root)
	var count = 0
	for q.IsEmpty() {
		count = 0
		current := []int{}
		leng := q.Lenth()
		for count < leng {
			x := q.Dequeue()
			current = append(current, x.data.data)
			if x.data.left != nil {
				q.Enqueue(x.data.left)
			}
			if x.data.right != nil {
				q.Enqueue(x.data.right)
			}
			count++
		}
		output = append(output, current)
	}
	return output
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

type Node struct {
	data  int
	left  *Node
	right *Node
}

type bst struct {
	root *Node
}

func (Q *Queue) Lenth() int {
	if Q.front == nil || Q.rear == nil {
		return 0
	}
	var l int
	temp := Q.front
	for temp != Q.rear {
		l++
		temp = temp.next
	}
	l++
	return l
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
