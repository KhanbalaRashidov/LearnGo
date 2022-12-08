package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	jsonData := []byte(`{"FirstName":"Khanbala","LastName":"Rashidov","Age":22}`)
	fmt.Println(jsonData)

	var person Person

	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(person)

	person.Deserializer(jsonData)

}

func (p *Person) Deserializer(arr []byte) {
	err := json.Unmarshal(arr, p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*p)
}
