// goroutine adalah function yang dijalankan secara paralel

package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	// GOMAXPROCS adalah variable yang digunakan untuk mengatur jumlah goroutine yang dijalankan secara paralel

	go print(5, "Hello")
	go print(5, "World")
	var input string
	fmt.Scanln(&input)
}
