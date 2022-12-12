package main

import "fmt"

func main() {

	c := make(chan string)
	go sendMessage("Hello", c)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Leaving...")
}

func sendMessage(message string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", message, i)
	}
}
