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

type human interface {
	info() string
}

func (p person) info() string {
	return fmt.Sprint("Name:", p.firstName, " Surname:", p.lastName, " Age:", p.age)
}

func (s secretAgent) info() string {
	return fmt.Sprint("Name:", s.firstName, " Surname:", s.lastName, " Age:", s.age, " Ltk:", s.ltk)
}

func print(h human) {
	switch h.(type) {
	case person:
		fmt.Println("person struct:")
		fmt.Println(h.info())
	case secretAgent:
		fmt.Println("secretAgent struct:")
		fmt.Println(h.info())
	}
}
const(
	Mondey iota 
	Tuesday 
	_

)
func main() {

	p := person{
		firstName: "Khanbala",
		lastName:  "Reshidov",
		age:       22,
	}

	s := secretAgent{
		person: p,
		ltk:    true,
	}

	//info Method
	fmt.Println(p.info())
	fmt.Println(s.info())

	//polymorphism

	print(p)
	print(s)

	//type
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", s)
}
