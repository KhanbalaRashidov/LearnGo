package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) print() {
	fmt.Println("I'm", p.firstName, p.lastName, " I'm", p.age, "years old.")
}
func main() {

	p := person{
		firstName: "Khanbala",
		lastName:  "Rashidov",
		age:       22,
	}
	p.print()
}
