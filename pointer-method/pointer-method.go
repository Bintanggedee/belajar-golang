package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() { // tambahkan "*" = *Man
	man.Name = "Mr. " + man.Name
	//fmt.Println("Di method", man.Name)
}

func main() {
	bintang := Man{"Bintang"}
	bintang.Married()

	fmt.Println(bintang.Name)
}
