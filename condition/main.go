package main

import (
	"fmt"
)

func main() {

	//if statement start
	var number1 int = 25
	var number2 int = 3

	if number1 > number2 {
		fmt.Println("number1 is greater than numbur2")
	} else if number1 == number2 {
		fmt.Println("number1 equals numbur2")
	} else {
		fmt.Println("number1 is less than numbur2")
	}

	// end

	//switch statement start
	var number int
	fmt.Print("Enter last exam score: ")
	fmt.Scanf("%v", &number)
	switch {
	case number < 51:
		fmt.Println("F")
	case number >= 51 && number < 61:
		fmt.Println("E")
	case number >= 61 && number < 71:
		fmt.Println("D")
	case number >= 71 && number < 81:
		fmt.Println("C")
	case number >= 81 && number < 91:
		fmt.Println("B")
	case number >= 91 && number <= 100:
		fmt.Println("A")
	}

	//end
}
