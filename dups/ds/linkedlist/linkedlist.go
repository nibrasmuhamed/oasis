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

func (l List) PrintReverse(node *Node) {
	temp := node
	if temp == nil {
		return
	}
	l.PrintReverse(temp.next)

	fmt.Printf("%d ->", temp.data)
}

func main() {
	myList := List{}
	myList.addBeginning(89)
	myList.addToList(80)
	myList.addToList(9)
	myList.addToList(8)
	myList.addBeginning(90)
	myList.delete(90)
	myList.delete(8)
	myList.addBeginning(14)
	myList.addBeginning(67)
	myList.insertAfter(89, 10)
	myList.insertBefore(10, 5)
	myList.insertAfter(9, 0)
	myList.insertAfter(67, 2)
	myList.insertBefore(67, 9)
	myList.Print()
	fmt.Println()
	myList.PrintReverse(myList.head)

}

func (l *List) insertBefore(beforeTo, data int) {
	newNode := new(Node)
	newNode.data = data
	temp := l.head
	prev := temp
	if l.head.data == beforeTo {
		newNode.next = l.head
		l.head = newNode
		return
	}
	for temp.data != beforeTo {
		prev = temp
		temp = temp.next
	}
	prev.next = newNode
	newNode.next = temp
}

func (l *List) insertAfter(afterTo, data int) {
	newNode := new(Node)
	newNode.data = data
	temp := l.head
	for temp.data != afterTo {
		temp = temp.next
	}
	newNode.next = temp.next
	temp.next = newNode
}

func (l *List) addBeginning(data int) {
	newNode := new(Node)
	newNode.data = data
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	newNode.next = l.head
	l.head = newNode
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

func (l *List) delete(data int) {
	if l.head.data == data {
		l.head = l.head.next
		return
	}
	temp := l.head
	prev := temp
	for temp.data != data {
		prev = temp
		temp = temp.next
	}
	if temp == l.tail {
		prev.next = nil
		return
	}
	prev.next = temp.next

}
