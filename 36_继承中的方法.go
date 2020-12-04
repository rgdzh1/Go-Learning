package main

import "fmt"

// 方法是可以继承的,如果匿名字段实现了一个方法,那么包含该匿名字段的结构体也能调用该方法
type Person9 struct {
	name string
	age  int
}
type Student9 struct {
	Person9 // 模拟继承, 匿名字段,提升字段
	school  string
}

func (receiver Person9) eat() {
	fmt.Println("父类方法吃窝窝头")
}
func (receiver Student9) study() {
	fmt.Println("字类新增方法,学生学习啦...")
}
func (receiver Student9) eat() {
	fmt.Println("子类重写父类的方法,吃炸鸡喝啤酒...")
}
func main() {
	person9 := Person9{
		name: "烽火",
		age:  50,
	}
	person9.eat()
	student9 := Student9{Person9{"小明", 20}, "北校"}
	student9.eat() // 可以调用父类的方法
	student9.study()
}
