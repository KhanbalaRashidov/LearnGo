package main

import "fmt"

func main() {

	test1 := closure()
	fmt.Println(test1())
	fmt.Println(test1())
	fmt.Println(test1())

	test2 := closure()
	fmt.Println(test2())
	fmt.Println(test2())
	fmt.Println(test2())
	fmt.Println(test2())
	fmt.Println(test2())
}
func closure() func() int {
	num := 0
	return func() int {
		num++
		return num
	}
}
