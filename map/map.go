package main

import "fmt"

func main() {

	person := map[string]string{
		"Name":    "Bintang",
		"Address": "Ciracas",
	}

	person["Title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["Name"])
	fmt.Println(person["Address"])

	var book map[string]string = make(map[string]string)
	book["Title"] = "Belajar Go-Lang"
	book["Author"] = "Eko"
	book["Ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "Ups")

	fmt.Println(book)
	fmt.Println(len(book))

}
