package main

import "fmt"

func main() {
	var name = "Joko"

	if name == "Bintang" {
		fmt.Println("Hello Bintang")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hi, Salken")
	}

	if length := len(name); length <= 4 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
