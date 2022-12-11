package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

/*
 * Complete the 'encryption' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */
func detectGridSize(l int) (int, int) {
	s := math.Sqrt(float64(l))

	f := int(math.Floor(s))
	if f*f >= l {
		return f, f
	}

	c := int(math.Ceil(s))
	if f*c >= l {
		return f, c
	}

	return c, c
}

func encodeString(m, n int, s string) string {
	var str string
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if k := i*n + j; k < len(s) {
				str += string(s[k])
			}
		}
		str += "\n"
	}
	return str
}

func encryption(s string) string {
	// Write your code here
	var length = len(s)
	var m, n int = detectGridSize(length)
	var str string = encodeString(m, n, s)
	return str
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := encryption(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
