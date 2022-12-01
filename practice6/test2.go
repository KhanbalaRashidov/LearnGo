package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum1 := sumVariadic(nums...)
	sum2 := sumSlice(nums)

	fmt.Println(sum1, sum2)

}
func sumVariadic(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func sumSlice(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}
