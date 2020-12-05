package main

import "fmt"

func main() {
	// 1. 省略switch后的变量相当于直接作用在true上
	switch {
	case true:
		fmt.Println("true...")
	case false:
		fmt.Println("false...")
	}
	// 写法一
	score := 60
	switch {
	case score < 60 && score >= 0:
		fmt.Println("不及格", score)
	case score >= 60 && score <= 70:
		fmt.Println("及格", score)
	default:
		fmt.Println("成绩有误...")
	}
	// 写法二
	switch "A" {
	case "A", "E", "I", "O", "U":
		fmt.Println("元音...")
	case "M", "N":
		fmt.Println("M或者H")
	}
	// 写法三,num只能在switch中使用
	switch num := 10; num {
	case 10, 4, 8:
		fmt.Println(num)
	}
	// break 可以使用在switch中, 也可以使用for循环中,强制结束case语句,从而结束switch分支.
	n := 2
	switch n {
	case 1:
		fmt.Println("我是熊大")
		fmt.Println("我是熊大")
		fmt.Println("我是熊大")
	case 2:
		fmt.Println("我是熊二")
		fmt.Println("我是熊二")
		break // 用于强制结束case,意味着switch语句被强制结束
		fmt.Println("我是熊二")
	}
	// fallthrough用于穿透switch,当siwtch中某个case匹配成功之后,就执行该case语句,如果遇到fallthrough,那么后面紧邻的case无需匹配,直接穿透执行
	// fallthrough只能在case语句中的最后一行
	m := 2
	switch m {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
		fallthrough // case最后一句,穿透执行
	case 3:
		fmt.Println("第三季度")
	case 4:
		fmt.Println("第四季度")
	}
}
