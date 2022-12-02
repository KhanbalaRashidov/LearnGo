package main

import "fmt"

func main() {
	callBack := func(arr []int) int {
		sum := 0
		for _, val := range arr {
			sum += val
		}
		return sum
	}
	fmt.Println(callBackFunc(callBack, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

func callBackFunc(callBack func(arr []int) int, nums []int) int {
	result := callBack(nums)
	result++
	return result
}
