package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (name ByName) Len() int { return len(name) }

func (name ByName) Swap(i, j int) { name[i], name[j] = name[j], name[i] }

func (name ByName) Less(i, j int) bool { return name[i].Name > name[j].Name }

func main() {

	p1 := Person{Name: "James", Age: 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}
	fmt.Println(people)

	sort.Sort(ByName(people))
	fmt.Println(people)

}
