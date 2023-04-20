package main

import "fmt" // 引入一个标准库包

func eqal(a, b int) {
	var c int
	c = a * b
	fmt.Printf(" %d * %d = %d\t", a, b, c)
	if a == b {
		fmt.Print("\n")
	}
}
func main() {

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			eqal(j, i)
		}
	}
}
