package main

import "fmt"

func main() {

	defer bye()

	hello()
}

func hello() {
	fmt.Println("Hello Gopher")
}

func bye() {
	fmt.Println("Good bye...")
}
