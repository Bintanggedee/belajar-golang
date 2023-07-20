package main

import "fmt"

func main() {

	var name [3]string

	name[0] = "Bintang"
	name[1] = "Gede"
	name[2] = "Hartono"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var value = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(value)
	fmt.Println(value[0])
	fmt.Println(value[1])
	fmt.Println(value[2])

	fmt.Println(len(name))
	fmt.Println(len(value))

	var lagi [10]string

	fmt.Println(len(lagi))
}
