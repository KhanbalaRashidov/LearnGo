package main

import "fmt"

var age int = 22
var firstName string = "Khanbala"

// custom type
type name string

var lastName name = "Rashidov"

func main() {

	//age variable
	fmt.Println("age:", age)
	fmt.Printf("%T\n", age)

	//firstName variable
	fmt.Println("firstName:", firstName)
	fmt.Printf("%T\n", firstName)

	//lastName variable
	fmt.Println("lastName:", lastName)
	fmt.Printf("%T\n", lastName)

	//firstName=stringlastName incorrect

	//correct
	firstName = string(lastName)
	fmt.Println("firstName:", firstName)
	fmt.Printf("%T\n", firstName)

}
