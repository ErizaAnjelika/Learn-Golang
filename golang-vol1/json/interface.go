package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `[{"name":"Eko","age":20},{"name":"Kurniawan","age":20}]`
	var jsonData = []byte(jsonString)

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	// unmarshal adalah function yang digunakan untuk mengubah json menjadi map
	// interface adalah tipe data khusus di golang

	fmt.Println(data1["name"], data1["age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)
	var decodedData = data2.(map[string]interface{})

	fmt.Println(decodedData["name"], decodedData["age"])
}
