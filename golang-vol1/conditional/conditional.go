package main

import "fmt"

func main() {
	var trueValue = true
	// conditional
	// if trueValue {
	// 	fmt.Println("true")
	// } else {
	// 	fmt.Println("false")
	// }
	switch trueValue {
	case true:
		fmt.Println("true")
	default:
		fmt.Println("false")
	}
}
