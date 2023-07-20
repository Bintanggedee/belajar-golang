package main

import "fmt"

func main() {

	var ujian = 80
	var absensi = 75

	var lulusUjian = ujian >= 85     //
	var lulusAbsensi = absensi >= 85 //
	fmt.Println(lulusUjian)          //
	fmt.Println(lulusAbsensi)        // step by step

	var lulus = lulusAbsensi && lulusUjian //
	fmt.Println(lulus)                     //

	fmt.Println(ujian >= 80 || absensi >= 80)
}
