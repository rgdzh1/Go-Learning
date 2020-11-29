package main

import "fmt"

func main() {
	// 第一种定义变量
	var a, b, c = 100, "你好", 1.02
	fmt.Println(a, b, c)
	// 第二种定义变量
	var (
		name = "小红"
		age  = 18
		sex  = "女"
	)
	fmt.Println(name, age, sex)
	// 第三种定义变量
	var i int
	i = 1222
	fmt.Println(i)
	// 第四种定义变量
	var p = "哈哈"
	fmt.Println(p)
	// 第五种定义变量
	l := 3.14159 // 这种定义变量的方式不支持全局变量.
	fmt.Println(l)
}
