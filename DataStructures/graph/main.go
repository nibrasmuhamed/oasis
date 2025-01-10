package main

import "fmt"

type QueueNode struct {
	data string
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
		fmt.Printf("%s - \n", temp.data)
		temp = temp.next
	}
}
func (q *Queue) IsEmpty() bool {
	if q.rear == nil && q.front == nil {
		return false
	}
	return true
}

func (q *Queue) Enqueue(data string) {
	newNode := &QueueNode{data, nil}
	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
		return
	}
	q.rear.next = newNode
	q.rear = newNode
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

type Graph map[string][]string

func (g *Graph) print() {
	for i, val := range *g {
		fmt.Println(i, " is ", val)
	}
}

func (g *Graph) addVertex(data string) {
	x := []string{}
	(*g)[data] = x
}

func (g *Graph) insert(vertex, edge string, isBidirectional bool) {
	_, ok := (*g)[vertex]
	if !ok {
		g.addVertex(vertex)
	}
	_, ok = (*g)[edge]
	if !ok {
		g.addVertex(edge)
	}
	(*g)[vertex] = append((*g)[vertex], edge)
	if isBidirectional {
		(*g)[edge] = append((*g)[edge], vertex)
	}
}

func main() {
	g := Graph{}
	g.insert("A", "B", true)
	g.insert("A", "F", true)
	g.insert("B", "F", true)
	g.insert("B", "C", true)
	g.insert("C", "E", true)
	g.insert("C", "D", true)
	g.insert("D", "E", true)
	g.insert("E", "F", true)
	g.print()
	fmt.Println(g.BFS("A"))
}

func (g *Graph) BFS(start string) []string {
	output := []string{}
	visited := map[string]bool{}
	q := Queue{}
	q.Enqueue(start)
	visited[start] = true
	for q.IsEmpty() {
		x := q.Dequeue()
		output = append(output, x.data)
		for _, val := range (*g)[x.data] {
			_, ok := visited[val]
			if !ok {
				q.Enqueue(val)
				visited[val] = true
			}
		}
	}
	return output
}
