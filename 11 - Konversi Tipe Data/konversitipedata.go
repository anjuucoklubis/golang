package main

import "fmt"

func main() {

	// Konversi Number
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// Konversi String
	myName := "Anju Ucok Lubis"

	var a = myName[0]       // byte / uint8
	var aString = string(a) // convert to string
	fmt.Println(a)
	fmt.Println(aString)
}
