package main

import "fmt"

const (
	year1 = 2019 + iota
	year2
	year3
	year4
)

func main() {

	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)
}
