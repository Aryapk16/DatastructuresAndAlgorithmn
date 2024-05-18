package main

import "fmt"

func InsertionSort(arr []int, n int) []int {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	return arr
}

func main() {
	arr := []int{1, 4, 2, 5, 6, 2, 5, 7}
	n := len(arr)
	fmt.Println(arr)
	result := InsertionSort(arr, n)
	fmt.Println("after sorting", result)
}
