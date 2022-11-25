package main

import "fmt"

func main() {
	year := 1972
	for {
		if year > 2023 {
			break
		}
		fmt.Printf("Year: %v\n", year)
		year++
	}
}
