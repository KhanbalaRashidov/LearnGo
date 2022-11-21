package main

import "fmt"

var str1 string
var str2 string

func main() {

	str1 = `
	Hello
	world
	Gopher`

	fmt.Println(str1)
	fmt.Printf("%T\n", str1)
	fmt.Println([]byte(str1))

	str2 = "Hello world Gopher"
	fmt.Println(str2)
	fmt.Printf("%T\n", str2)
	fmt.Println([]byte(str2))

}
