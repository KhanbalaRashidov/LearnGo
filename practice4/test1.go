package main

import "fmt"

func main() {

	array := [5]int{1, 2, 3, 4, 5}

	for index, value := range array {

		fmt.Println(index, value)
	}

	fmt.Printf("%T", array)
}
