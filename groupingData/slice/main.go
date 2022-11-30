package main

import (
	"fmt"
)

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice)

	//make function
	var array = make([]int, 10, 13)

	fmt.Println(array)
	fmt.Println(len(array))

	array = append(array, 25, 3, 03, 03, 25, 25, 25)
	fmt.Println(array)
	fmt.Println(len(array))

	dimensionalSlice := make([][]int, 10)

	for index := range dimensionalSlice {
		dimensionalSlice[index] = make([]int, 10)
		for index2 := range dimensionalSlice[index] {
			dimensionalSlice[index][index2] = index2 + 1
		}
	}

	for _, element := range dimensionalSlice {

		fmt.Printf("%v \n", element)
	}

}
