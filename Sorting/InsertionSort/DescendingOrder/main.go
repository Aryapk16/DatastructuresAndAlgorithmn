package main

import "fmt"

func InsertionSort(arr []int, n int) []int {
	for i := 0; i < n; i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] < temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
	return arr
}
func main() {
	arr := []int{4, 5, 2, 5, 7, 3, 24}
	n := len(arr)
	fmt.Println(arr)
	result := InsertionSort(arr, n)
	fmt.Println("After sorting", result)
}
