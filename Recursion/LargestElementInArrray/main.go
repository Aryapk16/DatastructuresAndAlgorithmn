package main

import "fmt"

func Largest(arr []int, largest, size int) int {
	if size == 0 {
		return largest
	} else {
		if largest < arr[size] {
			largest = arr[size]
		           }
	      }
	return Largest(arr, largest, size-1)
}
func main() {
	arr := []int{3, 5, 7, 1, 2, 6}
	largest := arr[0]
	size := len(arr) - 1
	result := Largest(arr, largest, size)
	fmt.Println("the largest element is:", result)
	fmt.Println(arr)
}
