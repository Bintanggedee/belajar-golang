package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 1 { // break menghentikan seluruh perulangan
			continue // continue menghentikan perulangan saat ini dan langsung dilanjutkan ke post statement dan dilanjutkan ke perulangan selanjutnya (skip)
		}
		fmt.Println("Perulangan ke", i)
	}
}
