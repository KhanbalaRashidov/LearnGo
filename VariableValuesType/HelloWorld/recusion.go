package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return factorial(n-1) * n
}

func main() {
	fmt.Println(factorial(25))

	var fibonachi func(n int) int

	fibonachi = func(n int) int {
		if n < 2 {
			return n
		}
		return fibonachi(n-1) + fibonachi(n-2)
	}
	fmt.Println(fibonachi(3))
}
