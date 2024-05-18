package main

import (
	"fmt"
)


func isAnagram(s, t string) bool {

	charCount := make(map[rune]int)

	
	for _, char := range s {
		charCount[char]++
	}

	
	for _, char := range t {
		charCount[char]--
	}

	for _, count := range charCount {
		if count != 0 {
			return false
		}
	}

	return true
}

func main() {
	s1 := "listen"
	s2 := "silent"
	fmt.Printf("%s and %s are anagrams: %t\n", s1, s2, isAnagram(s1, s2))

	s3 := "hello"
	s4 := "world"
	fmt.Printf("%s and %s are anagrams: %t\n", s3, s4, isAnagram(s3, s4))
}
