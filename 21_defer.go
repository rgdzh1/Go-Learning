package main

import "fmt"

func main() { //外围函数
	//deferFun("hello")
	//fmt.Println("123456")
	//defer deferFun("world") // 当main函数(外围函数)执行完前一刻才能执行deferFun()
	//fmt.Println("王二狗")

	// 多个defer语句时,遵循先进后出的原则.
	//defer deferFun("hello")
	//fmt.Println("123456")
	//defer deferFun("world") // 当main函数(外围函数)执行完前一刻才能执行deferFun()
	//fmt.Println("王二狗")

	i := 2
	fmt.Println(i)
	defer deferFun1(i) // defer只是延迟了函数的执行,并没有延迟函数的调用,所以i变量被传入到deferFun1中时值其实是2
	i++
	fmt.Println("main中:", i)

}
func deferFun1(i int) {
	fmt.Println("defer", i)
}
func deferFun(s string) {
	fmt.Println(s)
}

// defer函数注意点
/**
1. 当外围函数中的语句正常执行完毕时,只有其中所有的延迟函数都执行完毕,外围函数才会真正的结束执行
2. 当执行外围函数中的return语句时, 只有其中所有的延迟函数都执行完毕后,外围函数才会真正返回
3. 当外围函数中的代码引发运行恐慌时,只有其中所有的延迟函数都执行完毕,该运行时恐慌才会真正被扩展至调用函数
*/
