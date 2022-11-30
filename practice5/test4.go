package main

import "fmt"

func main() {

	person := struct {
		firstName string
		lastName  string
		favColors []string
	}{
		firstName: "Khanbala",
		lastName:  "Rashidov",
		favColors: []string{
			"black",
			"green",
			"white",
		},
	}

	fmt.Println(person.firstName)
	fmt.Println(person.lastName)
	for index, color := range person.favColors {
		fmt.Println(index, color)
	}
	score := 0     
}
