package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//sendMessage("Hello")
	go sendMessage2("Hello2")
	fmt.Println("listening...")
	time.Sleep(time.Second * 2)
	fmt.Println("Leaving")
}

func sendMessage(message string) {
	for i := 0; ; i++ {
		fmt.Println(message, i)
		time.Sleep(time.Second)
	}
}

func sendMessage2(message string) {
	for i := 0; ; i++ {
		fmt.Println(message, i)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}
