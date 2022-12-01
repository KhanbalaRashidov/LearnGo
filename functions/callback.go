package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sum(nums...), "\n\n")

	evenTotal := evenSum(sum, nums...)
	fmt.Println("even total = ", evenTotal)
	oddTotal := oddSum(sum, nums...)
	fmt.Println("odd total = ", oddTotal)

}

func sum(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

func evenSum(f func(n ...int) int, nums ...int) int {
	var arr []int
	for _, val := range nums {
		if val%2 == 0 {
			arr = append(arr, val)
		}
	}
	return f(arr...)
}

func oddSum(f func(n ...int) int, nums ...int) int {
	var arr []int
	for _, val := range nums {
		if val%2 != 0 {
			arr = append(arr, val)
		}
	}
	total := f(arr...)
	return total
}
