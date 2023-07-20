package main

import "fmt"

func main() {
	name := "Bintang"
	counter := 0

	increment := func() {
		name := "Arya"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
