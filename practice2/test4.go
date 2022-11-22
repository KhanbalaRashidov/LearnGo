package main

import "fmt"

func main() {
	number1 := 25
	fmt.Printf("decimal: %d\t\t binary: %b\t\t  hex: %#x\n", number1, number1, number1)
	number2 := number1 << 1
	fmt.Printf("decimal: %d\t\t binary: %b\t\t  hex: %#x\n", number2, number2, number2)

}
