package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}
type personInterface interface {
	print() string
}

// func (p person) print() {
// 	fmt.Printf("First name:%v\t Last name:%v\n", p.firstName, p.lastName)
// }

func (p *person) print() string {
	return p.firstName + p.lastName
}

func info(i personInterface) {
	fmt.Println(i.print())
}

func main() {
	p := person{
		firstName: "Khanbala",
		lastName:  "Rashidov",
	}
	p.print()

	info(&p)

}
