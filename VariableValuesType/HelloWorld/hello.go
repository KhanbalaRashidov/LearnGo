package main

import "fmt"

func main() {
	fmt.Println("Hello Gopher...")
	foo()
	fmt.Println("Somthing operation...")

	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			fmt.Println("i cutdur:", i)
		} else {
			fmt.Println("i tekdir", i)
		}
	}
}
func foo() {
	fmt.Println("I`m Khanbala Rashidov!")
}
