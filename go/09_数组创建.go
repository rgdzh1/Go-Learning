package main

import (
	"fmt"
)

func main() {
	// 1,创建数组
	var arr1 [4]int
	// 2,数组访问
	arr1[0] = 0
	arr1[1] = 1
	arr1[2] = 2
	arr1[3] = 3
	fmt.Println(arr1[0])
	fmt.Println(arr1[1])
	// 3,数组长度
	fmt.Println(len(arr1))
	// 4,数组容量
	fmt.Println(cap(arr1))

	// 数组的创建方法
	var arr2 [4]int
	fmt.Println(arr2) // 默认填充0
	var arr3 = [4]int{1, 2, 3, 4}
	fmt.Println(arr3)
	var arr4 = [5]int{1, 2, 3}
	fmt.Println(arr4)               // 不够填充0
	var arr5 = [4]int{1: 5, 3: 800} // 1:代表索引
	fmt.Println(arr5)
	var arr6 = [...]string{1: "第一", 10000: "无限"} // [...] 表示长度不限
	fmt.Println(arr6)
}
