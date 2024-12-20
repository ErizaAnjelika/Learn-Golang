// buffer channel adalah proses transfer data secara default dilakukan dengan cara un-buffered

package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	messages := make(chan int, 3)

	go func() {
		for {
			i := <-messages
			fmt.Println("receive data", i)
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i
	}
}
