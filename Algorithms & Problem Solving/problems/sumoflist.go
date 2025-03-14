package main

import (
	"fmt"
	"math"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) addToList(data int) {
	newNode := new(Node)
	newNode.data = data
	if l.head == nil {
		l.head = newNode
	} else if l.tail != nil {
		l.tail.next = newNode
	} else {
		temp := l.head
		for temp != nil {
			temp = temp.next
		}
		temp.next = newNode
	}
	l.tail = newNode
}
func (l List) Print() {
	if l.head == nil {
		fmt.Println("list is empty")
		return
	}
	temp := l.head
	for temp != nil {
		fmt.Printf("%d-> ", temp.data)
		temp = temp.next
	}
}

func main() {
	l := &List{}
	l.addToList(0)
	l.addToList(4)
	l.addToList(5)
	m := &List{}
	m.addToList(3)
	m.addToList(2)
	m.addToList(7)
	x := sum(m, l)
	x.Print()
}

func sum(f, s *List) *List {
	x := &List{}
	c := 0
	temp1 := f.head
	temp2 := s.head
	for temp1 != nil || temp2 != nil {
		s := temp1.data + temp2.data + c
		q := s % 10
		c = int(math.Floor(float64(s / 10)))
		x.addToList(q)
		temp1 = temp1.next
		temp2 = temp2.next
	}
	x.addToList(c)
	return x
}
