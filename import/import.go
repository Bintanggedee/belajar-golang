package main

import (
	"fmt"
	"latihan/helper"
)

func main() {
	helper.SayHello("Bintang")
	//helper.sayGoodbye("Bintang") // tidak akan bisa diakses jika diawali huruf kecil
	fmt.Println(helper.Application)
	//fmt.Println(helper.version) // error juga
}
