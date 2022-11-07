package main

import "fmt"

func seqInt() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := seqInt()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newNextInt := seqInt()
	fmt.Println(newNextInt())
}
