package main

import "fmt"

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
	l.addToList(1)
	l.addToList(1)
	l.addToList(2)
	l.addToList(3)
	l.addToList(4)
	l.addToList(4)
	l.addToList(4)
	l.addToList(5)
	l.addToList(5)
	l.addToList(5)
	l.addToList(5)
	l.addToList(5)
	l.addToList(5)

	removeDup(l)
	reverselist(l)
	l.Print()
}

func reverselist(l *List) {
	prev, curr, next := &Node{}, &Node{}, &Node{}
	prev = nil
	curr = l.head
	next = curr.next
	for next != nil {
		curr.next = prev
		prev = curr
		curr = next
		next = next.next
	}
	l.head = prev
}

func removeDuplicate(l *List) {
	temp := l.head
	for temp.next != nil {
		prev := temp
		if temp.data == temp.next.data {
			temp.next = temp.next.next
			temp = prev
			continue
		}
		temp = temp.next
	}
}

func removeDup(l *List) {
	temp := l.head
	for temp != nil {
		temp2 := temp.next
		for temp2 != nil && temp.data == temp2.data {
			temp2 = temp2.next
		}
		temp.next = temp2
		temp = temp.next
	}
}
