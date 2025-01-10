package main

import "fmt"

type Node struct {
	value    int
	priority int
}

//

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
	q.bubbleup()
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

func main() {

}

func (q *Queue) bubbleup() {
	n := q.rear.data
	for {

	}
}
