package main

import "fmt"

func main() {
	num1 := 25
	fmt.Println(num1)

	//closure scope
	//correct
	{
		num2 := 3
		fmt.Println(num2)
	}

	//incorrect
	//fmt.Println(num2) undecalred variable

	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	dec := decrement()
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())

}

func increment() func() int {
	var num int
	return func() int {
		num++
		return num
	}
}

func decrement() func() int {
	var num int
	return func() int {
		num--
		return num
	}
}
