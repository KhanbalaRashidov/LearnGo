package main

import "fmt"

func main() {

	firstName, lastName := "Khanbala", "Rashidov"

	print()

	func() {
		fmt.Println("Print anonymous function")
	}()

	func(firstName, lastName string) {
		fmt.Println(firstName, " ", lastName)
	}(firstName, lastName)

}

func print() {
	fmt.Println("Hello Gopher")
}
