package main

import "fmt"

func main() {
	array := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	array = append(array[:3], array[6:]...)
	fmt.Println(array)
}
