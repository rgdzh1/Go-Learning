package main

import "fmt"

func main() {
	// 数组,值类型
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := arr1
	fmt.Println(arr1)
	arr2[0] = 20 //修改arr2中的值对arr1中的值无影响
	fmt.Println(arr1)
	fmt.Println(arr2)
	// 切片,引用类型
	slice1 := []int{1, 2, 3}   // slice1内存地址对应的存储中存储的地址值指向的是一个底层数组
	slice2 := slice1           // slice1只是将内存地址对应的存储中存的地址值赋值给了slice2
	slice2[0] = 500            // 修改了切片2,切片1中的值也改变了
	fmt.Printf("%p\n", slice1) // 这里打印的是slice1所对应的内存地址中存储的内存地址值
	fmt.Printf("%p\n", slice2)
	fmt.Printf("%p\n", &slice1) // 这里打印的是slice1所对应的内存地址
	fmt.Printf("%p\n", &slice2)

	// 深拷贝:拷贝的是数据本身
	// 浅拷贝:拷贝的是数据地址
	s1 := []int{1, 2, 3, 4}
	s2 := []int{5, 6, 7}
	// copy(s1,s2) // 直接将s2拷贝到s1当中,从0索引开始拷贝
	fmt.Println(s1)
	fmt.Println(s2)
	// copy(s2,s1) // 将s1拷贝到s2中,从0索引开始
	fmt.Println(s1)
	fmt.Println(s2)
	copy(s1[2:], s2[1:]) // 将s2中从索引1处开始的值,拷贝复制到s1中,并且s1是从索引2处开始接收这些复制的
	fmt.Println(s1)

}
