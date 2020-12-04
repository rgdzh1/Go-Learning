package main

import (
	"errors"
	"fmt"
	"net"
	"os"
)

func main() {
	// 错误,可能出现问题的地方出现了问题,意料之中的问题称为错误
	// 异常,不应该出现问题的地方出现了问题,在意料之外的问题称为异常
	f, err := os.Open("test.txt")
	if err != nil {
		// log.Fatal(err) Fatal打断函数执行
		fmt.Println(err)
		return // 这样也是可以的
	}
	fmt.Println(f.Name(), "文件打开")

	// 创建err
	err = errors.New("创建错误")
	fmt.Println(err)
	fmt.Printf("%T \n", err)

	// 另一个创建err的方法
	err2 := fmt.Errorf("错误信息: %d \n", 100)
	fmt.Print(err2)
	fmt.Printf("%T \n", err2)
}

// 函数返回错误信息
func check() error {
	return errors.New("年龄不合法")
	return fmt.Errorf("错误码 %d\n", 10)
	return nil // 如果没有错误的话
}

// 对err进行断言
func checkErr() {
	host, err := net.LookupHost("www.baidu.com")
	if err != nil {
		if ins, ok := err.(*net.DNSError); ok {
			if ins.Timeout() {
				fmt.Println("请求超时")
			}
		}
		return
	}
	fmt.Println(host)
}
