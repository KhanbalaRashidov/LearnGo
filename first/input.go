package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Text Input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("Enter a number: ")
	numStr, _ := reader.ReadString('\n')
	number, err := strconv.ParseFloat(strings.TrimSpace(numStr), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(number)
	}

}
