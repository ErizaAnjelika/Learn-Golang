package main

import "fmt"

// Contoh function callback
type Callback func(int)

// function callback adalah function yang dipanggil oleh function lain

func performOperation(callback Callback) {
	// Lakukan operasi yang diperlukan
	result := 10 + 5

	// Panggil function callback dengan hasil operasi
	callback(result)
}

func main() {
	// Contoh penggunaan function callback
	performOperation(func(result int) {
		fmt.Println("Hasil operasi:", result)
	})
}
