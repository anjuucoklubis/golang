package main

import "fmt"

func main() {
	const name string = "Budi" //menggunakan type data
	const age = 20             // tidak menggunakan type data

	// error , karena const tidak bisa diubah lagi
	// 	name = "Budi"
	//	age= 21

	fmt.Println(name)
	fmt.Println(age)

	// mutiple function const
	const (
		myname string = "Budi suhartono"
		myage         = 98
	)

	fmt.Println(myname)
	fmt.Println(myage)

}
