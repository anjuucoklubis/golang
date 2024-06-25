package main

import "fmt"

func main() {
	var myName string
	myName = "Anju Ucok Lubis"
	fmt.Println(myName)

	myName = "Anju Ucok"
	fmt.Println(myName)

	//function variables tanpa menggunakan tpye data
	var age = 18
	fmt.Println(age)

	var name = "Anju Ucok Lubis"
	fmt.Println(name)

	// kata kunci var tidak wajib , asal menggunakan " := "
	nameMy := "Anju Ucok Lubis"
	fmt.Println(nameMy)

	ageMy := 18
	fmt.Println(ageMy)

	// jika variable dengan menggunakan " := " sudah digunakan,
	// dan ketika ingin mengganti value dari variable sebelumnya ?
	// harus menggunakan " = ", tidak boleh lagi menggunakan " := "
	// cukup 1x penggunakan dalam 1 variable

	mylastname := "lubis"
	fmt.Println(mylastname)

	mylastname = "lubis update"
	fmt.Println(mylastname)

	// mutiple variables
	var (
		nameeMy = "Anju"
		ageeMy  = 18
	)
	fmt.Println(nameeMy, ageeMy)
	fmt.Println("nama saya", nameeMy, ",", "umur saya", ageeMy)

}
