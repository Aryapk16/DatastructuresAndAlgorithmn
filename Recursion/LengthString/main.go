package main

import "fmt"

func Length(Str string) int {
	if len(Str) == 0 {
		return 0
	}
	return 1 + Length(Str[1:])
}
func main() {
	Str := "ARYA"
	result := Length(Str)
	fmt.Println("length of string is", result)
}
