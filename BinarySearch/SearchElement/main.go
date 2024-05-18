package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
func main() {
	arr := []int{1, 3, 5, 7, 9, 10}
	target := 5
	result := BinarySearch(arr, target)
	if result != -1 {
		fmt.Printf(" element  %d found at postion %d: ", target, result)
	} else {
		fmt.Println("element not found", target)
	}
}
