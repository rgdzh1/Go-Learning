package main

import "fmt"

func main() {
	// 匿名结构体
	student := struct {
		name string
		age  int
	}{
		name: "小白",
		age:  18,
	}
	fmt.Println(student)
	fmt.Println(student.name, student.age)
}

type Student struct {
	// 匿名字段,用数据类型替代了结构体中的变量名称, 这样每个结构体不能有重复的数据类型
	string
	int
}
