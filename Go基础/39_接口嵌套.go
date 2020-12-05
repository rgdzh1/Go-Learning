package main

type A1 interface {
	test1()
}
type B1 interface {
	test2()
}

type C interface {
	A1 // 接口嵌套
	B1
	test3()
}

type Cat struct {
	// Cat想要实现接口C,那么需要将接口A1与接口B1中的方法都实现
}

func (receiver Cat) test1() {

}
func (receiver Cat) test2() {

}
func (receiver Cat) test3() {

}

func main() {
	cat := Cat{}

	var a1 A1
	a1 = cat
	a1.test1()

	var a2 B1
	a2 = cat
	a2.test2()

	var a3 C
	a3 = cat
	a3.test2()
	a3.test1()
	a3.test3()
}
