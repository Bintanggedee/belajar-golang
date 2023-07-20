package main

import (
	"fmt"
	"latihan/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
