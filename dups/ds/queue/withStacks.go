package main

import "fmt"

type Stack []int

type queue struct {
	input  *Stack
	output *Stack
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func main() {
	myQueue := queue{}
	myQueue.enqueue(10)
	myQueue.enqueue(11)
	myQueue.enqueue(12)
	myQueue.dequeue()
	myQueue.Print()
}

func (q *queue) dequeue() {
	if q.output.IsEmpty() {
		for len(*q.input) < 0 {
			*q.output = append(*q.output, (*q.input)[len(*q.input)-1])
			q.input.Pop()
		}

	} else {
		q.output.Pop()
	}
}

func (q *queue) enqueue(data int) {
	(*q).input = append(*q.(*input), data)
}

func (q queue) Print() {
	if q.output.IsEmpty() {
		fmt.Println(q.input)
	}
	fmt.Println(q.output)
}
func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	} else {
		index := len(*s) - 1 // Get the index of the top most element.
		// element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]
	}
}
