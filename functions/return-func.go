package main

import "fmt"

func main() {
	hello := greeting()
	fmt.Printf("Type: %T \t return func: %v\n", hello, hello())

	bye := goodBye()
	fmt.Printf("Type: %T \t return func: %v\n", bye, bye())

	//or
	fmt.Println("\n------------------------------------------------------\n")
	fmt.Printf("Type: %T \t return func: %v\n", greeting(), greeting()())
	fmt.Printf("Type: %T \t return func: %v\n", goodBye(), goodBye()())

}

func greeting() func() string {
	return func() string {
		return "Hello Gopher"
	}
}

func goodBye() func() string {

	return func() string {
		return "Good bye"
	}
}
