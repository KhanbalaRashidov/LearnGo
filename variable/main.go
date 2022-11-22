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

	fmt.Println("\nstr1 output")
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%#U ", str1[i])
	}

	fmt.Println("\nstr2 output")
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%#U ", str2[i])
	}

	//for range
	fmt.Println("\n str1 output ")
	for i, v := range str1 {
		fmt.Println(i, v)
	}

	fmt.Println("\n str2 output ")
	for i, v := range str2 {
		fmt.Println(i, v)
	}

	fmt.Println("\n str1 output hex \n")
	for i, v := range str1 {
		fmt.Printf("index:%d  hex: %#x \n", i, v)
	}

	fmt.Println("\n str2 output hex \n")
	for i, v := range str2 {
		fmt.Printf("index:%d  hex: %#x \n ", i, v)
	}

}
