package main

import "fmt"

func main() {
	// goto: 可以无条件的转移到指定的行代码处执行
	a := 10
LOOP:
	for a < 20 {
		if a == 15 {
			a += 1
			goto LOOP
		}
		fmt.Println(a)
		a++
	}
}
