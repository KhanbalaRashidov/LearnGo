package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Age       int    `json:"age"`
}

func main() {

	p := person{
		FirstName: "Khanbala",
		LastName:  "Rashidov",
		Age:       22,
	}
	p2 := person{
		FirstName: "Khanbala",
		LastName:  "Rashidov",
		Age:       22,
	}

	people := []person{p, p2}

	fmt.Println(p)

	pJson, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(pJson))

	var Name string = "Khanbala"
	name, err := json.Marshal(Name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(name))
}
