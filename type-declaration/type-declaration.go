package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpBintang NoKTP = "23842374234209"
	var marriedStatus Married = false
	fmt.Println(noKtpBintang)
	fmt.Println(marriedStatus)
}
