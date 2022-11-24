package main

import "fmt"

func main() {

	for i := 0; i < 25; i++ {
		fmt.Printf("Loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("inner loop: %d\t \n", j)
		}
	}
}
