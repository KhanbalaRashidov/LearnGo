package main

import "fmt"

var age = 28

func main() {
	num1 := 3
	num2 := 25
	sum := num1 + num2

	name := "Khanbala Rashidov"

	fmt.Println("Sum=", sum)
	fmt.Println(name)

	num2 = 9
	num1 = 6
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(age)
	Age()

}

func Age() {
	fmt.Println("Age:", age)
}
