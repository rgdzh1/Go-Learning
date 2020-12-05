package main

import "fmt"

func main() {
	i := 2
	iii(&i)
	fmt.Println(i)
}

// 引用传递,传递的是地址值
func iii(i *int) {
	// * 取指针的值进行赋值
	*i = 100
}
