package main

import (
	"fmt"
)

func main() {

	fact := factorial(3)
	fmt.Println(fact)
	fib := fibonacci(25)
	fmt.Println(fib)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
