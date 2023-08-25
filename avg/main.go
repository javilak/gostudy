package main

import "fmt"

func main() {
	k := make([]float64, 20)
	var sum float64
	x := 10
	b := 0.0346
	b = 0
	for i, j := 1, 1.00; i < x; i++ {
		b = (4200*j + 122204) * 0.0346
		k[i] = (4200*j+122204)/(4200*(j+1)+122204) - 1
		j++
		fmt.Printf("%.3f\n", k[i])
		sum += k[i]
		fmt.Printf("%.3f\n", b)
	}
	fmt.Printf("%.3f/%d", sum, x+6)
}
