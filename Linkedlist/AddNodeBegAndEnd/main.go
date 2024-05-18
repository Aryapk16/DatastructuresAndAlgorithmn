package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

func (list *LinkedList) AddNodeAtBeginning(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
		return
	}
	newNode.next = list.head
	list.head = newNode
}

func (list *LinkedList) Print() {
	temp := list.head
	for temp != nil {
		fmt.Println(temp.data, "")
		temp = temp.next
	}
	fmt.Println()
}

func (list *LinkedList) AddNodeAtEnd(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
		return
	}
	newNode.next = nil
	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

func main() {
	myList := LinkedList{}
	myList.AddNodeAtBeginning(3)
	myList.AddNodeAtBeginning(5)
	fmt.Println("Add beginning:")
	myList.Print()
	myList.AddNodeAtEnd(10)
	myList.AddNodeAtEnd(6)
	fmt.Println("At end:")
	myList.Print()
	fmt.Println()
}
