package main

import "fmt"

func main() {
	print := func() {
		fmt.Println("Hello Gopher")
	}

	print()

	info := func(name string) {
		fmt.Println(name)
	}
	info("Khanbala")
}
