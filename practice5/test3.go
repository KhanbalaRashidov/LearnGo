package main

import "fmt"

type Vehicle struct {
	doors int
	color string
}

type Truck struct {
	Vehicle
	fourWheel bool
}

type Sedan struct {
	Vehicle
	luxury bool
}

func main() {
	truck := Truck{
		Vehicle: Vehicle{
			doors: 2,
			color: "black"},
		fourWheel: true,
	}

	sedan := Sedan{
		Vehicle: Vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: false,
	}

	fmt.Println(truck)
	fmt.Println(sedan)

	fmt.Println(truck.doors)
	fmt.Println(sedan.doors)
}
