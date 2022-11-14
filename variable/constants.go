package main

import "fmt"

type Company string

const (
	Google    Company = "Google"
	Facebook  Company = "Facebook"
	Microsoft Company = "Microsoft"
)

func PrintCompany(company Company) {
	fmt.Println(company)
}
func main() {
	PrintCompany(Google)
	PrintCompany(Facebook)
	PrintCompany(Microsoft)

}
