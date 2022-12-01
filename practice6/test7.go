package main

import "fmt"

var gFunc func()

func main() {

	f := func() {
		for i := 1; i <= 100; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("%T\n", f)
	gFunc = f
	gFunc()
	fmt.Printf("%T\n", gFunc)

	fmt.Println("Done!")

}
