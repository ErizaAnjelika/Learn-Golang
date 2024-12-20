// recover adalah function untuk menghandle panic
// recover adalah function yang dipanggil oleh function panic
package main

import (
	"errors"
	"fmt"
	"strings"
)

func catch() {
	if r := recover(); r != nil {
		fmt.Println("error:", r)
	} else {
		fmt.Println("aplication running perfectly")
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("input is empty")
	}
	return true, nil
}

func main() {
	defer catch() // defer adalah function yang dipanggil oleh function main
	var name string
	fmt.Print("input name:")
	fmt.Scanln(&name)
	if valid, err := validate(name); valid {
		fmt.Println("halo:", name)
	} else {
		panic(err.Error())
		fmt.Println("end program")
	}
}
