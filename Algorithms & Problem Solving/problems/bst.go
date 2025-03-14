package main

import "fmt"

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

func main() {
	b := &bst{}
	// b.insert(10)
	// // b.insert(11)
	// b.insert(11)
	// b.insert(12)
	// b.insert(1)
	// b.insert(7)
	b.insertArray([]int{1, 2, 3, 8, 9, 0})
	fmt.Println(b.bfs())
}

func (b *bst) insertArray(a []int) {
	q := Queue{}
	i := 0
	if b.root == nil {
		if len(a) <= 0 {
			return
		}
		node := &Node{a[i], nil, nil}
		i++
		b.root = node
	}
	q.Enqueue(b.root)
	for q.IsEmpty() {
		x := q.Dequeue()
		if x.data.left == nil {
			if i <= len(a) {
				node := &Node{a[i], nil, nil}
				x.data.left = node
			}
			i++
			if i == len(a) {
				return
			}
		}
		q.Enqueue(x.data.left)
		if x.data.right == nil {
			if i <= len(a) {
				node := &Node{a[i], nil, nil}
				x.data.right = node
			}
			i++
			if i == len(a) {
				return
			}
		}
		q.Enqueue(x.data.right)
	}

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

func findMin(current *Node) int {
	for current.left != nil {
		current = current.left
	}
	return current.data
}

func (b *bst) find(data int) bool {
	current := b.root
	if current.data == data {
		return true
	}
	for current != nil {
		if current.data == data {
			return true
		}
		if data < current.data {
			current = current.left
		} else {
			current = current.right
		}

	}
	return false
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
