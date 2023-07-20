package main

import "fmt"

func random() interface{} {
	return "Bintang" //interface kosong
}

func main() {
	var result interface{} = random()
	//var resultString string = result.(string) // jika interface kosong tidak sesuai/sama maka jadi panic
	//fmt.Println(resultString)					// lebih disarankan menggunakan switch statement

	// switch statement
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}
