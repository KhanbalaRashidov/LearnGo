package main

import "fmt"

func main() {

	for i := 33; i <= 122; i++ {
		fmt.Printf("decimal: %v\t\t hex:  %#x\t\t ascii: %#U  \t\t character: %#q  \n", i, i, i, i)
	}
}
