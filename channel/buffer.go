package main

import "fmt"

func main() {

	c := make(chan int, 2)

	c <- 25
	c <- 26

	fmt.Println(<-c)
	fmt.Println(<-c)

	/*does not run
	c := make(chan int, 1)

	c <- 25
	c <- 26

	fmt.Println(<-c)

	*/
}
