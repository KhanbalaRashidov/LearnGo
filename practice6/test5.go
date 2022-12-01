package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}
type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s.area())
}
func main() {

	c := circle{
		radius: 25,
	}

	s := square{
		length: 3,
	}

	info(c)
	info(s)
}
