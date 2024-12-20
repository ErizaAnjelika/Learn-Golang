// waitgroup adalah function yang dipanggil oleh function main
// concurrenct adalah kemampuan untuk menjalankan beberapa goroutine secara paralel

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func doPrint(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Println(message)
}

func main() {
	runtime.GOMAXPROCS(2)
	// GOMAXPROCS adalah variable yang digunakan untuk mengatur jumlah goroutine yang dijalankan secara paralel
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("Hello %d", i)
		wg.Add(1)
		go doPrint(&wg, data)
	}
	wg.Wait()
}
