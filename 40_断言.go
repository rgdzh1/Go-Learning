package main

import "fmt"

type Animal interface {
	eat()
}
type Dog struct {
}
type Fish struct {
}

func (receiver Dog) eat() {
	fmt.Println("小狗喝水")
}
func (receiver Fish) eat() {
	fmt.Println("鱼儿喝水")
}

// 断言 对象.(类型) 返回当前类型的对象和是否为当前类型的判断值
func getType(animal Animal) {
	if ins, ok := animal.(Dog); ok {
		ins.eat()
	}
	if ins, ok := animal.(Fish); ok {
		ins.eat()
	}
}

func main() {
	dog := Dog{}
	getType(dog)
	fish := Fish{}
	getType(fish)
}
