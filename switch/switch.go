package main

import "fmt"

func main() {

	name := "Aii"

	switch name {
	case "Bintang":
		fmt.Println("Hello Bintang")
		fmt.Println("Hello Bintang")
	case "Joko":
		fmt.Println("Hello Joko")
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Salken")
		fmt.Println("Salken")
	}

	//	switch length := len(name); length > 5 {
	//	case true:
	//		fmt.Println("Nama terlalu panjang")
	//	case false:
	//		fmt.Println("Nama sudah benar")
	//	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	case length <= 3:
		fmt.Println("Nama terlalu pendek")
	default:
		fmt.Println("Nama sudah benar")
	}
}
