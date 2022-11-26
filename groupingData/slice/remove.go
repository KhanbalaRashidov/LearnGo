package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)

	slice = remove(slice, 9)
	fmt.Println(slice)
}

func remove(slice []int, index int) []int {

	return append(slice[:index], slice[index+1:]...)

}
