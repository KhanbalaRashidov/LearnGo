package main

import "fmt"

func main() {
	maps := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	fmt.Println(maps)

	maps[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	for key, val := range maps {
		fmt.Println("This is key:", key)
		for index, value := range val {
			fmt.Println("\t", index, value)
		}
	}
}
