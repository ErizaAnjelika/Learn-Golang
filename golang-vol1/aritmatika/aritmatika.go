package main

import "fmt"

func main() {
	// aritmatika , relational, logical
	var a = 2
	var b = 6

	var c = (((a+b)%3)*4 - 2) / 3
	var isEqual = c == 1

	var value = true && true
	fmt.Println(c, isEqual, value)

}
