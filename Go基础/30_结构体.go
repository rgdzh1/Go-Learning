package main

import "fmt"

func main() {
	// 1
	var person1 Person
	person1.name = "小王"
	person1.age = 18
	person1.sex = "男"
	person1.address = "北京"
	fmt.Printf("姓名 %s,年龄 %d,性别 %s,地址 %s\n", person1.name, person1.age, person1.sex, person1.address)
	// 2
	person2 := Person{}
	person1.name = "小李"
	person1.age = 50
	person1.sex = "女"
	person1.address = "湖南"
	fmt.Printf("姓名 %s,年龄 %d,性别 %s,地址 %s\n", person2.name, person2.age, person2.sex, person2.address)
	// 3
	person3 := Person{
		name:    "小明",
		age:     14,
		sex:     "男",
		address: "广西",
	}
	fmt.Println(person3)
	// 4
	person4 := Person{"小天", 24, "男", "黑龙江"}
	fmt.Println(person4)

	// 结构体是值类型,一般通过指针去操作结构体
	// 结构体深拷贝
	person5 := Person{"小白", 24, "男", "黑龙江"}
	fmt.Println(person5)
	fmt.Printf("%p %T\n", &person5, person5)
	person6 := person5
	fmt.Println(person6)
	fmt.Printf("%p %T\n", &person6, person6)
	person6.name = "小龙" // 改变了person6中的内容并不会改变person5中的内容
	fmt.Println(person5)
	fmt.Println(person6)
	// 结构体浅拷贝
	var person7 *Person
	person7 = &person5
	fmt.Println(person7)                     // &{小白 24 男 黑龙江}
	fmt.Printf("%p %T\n", &person7, person7) // 0xc000006030 *main.Person
	fmt.Println(*person7)                    // {小白 24 男 黑龙江}

	// (*person7).name = "二狗"
	person7.name = "二狗" // 简化调用结构体,通过指针操作,person5中的数据就修改了
	fmt.Println(person7)
	fmt.Println(person5)

	// new()-> 返回的不是nil指针,而是指向了新分配的类型的内存空间,只不过储存的值为空
	// new 用来创建结构体指针,返回的是一个指针,他会开辟内存空间
	// new 可以创建任意类型的指针,make返回的不是指针.
	person8 := new(Person)
	person8.name = "王麻子"
	person8.address = "武清"
	person8.age = 18
	fmt.Printf("%p %T\n", &person8, person8) // 0xc000006038 *main.Person
	fmt.Println(person8)
}

// 定义结构体
type Person struct {
	name    string
	age     int
	sex     string
	address string
}
