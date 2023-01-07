package main

import "fmt"

func main() {

	/* does not run
	c := make(chan int)

	c <- 3
	fmt.Println(<-c)

	*/

	c := make(chan int)

	go func() {
		c <- 3
	}()

	fmt.Println(<-c)

	c1 := make(chan int, 3)

	c1 <- 25

	fmt.Println(<-c1)

}
