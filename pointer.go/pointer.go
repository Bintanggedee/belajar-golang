package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) { // 6. Address harus menjadi pointer "*"
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // 1. pass by value (tidak nge replace tapi duplikat)
	// 2. tambahkan operator AND "&", jadi address2 akan nge reference ke address1
	address2.City = "Bandung"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} // 3. jika Address hanya memakai operator AND, address1 tidak akan berubah
	// 4. pakai operator "*" di depan address2 untuk mengubah
	fmt.Println(address1)
	fmt.Println(address2) // 5. pointer "*"

	var address3 *Address = new(Address)
	address3.City = "Jakarta"
	fmt.Println(address3)

	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	var alamatPointer *Address = &alamat    // 8. tapi jika sudah ada pointer maka &alamat diganti menjadi alamatPointer
	ChangeCountryToIndonesia(alamatPointer) // 7. tambahkan operator "&" pada alamat
	fmt.Println(alamat)
}
