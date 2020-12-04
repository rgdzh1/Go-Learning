package main

import "fmt"

func main() {
	aar1 := [4]int{1, 2, 3, 4}
	var p1 *[4]int
	p1 = &aar1
	fmt.Println(p1)
	fmt.Printf("数组的指针 %p\n", p1)  //0xc000010360
	fmt.Printf("p1的指针 %p\n", &p1) //0xc000006028

	// 数组指针操作数组
	(*p1)[0] = 100
	fmt.Println(aar1)
	p1[0] = 3000 // 简化写法
	fmt.Println(aar1)

	// 指针数组
	// 首先是数组,存储的数据类型是指针
	a := 1
	b := 2
	c := 3
	d := 4
	aar2 := [4]int{a, b, c, d}
	aar3 := [4]*int{&a, &b, &c, &d}
	fmt.Println(aar2)
	fmt.Println(aar3) //[0xc00000a0d8 0xc00000a100 0xc00000a108 0xc00000a110] 指针数组
	aar2[0] = 100     //修改数组中的值并不会修改变量a的值
	fmt.Println(aar2)
	fmt.Println(a)

	*aar3[0] = 500 // 修改了数组指针中的元素,变量a的值也修改了
	fmt.Println(a)
	fmt.Println(*aar3[0])

	b = 20000 // 修改了变量b的值,数组中的值不会修改,但是指针数组中的指针对应的值会修改
	fmt.Println(aar2)
	fmt.Println(*aar3[1])

	/**
	*[5]int: 指针,数组的指针
	[3]*int: 数组,数组中存储了指针元素
	*[5]*string: 指针,指针对应的数组中存储了指针元素
	*/

}
