package main

import "fmt"

type Customer struct {
	Name, Address string // deklarasi struct
	Age           int
}

// struct function/struct method
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuu from", a.Name)
}

func main() {
	var bintang Customer
	bintang.Name = "Bintang" // jika deklarasi struct berbeda tidak akan error
	bintang.Address = "Indonesia"
	bintang.Age = 17

	bintang.sayHi("Bayu")
	bintang.sayHuuu()

	/**	fmt.Println(bintang.Name)
	fmt.Println(bintang.Address)
	fmt.Println(bintang.Age)

	hoshi := Customer{
		Name:    "hoshi", // jika deklarasi struct berbeda tidak akan error
		Address: "Jepang",
		Age:     17,
	}

	fmt.Println(hoshi)

	bayu := Customer{"Bayu", "Jakarta", 18} // rentan error jika mengubah deklarasi struct nya
	fmt.Println(bayu)*/
}
