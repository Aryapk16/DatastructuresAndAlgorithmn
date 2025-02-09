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

func (list *LinkedList) Append(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
		return
	}
	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}
func (list *LinkedList) Print() {
	temp := list.head
	for temp != nil {
		fmt.Println(temp.data, "")
		temp = temp.next
	}
}
func (list *LinkedList) DeleteBegin() {
	if list.head != nil {
		list.head = list.head.next
	}
}
func (list *LinkedList) DeleteEnd() {
	if list.head == nil {
		return
	}
	if list.head.next == nil {
		list.head = nil
		return
	}
	temp := list.head
	for temp.next.next != nil {
		temp = temp.next
	}
	temp.next = nil
}
func (list *LinkedList) DeleteAtPosition(pos int) {
	if list.head == nil {
		return
	}
	if pos == 0 {
		list.head = list.head.next
		return
	}
	temp := list.head
	for i := 0; i < pos-1 && temp != nil; i++ {
		temp = temp.next
	}
	if temp == nil || temp.next == nil {
		fmt.Println("Invalid position. Position exceeds the length of the list.")
		return
	}
	temp.next = temp.next.next
}
func main() {
	mylist := LinkedList{}
	mylist.Append(3)
	mylist.Append(5)
	mylist.Append(6)
	mylist.Append(10)
	mylist.Append(9)
	mylist.Print()
	fmt.Println("After delete at beginning:")
	mylist.DeleteBegin()
	mylist.Print()
	fmt.Println("After delete at end:")
	mylist.DeleteEnd()
	mylist.Print()
	fmt.Println("After delete at position:")
	mylist.DeleteAtPosition(2)
	mylist.Print()

}
