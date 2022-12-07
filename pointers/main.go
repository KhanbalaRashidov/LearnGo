package main

import "fmt"

func main() {
	num := 2503
	fmt.Println("num=", num)
	//or
	fmt.Println(*&num)
	//

	fmt.Println("Pointer: ", &num)

	var num2 *int = &num
	fmt.Println(num2)
	fmt.Println("num2=", *num2)
	*num2 = 25
	fmt.Println("num=", num)
	fmt.Println("num2=", *num2)

	num3 := &num
	fmt.Println(num3)

}
