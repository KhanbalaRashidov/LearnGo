package main

import "fmt"

func main() {

	defer deferPrint()
	fmt.Println("Hello Gopher")
}

func deferPrint() {
	defer func() {
		fmt.Println("Defer print")
	}()
	fmt.Println("Test...")
}
