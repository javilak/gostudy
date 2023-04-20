package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "how do you do"
	var wordc = make(map[string]int, 10)
	var words = strings.Split(s, " ")
	for _, word := range words {
		v, ok := wordc[word]
		if ok {
			wordc[word] = v + 1
		} else {
			wordc[word] = 1
		}
	}
	for k, v := range wordc {
		fmt.Println(k, v)
	}
}
