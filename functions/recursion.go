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
	fmt.Println(Fib3(24, make([]int, 26)))
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

func Fib3(n int, k []int) int {
	k[0] = 1
	k[1] = 1
	for i := 2; i <= n; i++ {
		k[i] = k[i-1] + k[i-2]
	}
	return k[n]

}
