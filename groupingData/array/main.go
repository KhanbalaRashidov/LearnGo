package main

import (
	"fmt"
)

func main() {

	var arr [100]int

	for i := 0; i < 100; i++ {
		arr[i] = i + 1
	}

	printArray(arr[:])
}

func printArray(arr []int) {

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v\t", arr[i])
		if arr[i]%10 == 0 {
			fmt.Println("")
		}
	}
}
