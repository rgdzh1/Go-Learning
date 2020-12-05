package main

import "fmt"

type Book struct {
	bookName string
	price    float64
}

type Student1 struct {
	name string
	age  int
	book Book
}
type Student2 struct {
	name string
	age  int
	book *Book // 嵌套结构体使用指针,这样可以避免值传递
}

func main() {
	s1 := Student1{
		name: "王大",
		age:  18,
		book: Book{
			price:    100,
			bookName: "西游记",
		},
	}
	fmt.Println(s1)
	fmt.Printf("姓名 %s 书名 %s", s1.name, s1.book.bookName)

	b1 := Book{
		bookName: "水浒",
		price:    88,
	}
	s2 := Student2{
		name: "王二狗",
		age:  10,
		book: &b1, // 使用地址值
	}
	fmt.Println(s2)
	fmt.Printf("姓名 %s 书名 %s", s2.name, s2.book.bookName)
}
