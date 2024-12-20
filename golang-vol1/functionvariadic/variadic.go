package main

import "fmt"

func main() {
	var avg = calculate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	var msg = fmt.Sprintf("average is %f", avg)
	fmt.Println(msg)
}

// function variadic adalah function yang bisa menerima banyak parameter

func calculate(values ...float64) float64 {
	var total float64 = 0
	for _, value := range values {
		total += value
	}
	return total / float64(len(values))
}