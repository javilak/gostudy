package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
		fmt.Println(x)
	}(x) //此处x非函数定义的x
	return 5
}
func main() { //加入注释判别输出
	fmt.Println("hellof1", f1())
	fmt.Println("f2", f2())
	fmt.Println("f3", f3())
	fmt.Println("f4", f4())
}
