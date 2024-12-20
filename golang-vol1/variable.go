package main

import "fmt"

func main() {
	var name string = "Eko"
	const age int = 20
	var bilanganBulat uint8 = 10
	// uint8 = 0 - 255

	// bilangan poin
	var bilA int = 2
	var bilB *int = &bilA
	*bilB = 10
	fmt.Println(name, age, bilanganBulat, bilA, *bilB)
}
