package main

import "fmt"

func main() {
	/**
	闭包:
	一个外层函数中有内层函数,该内层函数中会操作外层函数的局部变量(外层函数中的参数,或者外层函数中直接定义的局部变量)
	并且该外层函数的返回值就是这个内层函数,这个内层函数和外层函数的局部变量统称为闭包结构.
	局部变量的生命周期会发生改变,正常的局部变量随着函数的调用而创建,随着函数的结束而销毁,
	但是闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁,因为内层函数还要继续使用.
	*/
	f1 := increment() // 获取内层函数,该函数持有外层函数局部变量i
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())

	f2 := increment() // 当重新获取内层函数的时候,外层函数中的局部变量会重新创建
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())

}

func increment() func() int { // 外层函数, 该外层函数返回值也是一个函数
	//1 定义一个局部变量
	i := 0
	//2 定义了一个匿名函数,给外层函数的局部变量自增并返回
	fun := func() int {
		i++
		return i
	}
	//3 返回内层函数
	return fun
}
