package main

import "fmt"

/**

 */

func main() {
	// 函数地址所对应的内存空间中存储了函数体,当使用fun()调用函数后,就会将该地址对应的内存空间中的函数体从上到下执行一遍.
	fmt.Printf("%T \n", funxxx) // 函数类型
	fmt.Println("函数地址", funxxx) // 函数地址,

	// 直接定义函数类型的变量
	var fun func()   // 定义函数变量
	fmt.Println(fun) // 此时该函数类型变量还未初始化
	fun = funxxx     // 这一步是将funxxx函数的地址值赋值给fun
	fmt.Println(fun) // 打印后可以看到fun现在有地址值了
	fun()            // fun变量加了(),代表调用该函数.
	funxxx()
}

func funxxx() {
	fmt.Println("打印数据")
}
