package main

import (
	"fmt"
)

func main() {
	// 布尔值
	b1 := false
	b2 := true
	fmt.Println(b1, b2)
	b1 = true
	b2 = false
	fmt.Println(b1, b2)
	// ---整数型---
	// 有符号,最高位表示符号,其他表示数值
	// int8(-128 ~ 127)
	// int16
	// int32
	// int64
	// 无符号,所有位置都表示数值
	// uint8(0~255)
	// uint16
	// uint32
	// uint64
	// 平时使用时候就用int或者uint,具体多少位根据操作系统平台来确定的.32位操作系统就是int32,如此类推
	// 数据类型转换:Type(Value)
	var a1 int8 = 10
	var a2 int16
	a2 = int16(a1)
	fmt.Println(a2)

	// 100数据常数, 在需要的时候会自动转换
	a3 := 3.14
	a4 := a3 + 100
	fmt.Println(a4)

}
