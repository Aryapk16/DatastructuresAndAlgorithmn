package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}
type Linkedlist struct {
	head *Node
}

func main() {
	list := Linkedlist{}
	list.Appendlist(1)
	list.Appendlist(3)
	list.Appendlist(5)
	list.Displaylist()
	list.Traverselist()

}
func (list *Linkedlist) Appendlist(data int) {
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

func (list *Linkedlist) Displaylist() {
	temp := list.head
	for temp != nil {
		fmt.Printf("%d ", temp.data)
		temp = temp.next
	}
}
func (list *Linkedlist) Traverselist() {
	temp := list.head
	for temp != nil {

		fmt.Println("temp data is", temp.data)
		temp = temp.next

	}
}
