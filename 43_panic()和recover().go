package main

import "fmt"

// panic 恐慌
// recover 恢复
// go语言利用panic(),recover(),实现程序中极其特殊异常处理,
// panic() 让当前程序进入恐慌,中断程序的执行
// recover() 让程序恢复,必须在defer函数中执行

// 当外围函数的代码中发生了panic,只有其中所有的已经defer的函数全部执行完毕后,该panic才会被扩散到调用处

func funcM() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg, "但是,程序恢复了....")
		}
	}()
	defer fmt.Println("defer funcM...")
	for i := 0; i <= 100; i++ {
		if i == 5 {
			panic("引起程序恐慌") // panic() 中传入的参数最终会被recover()捕获并返回
		}
	}
}

func main() {
	defer fmt.Println("defer main...")
	funcM()
}
