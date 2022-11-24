package main

import "fmt"

func main() {
	number := 25
	for number < 100 {
		fmt.Println(number)
		number++
	}

	fmt.Println("done.")
}
