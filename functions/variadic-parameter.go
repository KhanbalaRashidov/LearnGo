package main

import "fmt"

func main() {
	printNums(25, 03, 2000, 2000)
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

func printNums(nums ...int) {
	fmt.Println(nums)
	fmt.Printf("%T\n", nums)
}
func sum(nums ...int) {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	fmt.Println("Sum: ", sum)
}
