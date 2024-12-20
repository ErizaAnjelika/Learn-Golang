package main

import "fmt"

func main() {
	var person struct {
		firstName string
		lastName  string
		age       int
	}
	// struct adalah tipe data khusus di golang, digunakan untuk membuat objek
	person.firstName = "Eko"
	person.lastName = "Kurniawan"
	person.age = 20

	var buah = make([]string, 3)

	buah[0] = "red"
	buah[1] = "green"
	buah[2] = "blue"

	fmt.Println(person, buah)
}
