package main

import (
	"fmt"
)

func main() {

	fact := factorial(3)
	fmt.Println(fact)
	fib := fibonacci(25)
	fmt.Println(fib)

	factLoop := factorialLoop(25)
	fmt.Println(factLoop)
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

func factorialLoop(n int) int {
	multiply := 1
	for ; n > 0; n-- {
		multiply *= n
	}
	return multiply
}
