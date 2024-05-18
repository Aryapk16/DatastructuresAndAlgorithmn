package main

import "fmt"

func Fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
func main() {
	n := 5
	for i := 1; i <= n; i++ {
		result := Fibonacci(i)
		fmt.Println(result)
	}

	pos := Fibonacci(n)
	fmt.Println("position is :", pos)
}
