package main

import "fmt"

func main() {

	person := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Khanbala",
		lastName:  "Rashidov",
		age:       22,
	}

	fmt.Println(person)

}
