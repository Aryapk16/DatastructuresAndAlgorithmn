package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Score int
}

func SortByName(students []Person) {
	n := len(students)
	for i := 0; i < n-1; i++ {
		var swap bool
		for j := 0; j < n-i-1; j++ {
			swap = false
			if students[j].Name > students[j+1].Name {
				students[j], students[j+1] = students[j+1], students[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}
func main() {
	students := []Person{
		{"Arya", 22, 16},
		{"Ajin", 21, 20},
		{"Arun", 20, 21},
	}
	fmt.Println("before:", students)
	SortByName(students)
	fmt.Println("After", students)
}
