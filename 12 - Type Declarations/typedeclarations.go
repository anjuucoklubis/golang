package main

import "fmt"

func main() {
	type nomorKtp string
	var ktp nomorKtp = "12345678901234"
	fmt.Println(ktp)

	var contoh string = "2222222222"
	var contohKTP nomorKtp = nomorKtp(contoh)
	fmt.Println(contohKTP)

}
