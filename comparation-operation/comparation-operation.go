package main

import "fmt"

func main() {

	var name1 = "Bintang"
	var name2 = "Javier"

	var result bool = name1 > name2 // perbandingan string (alfabet B lebih duluan = kecil, alfabet F setelah B = lebih besar)
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2) //tidak sama dengan
}
