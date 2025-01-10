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

func main() {
	a := []int{10, 35, 35, 65, 32, 55}
	// b := []int{}
	l := arrayToList(a)
	Print(l)
}

func arrayToList(a []int) *List {
	myList := new(List)

	for index, value := range a {
		newNode := new(Node)
		newNode.data = value
		if index == 0 {
			myList.head = newNode
			myList.tail = nil
			continue
		}
		temp := myList.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
		myList.tail = newNode
	}
	return myList
}

func Print(l *List) {
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
