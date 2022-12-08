package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := `Password123`

	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)

	password1 := `Password1234`
	err = bcrypt.CompareHashAndPassword(bs, []byte(password1))
	if err != nil {
		fmt.Println("Can't Login")
		return
	}
	fmt.Println("Successful...")

}
