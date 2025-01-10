package main

import "fmt"

type Node struct {
	prev *Node
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
		newNode.prev = l.tail
	} else {
		temp := l.head
		for temp != nil {
			temp = temp.next
		}
		newNode.prev = temp
		temp.next = newNode
	}
	l.tail = newNode
}

func (l List) PrintRev() {
	if l.tail == nil {
		fmt.Println("list is empty")
		return
	}
	temp := l.tail
	for temp != nil {
		fmt.Printf("%d-> ", temp.data)
		temp = temp.prev
	}
}

func main() {
	myList := List{}
	myList.addToList(80)
	myList.addToList(9)
	myList.addToList(8)
	myList.addToList(3)
	myList.addToList(4)
	myList.addToList(7)

	myList.Print()
	fmt.Println()
	myList.PrintRev()
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
