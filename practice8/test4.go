package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	texts := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(numbers)
	sort.Ints(numbers)
	fmt.Println(numbers)

	fmt.Println(texts)
	sort.Strings(texts)
	fmt.Println(texts)
}
