package main

import "fmt"

func plus(x []int) int {
	var b int
	for _, a := range x {
		b += a

	}
	return b
}

func main() {
	c := []int{1, 3, 5, 7, 8}

	fmt.Print(plus(c))
}
