package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%v modulus 4 = %v\n", i, i%4)
	}
}
