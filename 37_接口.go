package main

import "fmt"

/**
Go语言中,接口时一组方法的签名
Go语言中,接口和类型的实现是非侵入性的
*/
// 接口定义
type USB interface {
	start() // USB设备开始工作
	end()   // USB设备结束工作
}

// 实现类
type Mouse struct {
	name string
}
type FlashDisk struct {
	name string
}

// 实现接口方法
func (receiver Mouse) start() {
	fmt.Println("鼠标开始工作了")
}
func (receiver Mouse) end() {
	fmt.Println("鼠标结束工作了")
}

func (receiver FlashDisk) start() {
	fmt.Println("U盘开始工作了")
}
func (receiver FlashDisk) end() {
	fmt.Println("U盘结束工作了")
}

// 多态
func runDevice(usb USB) {
	usb.start()
	usb.end()
}

func main() {
	mouse := Mouse{
		"红色鼠标",
	}
	runDevice(mouse)
	disk := FlashDisk{name: "爱国者U盘"}
	runDevice(disk)

	//var usb USB 类型向上提升之后是无法访问子类中的数据的
	//usb = mouse
	//fmt.Println(usb.name)

	// 空接口
	// 不包含任何方法,正因为如此,所有的类型都实现了空接口,因此空接口可以存储任意类型的数值
	print(10)
	print3(mouse)

	map1 := make(map[string]interface{})
	map1["name"] = "张三"
	map1["age"] = 15
	map1["device"] = Mouse{
		"红色鼠标",
	}
	fmt.Println(map1)
}

type A interface {
	// 空接口
}

// 空接口作为参数
func print(a A) {
	fmt.Println(a)
}

// 匿名空接口
func print3(a interface{}) {
	fmt.Println(a)
}
