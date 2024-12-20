package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Eko"
	names[1] = "Kurniawan"
	names[2] = "Khannedy"

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// _ adalah variabel kosong (index array tidak di gunakan)
	for _, name := range names {
		fmt.Println(name)
	}

	var buah = []string{"apel", "jeruk", "mangga"}

	buah = append(buah, "pisang")
	var buah2 = buah
	buah2 = append(buah2, "jambu")
	// append menambah elemen pada array

	fmt.Println("jumlah buah", len(buah), buah2)

	// map adalah array asosiatif
	var ayam = map[string]int{"ayam": 10, "kambing": 5}
	fmt.Println(ayam["ayam"])

}
