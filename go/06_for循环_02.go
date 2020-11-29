package main

import "fmt"

func main() {
	/**
	break:彻底的结束循环
	continue:结束了某一次循环
	如果是在多层循环中, break和continue默认结束的是里层循环
	*/
	for m := 1; m <= 10; m++ {
		if m == 5 {
			// continue //结束某一次循环
			break // 彻底结束循环
		}
		fmt.Println(m)
	}
	// 多层循环中,break和continue结束外层循环可以使用贴标签的形式
out:
	for m := 1; m <= 10; m++ {
		for m := 1; m <= 10; m++ {
			if m == 5 {
				// continue out //贴标签的方式结束外层循环
				break out // 贴标签结束循环
			}
			fmt.Println(m)
		}
	}
}
