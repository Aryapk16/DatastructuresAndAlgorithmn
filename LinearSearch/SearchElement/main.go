package main

import "fmt"

func LinearSearch(arr []int, target int) int {
	for i := range arr {
		if arr[i] == target {
			return i
		}
	}
	return -1
}
func main() {
	arr := []int{5, 6, 7, 8, 9, 10, 4}
	target := 6
	result := LinearSearch(arr, target)
	if result != -1 {
		fmt.Printf("searched element %d found at postion %d\n : ", target, result)
	}
}
