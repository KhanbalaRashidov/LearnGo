package main

import "fmt"

func main() {

	mapArray := map[int]string{1: "one", 2: "two"}

	fmt.Println(mapArray)

	for key, value := range mapArray {
		fmt.Println(key, value)
	}
	if _, x := mapArray[3]; !x {
		mapArray[3] = "Three"
	}
	fmt.Println(mapArray)

	//create map make function
	mapArray2 := make(map[int]string, 10)
	fmt.Println(mapArray2)
}
