package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello"
	} else {
		return "Hello " + name
	}

}

func main() {
	result := getHello("Bintang")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
