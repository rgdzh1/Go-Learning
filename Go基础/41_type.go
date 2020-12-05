package main

import (
	"strconv"
	"strings"
)

// type 定义新的类型
type myInt int
type myString string

// type定义函数类型
type myFun func(a, b int) string

// myFun这个就是定义的函数类型
func funccc() myFun {
	return func(a, b int) string {
		s1 := strconv.Itoa(a)
		s2 := strconv.Itoa(b)
		// return	s1+s2
		return strings.Join([]string{s1, s2}, "")
	}
}

// 类型别名
type myInt2 = int // 这样不是重新定义类型, 只是为该类型起了一个别名
