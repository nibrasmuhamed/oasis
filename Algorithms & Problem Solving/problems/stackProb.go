package main

import (
	"fmt"
	"strconv"
)

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
	a := []string{"4", "13", "5", "/", "+"}
	fmt.Println(*usingStack(a))
}

func usingStack(a []string) *int {
	var s Stack
	for _, val := range a {
		i, err := strconv.Atoi(val)
		if err != nil {
			number2 := s.pop()
			number1 := s.pop()
			x := findOperator(*number1, *number2, val)
			s.push(x)
			continue
		}
		s.push(i)
	}
	return s.pop()
}

func findOperator(num1, num2 int, s string) int {
	switch s {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	}
	return 0
}
