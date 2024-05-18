package main

import "fmt"

	func Recursion(arr  []int)int{
		if len(arr)==0{
			return 0
		}
		return arr[0]+Recursion(arr[1:])
	}

	func main(){
		arr:= []int{1,2,3,4,5}
		result:= Recursion(arr)
		fmt.Println("array sum is:",result)
	}

