package main

import "fmt"

type Person struct {
	firstName  string
	lastName   string
	favFlavors []string
}

func main() {

	person1 := Person{
		firstName: "Khanbala",
		lastName:  "Rashidov",
		favFlavors: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	person2 := Person{
		firstName: "Mushfig",
		lastName:  "Rashidov",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	persons := map[string]Person{
		person1.firstName: person1,
		person2.firstName: person2,
	}

	for _, person := range persons {
		fmt.Println(person.firstName)
		fmt.Println(person.lastName)
		for index, flavor := range person.favFlavors {
			fmt.Println(index, flavor)
		}
		fmt.Println("-------------")
	}

}
