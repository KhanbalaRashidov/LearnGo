package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(array[:5])
	fmt.Println(array[5:])
	fmt.Println(array[2:7])
	fmt.Println(array[1:6])
	fmt.Println(array)
}
