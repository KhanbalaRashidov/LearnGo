package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.NumCPU())
	fmt.Println("Goroutune:", runtime.NumGoroutine())

	wg.Add(1)
	go Numbers()

	Numbers2()

	fmt.Println(runtime.NumCPU())
	fmt.Println("Goroutune:", runtime.NumGoroutine())
	wg.Wait()
}

func Numbers() {
	for i := 0; i < 25; i++ {
		fmt.Println("number:", i+1)
	}
	wg.Done()
}

func Numbers2() {
	for i := 0; i < 25; i++ {
		fmt.Println("number:", i+1)
	}
}
