package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"Eko", "Kurniawan Khannedy"}
	printMassage("halo", names)
	printMassage("lain", names)
}

func printMassage(message string, names []string) {
	var nameSting = strings.Join(names, " ")
	fmt.Println(strings.ToLower(message), nameSting)
}
