package main

import "fmt"

func VowelCount(str string) int {
	c := 0
	for _, i := range str {
		switch i {
		case 'A', 'E', 'I', 'O', 'U':
			c++
		}

	}
	return c
}
func main() {
	str := "AJIN"
	V := VowelCount(str)
	fmt.Println("vowels count", V)
}
