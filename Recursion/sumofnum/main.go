package main

import "fmt"

func Sum(n int) int {
	if n == 0 {
		return 0
	}
	return n + Sum(n-1)

}
func EvenSum(n int) int {
	if n == 0 {
		return 0
	}
	if n%2 == 0 {
		return n + EvenSum(n-2) //2 because sum of the even number is also added
	}
	return EvenSum(n - 1)
}
func OddSum(n int) int {
	if n == 0 {
		return 0
	}
	if n%2 != 0 {
		return n + OddSum(n-1) // 1 because sum of the odd numbers will be even and it is not included.
	}
	return OddSum(n - 1)
}

func main() {
	n := 10
	result := Sum(n)
	fmt.Println("result of natural numbers is :", result)
	EvenResult := EvenSum(n)
	fmt.Println("sum of even numbers:", EvenResult)
	OddResult := OddSum(n)
	fmt.Println("sum of odd numbers:", OddResult)

}
