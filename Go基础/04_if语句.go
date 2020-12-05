package main

import "fmt"

func main() {
	// 普通if写法
	num := 10
	if num > 8 {
		fmt.Println("num大于8")
	}
	fmt.Println("num等于10")
	// 另外一种写法,先初始化数据num,并且该num只能在该if当中使用,出了if就无法使用了
	if num := 1; num < 10 {
		fmt.Println("正数...", num)
	} else {
		fmt.Println("负数...", num)
	}
}
