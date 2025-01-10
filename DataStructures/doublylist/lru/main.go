package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	key        int
	val        int
	next, prev *node
}
type ll struct {
	head, tail *node
}
type LRUCache struct {
	capacity int
	size     int
	ll       *ll
	m        map[int]*node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		ll:       &ll{},
		m:        make(map[int]*node),
		size:     0,
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.m[key]
	if ok {
		return v.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.m[key]
	capacityExceeded := false
	n := &node{val: value, key: key}
	if ok {
		this.ll.deleteNode(v)
		this.ll.addEnd(n)
		this.m[key] = n
	} else {
		if this.size == this.capacity {
			delete(this.m, this.ll.head.key)
			if this.ll.head == this.ll.tail {
				// If there's only one node, set both head and tail to nil
				this.ll.head = nil
				this.ll.tail = nil
			} else {
				this.ll.deleteNode(this.ll.head)
			}
			capacityExceeded = true
		}
		this.ll.addEnd(n)
		this.m[key] = n
		if !capacityExceeded {
			this.size++
		}
	}
}

func (l *ll) deleteNode(n *node) {
	if n == nil {
		return // If the node is nil, do nothing
	}

	// Adjust pointers for the previous node
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		// If the head is the node to be deleted
		l.head = n.next
	}

	// Adjust pointers for the next node
	if n.next != nil {
		n.next.prev = n.prev
	} else {
		// If the tail is the node to be deleted
		l.tail = n.prev
	}
}

func (l *ll) addEnd(n *node) {
	if l.tail == nil {
		// The list is empty
		l.head = n
		l.tail = n
	} else {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	}
}
func main() {
	cache := Constructor(1) // Set the capacity of the LRUCache

	for {
		fmt.Print("Enter command (i/d/g key value): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		fields := strings.Fields(input)
		if len(fields) < 1 {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		command := fields[0]

		switch command {
		case "i":
			if len(fields) != 3 {
				fmt.Println("Invalid input. Usage: i key value")
				continue
			}

			key, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Println("Invalid key. Please enter a valid integer.")
				continue
			}

			value, err := strconv.Atoi(fields[2])
			if err != nil {
				fmt.Println("Invalid value. Please enter a valid integer.")
				continue
			}

			cache.Put(key, value)
			fmt.Println("Inserted:", key, value)

		case "d":
			if len(fields) != 2 {
				fmt.Println("Invalid input. Usage: d key")
				continue
			}

			key, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Println("Invalid key. Please enter a valid integer.")
				continue
			}

			cache.ll.deleteNode(cache.m[key])
			delete(cache.m, key)
			fmt.Println("Deleted:", key)

		case "g":
			if len(fields) != 2 {
				fmt.Println("Invalid input. Usage: g key")
				continue
			}

			key, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Println("Invalid key. Please enter a valid integer.")
				continue
			}

			result := cache.Get(key)
			fmt.Println("Get:", key, "Result:", result)

		default:
			fmt.Println("Invalid command. Please try again.")
		}
	}
}
