package main

import "fmt"

func main() {
	// 函数指针,在go中函数默认看作一个指针,没有*
	var a func() // 定义一个函数变量, 该变量其实就是一个指针
	a = funmmm   // 将函数指针赋值给a
	a()          // 调用函数

	// 普通函数返回的值的指针是不一样的
	arr1 := funll()
	fmt.Printf("普通函数调用返回的数组的指针%p\n", &arr1)
	// 指针函数,返回的是指针
	arr2 := funzzz()
	fmt.Printf("函数指针调用返回的数组的指针%p\n", &arr2)
	// 函数返回切片
	slice1 := funqqq()
	fmt.Printf("函数中返回切片的指针%p\n", slice1)
}

// 函数返回切片其实最终返回的是该切片的指针
func funqqq() []int {
	slice1 := []int{1, 2, 3, 4}
	fmt.Printf("函数中切片的指针%p\n", slice1)
	return slice1
}

// 指针函数
func funzzz() *[4]int {
	arr := [4]int{1, 2, 3, 4}
	fmt.Printf("函数中数组的指针%p\n", &arr)
	return &arr
}

// 普通函数,
func funll() [4]int {
	arr := [4]int{1, 2, 3, 4}
	fmt.Printf("函数中数组的指针%p\n", &arr)
	return arr
}

func funmmm() {
	fmt.Println("通过函数指针调用函数")
}
