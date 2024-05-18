package main

import "fmt"

func LinearSearch(arr []int, target int) []int {
	var ar []int
	for i := range arr {
		if arr[i] == target {
			ar = append(ar, i)
		}

	}
	return ar
}
func main() {
	arr := []int{2, 3, 1, 4, 5, 1, 7, 1}
	target := 1
	result := LinearSearch(arr, target)
	if len(result) > 0 {
		fmt.Printf("Element %d Found at these position %d", target, result)
	} else {
		fmt.Printf("Element %d Not Found", target)
	}
}
