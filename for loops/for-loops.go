package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Bintang", "Gede", "Hartono"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
		//fmt.Println(value) //value diganti dengan "_"
	}

	person := make(map[string]string)
	person["Name"] = "Bintang"
	person["Title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)

	}
	fmt.Println("tambah data")

}
