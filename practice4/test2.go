package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, value := range array {

		fmt.Println(index, value)
	}

	fmt.Printf("%T", array)
}
