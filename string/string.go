package main

import "fmt"

func main() {
	var a = [...]string{"北京", "上海", "深圳"}
	for index, value := range a {
		fmt.Println(index, value)
	}
}
