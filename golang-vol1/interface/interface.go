// interface adalah tipe data khusus di golang, digunakan untuk membuat objek

package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func main() {
	var eko Person
	eko.Name = "Eko Kurniawan"

	sayHello(eko)

}
