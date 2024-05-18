package main

import (
	"fmt"
)

// Node represents a node in the doubly linked list
type Node struct {
	data int
	next *Node
	prev *Node
}

// DoublyLinkedList represents the doubly linked list
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// AddNode adds a new node with the given data to the end of the doubly linked list
func (dll *DoublyLinkedList) AddNode(data int) {
	newNode := &Node{data: data, next: nil, prev: nil}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	}
}

// DisplayList prints the elements of the doubly linked list
func (dll *DoublyLinkedList) DisplayList() {
	current := dll.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	dll := DoublyLinkedList{}

	// Adding nodes to the doubly linked list
	dll.AddNode(1)
	dll.AddNode(2)
	dll.AddNode(3)
	dll.AddNode(4)
	dll.AddNode(5)

	// Displaying the doubly linked list
	fmt.Println("Doubly Linked List:")
	dll.DisplayList()
}
