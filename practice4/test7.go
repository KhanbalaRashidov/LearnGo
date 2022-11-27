package main

import "fmt"

func main() {

	array1 := []string{"Hello", "Gopher"}
	array2 := []string{"Salam", "Gopher"}
	fmt.Println(array1)
	fmt.Println(array2)

	arrayDimentional := [][]string{array1, array2}
	fmt.Println(arrayDimentional)

	for i, array := range arrayDimentional {
		fmt.Println("record:", i)
		for index, value := range array {
			fmt.Printf("\tindex position: %v\t value: %v\t \n", index, value)
		}
	}
}
