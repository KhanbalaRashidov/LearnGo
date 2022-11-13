package main

import "fmt"

func main() {

	//+ operator
	sum := 25 + 3
	fmt.Println(sum)

	//- operator
	subtracts := 25 - 3
	fmt.Println(subtracts)

	//* operator
	multiple := 25 * 3
	fmt.Println(multiple)

	// / operator
	divide := 25 / 3
	fmt.Println(divide)

	// % operator
	modules := 25 % 3
	fmt.Println(modules)

	//++ operator
	num := 25
	fmt.Println(num)
	num++
	fmt.Println(num)

	//-- operator
	num--
	fmt.Println(num)

	//&& operator
	fmt.Println(true && false)

	// || operator
	fmt.Println(true || false)

	// ! operator
	fmt.Println(!(true && false))

	//*= operator
	num *= 3
	fmt.Println(num)

	// /= operator
	num /= 3
	fmt.Println(num)

}
