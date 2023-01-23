package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(<-chan int) //receive
	ch3 := make(chan<- int) //send

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
	fmt.Println(ch3)
}
