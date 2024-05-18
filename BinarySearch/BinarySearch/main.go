package main

import "fmt"

func SquareRoot(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left, right := 1, x
	result := 0
	for left <= right {
		mid := (left + right) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid <= x {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
func main() {
	fmt.Println("square root of 49 is:", SquareRoot(49))
	fmt.Println("square root of 25 is:", SquareRoot(25))
}
