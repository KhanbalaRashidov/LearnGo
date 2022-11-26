package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)

	slice = append(slice, 11, 12, 13, 14, 15)
	fmt.Println(slice)

	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice2)

	slice3 := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(slice3)
	slice2 = append(slice2, slice3...)
	fmt.Println(slice2)

}
