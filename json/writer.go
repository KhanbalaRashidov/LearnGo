package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello Gopher")

	fmt.Fprintln(os.Stdout, "Hello Gopher")

	io.WriteString(os.Stdout, "Hello Gopher")
}
