package main

import "fmt"

// type Node struct {
// 	data int
// 	next *Node
// }
// type LinkedList struct {
// 	head *Node
// }

// func (list *LinkedList) ArrayConvert(arr []int) {
// 	for _, i := range arr {
// 		newNode := &Node{data: i, next: nil}
// 		if list.head == nil {
// 			list.head = newNode
// 		} else {
// 			temp := list.head
// 			for temp.next != nil {
// 				temp = temp.next
// 			}
// 			temp.next = newNode
// 		}
// 	}
// }
// func (list *LinkedList) Print() {
// 	temp := list.head
// 	for temp != nil {
// 		fmt.Printf("%d ", temp.data)
// 		temp = temp.next
// 	}

// }

//	func main() {
//		list := LinkedList{}
//		arr := []int{21, 22, 23, 24, 25, 27}
//		fmt.Println("Array:", arr)
//		list.ArrayConvert(arr)
//		fmt.Print("LinkedList:")
//		list.Print()
//	}
type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

func (list *LinkedList) ArrayLinked(arr []int) {
	for _, i := range arr {
		newNode := &Node{data: i, next: nil}
		if list.head == nil {
			list.head = newNode
		} else {
			temp := list.head
			for temp.next != nil {
				temp = temp.next
			}
			temp.next = newNode
		}
	}
}
func (list *LinkedList) Print() {
	temp := list.head
	for temp != nil {
		fmt.Println(temp.data, "")
		temp = temp.next
	}
	fmt.Println()
}
func main() {
	list := LinkedList{}
	arr := []int{1, 2, 3, 4, 4, 5, 6, 7}
	fmt.Println("array", arr)
	fmt.Println("array", arr)
	list.ArrayLinked(arr)
	fmt.Println("linkedlist")
	list.Print()

}
