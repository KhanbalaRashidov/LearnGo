package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)

	slice2 := slice[:5]
	fmt.Println(slice2)

	slice3 := slice[5:]
	fmt.Println(slice3)

	//point slice array ==> slice[5]=7
	slice3[0] = 7

	fmt.Println(slice)

}
