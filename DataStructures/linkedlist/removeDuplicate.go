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
	mylist := &List{}
	mylist.addToList(2)
	mylist.addToList(34)
	mylist.addToList(55)
	mylist.addToList(76)
	mylist.addToList(45)
	mylist.addToList(443)
	mylist.addToList(443)
	mylist.addToList(443)
	mylist.addToList(12)
	mylist.addToList(6)
	mylist.addToList(433)
	removeDuplicate(mylist)
	mylist.Print()
}

func removeDuplicate(l *List) {
	// program to remove duplicate from linked list
	// 1. create a temp node and assign head to it
	// 2. create a prev node and assign head to it
	// 3. traverse the list and check if temp.data == temp.next.data

	// 4. if yes then temp.next = temp.next.next
	// 5. if no then prev = temp and temp = temp.next
	// 6. if prev.data == temp.next.data then temp.next = temp.next.next
	// 7. else prev = temp and temp = temp.next
	// 8. if temp.next == nil then break
	// 9. print the list
	temp := l.head
	temp1 := l.head
	prev := temp
	for temp1.next != nil {
		for temp.next != nil {
			if temp.data == temp.next.data {
				temp.next = temp.next.next
				temp = prev
				// continue
			}
			if prev.data == temp.next.data {
				temp.next = temp.next.next
			}
			prev = temp
			temp = temp.next
		}
		temp1 = temp1.next
	}
}

func (l *List) removeDuplicate() {
	temp := l.head
	prev := temp
	for temp.next != nil {
		if temp.data == temp.next.data {

			temp.next = temp.next.next
			temp = prev
			// continue
		}
		if prev.data == temp.next.data {
			fmt.Println("hi")
		}
		prev = temp
		temp = temp.next
	}
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
