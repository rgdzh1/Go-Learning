package main

import "fmt"

func main() {
	sum := getSum(1, 2, 3)
	fmt.Println(sum)
	result, _ := fun1() // _表示忽略的作用
	fmt.Println(result)
}

// 带返回值的函数
func getSum(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

// (sum int)这种情况, 就相当于指明道姓了 retun 后面不用跟变量名了.
func getSum1() (sum int) {
	sum = 10
	return
}

// 多返回值
func fun1() (int, int) {
	return 1, 2
}

// 多返回值
func fun2() (a int, b int) {
	return 1, 2
}

// return: 第一个作用是用来返回结果,第二用来结束函数执行
