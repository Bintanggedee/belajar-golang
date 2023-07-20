package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	//person := NewMap("Bintang")
	//fmt.Println(person)

	var person map[string]string = NewMap("Bintang") // jika tidak diisi, jadi Data Kosong

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}
}
