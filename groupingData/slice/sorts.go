package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 5, 6, 3, 7, 2, 9, 11, 4}
	fmt.Println("unsorted:", numbers)
	sort.Ints(numbers)
	fmt.Println("sorted:", numbers)

	names := []string{"Khanbala", "Rashidov", "Mushfig"}
	fmt.Println("unsorted:", names)
	sort.Strings(names)
	fmt.Println("sorted:", names)
}
