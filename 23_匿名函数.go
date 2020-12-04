package main

import "fmt"

func main() {
	// 匿名函数,没有名字的函数
	// 匿名函数直接进行调用,通常只能使用一次
	func() {
		fmt.Println("我是匿名函数1")
	}() // 直接调用

	// 定义一个匿名函数,将函数地址值赋值给一个变量,利用该变量进行调用
	funY := func() {
		fmt.Println("我是匿名函数2")
	}
	funY()
	// 带参数的匿名函数
	func(a int, b int) {
		fmt.Println("我是匿名函数2", a, b)
	}(10, 20)
	// 带返回值的匿名函数
	i := func(a, b int) int {
		return a + b
	}(10, 20)
	fmt.Println(i)
	// 匿名函数的作用
	// 1. 将匿名函数作为另一个函数的参数,回调函数
	// 2. 将匿名函数作为另一个函数的返回值,可以形成闭包结构
}
