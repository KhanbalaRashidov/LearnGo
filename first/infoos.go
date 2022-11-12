package main

import (
	"fmt"
	"os"
)

func main() {

	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }

	userName := os.Getenv("USERNAME")
	goPath := os.Getenv("GOPATH")
	homePath := os.Getenv("HOMEPATH")
	userDomain := os.Getenv("USERDOMAIN")
	processorArhitecture := os.Getenv("PROCESSOR_ARCHITECTURE")
	processorLevel := os.Getenv("PROCESSOR_LEVEL")
	goRoot := os.Getenv("GOROOT")

	fmt.Println("Username:", userName)
	fmt.Println("Gopath:", goPath)
	fmt.Println("HomePath:", homePath)
	fmt.Println("UserDomain:", userDomain)
	fmt.Println("ProcessorArhitecture:", processorArhitecture)
	fmt.Println("ProcessorLevel:", processorLevel)
	fmt.Println("GoRoot:", goRoot)
}
