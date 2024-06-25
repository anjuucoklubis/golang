package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var d = 5
	var e = 2
	c := a + b - d*e
	fmt.Println(c)

	// augmentation assignment
	var i = 10
	i += 10
	fmt.Println(i)
	i -= 5
	fmt.Println(i)
	i *= 2
	fmt.Println(i)

	// unary operation
	var j = 1
	j++ // naikkan 1, jadi 2
	fmt.Println(j)
	j++ // naikkan 1 lagi , jadi 3
	fmt.Println(j)
	j-- // turunkan 1, jadi 2
	fmt.Println(j)

}
