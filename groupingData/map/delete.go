package main

import "fmt"

func main() {

	mapArray := map[int]string{1: "one", 2: "two"}

	fmt.Println(mapArray)

	for key, value := range mapArray {
		fmt.Println(key, value)
	}

	delete(mapArray, 2)
	fmt.Println(mapArray)

}
