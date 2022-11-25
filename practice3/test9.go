package main

import "fmt"

func main() {
	name := "Khanbala"

	switch name {
	case "Rashidov":
		fmt.Println("Rashidov")
	case "Khanbala":
		fmt.Println("true")
	default:
		fmt.Println("default")
	}
}
