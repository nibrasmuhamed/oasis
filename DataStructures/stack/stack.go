package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) pop() *int {
	if s.top == nil {
		return nil
	}
	x := s.top
	s.top = s.top.next
	return &x.data
}

func (s *Stack) push(data int) {
	node := new(Node)
	node.data = data
	if s.top == nil {
		s.top = node
		return
	}
	node.next = s.top
	s.top = node
}

func (s Stack) Print() {
	if s.top == nil {
		fmt.Println("stack empty")
		return
	}
	temp := s.top
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
	fmt.Println()
}

func main() {
	myStack := Stack{}
	myStack.push(2)
	myStack.push(5)
	myStack.push(8)
	myStack.push(9)
	myStack.pop()
	myStack.pop()
	myStack.Print()
}
