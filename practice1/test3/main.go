package main

import "fmt"

var x int = 25
var y = "Khanbala Rashidov"
var z = true

func main() {

	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)

	fmt.Println(s)
}
