package main

import "fmt"

func main() {

	hello()
	helloName("Khanbala")
	male := helloMale("Khanbala")
	print(male)
	female := helloFemale("Zeyna")
	print(female)
	fullName := printFullName("Khanbala", "Rashidov")
	print(fullName)

}

func hello() {
	fmt.Println("Hello Gopher")
}

func helloName(name string) {
	fmt.Println("Hello:", name)
}

func helloMale(name string) string {
	return fmt.Sprint("Mr.", name)
}
func helloFemale(name string) string {
	return fmt.Sprint("Mrs.", name)
}

func print(name string) {
	fmt.Println(name)
}

func printFullName(fn, ln string) string {
	return fmt.Sprint(fn, " ", ln)
}
