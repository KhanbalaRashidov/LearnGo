package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Age   int
}

func main() {

	jsonData := `[{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"M","Age":54}]`

	var people []Person

	err := json.Unmarshal([]byte(jsonData), people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

}
