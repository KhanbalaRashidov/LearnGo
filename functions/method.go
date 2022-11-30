package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) print() {
	fmt.Println("I'm", s.firstName, s.lastName)
	fmt.Println("I'm ", s.age)
}

func main() {

	agent := secretAgent{
		person: person{
			firstName: "Khanbala",
			lastName:  "Reshidov",
			age:       22},
		ltk: true,
	}
	agent.print()

}
