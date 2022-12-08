package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "Khanbala",
		Age:   32,
	}

	u2 := user{
		First: "Mushfig",
		Age:   50,
	}

	u3 := user{
		First: "Rashid",
		Age:   70,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	usersJson, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(usersJson))
}
