package main

import (
	"fmt"
	"package/lib/math"
)

func main() {
	xs := []float64{1, 2, 3, 4, 5}
	avg := math.Average(xs)
	fmt.Println(avg)
}
