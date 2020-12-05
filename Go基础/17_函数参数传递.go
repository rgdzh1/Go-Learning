package main

import "fmt"

func main() {
	// 值传递
	aar1 := [4]int{1, 2, 3, 4}
	fmt.Println("函数调用前的数组", aar1)
	print1(aar1)
	fmt.Println("函数调用后的数组", aar1)
	// 引用传递
	slice1 := []int{1, 2, 3, 4}
	fmt.Println("函数调用前的切片", slice1)
	print2(slice1)
	fmt.Println("函数调用后的切片", slice1)
}

// 值类型的数据都是值传递,基本数据类型, 数组和结构体
func print1(arr [4]int) {
	fmt.Println("函数中的数组", arr)
	arr[0] = 100
	fmt.Println("函数中,修改后的数组", arr)
}

// 引用传递,传递的是数据的地址,它会导致多个变量指向同一块内存.
// 默认引用数据类型的数据,默认都是引用传递
func print2(slice []int) {
	fmt.Println("函数中的切片", slice)
	slice[0] = 100
	fmt.Println("函数中,修改后的切换", slice)
}
