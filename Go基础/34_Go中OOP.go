package main

/**
Go语言模拟继承关系
type A struct{
	field
}
type B struct{
	A // 匿名字段
}
模拟聚合关系
type A struct{
	field
}
type B struct{
	a A // 聚合关系
}
*/
type Person3 struct {
}
type Student3 struct {
	Person3 // 模拟继承关系
}

type Student4 struct {
	p Person // 模拟聚合关系
}
