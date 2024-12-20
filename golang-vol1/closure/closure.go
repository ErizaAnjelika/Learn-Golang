package main

import "fmt"

func main() {
	// function closure adalah function yang memiliki akses ke variabel di luar function
	
	var calculate = func(value int) int {
		return value * 2
	}
	var result = calculate(10)

	fmt.Println(result)

}
