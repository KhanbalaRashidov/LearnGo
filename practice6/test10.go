package main

func main() {

}
func closure() func() int {
	num := 0
	return func() int {
		num++
		return num
	}
}
