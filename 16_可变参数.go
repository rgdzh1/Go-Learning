package main

import "fmt"

func main() {
	getSun()              //参数为空,和等于0
	getSun(1, 2, 3, 4, 5) //可变参数,参数多个,和等于15
	s1 := []int{1, 2, 3, 4, 5}
	getSun(s1...) // 将数组中的元素作为函数可变参数的实参
}

// 可变参数: 参数类型确定,但是参数个数不确定
// 可变参数相当于切片
// 注意:
// 1, 如果一个函数的参数是可变参数,同时还有其他参数,可变参数要放在参数列表的最后
// 2, 一个函数的参数列表中最多只能有一个可变参数
func getSun(nums ...int) {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println("总和等于: ", sum)
}
