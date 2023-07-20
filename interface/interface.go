package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	var bintang Person
	bintang.Name = "Bintang"
	SayHello(bintang)

	cat := Animal{
		Name: "Miaw",
	}
	SayHello(cat)
}
