package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func main() {
	sayHello()
	sayHello()

	for i := 0; i < 10; i++ {
		sayHello()
	}
}
