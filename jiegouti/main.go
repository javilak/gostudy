package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	fmt.Println(m)
	for a, stu := range stus {
		m[stu.name] = &stu //将stu中的name元素放到里面变成键值，然后值则是附上函数中的stu的地址。所以会出现一个问题，stu的地址是相同的。
		fmt.Println(m, a)
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.age)
	}
}
