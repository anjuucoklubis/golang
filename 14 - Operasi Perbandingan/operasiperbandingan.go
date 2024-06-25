package main

import "fmt"

func main() {
	var name1 = "anju"
	var name2 = "lubis"
	var name3 = "anju"
	var result bool = name1 == name2
	var result2 bool = name1 == name3
	var result3 bool = name1 != name2
	var result4 bool = name1 != name3
	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	var number1 = 10
	var number2 = 20
	var number3 = 20
	var result5 bool = number1 < number2
	var result6 bool = number1 > number2
	var result7 bool = number2 <= number3
	var result8 bool = number2 >= number3

	fmt.Println(result5)
	fmt.Println(result6)
	fmt.Println(result7)
	fmt.Println(result8)

	var a = 'a'
	var b = byte(a)
	fmt.Println(b)
}
