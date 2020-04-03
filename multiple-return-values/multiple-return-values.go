package main

import "fmt"

func vals() (int, int, int) {
	return 3, 8, 7
}

func main() {
	a, b, _ := vals()
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	_, c, _ := vals()
	fmt.Println("c", c)
}
