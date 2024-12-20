package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("input is empty")
	}
	return true, nil
}

func main() {
	var name string
	fmt.Println("input name:")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo:", name)
	} else {
		fmt.Println(err.Error())
	}
	// var input string
	// fmt.Println("input:")
	// fmt.Scanln(&input)
	// var number (int)
	// var err error

	// number, err = strconv.Atoi(input)
	// if err == nil {
	// 	fmt.Println(number, "is number")
	// } else {
	// 	fmt.Println(input, "is not number")
	// 	fmt.Println(err.Error())
	// }

}
