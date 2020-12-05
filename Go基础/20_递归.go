package main

import "fmt"

func main() {
	x := getSumX(10)
	fmt.Println(x)
}

func getSumX(n int) int {
	if n == 1 {
		// 递归调用的出口
		return 1
	}
	return getSumX(n-1) + n
}

// 递归调用斐波拉契
func getFibonacci(n int) int {
	if n == 1 || n == 2 {
		// 递归调用的出口
		return 1
	}
	return getFibonacci(n-1) + getFibonacci(n-1)
}
