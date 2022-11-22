package main

import "fmt"

const (
	a = iota
	b
	c
)

const (
	d = iota + 1
	e
	f
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Printf("%T", a)

}
