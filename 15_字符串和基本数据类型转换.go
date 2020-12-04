package main

import (
	"fmt"
	"strconv"
)

func main() {
	// bool类型
	s1 := "true"
	b1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T %t\n", b1, b1)
	// 整数
	s2 := "100"
	i2, err := strconv.ParseInt(s2, 10, 64) // base:表示几进制,bitSize:表示最大位数
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T %d\n", i2, i2)
	// 整数转为字符串
	ss2 := strconv.FormatInt(i2, 10) // base:10进制
	fmt.Printf("%T %d\n", ss2, ss2)

	// itoa atoi
	i3, err := strconv.Atoi("-42") //转为int类型
	fmt.Printf("%T %d\n", i3, i3)

	ss3 := strconv.Itoa(7894) //转为字符串
	fmt.Printf(ss3)
}
