package main

import "fmt"

type BLacklist func(string) bool

func registerUser(name string, blacklist BLacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

//func blacklistAdmin(name string) bool {
// return name = "admin"
//}
func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("Bintang", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})

	registerUser("Bintang", func(name string) bool {
		return name == "root"
	})
}
