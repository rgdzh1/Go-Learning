package main

import "fmt"

func main() {
	// &表示获取指针
	// *表示获取指针的值
	a := 10
	fmt.Println("a的数值", a)         //10
	fmt.Printf("a的类型 %T\n", a)     //int
	fmt.Printf("a的内存地址值 %p\n", &a) //0xc00000a0b0
	// 创建一个指针变量,用于存储变量a的内存地址
	var p1 *int
	fmt.Println(p1)                       //<nil> 空指针
	p1 = &a                               //指针p1中存储变量a的地址
	fmt.Println(p1)                       // 0xc00000a0b0 打印指针p1中存储的变量a的地址值
	fmt.Printf("p1内存地址值 %p\n", &p1)       // 0xc000006030 p1的内存地址值
	fmt.Printf("p1存储的值为a的地址值%d \n ", *p1) // 10 *获取指针对应的值,指针p1中存储变量a的地址,*p1类似与*&a,结果就是变量a的值

	// 操作变量,更改数值,并不会改变指针
	a = 100
	fmt.Println("a的数值", a)        //10
	fmt.Printf("a的内存地址 %p\n", &a) //0xc00000a0b0
	// 通过指针,改变变量的数值
	*p1 = 200
	fmt.Println("a的数值", a) //200 指针p1中存储的是a的内存地址
	// 二级指针
	var p2 **int
	fmt.Println(p2)                                                   //<nil>
	p2 = &p1                                                          //p1本身是指针类型,再取p1的指针,赋值给二级指针类型p2
	fmt.Printf("%T %T %T\n", a, p1, p2)                               // int *int **int
	fmt.Println("p2的数值", p2)                                          //0xc0000cc020 二级指针p2中存储的就是指针p1的内存地址值
	fmt.Printf("p2的内存地址 %p\n", &p2)                                   // 0xc0000cc028
	fmt.Println("p2中存储的是p1的内存地址值,p1中存储的是a的内存地址值,所以*p2获取a的内存地址值", *p2) // a的内存地址值
	fmt.Println("p2中存储p1的内存地址值,p1中存储a的内存地址值,所以**p2获取的是a的值", **p2)     // 200

}
