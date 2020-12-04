package main

import "fmt"

func main() {
	//
	i := oper(10, 20, add) // add 代表方法的内存地址值
	fmt.Println(i)
	i2 := oper(10, 4, sub)
	fmt.Println(i2)

	// 第三种写法
	i3 := oper(10, 20, func(i int, i2 int) int {
		return i + i2
	})
	fmt.Println(i3)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

// 高阶函数 fun参数只能接收方法的内存地址值
func oper(a, b int, fun func(int, int) int) int {
	return fun(a, b)
}
