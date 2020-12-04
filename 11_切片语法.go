package main

import "fmt"

func main() {
	// 数组,定长
	arr := [4]int{0, 1, 2, 3}
	fmt.Println(arr)
	// 切片,变长
	slice := []int{0, 1, 2, 3, 4}
	fmt.Println(slice)
	// 打印各自类型
	fmt.Printf("%T,%T\n", arr, slice) // [4]int,[]int
	// make 函数, 专门用来创建引用类型数据,slice map chan
	slice1 := make([]int, 3, 8) //3:表示长度,8:表示容量
	fmt.Println(slice1)
	fmt.Printf("容量 %d, 长度 %d\n 类型%p", cap(slice1), len(slice1), slice1)
	// 操作切片
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	fmt.Println(slice1)
	// append 在切片末尾追加元素,因为切片是可变的,所以对切片append之后需要重新赋值给切片名变量
	slice2 := make([]int, 2, 8)
	fmt.Println(slice2)
	slice2[0] = 10
	slice2[1] = 11
	slice2 = append(slice2, 1, 2, 3, 4)
	fmt.Println(slice2)
	// 使用append将切片1中元素添加到切片2中
	slice2 = append(slice2, slice1...)
	fmt.Println(slice2)
	// 切片遍历
	for i := 0; i < len(slice2); i++ {
		fmt.Println(i)
	}
	for i, v := range slice2 {
		fmt.Printf("%d %d\n", i, v)
	}
	// 从数组中创建切片 使用:
	arr1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice3 := arr1[:5]  // 1-5
	slice4 := arr1[3:8] // 4-8
	slice5 := arr1[5:]  // 6-10
	slice6 := arr1[:]   // 1-10
	// slice3,slice4,slice5,slice6 切片内存地址指向的其实是数组中的数据,改变任意切片中的数据就会改变其他切片中的数据和数组中的数据
	// 如果对slice3切片追加数据,该切片扩容之后,slice3所对应的内存地址就会重新分配.
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(slice6)
}
