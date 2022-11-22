package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
	tb = 1 << (iota * 10)
)

func main() {
	number1 := 25
	fmt.Printf("%v\t\t%b\n", number1, number1)

	number2 := number1 >> 3
	fmt.Printf("%v\t\t%b\n", number2, number2)

	number3 := number1 << 3
	fmt.Printf("%v\t\t%b\n", number3, number3)

	number4 := 3
	fmt.Printf("%v\t\t%b\n", number4, number4)

	// kb := 1024
	// fmt.Printf("%v\t\t%b\n", kb, kb)

	// mb := 1024 * kb
	// fmt.Printf("%v\t\t%b\n", mb, mb)

	// gb := 1024 * mb
	// fmt.Printf("%v\t\t%b\n", gb, gb)

	// tb := 1024 * gb
	// fmt.Printf("%v\t\t%b\n", tb, tb)

	fmt.Printf("%v\t\t%b\n", kb, kb)
	fmt.Printf("%v\t\t%b\n", mb, mb)
	fmt.Printf("%v\t\t%b\n", gb, gb)
	fmt.Printf("%v\t\t%b\n", tb, tb)
}
