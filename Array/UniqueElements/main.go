package main

import "fmt"

func Unique(arr []int) []int {
	unique := make(map[int]bool)
	r := []int{}
	for _, num := range arr {
		if !unique[num] {
			unique[num] = true
			r = append(r, num)
		}
	}
	return r
}
func main() {
	arr := []int{2, 4, 5, 5, 6, 7, 8, 8, 9}
	result := Unique(arr)
	fmt.Println(result)
}
