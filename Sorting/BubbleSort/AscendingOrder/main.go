package main

import (
	"fmt"
)

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		flag := 0
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = 1
			}
		}
		if flag == 0 {
			break
		}

	}
	return arr
}

func main() {
	arr := []int{2, 3, 6, 8, 11, 1, 15}
	fmt.Println(arr)
	BubbleSort(arr)
	fmt.Println("array after sorting", arr)
}
