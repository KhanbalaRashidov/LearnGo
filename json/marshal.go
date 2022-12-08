package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"FirstName"`
	Last  string `json:"LastName"`
	Age   int    `json:"Age"`
}

func main() {
	p1 := Person{
		First: "Khanbala",
		Last:  "Rashidov",
		Age:   22,
	}
	fmt.Println(p1)
	ps, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(ps))

	//Custom Method
	p1.Serializer()
}

func (p Person) Serializer() {
	ps, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(ps))
}
