package main

import "fmt"

func getFullName() (string, string, string) {
	return "Bintang", "Gede", "Hartono"
}

func main() {
	firstName, middleName, _ := getFullName() // "_" untuk menghiraukan
	fmt.Println(firstName)
	fmt.Println(middleName)
	//fmt.Println(lastName)
}
