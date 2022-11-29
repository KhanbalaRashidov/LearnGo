package main

import "fmt"

type Person struct {
	firstName  string
	lastName   string
	favFlavors []string
}

func main() {

	person := Person{
		firstName:  "Khanbala",
		lastName:   "Rashidov",
		favFlavors: []string{"chocolate", "vanilia"},
	}

	fmt.Println(person.firstName)
	fmt.Println(person.lastName)

	for index, flavor := range person.favFlavors {
		fmt.Println(index, flavor)
	}

}
