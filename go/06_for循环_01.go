package main

import "fmt"

func main() {
	// 第一种写法,i的作用域在for循环中,这是标准写法
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	// 第二种写法,i的作用域不局限于for循环中
	i := 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	// 第三种写法,类似while(true)
	for {
		fmt.Println(i)
		i++
	}
	/**
	break:彻底的结束循环
	continue:结束了某一次循环
	*/
	m := 1
	for {
		if m == 10 {
			continue
		}
		if m == 20 {
			break
		}
		fmt.Println(m)
	}
}
