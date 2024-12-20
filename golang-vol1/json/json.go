package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"name"`
	Age      int
}

func main() {
	// var jsonString = `[{"name":"Eko","age":20},{"name":"Kurniawan","age":20}]`
	var object = []User{{"Eko", 20}, {"Kurniawan", 20}}

	var jsonData, err = json.Marshal(object)
	// marshal adalah function yang digunakan untuk mengubah map menjadi json
	
	// var jsonData = []byte(jsonString)
	// var data []User
	// var err = json.Unmarshal([]byte(jsonString), &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)

	// fmt.Println(data[0].FullName, data[0].Age)
	// fmt.Println(data[1].FullName, data[1].Age)
}
