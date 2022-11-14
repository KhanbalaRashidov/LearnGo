package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Convert
	var str string = "2503"

	num, _ := strconv.Atoi(str)

	fmt.Printf("%T, %v\n", num, num)

	num2 := 25
	str2 := strconv.Itoa(num2)

	fmt.Printf("%T, %v\n", str2, str2)

	//Parsing
	var strNum string = "325"
	numStr, _ := strconv.ParseInt(strNum, 10, 32)
	fmt.Println(numStr)
	fmt.Printf("%T, %v\n", numStr, numStr)

	//Casting

	var numInt int = 2503
	var numFloat float32 = float32(numInt)
	fmt.Println(numFloat)
	fmt.Printf("%T, %v\n", numFloat, numFloat)

	numInt = int(numFloat)
	fmt.Printf("%T, %v\n", numInt, numInt)

	//var numUInt uint8 = 258 throw integer overflow
	//fmt.Println(numUInt)
}
