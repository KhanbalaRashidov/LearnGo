package main

import "fmt"

func main() {

	number := 25
	for {

		if number > 100 {
			break
		}

		fmt.Println(number)
		number++

	}
}
