package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) enqueue(data int) {
	newNode := &Node{data, nil}
	// rep := new(Node)
	// rep.data = 19
	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
		return
	}
	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue) dequeue() {
	if q.front == nil {
		fmt.Println("queue empty")
		return
	}
	q.front = q.front.next
}
