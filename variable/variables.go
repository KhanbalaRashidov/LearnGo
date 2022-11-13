package main

import (
	"fmt"
)

func main() {

	//string type start
	var message1 string
	message1 = "Hello Go"
	fmt.Println(message1)

	var message2 string = "Hello Go"
	fmt.Println(message2)

	var message3 = "Hello Go"
	fmt.Println(message3)

	//end

	//int type start
	var num1 int
	num1 = 3
	fmt.Println(num1)

	var num2 int = 25
	fmt.Println(num2)

	var num3 = 2503
	fmt.Println(num3)

	var num4, num5, num6 int // zero value = 0
	fmt.Println(num4)
	fmt.Println(num5)
	fmt.Println(num6)

	var num7, num8, num9 int = 7, 8, 9
	fmt.Println(num7)
	fmt.Println(num8)
	fmt.Println(num9)

	//end

	//bool type start
	var check1 bool = true
	fmt.Println(check1)

	var check2 = true
	fmt.Println(check2)

	//end

	//multiple type assign
	var num, str, check = 3, "string", true
	fmt.Println(num, str, check)

	//short declaration
	number, str2, check3 := 25, "test", false
	fmt.Println(number, str2, check3)

}
