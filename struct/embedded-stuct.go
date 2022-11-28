package main

import "fmt"

type Person struct {
	firsName string
	lastName string
	age      int
}

type SecretAgent struct {
	Person
	ltk bool
}

func (agent SecretAgent) FullPrint() {
	fmt.Printf("Agent Name: \t %v %v \t Age: \t %v \t ltk:\t %v\n", agent.Person.firsName, agent.Person.lastName, agent.Person.age, agent.ltk)
}

func (agent SecretAgent) ShortPrint() {
	fmt.Printf("Agent Name: \t %v %v \t Age: \t %v \t ltk:\t %v\n", agent.firsName, agent.lastName, agent.age, agent.ltk)
}

func main() {
	person1 := Person{
		firsName: "Khanbala",
		lastName: "Rashidov",
		age:      22,
	}
	person2 := Person{
		firsName: "Mushfig",
		lastName: "Rashidov",
		age:      50,
	}

	agent1 := SecretAgent{
		Person: person1,
		ltk:    true,
	}

	agent2 := SecretAgent{
		Person: person2,
		ltk:    false,
	}

	fmt.Println(agent1)
	fmt.Println(agent2)

	agent1.FullPrint()
	agent1.ShortPrint()

	agent2.FullPrint()
	agent2.ShortPrint()
}
