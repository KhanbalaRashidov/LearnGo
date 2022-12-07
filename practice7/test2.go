package main

import "fmt"

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "Mr" + p.name
	//or
	//(*p).name = "Mr" + (*p).name

}

func main() {

	p := person{
		name: "Khanbala",
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}
