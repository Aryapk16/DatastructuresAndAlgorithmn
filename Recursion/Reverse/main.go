package main

import "fmt"
func Reverse(str string )string{
	if len(str)<= 1{
		return str
	}
	first:=str[0]
	last:=str[len(str)-1]
	reverse := string(last) + Reverse(str[1:len(str)-1]) + string(first)
	return reverse
}
func main(){
	string:="ajin"
	result:= Reverse(string)
	fmt.Println("string after reverse",result)
}