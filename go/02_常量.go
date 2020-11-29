package main

import "fmt"

// 常量定义之后就算没有用过也不会报错
func main() {
	// 常量定义
	const PATH = "www.kkk.com"
	// 定义一组常量
	const C1, C2, C3 = 100, 3.14159, "X"
	const (
		NALE   = 0
		FEMALE = 1
	)
	// 一组常量中,如果某个常量没有初始值,默认和上一行一致
	const (
		a = 100
		b
	)
	fmt.Println(a, b)
	// iota
	// 相对于常量组来讲,定义一个常量自增1,在定义一个常量自增1.
	// 当出现新的常量组,iota 归零
	const (
		A = iota   // 0
		B          //iota = 1
		C          //iota = 2
		D = "haha" //iota = 3
		E          //"haha" iota = 4
		F = 100    //iota = 5
		G          //100 iota = 6
		H = iota   //iota = 7
		I          //iota = 8
	)
	const (
		J = iota //iota = 0
	)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	fmt.Println(I)
	fmt.Println(J)
}
