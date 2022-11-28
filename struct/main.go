package main

import "fmt"

type Person struct {
	firsName string
	lastName string
	age      int
}

func main() {
	person1 := Person{
		firsName: "Khanbala",
		lastName: "Rashidov",
		age:      22,
	}
	person2 := Person{
		firsName: "Mushfig",
		lastName: "Rashidov",
		age:      50,
	}

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.firsName, person1.lastName, person1.age)
	fmt.Println(person2.firsName, person2.lastName, person2.age)
}
