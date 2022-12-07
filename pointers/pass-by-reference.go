package main

import "fmt"

func main() {
	num := 3
	printNum(num)

	fmt.Println(num)

	printNumPointer(&num)

	fmt.Println(num)

}

func printNum(number int) {
	fmt.Println(number)
	number = 25
	fmt.Println(number)
}

func printNumPointer(number *int) {
	fmt.Println("before:", *number)
	*number = 25
	fmt.Println("after", *number)
}
