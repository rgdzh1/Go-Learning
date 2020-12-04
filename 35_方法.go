package main

import "fmt"

/**
方法:method
	一个方法就是一个包含了接受者的函数,接受者可以是命名类型或者结构体类型的一个值或者指针.
方法与函数对比:
意义:
方法: 某个结构体的行为功能,需要指定的接受者调用
函数: 一段独立功能的代码,可以直接调用
语法:
方法:方法名可以相同,只要接受者不同
函数:命名不能冲突
*/
type Student8 struct {
	name string
	age  int
	sex  string
}

func (receiver Student8) work() {
	fmt.Println(receiver.age, receiver.name, receiver.sex, "在工作")
}

func (receiver *Student8) work1() {
	fmt.Println(receiver.age, receiver.name, receiver.sex, "在工作")
}
func main() {
	s := new(Student8)
	s.name = "张三"
	s.sex = "男"
	s.age = 20
	s.work() //结构体指针调用work方法
	s.work1()
	s1 := Student8{
		name: "小刘",
		age:  17,
		sex:  "男",
	}
	s1.work()  //结构体调用
	s1.work1() //方法接收者是指针类型时候也可以使用结构体调用方法
}
