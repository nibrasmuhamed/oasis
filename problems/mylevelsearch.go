/**
 * Definition for a binary tree node.
 *
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val  *TreeNode
	Next *Node
}

type Queue struct {
	front *Node
	rear  *Node
}

func main() {
	a := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, &TreeNode{15, nil, nil}, &TreeNode{17, nil, nil}}}
	fmt.Println(levelOrder(a))
}

func levelOrder(root *TreeNode) [][]int {
	output := [][]int{}
	if root == nil {
		return output
	}
	q := Queue{}
	q.Enque(root)
	var count = 0
	for q.isEmpty() {
		count = 0
		current := []int{}
		leng := q.Lenth()
		for count < leng {
			x := q.Dequeue()
			current = append(current, x.Val)
			if x.Left != nil {
				q.Enque(x.Left)
			}
			if x.Right != nil {
				q.Enque(x.Right)
			}
			count++
		}
		output = append(output, current)
	}
	return output
}

func (q *Queue) lookUp() *TreeNode {
	if q.front == nil {
		return nil
	}
	return q.front.Val
}

func (q *Queue) isEmpty() bool {
	if q.rear != nil && q.front != nil {
		return false
	}
	return true
}

func (q *Queue) Enque(r *TreeNode) {
	v := &Node{r, nil}
	if q.front == nil && q.rear == nil {
		q.front = v
		q.rear = v
		return
	}
	q.rear.Next = v
	q.rear = q.rear.Next
}

func (q *Queue) Dequeue() *TreeNode {
	if q.isEmpty() {
		return nil
	}
	x := q.front
	if q.front == q.rear {
		q.front = nil
		q.rear = nil
		return x.Val
	}

	q.front = q.front.Next
	return x.Val
}

func (Q *Queue) Lenth() int {
	if Q.front == nil || Q.rear == nil {
		return 0
	}
	var l int
	temp := Q.front
	for temp != Q.rear {
		l++
		temp = temp.Next
	}
	l++
	return l
}
