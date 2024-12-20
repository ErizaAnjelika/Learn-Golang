// channel adalah tipe data khusus di golang yang digunakan untuk mengirimkan data

package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	// GOMAXPROCS adalah variable yang digunakan untuk mengatur jumlah goroutine yang dijalankan secara paralel
	var message = make(chan string)
	var sayHello = func(who string) {
		var data = fmt.Sprintf("Hello %s", who)
		message <- data
	}

	go sayHello("Eko")
	go sayHello("Kurniawan")
	go sayHello("Khannedy")

	var message1 = <-message
	fmt.Println(message1)
	var message2 = <-message
	fmt.Println(message2)
	var message3 = <-message
	fmt.Println(message3)
}
