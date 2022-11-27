package main

import "fmt"

func main() {

	array := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	array = append(array, 52)
	fmt.Println(array)

	array = append(array, 53, 54, 55)
	fmt.Println(array)

	array2 := []int{56, 57, 58, 59, 60}

	array = append(array, array2...)
	fmt.Println(array)

}
